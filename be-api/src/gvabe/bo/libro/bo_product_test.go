package libro

import (
	"encoding/json"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/btnguyen2k/henge"
	"main/src/gvabe/bo"
)

func TestNewProduct(t *testing.T) {
	testName := "TestNewProduct"
	teardownTest := setupTest(t, testName, func(t *testing.T, testName string) {
		bo.UboTimestampRouding = henge.TimestampRoundSettingNone
	}, nil)
	defer teardownTest(t)

	tstart := time.Now()
	_tagVersion := uint64(1337)
	_id := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_numTopics := rand.Intn(1024)
	prod := NewProduct(_tagVersion, _id, _name, _desc, _isPublished)
	if prod == nil {
		t.Fatalf("%s failed: nil", testName)
	}
	prod.SetNumTopics(_numTopics)
	if tagVersion := prod.GetTagVersion(); tagVersion != _tagVersion {
		t.Fatalf("%s failed: expected tag-version to be %#v but received %#v", testName, _tagVersion, tagVersion)
	}
	if id := prod.GetId(); id != _id {
		t.Fatalf("%s failed: expected bo's id to be %#v but received %#v", testName, _id, id)
	}
	if prodName := prod.GetName(); prodName != _name {
		t.Fatalf("%s failed: expected bo's testName to be %#v but received %#v", testName, _name, prodName)
	}
	if desc := prod.GetDescription(); desc != _desc {
		t.Fatalf("%s failed: expected bo's desc to be %#v but received %#v", testName, _desc, desc)
	}
	if isPublished := prod.IsPublished(); isPublished != _isPublished {
		t.Fatalf("%s failed: expected bo's is-published to be %#v but received %#v", testName, _isPublished, isPublished)
	}
	if numTopics := prod.GetNumTopics(); numTopics != _numTopics {
		t.Fatalf("%s failed: expected num-topics to be %#v but received %#v", testName, _numTopics, numTopics)
	}

	if contacts, expected := prod.GetContacts(), map[string]string{}; !reflect.DeepEqual(contacts, expected) {
		t.Fatalf("%s failed: expected contacts to be %#v but received %#v", testName, expected, contacts)
	}
	prod.AddContact("github", "btnguyen2k/libro")
	if contacts, expected := prod.GetContacts(), map[string]string{"github": "btnguyen2k/libro"}; !reflect.DeepEqual(contacts, expected) {
		t.Fatalf("%s failed: expected contacts to be %#v but received %#v", testName, expected, contacts)
	}
	prod.AddContact("website", "https://github.com/btnguyen2k/libro")
	if contacts, expected := prod.GetContacts(), map[string]string{"github": "btnguyen2k/libro", "website": "https://github.com/btnguyen2k/libro"}; !reflect.DeepEqual(contacts, expected) {
		t.Fatalf("%s failed: expected contacts to be %#v but received %#v", testName, expected, contacts)
	}
	prod.SetContacts(map[string]string{"github": "btnguyen2k/libro", "fb": "fb/libro"})
	if contacts, expected := prod.GetContacts(), map[string]string{"github": "btnguyen2k/libro", "fb": "fb/libro"}; !reflect.DeepEqual(contacts, expected) {
		t.Fatalf("%s failed: expected contacts to be %#v but received %#v", testName, expected, contacts)
	}
	prod.SetContacts(nil)
	if contacts, expected := prod.GetContacts(), map[string]string{}; !reflect.DeepEqual(contacts, expected) {
		t.Fatalf("%s failed: expected contacts to be %#v but received %#v", testName, expected, contacts)
	}

	tend := time.Now()
	if prod.GetTimeCreated().Before(tstart) || prod.GetTimeCreated().After(tend) {
		t.Fatalf("%s failed: timestamp-created is invalid\nStart: %s / Created: %s / End: %s", testName, tstart, prod.GetTimeCreated(), tend)
	}
	if prod.GetTimeUpdated().Before(tstart) || prod.GetTimeUpdated().After(tend) || prod.GetTimeUpdated().Before(prod.GetTimeCreated()) {
		t.Fatalf("%s failed: timestamp-updated is invalid\nStart: %s / Updated: %s / End: %s", testName, tstart, prod.GetTimeUpdated(), tend)
	}
}

