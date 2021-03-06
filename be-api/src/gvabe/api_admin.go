package gvabe

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/btnguyen2k/consu/reddo"
	"github.com/btnguyen2k/consu/semita"
	"github.com/btnguyen2k/henge"
	"main/src/goapi"
	"main/src/gvabe/bo"
	"main/src/gvabe/bo/libro"
	"main/src/gvabe/bo/user"
	"main/src/itineris"
	"main/src/respicite"
	"main/src/utils"
)

func authenticateAdminApiCall(ctx *itineris.ApiContext) (*user.User, *itineris.ApiResult) {
	_, user, err := _currentUserFromContext(ctx)
	if err != nil {
		return nil, itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if user == nil {
		return nil, itineris.NewApiResult(itineris.StatusNoPermission).SetMessage("action denied")
	}
	return user, nil
}

/*----------------------------------------------------------------------*/

// apiAdminGetStats handles API call "adminGetStats"
func apiAdminGetStats(ctx *itineris.ApiContext, _ *itineris.ApiAuth, _ *itineris.ApiParams) *itineris.ApiResult {
	_, authResult := authenticateAdminApiCall(ctx)
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

/*----------------------------------------------------------------------*/

var funcProductToMapTransform = func(m map[string]interface{}) map[string]interface{} {
	s := semita.NewSemita(m)

	// transform input map
	result := map[string]interface{}{
		"id":        m[henge.FieldId],
		"t_created": m[henge.FieldTimeCreated],
		"t_updated": m[henge.FieldTimeUpdated],
		"domains":   make([]string, 0),
	}
	result["is_published"], _ = s.GetValueOfType(fmt.Sprintf("%s.%s", bo.SerKeyAttrs, libro.ProdAttrIsPublished), reddo.TypeBool)
	result["name"], _ = s.GetValueOfType(fmt.Sprintf("%s.%s", bo.SerKeyAttrs, libro.ProdAttrName), reddo.TypeString)
	result["desc"], _ = s.GetValueOfType(fmt.Sprintf("%s.%s", bo.SerKeyAttrs, libro.ProdAttrDesc), reddo.TypeString)
	result["num_topics"], _ = s.GetValueOfType(fmt.Sprintf("%s.%s", bo.SerKeyAttrs, libro.ProdAttrNumTopics), reddo.TypeInt)
	result["contacts"], _ = s.GetValueOfType(fmt.Sprintf("%s.%s", bo.SerKeyAttrs, libro.ProdAttrContacts), libro.TypContactsMap)

	// convert "timestamp" to UTC
	if t, ok := result["t_created"].(time.Time); ok {
		result["t_created"] = t.In(time.UTC)
	}
	if t, ok := result["t_updated"].(time.Time); ok {
		result["t_updated"] = t.In(time.UTC)
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
	_, authResult := authenticateAdminApiCall(ctx)
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
	_, authResult := authenticateAdminApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	id := _extractParam(params, "id", reddo.TypeString, utils.UniqueIdSmall(), nil)
	prod, err := productDao.Get(id.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).
			SetMessage(fmt.Sprintf("error getting product [%s] (error: %s)", id, err))
	}
	if prod != nil {
		return itineris.NewApiResult(itineris.StatusErrorClient).
			SetMessage(fmt.Sprintf("product [%s] already existed", id))
	}
	id = strings.ToLower(id.(string))

	// extract params
	isPublished := _extractParam(params, "is_published", reddo.TypeBool, false, nil)
	name := _extractParam(params, "name", reddo.TypeString, "", nil)
	if name == "" {
		return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("name is empty")
	}
	desc := _extractParam(params, "desc", reddo.TypeString, "", nil)
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

	contactsMap := map[string]string{
		"email":    _extractParam(params, "contacts.email", reddo.TypeString, "", nil).(string),
		"website":  _extractParam(params, "contacts.website", reddo.TypeString, "", nil).(string),
		"github":   _extractParam(params, "contacts.github", reddo.TypeString, "", nil).(string),
		"facebook": _extractParam(params, "contacts.facebook", reddo.TypeString, "", nil).(string),
		"linkedin": _extractParam(params, "contacts.linkedin", reddo.TypeString, "", nil).(string),
		"slack":    _extractParam(params, "contacts.slack", reddo.TypeString, "", nil).(string),
		"twitter":  _extractParam(params, "contacts.twitter", reddo.TypeString, "", nil).(string),
	}

	// create product
	prod = libro.NewProduct(goapi.AppVersionNumber, utils.UniqueIdSmall(), name.(string), desc.(string), isPublished.(bool))
	prod.SetId(id.(string))
	prod.SetContacts(contactsMap)
	result, err := productDao.Create(prod)
	if err != nil || !result {
		return itineris.NewApiResult(itineris.StatusErrorServer).
			SetMessage(fmt.Sprintf("cannot create product %s (error: %s)", name, err))
	}

	// map domains
	for _, domain := range domainList {
		result, err := domainProductMappingDao.Set(domain, prod.GetId())
		if err != nil || !result {
			return itineris.NewApiResult(201).
				SetMessage(fmt.Sprintf("Product %s created, but cannot map domain %s to product (error: %s)", name, domain, err))
		}
	}

	return itineris.NewApiResult(itineris.StatusOk)
}

// apiAdminGetProduct handles API call "adminGetProduct"
func apiAdminGetProduct(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	_, authResult := authenticateAdminApiCall(ctx)
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

// apiAdminUpdateProduct handles API call "adminUpdateProduct"
func apiAdminUpdateProduct(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	_, authResult := authenticateAdminApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	id := _extractParam(params, "id", reddo.TypeString, "", nil)
	prod, err := productDao.Get(id.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if prod == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).SetMessage("product not found")
	}

	// extract params
	isPublished := _extractParam(params, "is_published", reddo.TypeBool, false, nil)
	name := _extractParam(params, "name", reddo.TypeString, "", nil)
	if name == "" {
		return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("name is empty")
	}
	desc := _extractParam(params, "desc", reddo.TypeString, "", nil)

	contactsMap := map[string]string{
		"email":    _extractParam(params, "contacts.email", reddo.TypeString, "", nil).(string),
		"website":  _extractParam(params, "contacts.website", reddo.TypeString, "", nil).(string),
		"github":   _extractParam(params, "contacts.github", reddo.TypeString, "", nil).(string),
		"facebook": _extractParam(params, "contacts.facebook", reddo.TypeString, "", nil).(string),
		"linkedin": _extractParam(params, "contacts.linkedin", reddo.TypeString, "", nil).(string),
		"slack":    _extractParam(params, "contacts.slack", reddo.TypeString, "", nil).(string),
		"twitter":  _extractParam(params, "contacts.twitter", reddo.TypeString, "", nil).(string),
	}

	// update product
	prod.SetPublished(isPublished.(bool)).SetName(name.(string)).SetDescription(desc.(string))
	prod.SetContacts(contactsMap)
	result, err := productDao.Update(prod)
	if err != nil || !result {
		return itineris.NewApiResult(itineris.StatusErrorServer).
			SetMessage(fmt.Sprintf("cannot update product [%s/%s] (error: %s)", prod.GetId(), prod.GetName(), err))
	}

	return itineris.NewApiResult(itineris.StatusOk)
}

// apiAdminDeleteProduct handles API call "adminDeleteProduct"
func apiAdminDeleteProduct(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	_, authResult := authenticateAdminApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	id := _extractParam(params, "id", reddo.TypeString, "", nil)
	prod, err := productDao.Get(id.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if prod == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).SetMessage("product not found")
	}

	_, err = productDao.Delete(prod)
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}

	// TODO: post-production-deletion clean-up via event-driven manner
	go func(prod *libro.Product) {
		// unmap domains
		if domainProductMappings, err := domainProductMappingDao.Rget(prod.GetId()); err != nil {
			log.Printf("[WARN] Post-delete product [%s] - Error getting mapped domain names: %e", prod.GetId(), err)
		} else {
			for _, mapping := range domainProductMappings {
				if _, err := domainProductMappingDao.Remove(mapping.Src, mapping.Dest); err != nil {
					log.Printf("[WARN] Post-delete product [%s] - Error unmapping domain names [%s]: %e", prod.GetId(), mapping.Src, err)
				}
			}
		}

		// delete topics and pages
		if topics, err := topicDao.GetAll(prod, nil, nil); err != nil {
			log.Printf("[WARN] Post-delete product [%s] - Error getting all topics for product: %e", prod.GetId(), err)
		} else {
			for _, topic := range topics {
				if pages, err := pageDao.GetAll(topic, nil, nil); err != nil {
					log.Printf("[WARN] Post-delete product [%s] - Error getting all pages for topic [%s]: %e", prod.GetId(), topic.GetId(), err)
				} else {
					for _, page := range pages {
						if _, err = pageDao.Delete(page); err != nil {
							log.Printf("[WARN] Post-delete product [%s] - Error deleting page [%s]: %e", prod.GetId(), page.GetId(), err)
						}
					}
				}
				if _, err = topicDao.Delete(topic); err != nil {
					log.Printf("[WARN] Post-delete product [%s] - Error deleting topic [%s]: %e", prod.GetId(), topic.GetId(), err)
				}
			}
		}
	}(prod)

	return itineris.NewApiResult(itineris.StatusOk)
}

