package gvabe

import (
	"time"

	"github.com/btnguyen2k/henge"
	"main/src/gvabe/bo/product"
	"main/src/gvabe/bo/user"
	"main/src/itineris"
)

func authenticateApiCall(ctx *itineris.ApiContext) (*user.User, *itineris.ApiResult) {
	_, user, err := _currentUserFromContext(ctx)
	if err != nil {
		return nil, itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if user == nil {
		return nil, itineris.NewApiResult(itineris.StatusNoPermission).SetMessage("user not found")
	}
	return user, nil
}

var funcAppToMapTransform = func(m map[string]interface{}) map[string]interface{} {
	// transform input map
	result := map[string]interface{}{
		"id":           m[henge.FieldId],
		"t_created":    m[henge.FieldTimeCreated],
		"is_published": m[product.ProdAttrIsPublished],
		"name":         m[product.ProdAttrName],
		"desc":         m[product.ProdAttrDesc],
		"domains":      make([]string, 0),
	}

	// convert "creation timestamp" to UTC
	if t, ok := result["t_created"].(time.Time); ok {
		result["t_created"] = t.In(time.UTC)
	}

	// populate "domains" field
	if id, ok := result["id"].(string); ok {
		domainAppMappings, _ := domainAppMappingDao.Rget(id)
		domainList := make([]string, 0)
		for _, mapping := range domainAppMappings {
			domainList = append(domainList, mapping.Src)
		}
		result["domains"] = domainList
	}
	return result
}

// apiAdminProductList handles API call "adminProductList"
func apiAdminProductList(ctx *itineris.ApiContext, _ *itineris.ApiAuth, _ *itineris.ApiParams) *itineris.ApiResult {
	_, authResult := authenticateApiCall(ctx)
	if authResult != nil {
		return authResult
	}
	appList, err := appDao.GetAll(nil, nil)
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	data := make([]map[string]interface{}, 0)
	for _, a := range appList {
		data = append(data, a.ToMap(funcAppToMapTransform))
	}
	return itineris.NewApiResult(itineris.StatusOk).SetData(data)
}

// apiAdminStats handles API call "adminStats"
func apiAdminStats(ctx *itineris.ApiContext, _ *itineris.ApiAuth, _ *itineris.ApiParams) *itineris.ApiResult {
	_, authResult := authenticateApiCall(ctx)
	if authResult != nil {
		return authResult
	}
	appList, err := appDao.GetAll(nil, nil)
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	numApps := len(appList)
	numTopics, numPages := 0, 0
	for _, app := range appList {
		topicList, err := topicDao.GetAll(app, nil, nil)
		if err != nil {
			return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
		}
		numTopics += len(topicList)
		for _, topic := range topicList {
			numPages += topic.GetNumPages()
		}
	}
	data := map[string]interface{}{
		"num_apps":   numApps,
		"num_topics": numTopics,
		"num_pages":  numPages,
	}
	return itineris.NewApiResult(itineris.StatusOk).SetData(data)
}