func TestNewProductFromUbo(t *testing.T) {
	testName := "TestNewProductFromUbo"
	teardownTest := setupTest(t, testName, func(t *testing.T, testName string) {
		bo.UboTimestampRouding = henge.TimestampRoundSettingNone
	}, nil)
	defer teardownTest(t)

	if NewProductFromUbo(nil) != nil {
		t.Fatalf("%s failed: NewProductFromUbo(nil) should return nil", testName)
	}

	tstart := time.Now()
	_tagVersion := uint64(1337)
	_id := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_numTopics := rand.Intn(1024)
	ubo := henge.NewUniversalBo(_id, _tagVersion, henge.UboOpt{TimeLayout: bo.UboTimeLayout, TimestampRounding: bo.UboTimestampRouding})
	ubo.SetDataAttr(ProdAttrName, _name)
	ubo.SetDataAttr(ProdAttrDesc, _desc)
	ubo.SetDataAttr(ProdAttrIsPublished, _isPublished)
	ubo.SetDataAttr(ProdAttrNumTopics, _numTopics)

	prod := NewProductFromUbo(ubo)
	if prod == nil {
		t.Fatalf("%s failed: nil", testName)
	}
	if tagVersion := prod.GetTagVersion(); tagVersion != _tagVersion {
		t.Fatalf("%s failed: expected tag-version to be %#v but received %#v", testName, _tagVersion, tagVersion)
	}
	if id := prod.GetId(); id != _id {
		t.Fatalf("%s failed: expected bo's id to be %#v but received %#v", testName, _id, id)
	}
	if prodName := prod.GetName(); prodName != _name {
		t.Fatalf("%s failed: expected bo's testName to be %#v but received %#v", testName, _name, prodName)
	}
	if desc := prod.GetDescription(); desc != _desc {
		t.Fatalf("%s failed: expected bo's desc to be %#v but received %#v", testName, _desc, desc)
	}
	if isPublished := prod.IsPublished(); isPublished != _isPublished {
		t.Fatalf("%s failed: expected bo's is-published to be %#v but received %#v", testName, _isPublished, isPublished)
	}
	if numTopics := prod.GetNumTopics(); numTopics != _numTopics {
		t.Fatalf("%s failed: expected num-topics to be %#v but received %#v", testName, _numTopics, numTopics)
	}
	if contacts, expected := prod.GetContacts(), map[string]string{}; !reflect.DeepEqual(contacts, expected) {
		t.Fatalf("%s failed: expected contacts to be %#v but received %#v", testName, expected, contacts)
	}

	ubo.SetDataAttr(ProdAttrContacts, map[string]string{"github": "btnguyen2k/libro", "fb": "fb/libro"})
	prod = NewProductFromUbo(ubo)
	if contacts, expected := prod.GetContacts(), map[string]string{"github": "btnguyen2k/libro", "fb": "fb/libro"}; !reflect.DeepEqual(contacts, expected) {
		t.Fatalf("%s failed: expected contacts to be %#v but received %#v", testName, expected, contacts)
	}
	prod.SetNumTopics(_numTopics) // force change 'timestamp-updated'

	tend := time.Now()
	if prod.GetTimeCreated().Before(tstart) || prod.GetTimeCreated().After(tend) {
		t.Fatalf("%s failed: timestamp-created is invalid\nStart: %s / Created: %s / End: %s", testName, tstart, prod.GetTimeCreated(), tend)
	}
	if prod.GetTimeUpdated().Before(tstart) || prod.GetTimeUpdated().After(tend) || prod.GetTimeUpdated().Before(prod.GetTimeCreated()) {
		t.Fatalf("%s failed: timestamp-updated is invalid\nStart: %s / Updated: %s / End: %s", testName, tstart, prod.GetTimeUpdated(), tend)
	}
}