/*----------------------------------------------------------------------*/

// apiAdminMapDomain handles API call "adminMapDomain"
func apiAdminMapDomain(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	_, authResult := authenticateAdminApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	// extract params
	productId := _extractParam(params, "pid", reddo.TypeString, "", nil)
	domainName := _extractParam(params, "domain", reddo.TypeString, "", nil)
	if productId == "" || domainName == "" {
		return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("name is empty")
	}
	product, err := productDao.Get(productId.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if product == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).SetMessage("product not found")
	}

	domainName = strings.ToLower(domainName.(string))
	result, err := domainProductMappingDao.Set(domainName.(string), product.GetId())
	if err == respicite.ErrDuplicated {
		return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage(fmt.Sprintf("Domain '%s' has already been mapped to a product.", domainName))
	}
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if !result {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(fmt.Sprintf("Unknown error while mapping domain '%s' to product.", domainName))
	}

	mappings, err := domainProductMappingDao.Rget(product.GetId())
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	domainList := make([]string, 0)
	for _, mapping := range mappings {
		domainList = append(domainList, mapping.Src)
	}

	return itineris.NewApiResult(itineris.StatusOk).SetData(domainList)
}

// apiAdminUnmapDomain handles API call "adminUnmapDomain"
func apiAdminUnmapDomain(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	_, authResult := authenticateAdminApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	// extract params
	productId := _extractParam(params, "pid", reddo.TypeString, "", nil)
	domainName := _extractParam(params, "domain", reddo.TypeString, "", nil)
	if productId == "" || domainName == "" {
		return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("name is empty")
	}
	product, err := productDao.Get(productId.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if product == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).SetMessage("product not found")
	}

	domainName = strings.ToLower(domainName.(string))
	_, err = domainProductMappingDao.Remove(domainName.(string), product.GetId())
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}

	mappings, err := domainProductMappingDao.Rget(product.GetId())
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	domainList := make([]string, 0)
	for _, mapping := range mappings {
		domainList = append(domainList, mapping.Src)
	}

	return itineris.NewApiResult(itineris.StatusOk).SetData(domainList)
}

