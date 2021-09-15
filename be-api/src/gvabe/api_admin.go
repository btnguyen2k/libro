package gvabe

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/btnguyen2k/consu/reddo"
	"github.com/btnguyen2k/consu/semita"
	"github.com/btnguyen2k/henge"
	"main/src/goapi"
	"main/src/gvabe/bo"
	"main/src/gvabe/bo/product"
	"main/src/gvabe/bo/user"
	"main/src/itineris"
	"main/src/utils"
)

func authenticateApiCall(ctx *itineris.ApiContext) (*user.User, *itineris.ApiResult) {
	_, user, err := _currentUserFromContext(ctx)
	if err != nil {
		return nil, itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if user == nil {
		return nil, itineris.NewApiResult(itineris.StatusNoPermission).SetMessage("action denied")
	}
	return user, nil
}

// apiAdminGetStats handles API call "adminGetStats"
func apiAdminGetStats(ctx *itineris.ApiContext, _ *itineris.ApiAuth, _ *itineris.ApiParams) *itineris.ApiResult {
	_, authResult := authenticateApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	prodList, err := productDao.GetAll(nil, nil)
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	numProds := len(prodList)
	numTopics, numPages := 0, 0
	for _, prod := range prodList {
		topicList, err := topicDao.GetAll(prod, nil, nil)
		if err != nil {
			return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
		}
		numTopics += len(topicList)
		for _, topic := range topicList {
			numPages += topic.GetNumPages()
		}
	}
	data := map[string]interface{}{
		"num_products": numProds,
		"num_topics":   numTopics,
		"num_pages":    numPages,
	}
	return itineris.NewApiResult(itineris.StatusOk).SetData(data)
}

var funcProductToMapTransform = func(m map[string]interface{}) map[string]interface{} {
	s := semita.NewSemita(m)

	// transform input map
	result := map[string]interface{}{
		"id":        m[henge.FieldId],
		"t_created": m[henge.FieldTimeCreated],
		"domains":   make([]string, 0),
	}
	result["is_published"], _ = s.GetValueOfType(fmt.Sprintf("%s.%s", bo.SerKeyAttrs, product.ProdAttrIsPublished), reddo.TypeBool)
	result["name"], _ = s.GetValueOfType(fmt.Sprintf("%s.%s", bo.SerKeyAttrs, product.ProdAttrName), reddo.TypeString)
	result["desc"], _ = s.GetValueOfType(fmt.Sprintf("%s.%s", bo.SerKeyAttrs, product.ProdAttrDesc), reddo.TypeString)

	// convert "creation timestamp" to UTC
	if t, ok := result["t_created"].(time.Time); ok {
		result["t_created"] = t.In(time.UTC)
	}

	// populate "domains" field
	if id, ok := result["id"].(string); ok {
		domainProductMappings, _ := domainProductMappingDao.Rget(id)
		domainList := make([]string, 0)
		for _, mapping := range domainProductMappings {
			domainList = append(domainList, mapping.Src)
		}
		result["domains"] = domainList
	}

	return result
}

// apiAdminGetProductList handles API call "adminGetProductList"
func apiAdminGetProductList(ctx *itineris.ApiContext, _ *itineris.ApiAuth, _ *itineris.ApiParams) *itineris.ApiResult {
	_, authResult := authenticateApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	prodList, err := productDao.GetAll(nil, nil)
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	data := make([]map[string]interface{}, 0)
	for _, prod := range prodList {
		data = append(data, prod.ToMap(funcProductToMapTransform))
	}
	return itineris.NewApiResult(itineris.StatusOk).SetData(data)
}

// apiAdminAddProduct handles API call "adminAddProduct"
func apiAdminAddProduct(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	_, authResult := authenticateApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	// extract params
	isPublished := _extractParam(params, "is_published", reddo.TypeBool, false, nil)
	name := _extractParam(params, "name", reddo.TypeString, "", nil)
	if name == "" {
		return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("name is empty")
	}
	desc := _extractParam(params, "description", reddo.TypeString, "", nil)
	domains := _extractParam(params, "domains", reddo.TypeString, "", nil)
	domains = strings.ToLower(domains.(string))

	domainList := regexp.MustCompile(`[,\s]+`).Split(domains.(string), -1)
	for _, domain := range domainList {
		mapping, err := domainProductMappingDao.Get(domain)
		if err != nil {
			return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
		}
		if mapping != nil {
			return itineris.NewApiResult(itineris.StatusNoPermission).SetMessage(fmt.Sprintf("domain %s has been used", domains))
		}
	}

	// create product
	product := product.NewProduct(goapi.AppVersionNumber, utils.UniqueIdSmall(), name.(string), desc.(string), isPublished.(bool))
	result, err := productDao.Create(product)
	if err != nil || !result {
		return itineris.NewApiResult(itineris.StatusErrorServer).
			SetMessage(fmt.Sprintf("cannot create product %s (error: %s)", name, err))
	}

	// map domains
	for _, domain := range domainList {
		result, err := domainProductMappingDao.Set(domain, product.GetId())
		if err != nil || !result {
			return itineris.NewApiResult(201).
				SetMessage(fmt.Sprintf("Product %s created, but cannot map domain %s to product (error: %s)", name, domain, err))
		}
	}

	return itineris.NewApiResult(itineris.StatusOk)
}

// apiAdminGetProduct handles API call "adminGetProduct"
func apiAdminGetProduct(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	_, authResult := authenticateApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	id := _extractParam(params, "id", reddo.TypeString, "", nil)
	product, err := productDao.Get(id.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if product == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).SetMessage("product not found")
	}
	return itineris.NewApiResult(itineris.StatusOk).SetData(product.ToMap(funcProductToMapTransform))
}

// apiAdminDeleteProduct handles API call "adminDeleteProduct"
func apiAdminDeleteProduct(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	_, authResult := authenticateApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	id := _extractParam(params, "id", reddo.TypeString, "", nil)
	product, err := productDao.Get(id.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if product == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).SetMessage("product not found")
	}

	domainProductMappings, err := domainProductMappingDao.Rget(product.GetId())
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	for _, mapping := range domainProductMappings {
		_, err := domainProductMappingDao.Remove(mapping.Src, mapping.Dest)
		if err != nil {
			msg := fmt.Sprintf("error while unmapping domain %s (product has not been deleted): %s", mapping.Src, err)
			return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(msg)
		}
	}

	_, err = productDao.Delete(product)
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	// if !ok {
	// 	return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage("cannot delete product")
	// }
	return itineris.NewApiResult(itineris.StatusOk)
}
