package gvabe

import (
	"reflect"
	"regexp"
	"strings"

	"main/src/gvabe/bo/user"
	"main/src/itineris"
)

/*
Setup API handlers: application register its api-handlers by calling router.SetHandler(apiName, apiHandlerFunc)
  - api-handler function must have the following signature: func (itineris.ApiContext, itineris.ApiAuth, itineris.ApiParams) *itineris.ApiResult
*/
func initApiHandlers(router *itineris.ApiRouter) {
	// pubic APIs
	router.SetHandler("info", apiInfo)
	router.SetHandler("login", apiLogin)
	router.SetHandler("verifyLoginToken", apiVerifyLoginToken)
	router.SetHandler("systemInfo", apiSystemInfo)

	// frontend APIs
	router.SetHandler("feGetProduct", apiFeGetProduct)
	router.SetHandler("feGetTopic", apiFeGetTopic)
	router.SetHandler("feGetUserProfile", apiFeGetUserProfile)
	router.SetHandler("adminUpdateUserProfile", apiAdminUpdateUserProfile)
	router.SetHandler("adminUpdateUserPassword", apiAdminUpdateUserPassword)

	// admin APIs
	router.SetHandler("adminGetStats", apiAdminGetStats)
	router.SetHandler("adminGetProductList", apiAdminGetProductList)
	router.SetHandler("adminAddProduct", apiAdminAddProduct)
	router.SetHandler("adminGetProduct", apiAdminGetProduct)
	router.SetHandler("adminUpdateProduct", apiAdminUpdateProduct)
	router.SetHandler("adminDeleteProduct", apiAdminDeleteProduct)
	router.SetHandler("adminMapDomain", apiAdminMapDomain)
	router.SetHandler("adminUnmapDomain", apiAdminUnmapDomain)
	router.SetHandler("adminGetProductTopics", apiAdminGetProductTopics)
	router.SetHandler("adminAddProductTopic", apiAdminAddProductTopic)
	router.SetHandler("adminDeleteProductTopic", apiAdminDeleteProductTopic)
	router.SetHandler("adminModifyProductTopic", apiAdminModifyProductTopic)
	router.SetHandler("adminUpdateProductTopic", apiAdminUpdateProductTopic)
	router.SetHandler("adminGetTopicPages", apiAdminGetTopicPages)
	router.SetHandler("adminAddTopicPage", apiAdminAddTopicPage)
	router.SetHandler("adminDeleteTopicPage", apiAdminDeleteTopicPage)
	router.SetHandler("adminModifyTopicPage", apiAdminModifyTopicPage)
	router.SetHandler("adminUpdateTopicPage", apiAdminUpdateTopicPage)
}

/*------------------------------ shared variables and functions ------------------------------*/

var (
	// those APIs will not need authentication.
	// "false" means client, however, needs to send app-id along with the API call
	// "true" means the API is free for public call
	publicApis = map[string]bool{
		"login":            false,
		"info":             true,
		"getApp":           false,
		"verifyLoginToken": true,
		"loginChannelList": true,
		"feGetProduct":     false,
		"feGetTopic":       false,
	}
)

// available since template-v0.2.0
func _currentUserFromContext(ctx *itineris.ApiContext) (*SessionClaims, *user.User, error) {
	sessClaims, ok := ctx.GetContextValue(ctxFieldSession).(*SessionClaims)
	if !ok || sessClaims == nil {
		return nil, nil, nil
	}
	user, err := userDao.Get(sessClaims.UserId)
	return sessClaims, user, err
}

// available since template-v0.2.0
func _extractParam(params *itineris.ApiParams, paramName string, typ reflect.Type, defValue interface{}, regexp *regexp.Regexp) interface{} {
	v, _ := params.GetParamAsType(paramName, typ)
	if v == nil {
		v = defValue
	}
	if v != nil {
		if _, ok := v.(string); ok {
			v = strings.TrimSpace(v.(string))
			if v.(string) == "" {
				v = defValue
			}
			if regexp != nil && !regexp.Match([]byte(v.(string))) {
				return nil
			}
		}
	}
	return v
}