/*----------------------------------------------------------------------*/

var funcTopicToMapTransform = func(m map[string]interface{}) map[string]interface{} {
	s := semita.NewSemita(m)

	// transform input map
	result := map[string]interface{}{
		"id":        m[henge.FieldId],
		"t_created": m[henge.FieldTimeCreated],
		"t_updated": m[henge.FieldTimeUpdated],
	}
	result["product_id"], _ = s.GetValueOfType(fmt.Sprintf("%s.%s", bo.SerKeyFields, libro.TopicFieldProductId), reddo.TypeString)
	result["pos"], _ = s.GetValueOfType(fmt.Sprintf("%s.%s", bo.SerKeyFields, libro.TopicFieldPosition), reddo.TypeInt)
	result["title"], _ = s.GetValueOfType(fmt.Sprintf("%s.%s", bo.SerKeyAttrs, libro.TopicAttrTitle), reddo.TypeString)
	result["icon"], _ = s.GetValueOfType(fmt.Sprintf("%s.%s", bo.SerKeyAttrs, libro.TopicAttrIcon), reddo.TypeString)
	result["summary"], _ = s.GetValueOfType(fmt.Sprintf("%s.%s", bo.SerKeyAttrs, libro.TopicAttrSummary), reddo.TypeString)
	result["num_pages"], _ = s.GetValueOfType(fmt.Sprintf("%s.%s", bo.SerKeyAttrs, libro.TopicAttrNumPages), reddo.TypeInt)

	// convert "timestamp" to UTC
	if t, ok := result["t_created"].(time.Time); ok {
		result["t_created"] = t.In(time.UTC)
	}
	if t, ok := result["t_updated"].(time.Time); ok {
		result["t_updated"] = t.In(time.UTC)
	}

	return result
}

// apiAdminGetProductTopics handles API call "adminGetProductTopics"
func apiAdminGetProductTopics(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	_, authResult := authenticateAdminApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	pid := _extractParam(params, "pid", reddo.TypeString, "", nil)
	product, err := productDao.Get(pid.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if product == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).SetMessage("product not found")
	}

	topicList, err := topicDao.GetAll(product, nil, nil)
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	data := make([]map[string]interface{}, 0)
	for _, topic := range topicList {
		data = append(data, topic.ToMap(funcTopicToMapTransform))
	}
	return itineris.NewApiResult(itineris.StatusOk).SetData(data)
}

// apiAdminAddProductTopic handles API call "adminAddProductTopic"
func apiAdminAddProductTopic(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	_, authResult := authenticateAdminApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	pid := _extractParam(params, "pid", reddo.TypeString, "", nil)
	product, err := productDao.Get(pid.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).
			SetMessage(fmt.Sprintf("error getting product [%s] (error: %s)", pid, err))
	}
	if product == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).
			SetMessage(fmt.Sprintf("product not found [%s]", pid))
	}

	id := _extractParam(params, "id", reddo.TypeString, utils.UniqueIdSmall(), nil)
	topic, err := topicDao.Get(id.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).
			SetMessage(fmt.Sprintf("error getting topic [%s] (error: %s)", id, err))
	}
	if topic != nil {
		return itineris.NewApiResult(itineris.StatusErrorClient).
			SetMessage(fmt.Sprintf("topic [%s] already existed", id))
	}
	id = strings.ToLower(id.(string))

	icon := _extractParam(params, "icon", reddo.TypeString, "", nil)
	icon = strings.ToLower(icon.(string))
	title := _extractParam(params, "title", reddo.TypeString, "", nil)
	summary := _extractParam(params, "summary", reddo.TypeString, "", nil)

	if title == "" {
		return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("title is empty")
	}

	topic = libro.NewTopic(goapi.AppVersionNumber, product, title.(string), icon.(string), summary.(string))
	topic.SetId(id.(string))
	result, err := topicDao.Create(topic)
	if err != nil || !result {
		return itineris.NewApiResult(itineris.StatusErrorServer).
			SetMessage(fmt.Sprintf("cannot create topic [%s/%s] (error: %s)", id, title, err))
	}

	// TODO: update product's stats via event-driven manner
	go func(topic *libro.Topic) {
		prod, err := productDao.Get(topic.GetProductId())
		if err != nil {
			log.Printf("[WARN] Post-add topic [%s] - Error getting product [%s]: %e", topic.GetId(), topic.GetProductId(), err)
			return
		}

		topics, err := topicDao.GetAll(prod, nil, nil)
		if err != nil {
			log.Printf("[WARN] Post-add topic [%s] - Error getting all topics for product [%s]: %e", topic.GetId(), topic.GetProductId(), err)
			return
		}
		prod.SetNumTopics(len(topics))
		ok, err := productDao.Update(prod)
		if err != nil || !ok {
			log.Printf("[WARN] Post-add topic [%s] - Cannot update product stats [%s]: %#v / %e", topic.GetId(), topic.GetProductId(), ok, err)
		}
	}(topic)

	return itineris.NewApiResult(itineris.StatusOk)
}

