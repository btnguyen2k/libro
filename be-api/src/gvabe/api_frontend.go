package gvabe

import (
	"log"
	"strings"

	"github.com/btnguyen2k/consu/reddo"
	"main/src/gvabe/bo/libro"
	"main/src/itineris"
)

func authenticateFeApiCall(ctx *itineris.ApiContext) *itineris.ApiResult {
	return nil
}

/*----------------------------------------------------------------------*/

func _fetchProductForDomain(domain string) (prod *libro.Product, err error) {
	mapping, err := domainProductMappingDao.Get(domain)
	log.Printf("\t[DEBUG] - _fetchProductForDomain, domain %s / mapping %#v", domain, mapping)
	if err != nil {
		return nil, err
	}
	if mapping != nil {
		return productDao.Get(mapping.Dest)
	}
	return nil, nil
}

// apiFeGetProduct handles API call "feGetProduct"
func apiFeGetProduct(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	authResult := authenticateFeApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	domain := _extractParam(params, "domain", reddo.TypeString, "", nil)
	log.Printf("[DEBUG] - apiFeGetProduct, domain %s", domain)
	prod, err := _fetchProductForDomain(domain.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if tokens := strings.Split(domain.(string), ":"); prod == nil && len(tokens) > 1 {
		log.Printf("[DEBUG] - apiFeGetProduct, domain (2nd) %s", tokens[0])
		// handle case <host>:<port>
		prod, err = _fetchProductForDomain(tokens[0])
		if err != nil {
			return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
		}
	}
	if prod == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).SetMessage("product not found")
	}
	topicList, err := topicDao.GetAll(prod, nil, nil)
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	topicMapList := make([]map[string]interface{}, len(topicList))
	for i, topic := range topicList {
		topicMapList[i] = topic.ToMap(funcTopicToMapTransform)
	}
	prodMap := prod.ToMap(funcProductToMapTransform)
	prodMap["topics"] = topicMapList
	return itineris.NewApiResult(itineris.StatusOk).SetData(prodMap)
}

// apiFeGetTopic handles API call "feGetTopic"
func apiFeGetTopic(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	authResult := authenticateFeApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	domain := _extractParam(params, "domain", reddo.TypeString, "", nil)
	log.Printf("[DEBUG] - apiFeGetTopic, domain %s", domain)
	prod, err := _fetchProductForDomain(domain.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if tokens := strings.Split(domain.(string), ":"); prod == nil && len(tokens) > 1 {
		log.Printf("[DEBUG] - apiFeGetProduct, domain (2nd) %s", tokens[0])
		// handle case <host>:<port>
		prod, err = _fetchProductForDomain(tokens[0])
		if err != nil {
			return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
		}
	}
	if prod == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).SetMessage("product not found")
	}

	tid := _extractParam(params, "tid", reddo.TypeString, "", nil)
	topic, err := topicDao.Get(tid.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if topic == nil || topic.GetProductId() != prod.GetId() {
		return itineris.NewApiResult(itineris.StatusNotFound).SetMessage("topic not found")
	}

	pageList, err := pageDao.GetAll(topic, nil, nil)
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	pageMapList := make([]map[string]interface{}, len(pageList))
	for i, page := range pageList {
		pageMapList[i] = page.ToMap(funcPageToMapTransform)
	}
	topicMap := topic.ToMap(funcTopicToMapTransform)
	topicMap["pages"] = pageMapList
	return itineris.NewApiResult(itineris.StatusOk).SetData(topicMap)
}
