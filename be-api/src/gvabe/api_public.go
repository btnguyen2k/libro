package gvabe

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/btnguyen2k/consu/reddo"
	"main/src/goapi"
	"main/src/itineris"
)

// API handler "info"
func apiInfo(_ *itineris.ApiContext, _ *itineris.ApiAuth, _ *itineris.ApiParams) *itineris.ApiResult {
	var publicPEM []byte
	if pubDER, err := x509.MarshalPKIXPublicKey(rsaPubKey); err == nil {
		pubBlock := pem.Block{
			Type:    "PUBLIC KEY",
			Headers: nil,
			Bytes:   pubDER,
		}
		publicPEM = pem.EncodeToMemory(&pubBlock)
	} else {
		publicPEM = []byte(err.Error())
	}

	// var m runtime.MemStats
	result := map[string]interface{}{
		"app": map[string]interface{}{
			"name":        goapi.AppConfig.GetString("app.name"),
			"shortname":   goapi.AppConfig.GetString("app.shortname"),
			"version":     goapi.AppConfig.GetString("app.version"),
			"description": goapi.AppConfig.GetString("app.desc"),
		},
		"exter": map[string]interface{}{
			"app_id":   exterAppId,
			"base_url": exterBaseUrl,
		},
		"rsa_public_key": string(publicPEM),
		// "memory": map[string]interface{}{
		// 	"alloc":     m.Alloc,
		// 	"alloc_str": strconv.FormatFloat(float64(m.Alloc)/1024.0/1024.0, 'f', 1, 64) + " MiB",
		// 	"sys":       m.Sys,
		// 	"sys_str":   strconv.FormatFloat(float64(m.Sys)/1024.0/1024.0, 'f', 1, 64) + " MiB",
		// 	"gc":        m.NumGC,
		// },
	}
	return itineris.NewApiResult(itineris.StatusOk).SetData(result)
}

// API handler "systemInfo"
func apiSystemInfo(_ *itineris.ApiContext, _ *itineris.ApiAuth, _ *itineris.ApiParams) *itineris.ApiResult {
	data := lastSystemInfo()
	return itineris.NewApiResult(itineris.StatusOk).SetData(data)
}

func _doLoginExter(ctx *itineris.ApiContext, params *itineris.ApiParams) *itineris.ApiResult {
	token := _extractParam(params, "token", reddo.TypeString, "", nil)
	if token == "" {
		return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("empty token")
	}
	if DEBUG && exterRsaPubKey != nil {
		exterToken, err := parseExterJwt(token.(string))
		if err != nil {
			log.Printf("[DEBUG] Error parsing submitted JWT: %e", err)
		} else {
			log.Printf("[DEBUG] Submitted JWT: {Id: %s / Type: %s / AppId: %s / UserId: %s / UserName: %s}",
				exterToken.Id, exterToken.Type, exterToken.AppId, exterToken.UserId, exterToken.UserName)
		}
	}
	if exterClient == nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage("Exter login is not enabled")
	}
	resp, err := exterClient.VerifyLoginToken(token.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if resp.Status != 200 {
		return itineris.NewApiResult(itineris.StatusNoPermission).
			SetMessage(fmt.Sprintf("Exter login failed (%d): %s", resp.Status, resp.Message))
	}
	if exterRsaPubKey == nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).
			SetMessage(fmt.Sprintf("Exter login failed, please retry"))
	}
	exterJwt := resp.GetString("data")
	exterToken, err := parseExterJwt(exterJwt)
	if DEBUG {
		if err != nil {
			log.Printf("[DEBUG] Error parsing returned JWT: %e", err)
		} else {
			log.Printf("[DEBUG] Submitted JWT: {Id: %s / Type: %s / AppId: %s / UserId: %s / UserName: %s}",
				exterToken.Id, exterToken.Type, exterToken.AppId, exterToken.UserId, exterToken.UserName)
		}
	}
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if exterToken.Type != "login" {
		return itineris.NewApiResult(itineris.StatusNoPermission).
			SetMessage(fmt.Sprintf("Exter login failed, please retry"))
	}
	user, err := createUserFromExterToken(exterToken)
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if user == nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage("can not create user account, please try again")
	}
	claims, err := genLoginClaims(ctx.GetId(), &Session{
		ClientRef:   ctx.GetId(),
		Channel:     loginChannelExter,
		UserId:      user.GetId(),
		DisplayName: user.GetDisplayName(),
		CreatedAt:   time.Now(),
		ExpiredAt:   time.Unix(exterToken.ExpiresAt, 0),
		Data:        []byte(exterJwt),
	})
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	jwt, err := genJws(claims)
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	return itineris.NewApiResult(itineris.StatusOk).SetData(jwt)
}

func _doLoginForm(ctx *itineris.ApiContext, params *itineris.ApiParams) *itineris.ApiResult {
	username := _extractParam(params, "username", reddo.TypeString, "", nil)
	if username == "" {
		return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("empty username")
	}
	resultLoginFailed := itineris.NewApiResult(itineris.StatusNoPermission).SetMessage("login failed")
	password := _extractParam(params, "password", reddo.TypeString, "", nil)
	if password == "" {
		return resultLoginFailed
	}
	user, err := userDao.Get(username.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if user == nil {
		return resultLoginFailed
	}
	if encryptPassword(user.GetId(), password.(string)) != user.GetPassword() {
		return resultLoginFailed
	}
	now := time.Now()
	claims, err := genLoginClaims(ctx.GetId(), &Session{
		ClientRef:   ctx.GetId(),
		Channel:     loginChannelForm,
		UserId:      user.GetId(),
		DisplayName: user.GetDisplayName(),
		CreatedAt:   now,
		ExpiredAt:   now.Add(3600 * time.Second),
		Data:        nil,
	})
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	jwt, err := genJws(claims)
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	return itineris.NewApiResult(itineris.StatusOk).SetData(jwt)
}

/*
apiLogin handles API call "login".
  - Upon login successfully, this API returns the login-token as JWT.
*/
func apiLogin(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	mode := _extractParam(params, "mode", reddo.TypeString, "form", nil)
	switch strings.ToLower(mode.(string)) {
	case "exter":
		return _doLoginExter(ctx, params)
	default:
		return _doLoginForm(ctx, params)
	}
}

/*
apiVerifyLoginToken handles API call "verifyLoginToken".
  - Upon successful, this API returns the login-token as JWT.
*/
func apiVerifyLoginToken(_ *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	// firstly extract JWT token from request and convert it into claims
	token := _extractParam(params, "token", reddo.TypeString, "", nil)
	if token == "" {
		return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("empty token")
	}
	claims, err := parseLoginToken(token.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusNoPermission).SetMessage(err.Error())
	}
	if claims.isExpired() {
		return itineris.NewApiResult(itineris.StatusNoPermission).SetMessage(errorExpiredJwt.Error())
	}

	// lastly return the login-token encoded as JWT
	jwt, err := genJws(claims)
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	return itineris.NewApiResult(itineris.StatusOk).SetData(jwt)
}