// apiAdminDeleteProductTopic handles API call "adminDeleteProductTopic"
func apiAdminDeleteProductTopic(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	_, authResult := authenticateAdminApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	pid := _extractParam(params, "pid", reddo.TypeString, "", nil)
	prod, err := productDao.Get(pid.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).
			SetMessage(fmt.Sprintf("error getting product [%s] (error: %s)", pid, err))
	}
	if prod == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).
			SetMessage(fmt.Sprintf("product not found [%s]", pid))
	}

	id := _extractParam(params, "id", reddo.TypeString, "", nil)
	topic, err := topicDao.Get(id.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).
			SetMessage(fmt.Sprintf("error getting topic [%s] (error: %s)", id, err))
	}
	if topic == nil || topic.GetProductId() != prod.GetId() {
		return itineris.NewApiResult(itineris.StatusNotFound).
			SetMessage(fmt.Sprintf("topic not found [%s]", id))
	}

	_, err = topicDao.Delete(topic)
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}

	// TODO: delete pages & update product's stats via event-driven manner
	go func(topic *libro.Topic) {
		prod, err := productDao.Get(topic.GetProductId())
		if err != nil {
			log.Printf("[WARN] Post-delete topic [%s] - Error getting product [%s]: %e", topic.GetId(), topic.GetProductId(), err)
			return
		}

		topics, err := topicDao.GetAll(prod, nil, nil)
		if err != nil {
			log.Printf("[WARN] Post-delete topic [%s] - Error getting all topics for product [%s]: %e", topic.GetId(), topic.GetProductId(), err)
			return
		}
		prod.SetNumTopics(len(topics))
		ok, err := productDao.Update(prod)
		if err != nil || !ok {
			log.Printf("[WARN] Post-delete topic [%s] - Cannot update product stats [%s]: %#v / %e", topic.GetId(), topic.GetProductId(), ok, err)
		}

		pages, err := pageDao.GetAll(topic, nil, nil)
		if err != nil {
			log.Printf("[WARN] Post-delete topic [%s] - Error getting all pages for topic [%s]: %e", topic.GetId(), topic.GetId(), err)
			return
		}
		for _, page := range pages {
			_, err := pageDao.Delete(page)
			if err != nil {
				log.Printf("[WARN] Post-delete topic [%s] - Error deleting page [%s]: %e", topic.GetId(), page.GetId(), err)
			}
		}
	}(topic)

	return itineris.NewApiResult(itineris.StatusOk)
}

// apiAdminModifyProductTopic handles API call "adminModifyProductTopic"
func apiAdminModifyProductTopic(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	_, authResult := authenticateAdminApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	pid := _extractParam(params, "pid", reddo.TypeString, "", nil)
	prod, err := productDao.Get(pid.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).
			SetMessage(fmt.Sprintf("error getting product [%s] (error: %s)", pid, err))
	}
	if prod == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).
			SetMessage(fmt.Sprintf("product not found [%s]", pid))
	}

	topicList, err := topicDao.GetAll(prod, nil, nil)
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).
			SetMessage(fmt.Sprintf("error getting topics for product [%s/%s] (error: %s)", pid, prod.GetName(), err))
	}

	id := _extractParam(params, "id", reddo.TypeString, "", nil)
	id = strings.ToLower(id.(string))
	found := -1
	for i, topic := range topicList {
		if topic.GetId() == id {
			found = i
			break
		}
	}

	if found < 0 {
		return itineris.NewApiResult(itineris.StatusNotFound).
			SetMessage(fmt.Sprintf("topic not found [%s]", id))
	}

	modifyAction := _extractParam(params, "action", reddo.TypeString, "", nil)
	switch modifyAction {
	case "move_up":
		if found == 0 {
			// at top, can not be moved up
			break
		}
		prev, curr := topicList[found-1], topicList[found]
		pCurr := curr.GetPosition()
		curr.SetPosition(pCurr - 1)
		prev.SetPosition(pCurr)
		_, eCurr := topicDao.Update(curr)
		_, ePrev := topicDao.Update(prev)
		if eCurr != nil || ePrev != nil {
			return itineris.NewApiResult(itineris.StatusErrorServer).
				SetMessage(fmt.Sprintf("error updating topics [%s/%s] (error: %s/%s)", prev.GetId(), curr.GetId(), ePrev, eCurr))
		}
	case "move_down":
		if found == len(topicList)-1 {
			// at bottom, can not be moved down
			break
		}
		curr, next := topicList[found], topicList[found+1]
		pCurr := curr.GetPosition()
		curr.SetPosition(pCurr + 1)
		next.SetPosition(pCurr)
		_, eCurr := topicDao.Update(curr)
		_, eNext := topicDao.Update(next)
		if eCurr != nil || eNext != nil {
			return itineris.NewApiResult(itineris.StatusErrorServer).
				SetMessage(fmt.Sprintf("error updating topics [%s/%s] (error: %s/%s)", curr.GetId(), next.GetId(), eCurr, eNext))
		}
	default:
		return itineris.NewApiResult(itineris.StatusErrorClient).
			SetMessage(fmt.Sprintf("invalid action: %s", modifyAction))
	}

	return itineris.NewApiResult(itineris.StatusOk)
}

// apiAdminUpdateProductTopic handles API call "adminUpdateProductTopic"
func apiAdminUpdateProductTopic(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	_, authResult := authenticateAdminApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	pid := _extractParam(params, "pid", reddo.TypeString, "", nil)
	prod, err := productDao.Get(pid.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).
			SetMessage(fmt.Sprintf("error getting product [%s] (error: %s)", pid, err))
	}
	if prod == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).
			SetMessage(fmt.Sprintf("product not found [%s]", pid))
	}

	id := _extractParam(params, "id", reddo.TypeString, "", nil)
	topic, err := topicDao.Get(id.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).
			SetMessage(fmt.Sprintf("error getting topic [%s] (error: %s)", id, err))
	}
	if topic == nil || topic.GetProductId() != prod.GetId() {
		return itineris.NewApiResult(itineris.StatusNotFound).
			SetMessage(fmt.Sprintf("topic not found [%s]", id))
	}

	icon := _extractParam(params, "icon", reddo.TypeString, "", nil)
	icon = strings.ToLower(icon.(string))
	title := _extractParam(params, "title", reddo.TypeString, "", nil)
	summary := _extractParam(params, "summary", reddo.TypeString, "", nil)

	if title == "" {
		return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("title is empty")
	}

	topic.SetIcon(icon.(string)).SetTitle(title.(string)).SetSummary(summary.(string))
	result, err := topicDao.Update(topic)
	if err != nil || !result {
		return itineris.NewApiResult(itineris.StatusErrorServer).
			SetMessage(fmt.Sprintf("cannot update topic [%s/%s] (error: %s)", topic.GetId(), topic.GetTitle(), err))
	}

	return itineris.NewApiResult(itineris.StatusOk)
}