// no contacts info
func TestProduct_ToMap(t *testing.T) {
	testName := "TestProduct_ToMap"
	teardownTest := setupTest(t, testName, nil, nil)
	defer teardownTest(t)

	_tagVersion := uint64(1337)
	_id := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_numTopics := rand.Intn(1024)
	prod := NewProduct(_tagVersion, _id, _name, _desc, _isPublished)
	if prod == nil {
		t.Fatalf("%s failed: nil", testName)
	}
	prod.SetNumTopics(_numTopics)

	m := prod.ToMap(nil)
	expected := map[string]interface{}{
		henge.FieldId:          _id,
		henge.FieldTimeCreated: prod.GetTimeCreated(),
		henge.FieldTimeUpdated: prod.GetTimeUpdated(),
		bo.SerKeyAttrs: map[string]interface{}{
			ProdAttrName:        _name,
			ProdAttrDesc:        _desc,
			ProdAttrIsPublished: _isPublished,
			ProdAttrNumTopics:   _numTopics,
			ProdAttrContacts:    map[string]string{},
		},
	}
	if !reflect.DeepEqual(m, expected) {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, expected, m)
	}

	m = prod.ToMap(func(input map[string]interface{}) map[string]interface{} {
		return map[string]interface{}{
			"FieldId":          input[henge.FieldId],
			"FieldTimeCreated": input[henge.FieldTimeCreated],
			"FieldTimeUpdated": input[henge.FieldTimeUpdated],
			"SerKeyAttrs":      input[bo.SerKeyAttrs],
		}
	})
	expected = map[string]interface{}{
		"FieldId":          _id,
		"FieldTimeCreated": prod.GetTimeCreated(),
		"FieldTimeUpdated": prod.GetTimeUpdated(),
		"SerKeyAttrs": map[string]interface{}{
			ProdAttrName:        _name,
			ProdAttrDesc:        _desc,
			ProdAttrIsPublished: _isPublished,
			ProdAttrNumTopics:   _numTopics,
			ProdAttrContacts:    map[string]string{},
		},
	}
	if !reflect.DeepEqual(m, expected) {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, expected, m)
	}
}

// some contacts info
func TestProduct_ToMap2(t *testing.T) {
	testName := "TestProduct_ToMap2"
	teardownTest := setupTest(t, testName, nil, nil)
	defer teardownTest(t)

	_tagVersion := uint64(1337)
	_id := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_numTopics := rand.Intn(1024)
	prod := NewProduct(_tagVersion, _id, _name, _desc, _isPublished)
	if prod == nil {
		t.Fatalf("%s failed: nil", testName)
	}
	prod.SetNumTopics(_numTopics)
	prod.AddContact("github", "btnguyen2k/libro")

	m := prod.ToMap(nil)
	expected := map[string]interface{}{
		henge.FieldId:          _id,
		henge.FieldTimeCreated: prod.GetTimeCreated(),
		henge.FieldTimeUpdated: prod.GetTimeUpdated(),
		bo.SerKeyAttrs: map[string]interface{}{
			ProdAttrName:        _name,
			ProdAttrDesc:        _desc,
			ProdAttrIsPublished: _isPublished,
			ProdAttrNumTopics:   _numTopics,
			ProdAttrContacts:    map[string]string{"github": "btnguyen2k/libro"},
		},
	}
	if !reflect.DeepEqual(m, expected) {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, expected, m)
	}

	m = prod.ToMap(func(input map[string]interface{}) map[string]interface{} {
		return map[string]interface{}{
			"FieldId":          input[henge.FieldId],
			"FieldTimeCreated": input[henge.FieldTimeCreated],
			"FieldTimeUpdated": input[henge.FieldTimeUpdated],
			"SerKeyAttrs":      input[bo.SerKeyAttrs],
		}
	})
	expected = map[string]interface{}{
		"FieldId":          _id,
		"FieldTimeCreated": prod.GetTimeCreated(),
		"FieldTimeUpdated": prod.GetTimeUpdated(),
		"SerKeyAttrs": map[string]interface{}{
			ProdAttrName:        _name,
			ProdAttrDesc:        _desc,
			ProdAttrIsPublished: _isPublished,
			ProdAttrNumTopics:   _numTopics,
			ProdAttrContacts:    map[string]string{"github": "btnguyen2k/libro"},
		},
	}
	if !reflect.DeepEqual(m, expected) {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, expected, m)
	}
}

