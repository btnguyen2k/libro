package product

import (
	"encoding/json"
	"math/rand"
	"reflect"
	"testing"

	"github.com/btnguyen2k/henge"
	"main/src/gvabe/bo"
)

func TestNewProduct(t *testing.T) {
	name := "TestNewProduct"
	_tagVersion := uint64(1337)
	_id := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_numTopics := rand.Intn(1024)
	prod := NewProduct(_tagVersion, _id, _name, _desc, _isPublished)
	if prod == nil {
		t.Fatalf("%s failed: nil", name)
	}
	prod.SetNumTopics(_numTopics)
	if tagVersion := prod.GetTagVersion(); tagVersion != _tagVersion {
		t.Fatalf("%s failed: expected tag-version to be %#v but received %#v", name, _tagVersion, tagVersion)
	}
	if id := prod.GetId(); id != _id {
		t.Fatalf("%s failed: expected bo's id to be %#v but received %#v", name, _id, id)
	}
	if prodName := prod.GetName(); prodName != _name {
		t.Fatalf("%s failed: expected bo's name to be %#v but received %#v", name, _name, prodName)
	}
	if desc := prod.GetDescription(); desc != _desc {
		t.Fatalf("%s failed: expected bo's desc to be %#v but received %#v", name, _desc, desc)
	}
	if isPublished := prod.IsPublished(); isPublished != _isPublished {
		t.Fatalf("%s failed: expected bo's is-published to be %#v but received %#v", name, _isPublished, isPublished)
	}
	if numTopics := prod.GetNumTopics(); numTopics != _numTopics {
		t.Fatalf("%s failed: expected num-topics to be %#v but received %#v", name, _numTopics, numTopics)
	}
}

func TestNewProductFromUbo(t *testing.T) {
	name := "TestNewProductFromUbo"

	if NewProductFromUbo(nil) != nil {
		t.Fatalf("%s failed: NewProductFromUbo(nil) should return nil", name)
	}
	_tagVersion := uint64(1337)
	_id := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_numTopics := rand.Intn(1024)
	ubo := henge.NewUniversalBo(_id, _tagVersion)
	ubo.SetDataAttr(ProdAttrName, _name)
	ubo.SetDataAttr(ProdAttrDesc, _desc)
	ubo.SetDataAttr(ProdAttrIsPublished, _isPublished)
	ubo.SetDataAttr(ProdAttrNumTopics, _numTopics)

	prod := NewProductFromUbo(ubo)
	if prod == nil {
		t.Fatalf("%s failed: nil", name)
	}
	if tagVersion := prod.GetTagVersion(); tagVersion != _tagVersion {
		t.Fatalf("%s failed: expected tag-version to be %#v but received %#v", name, _tagVersion, tagVersion)
	}
	if id := prod.GetId(); id != _id {
		t.Fatalf("%s failed: expected bo's id to be %#v but received %#v", name, _id, id)
	}
	if prodName := prod.GetName(); prodName != _name {
		t.Fatalf("%s failed: expected bo's name to be %#v but received %#v", name, _name, prodName)
	}
	if desc := prod.GetDescription(); desc != _desc {
		t.Fatalf("%s failed: expected bo's desc to be %#v but received %#v", name, _desc, desc)
	}
	if isPublished := prod.IsPublished(); isPublished != _isPublished {
		t.Fatalf("%s failed: expected bo's is-published to be %#v but received %#v", name, _isPublished, isPublished)
	}
	if numTopics := prod.GetNumTopics(); numTopics != _numTopics {
		t.Fatalf("%s failed: expected num-topics to be %#v but received %#v", name, _numTopics, numTopics)
	}
}

func TestProduct_ToMap(t *testing.T) {
	name := "TestProduct_ToMap"
	_tagVersion := uint64(1337)
	_id := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_numTopics := rand.Intn(1024)
	prod := NewProduct(_tagVersion, _id, _name, _desc, _isPublished)
	if prod == nil {
		t.Fatalf("%s failed: nil", name)
	}
	prod.SetNumTopics(_numTopics)

	m := prod.ToMap(nil)
	expected := map[string]interface{}{
		henge.FieldId: _id,
		bo.SerKeyAttrs: map[string]interface{}{
			ProdAttrName:        _name,
			ProdAttrDesc:        _desc,
			ProdAttrIsPublished: _isPublished,
			ProdAttrNumTopics:   _numTopics,
		},
	}
	if !reflect.DeepEqual(m, expected) {
		t.Fatalf("%s failed: expected %#v but received %#v", name, expected, m)
	}

	m = prod.ToMap(func(input map[string]interface{}) map[string]interface{} {
		return map[string]interface{}{
			"FieldId":     input[henge.FieldId],
			"SerKeyAttrs": input[bo.SerKeyAttrs],
		}
	})
	expected = map[string]interface{}{
		"FieldId": _id,
		"SerKeyAttrs": map[string]interface{}{
			ProdAttrName:        _name,
			ProdAttrDesc:        _desc,
			ProdAttrIsPublished: _isPublished,
			ProdAttrNumTopics:   _numTopics,
		},
	}
	if !reflect.DeepEqual(m, expected) {
		t.Fatalf("%s failed: expected %#v but received %#v", name, expected, m)
	}
}

func TestProduct_json(t *testing.T) {
	name := "TestProduct_json"
	_tagVersion := uint64(1337)
	_id := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_numTopics := rand.Intn(1024)
	prod1 := NewProduct(_tagVersion, _id, _name, _desc, _isPublished)
	if prod1 == nil {
		t.Fatalf("%s failed: nil", name)
	}
	prod1.SetNumTopics(_numTopics)
	js1, _ := json.Marshal(prod1)

	var prod2 *Product
	err := json.Unmarshal(js1, &prod2)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	if prod1.GetTagVersion() != prod2.GetTagVersion() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, prod1.GetTagVersion(), prod2.GetTagVersion())
	}
	if prod1.GetId() != prod2.GetId() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, prod1.GetId(), prod2.GetId())
	}
	if prod1.GetName() != prod2.GetName() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, prod1.GetName(), prod2.GetName())
	}
	if prod1.GetDescription() != prod2.GetDescription() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, prod1.GetDescription(), prod2.GetDescription())
	}
	if prod1.IsPublished() != prod2.IsPublished() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, prod1.IsPublished(), prod2.IsPublished())
	}
	if prod1.GetNumTopics() != prod2.GetNumTopics() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, prod1.GetNumTopics(), prod2.GetNumTopics())
	}
	if prod1.GetChecksum() != prod2.GetChecksum() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, prod1.GetChecksum(), prod2.GetChecksum())
	}
}