/*----------------------------------------------------------------------*/

var funcPageToMapTransform = func(m map[string]interface{}) map[string]interface{} {
	s := semita.NewSemita(m)

	// transform input map
	result := map[string]interface{}{
		"id":        m[henge.FieldId],
		"t_created": m[henge.FieldTimeCreated],
		"t_updated": m[henge.FieldTimeUpdated],
	}
	result["product_id"], _ = s.GetValueOfType(fmt.Sprintf("%s.%s", bo.SerKeyFields, libro.PageFieldProductId), reddo.TypeString)
	result["topic_id"], _ = s.GetValueOfType(fmt.Sprintf("%s.%s", bo.SerKeyFields, libro.PageFieldTopicId), reddo.TypeString)
	result["pos"], _ = s.GetValueOfType(fmt.Sprintf("%s.%s", bo.SerKeyFields, libro.PageFieldPosition), reddo.TypeInt)
	result["title"], _ = s.GetValueOfType(fmt.Sprintf("%s.%s", bo.SerKeyAttrs, libro.PageAttrTitle), reddo.TypeString)
	result["icon"], _ = s.GetValueOfType(fmt.Sprintf("%s.%s", bo.SerKeyAttrs, libro.PageAttrIcon), reddo.TypeString)
	result["summary"], _ = s.GetValueOfType(fmt.Sprintf("%s.%s", bo.SerKeyAttrs, libro.PageAttrSummary), reddo.TypeString)
	result["content"], _ = s.GetValueOfType(fmt.Sprintf("%s.%s", bo.SerKeyAttrs, libro.PageAttrContent), reddo.TypeString)

	// convert "timestamp" to UTC
	if t, ok := result["t_created"].(time.Time); ok {
		result["t_created"] = t.In(time.UTC)
	}
	if t, ok := result["t_updated"].(time.Time); ok {
		result["t_updated"] = t.In(time.UTC)
	}

	return result
}

// apiAdminGetTopicPages handles API call "adminGetTopicPages"
func apiAdminGetTopicPages(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	_, authResult := authenticateAdminApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	tid := _extractParam(params, "tid", reddo.TypeString, "", nil)
	topic, err := topicDao.Get(tid.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if topic == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).SetMessage("topic not found")
	}

	pageList, err := pageDao.GetAll(topic, nil, nil)
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	data := make([]map[string]interface{}, 0)
	for _, page := range pageList {
		data = append(data, page.ToMap(funcPageToMapTransform))
	}
	return itineris.NewApiResult(itineris.StatusOk).SetData(data)
}

// apiAdminAddTopicPage handles API call "adminAddTopicPage"
func apiAdminAddTopicPage(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	_, authResult := authenticateAdminApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	tid := _extractParam(params, "tid", reddo.TypeString, "", nil)
	topic, err := topicDao.Get(tid.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if topic == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).SetMessage("topic not found")
	}

	id := _extractParam(params, "id", reddo.TypeString, utils.UniqueIdSmall(), nil)
	page, err := pageDao.Get(id.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).
			SetMessage(fmt.Sprintf("error getting page [%s] (error: %s)", id, err))
	}
	if page != nil {
		return itineris.NewApiResult(itineris.StatusErrorClient).
			SetMessage(fmt.Sprintf("page [%s] already existed", id))
	}
	id = strings.ToLower(id.(string))

	icon := _extractParam(params, "icon", reddo.TypeString, "", nil)
	icon = strings.ToLower(icon.(string))
	title := _extractParam(params, "title", reddo.TypeString, "", nil)
	summary := _extractParam(params, "summary", reddo.TypeString, "", nil)
	content := _extractParam(params, "content", reddo.TypeString, "", nil)

	if title == "" {
		return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("title is empty")
	}

	page = libro.NewPage(goapi.AppVersionNumber, topic, title.(string), icon.(string), summary.(string), content.(string))
	page.SetId(id.(string))
	result, err := pageDao.Create(page)
	if err != nil || !result {
		return itineris.NewApiResult(itineris.StatusErrorServer).
			SetMessage(fmt.Sprintf("cannot create page [%s/%s] (error: %s)", id, title, err))
	}

	// TODO: update topic's stats via event-driven manner
	go func(page *libro.Page) {
		topic, err := topicDao.Get(page.GetTopicId())
		if err != nil {
			log.Printf("[WARN] Post-add page [%s] - Error getting topic [%s]: %e", page.GetId(), page.GetTopicId(), err)
			return
		}

		pages, err := pageDao.GetAll(topic, nil, nil)
		if err != nil {
			log.Printf("[WARN] Post-add page [%s] - Error getting all pages for topic [%s]: %e", page.GetId(), page.GetTopicId(), err)
			return
		}
		topic.SetNumPages(len(pages))
		ok, err := topicDao.Update(topic)
		if err != nil || !ok {
			log.Printf("[WARN] Post-add page [%s] - Cannot update topic stats [%s]: %#v / %e", page.GetId(), page.GetTopicId(), ok, err)
		}
	}(page)

	return itineris.NewApiResult(itineris.StatusOk)
}