// no contacts info
func TestProduct_json(t *testing.T) {
	testName := "TestProduct_json"
	teardownTest := setupTest(t, testName, nil, nil)
	defer teardownTest(t)

	_tagVersion := uint64(1337)
	_id := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_numTopics := rand.Intn(1024)
	prod1 := NewProduct(_tagVersion, _id, _name, _desc, _isPublished)
	if prod1 == nil {
		t.Fatalf("%s failed: nil", testName)
	}
	prod1.SetNumTopics(_numTopics)
	prod1.AddContact("github", "btnguyen2k/libro")
	js1, _ := json.Marshal(prod1)

	var prod2 *Product
	err := json.Unmarshal(js1, &prod2)
	if err != nil {
		t.Fatalf("%s failed: %e", testName, err)
	}
	if prod1.GetTagVersion() != prod2.GetTagVersion() {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, prod1.GetTagVersion(), prod2.GetTagVersion())
	}
	if prod1.GetId() != prod2.GetId() {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, prod1.GetId(), prod2.GetId())
	}
	if prod1.GetName() != prod2.GetName() {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, prod1.GetName(), prod2.GetName())
	}
	if prod1.GetDescription() != prod2.GetDescription() {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, prod1.GetDescription(), prod2.GetDescription())
	}
	if prod1.IsPublished() != prod2.IsPublished() {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, prod1.IsPublished(), prod2.IsPublished())
	}
	if prod1.GetNumTopics() != prod2.GetNumTopics() {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, prod1.GetNumTopics(), prod2.GetNumTopics())
	}
	if !reflect.DeepEqual(prod1.GetContacts(), prod2.GetContacts()) {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, prod1.GetContacts(), prod2.GetContacts())
	}
	if prod1.GetChecksum() != prod2.GetChecksum() {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, prod1.GetChecksum(), prod2.GetChecksum())
	}

}

// some contacts info
func TestProduct_json2(t *testing.T) {
	testName := "TestProduct_json2"
	teardownTest := setupTest(t, testName, nil, nil)
	defer teardownTest(t)

	_tagVersion := uint64(1337)
	_id := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_numTopics := rand.Intn(1024)
	prod1 := NewProduct(_tagVersion, _id, _name, _desc, _isPublished)
	if prod1 == nil {
		t.Fatalf("%s failed: nil", testName)
	}
	prod1.SetNumTopics(_numTopics)
	prod1.SetContacts(map[string]string{"github": "btnguyen2k/libro", "fb": "Libro"})
	js1, _ := json.Marshal(prod1)

	var prod2 *Product
	err := json.Unmarshal(js1, &prod2)
	if err != nil {
		t.Fatalf("%s failed: %e", testName, err)
	}
	if prod1.GetTagVersion() != prod2.GetTagVersion() {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, prod1.GetTagVersion(), prod2.GetTagVersion())
	}
	if prod1.GetId() != prod2.GetId() {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, prod1.GetId(), prod2.GetId())
	}
	if prod1.GetName() != prod2.GetName() {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, prod1.GetName(), prod2.GetName())
	}
	if prod1.GetDescription() != prod2.GetDescription() {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, prod1.GetDescription(), prod2.GetDescription())
	}
	if prod1.IsPublished() != prod2.IsPublished() {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, prod1.IsPublished(), prod2.IsPublished())
	}
	if prod1.GetNumTopics() != prod2.GetNumTopics() {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, prod1.GetNumTopics(), prod2.GetNumTopics())
	}
	if !reflect.DeepEqual(prod1.GetContacts(), prod2.GetContacts()) {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, prod1.GetContacts(), prod2.GetContacts())
	}
	if prod1.GetChecksum() != prod2.GetChecksum() {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, prod1.GetChecksum(), prod2.GetChecksum())
	}
}
