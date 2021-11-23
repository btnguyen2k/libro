package gvabe

import (
	"log"
	"strings"

	"github.com/btnguyen2k/consu/reddo"
	"main/src/gvabe/bo/product"
	"main/src/itineris"
)

func authenticateFeApiCall(ctx *itineris.ApiContext) *itineris.ApiResult {
	return nil
}

/*----------------------------------------------------------------------*/

func _fetchProductForDomain(domain string) (prod *product.Product, err error) {
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