// apiAdminDeleteTopicPage handles API call "adminDeleteTopicPage"
func apiAdminDeleteTopicPage(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	_, authResult := authenticateAdminApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	tid := _extractParam(params, "tid", reddo.TypeString, "", nil)
	topic, err := topicDao.Get(tid.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if topic == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).SetMessage("topic not found")
	}

	id := _extractParam(params, "id", reddo.TypeString, "", nil)
	page, err := pageDao.Get(id.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).
			SetMessage(fmt.Sprintf("error getting page [%s] (error: %s)", id, err))
	}
	if page == nil || page.GetTopicId() != topic.GetId() {
		return itineris.NewApiResult(itineris.StatusNotFound).
			SetMessage(fmt.Sprintf("page not found [%s]", id))
	}

	_, err = pageDao.Delete(page)
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}

	// TODO: update topic's stats via event-driven manner
	go func(page *libro.Page) {
		topic, err := topicDao.Get(page.GetTopicId())
		if err != nil {
			log.Printf("[WARN] Post-delete page [%s] - Error getting topic [%s]: %e", page.GetId(), page.GetTopicId(), err)
			return
		}

		pages, err := pageDao.GetAll(topic, nil, nil)
		if err != nil {
			log.Printf("[WARN] Post-delete page [%s] - Error getting all pages for topic [%s]: %e", page.GetId(), page.GetTopicId(), err)
			return
		}
		topic.SetNumPages(len(pages))
		ok, err := topicDao.Update(topic)
		if err != nil || !ok {
			log.Printf("[WARN] Post-delete page [%s] - Cannot update topic stats [%s]: %#v / %e", page.GetId(), page.GetTopicId(), ok, err)
		}
	}(page)

	return itineris.NewApiResult(itineris.StatusOk)
}

// apiAdminModifyTopicPage handles API call "adminModifyTopicPage"
func apiAdminModifyTopicPage(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	_, authResult := authenticateAdminApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	tid := _extractParam(params, "tid", reddo.TypeString, "", nil)
	topic, err := topicDao.Get(tid.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if topic == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).SetMessage("topic not found")
	}

	pageList, err := pageDao.GetAll(topic, nil, nil)
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).
			SetMessage(fmt.Sprintf("error getting pages for topic [%s/%s] (error: %s)", tid, topic.GetTitle(), err))
	}

	id := _extractParam(params, "id", reddo.TypeString, "", nil)
	id = strings.ToLower(id.(string))
	found := -1
	for i, page := range pageList {
		if page.GetId() == id {
			found = i
			break
		}
	}

	if found < 0 {
		return itineris.NewApiResult(itineris.StatusNotFound).
			SetMessage(fmt.Sprintf("page not found [%s]", id))
	}

	modifyAction := _extractParam(params, "action", reddo.TypeString, "", nil)
	switch modifyAction {
	case "move_up":
		if found == 0 {
			// at top, can not be moved up
			break
		}
		prev, curr := pageList[found-1], pageList[found]
		pCurr := curr.GetPosition()
		curr.SetPosition(pCurr - 1)
		prev.SetPosition(pCurr)
		_, eCurr := pageDao.Update(curr)
		_, ePrev := pageDao.Update(prev)
		if eCurr != nil || ePrev != nil {
			return itineris.NewApiResult(itineris.StatusErrorServer).
				SetMessage(fmt.Sprintf("error updating pages [%s/%s] (error: %s/%s)", prev.GetId(), curr.GetId(), ePrev, eCurr))
		}
	case "move_down":
		if found == len(pageList)-1 {
			// at bottom, can not be moved down
			break
		}
		curr, next := pageList[found], pageList[found+1]
		pCurr := curr.GetPosition()
		curr.SetPosition(pCurr + 1)
		next.SetPosition(pCurr)
		_, eCurr := pageDao.Update(curr)
		_, eNext := pageDao.Update(next)
		if eCurr != nil || eNext != nil {
			return itineris.NewApiResult(itineris.StatusErrorServer).
				SetMessage(fmt.Sprintf("error updating pages [%s/%s] (error: %s/%s)", curr.GetId(), next.GetId(), eCurr, eNext))
		}
	default:
		return itineris.NewApiResult(itineris.StatusErrorClient).
			SetMessage(fmt.Sprintf("invalid action: %s", modifyAction))
	}

	return itineris.NewApiResult(itineris.StatusOk)
}

// apiAdminUpdateTopicPage handles API call "adminUpdateTopicPage"
func apiAdminUpdateTopicPage(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	_, authResult := authenticateAdminApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	tid := _extractParam(params, "tid", reddo.TypeString, "", nil)
	topic, err := topicDao.Get(tid.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if topic == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).SetMessage("topic not found")
	}

	id := _extractParam(params, "id", reddo.TypeString, "", nil)
	page, err := pageDao.Get(id.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).
			SetMessage(fmt.Sprintf("error getting page [%s] (error: %s)", id, err))
	}
	if page == nil || page.GetTopicId() != topic.GetId() {
		return itineris.NewApiResult(itineris.StatusNotFound).
			SetMessage(fmt.Sprintf("page not found [%s]", id))
	}

	icon := _extractParam(params, "icon", reddo.TypeString, "", nil)
	icon = strings.ToLower(icon.(string))
	title := _extractParam(params, "title", reddo.TypeString, "", nil)
	summary := _extractParam(params, "summary", reddo.TypeString, "", nil)
	content := _extractParam(params, "content", reddo.TypeString, "", nil)

	if title == "" {
		return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("title is empty")
	}

	page.SetIcon(icon.(string)).SetTitle(title.(string)).SetSummary(summary.(string)).SetContent(content.(string))
	result, err := pageDao.Update(page)
	if err != nil || !result {
		return itineris.NewApiResult(itineris.StatusErrorServer).
			SetMessage(fmt.Sprintf("cannot update page [%s/%s] (error: %s)", page.GetId(), page.GetTitle(), err))
	}

	return itineris.NewApiResult(itineris.StatusOk)
}

/*----------------------------------------------------------------------*/

var funcUserToMapTransform = func(m map[string]interface{}) map[string]interface{} {
	s := semita.NewSemita(m)

	// transform input map
	result := map[string]interface{}{
		"id":        m[henge.FieldId],
		"t_created": m[henge.FieldTimeCreated],
		"t_updated": m[henge.FieldTimeUpdated],
	}
	result["mid"], _ = s.GetValueOfType(fmt.Sprintf("%s.%s", bo.SerKeyFields, user.UserFieldMaskId), reddo.TypeString)
	result["name"], _ = s.GetValueOfType(fmt.Sprintf("%s.%s", bo.SerKeyAttrs, user.UserAttrDisplayName), reddo.TypeString)
	result["is_admin"], _ = s.GetValueOfType(fmt.Sprintf("%s.%s", bo.SerKeyAttrs, user.UserAttrIsAdmin), reddo.TypeBool)

	// convert "timestamp" to UTC
	if t, ok := result["t_created"].(time.Time); ok {
		result["t_created"] = t.In(time.UTC)
	}
	if t, ok := result["t_updated"].(time.Time); ok {
		result["t_updated"] = t.In(time.UTC)
	}

	return result
}

// apiAdminGetUserList handles API call "adminGetUserList"
func apiAdminGetUserList(ctx *itineris.ApiContext, _ *itineris.ApiAuth, _ *itineris.ApiParams) *itineris.ApiResult {
	curUser, authResult := authenticateAdminApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	userList, err := userDao.GetAll(nil, nil)
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	data := make([]map[string]interface{}, 0)
	for _, user := range userList {
		if curUser.IsAdmin() || curUser.GetId() == user.GetId() {
			// admin can see all user accounts, otherwise can only see myself
			data = append(data, user.ToMap(funcUserToMapTransform))
		}
	}
	return itineris.NewApiResult(itineris.StatusOk).SetData(data)
}

// apiAdminAddUser handles API call "adminAddUser"
func apiAdminAddUser(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	curUser, authResult := authenticateAdminApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	if !curUser.IsAdmin() {
		// only admin can create new user account
		return itineris.NewApiResult(itineris.StatusNoPermission).SetMessage(fmt.Sprintf("creating new user account is denied"))
	}

	id := _extractParam(params, "id", reddo.TypeString, "", nil)
	id = strings.ToLower(id.(string))
	if id == "" {
		return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("user-id is empty")
	}
	newUser, err := userDao.Get(id.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(fmt.Sprintf("error getting user [%s] (error: %s)", id, err))
	}
	if newUser != nil {
		return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage(fmt.Sprintf("user [%s] already existed", id))
	}

	isAdmin := _extractParam(params, "is_admin", reddo.TypeBool, false, nil)
	maskId := strings.ToLower(utils.UniqueId())
	name := _extractParam(params, "name", reddo.TypeString, maskId, nil)

	newPwdEnc := _extractParam(params, "new_pwd", reddo.TypeString, "", nil)
	newPwdRaw, err := RsaDecryptFromBase64(RsaModePKCS1v15, newPwdEnc.(string), rsaPrivKey)
	if err != nil || newPwdRaw == nil {
		return itineris.NewApiResult(itineris.StatusErrorClient).
			SetMessage(fmt.Sprintf("cannot decrypt data (error: %s)", err))
	}
	confirmedPwdEnc := _extractParam(params, "confirmed_pwd", reddo.TypeString, "", nil)
	confirmedPwdRaw, err := RsaDecryptFromBase64(RsaModePKCS1v15, confirmedPwdEnc.(string), rsaPrivKey)
	if err != nil || confirmedPwdRaw == nil {
		return itineris.NewApiResult(itineris.StatusErrorClient).
			SetMessage(fmt.Sprintf("cannot decrypt data (error: %s)", err))
	}
	if strings.TrimSpace(string(newPwdRaw)) != strings.TrimSpace(string(confirmedPwdRaw)) {
		return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("password does not match the confirmed one")
	}

	newUser = user.NewUser(goapi.AppVersionNumber, id.(string), maskId)
	newUser.
		SetAdmin(isAdmin.(bool)).
		SetDisplayName(name.(string)).
		SetPassword(encryptPassword(newUser.GetId(), strings.TrimSpace(string(newPwdRaw))))
	result, err := userDao.Create(newUser)
	if err != nil || !result {
		return itineris.NewApiResult(itineris.StatusErrorServer).
			SetMessage(fmt.Sprintf("cannot create user account %s (error: %s)", id, err))
	}

	return itineris.NewApiResult(itineris.StatusOk)
}

// apiAdminUpdateMyProfile handles API call "adminUpdateMyProfile"
func apiAdminUpdateMyProfile(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	curUser, authResult := authenticateAdminApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	uid := _extractParam(params, "uid", reddo.TypeString, "", nil)
	user, err := userDao.Get(uid.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if user == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).SetMessage("user not found")
	}

	if curUser.GetId() != user.GetId() {
		// can only update profile of the currently logged-in user
		return itineris.NewApiResult(itineris.StatusNoPermission).SetMessage("no permission")
	}

	// for now, only display-name can be changed
	displayName := _extractParam(params, "name", reddo.TypeString, user.GetMaskId(), nil)
	user.SetDisplayName(displayName.(string))

	result, err := userDao.Update(user)
	if err != nil || !result {
		return itineris.NewApiResult(itineris.StatusErrorServer).
			SetMessage(fmt.Sprintf("cannot update user [%s] (error: %s)", user.GetId(), err))
	}

	return itineris.NewApiResult(itineris.StatusOk)
}

// apiAdminUpdateMyPassword handles API call "adminUpdateMyPassword"
func apiAdminUpdateMyPassword(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	curUser, authResult := authenticateAdminApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	uid := _extractParam(params, "uid", reddo.TypeString, "", nil)
	user, err := userDao.Get(uid.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if user == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).SetMessage("user not found")
	}

	if curUser.GetId() != user.GetId() {
		// can only change password of the currently logged-in user
		return itineris.NewApiResult(itineris.StatusNoPermission).SetMessage("no permission")
	}

	currentPwdEnc := _extractParam(params, "current_pwd", reddo.TypeString, "", nil)
	currentPwdRaw, err := RsaDecryptFromBase64(RsaModePKCS1v15, currentPwdEnc.(string), rsaPrivKey)
	if err != nil || currentPwdRaw == nil {
		return itineris.NewApiResult(itineris.StatusErrorClient).
			SetMessage(fmt.Sprintf("cannot decrypt data (error: %s)", err))
	}
	newPwdEnc := _extractParam(params, "new_pwd", reddo.TypeString, "", nil)
	newPwdRaw, err := RsaDecryptFromBase64(RsaModePKCS1v15, newPwdEnc.(string), rsaPrivKey)
	if err != nil || newPwdRaw == nil {
		return itineris.NewApiResult(itineris.StatusErrorClient).
			SetMessage(fmt.Sprintf("cannot decrypt data (error: %s)", err))
	}
	confirmedPwdEnc := _extractParam(params, "confirmed_pwd", reddo.TypeString, "", nil)
	confirmedPwdRaw, err := RsaDecryptFromBase64(RsaModePKCS1v15, confirmedPwdEnc.(string), rsaPrivKey)
	if err != nil || confirmedPwdRaw == nil {
		return itineris.NewApiResult(itineris.StatusErrorClient).
			SetMessage(fmt.Sprintf("cannot decrypt data (error: %s)", err))
	}
	if !verifyPassword(user, string(currentPwdRaw)) {
		return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("current password mismatched")
	}
	if strings.TrimSpace(string(newPwdRaw)) != strings.TrimSpace(string(confirmedPwdRaw)) {
		return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("new password does not match the confirmed one")
	}

	if DEVMODE {
		// prevent changing password of the default user account in DEVMODE
		initAdminUserId := goapi.AppConfig.GetString("gvabe.init.admin_user_id")
		if initAdminUserId == user.GetId() {
			newPwdRaw = []byte(goapi.AppConfig.GetString("gvabe.init.admin_user_pwd"))
		}
	}
	user.SetPassword(encryptPassword(user.GetId(), string(newPwdRaw)))
	result, err := userDao.Update(user)
	if err != nil || !result {
		return itineris.NewApiResult(itineris.StatusErrorServer).
			SetMessage(fmt.Sprintf("cannot update user [%s] (error: %s)", user.GetId(), err))
	}

	return itineris.NewApiResult(itineris.StatusOk)
}

// apiAdminUpdateUserProfile handles API call "adminUpdateUserProfile"
func apiAdminUpdateUserProfile(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	curUser, authResult := authenticateAdminApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	uid := _extractParam(params, "uid", reddo.TypeString, "", nil)
	user, err := userDao.Get(uid.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if user == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).SetMessage("user not found")
	}

	if curUser.GetId() == user.GetId() {
		// can not update your own profile via this API
		return itineris.NewApiResult(itineris.StatusNoPermission).SetMessage("no permission")
	}

	isAdmin := _extractParam(params, "is_admin", reddo.TypeBool, false, nil)
	displayName := _extractParam(params, "name", reddo.TypeString, user.GetMaskId(), nil)
	user.SetDisplayName(displayName.(string)).SetAdmin(isAdmin.(bool))

	result, err := userDao.Update(user)
	if err != nil || !result {
		return itineris.NewApiResult(itineris.StatusErrorServer).
			SetMessage(fmt.Sprintf("cannot update user [%s] (error: %s)", user.GetId(), err))
	}

	return itineris.NewApiResult(itineris.StatusOk)
}

// apiAdminDeleteUserProfile handles API call "adminDeleteUserProfile"
func apiAdminDeleteUserProfile(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	curUser, authResult := authenticateAdminApiCall(ctx)
	if authResult != nil {
		return authResult
	}

	uid := _extractParam(params, "uid", reddo.TypeString, "", nil)
	user, err := userDao.Get(uid.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if user == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).SetMessage("user not found")
	}

	if curUser.GetId() == user.GetId() || !curUser.IsAdmin() {
		// can not delete your own account, or you are not admin
		return itineris.NewApiResult(itineris.StatusNoPermission).SetMessage("no permission")
	}

	result, err := userDao.Delete(user)
	if err != nil || !result {
		return itineris.NewApiResult(itineris.StatusErrorServer).
			SetMessage(fmt.Sprintf("cannot delete user [%s] (error: %s)", user.GetId(), err))
	}

	return itineris.NewApiResult(itineris.StatusOk)
}
