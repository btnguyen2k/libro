package app

import (
	"encoding/json"
	"math/rand"
	"reflect"
	"testing"

	"github.com/btnguyen2k/henge"
	"main/src/gvabe/bo"
)

func TestNewApp(t *testing.T) {
	name := "TestNewApp"
	_tagVersion := uint64(1337)
	_id := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_numTopics := rand.Intn(1024)
	app := NewApp(_tagVersion, _id, _name, _desc, _isPublished)
	if app == nil {
		t.Fatalf("%s failed: nil", name)
	}
	app.SetNumTopics(_numTopics)
	if tagVersion := app.GetTagVersion(); tagVersion != _tagVersion {
		t.Fatalf("%s failed: expected tag-version to be %#v but received %#v", name, _tagVersion, tagVersion)
	}
	if id := app.GetId(); id != _id {
		t.Fatalf("%s failed: expected bo's id to be %#v but received %#v", name, _id, id)
	}
	if appName := app.GetName(); appName != _name {
		t.Fatalf("%s failed: expected bo's name to be %#v but received %#v", name, _name, appName)
	}
	if desc := app.GetDescription(); desc != _desc {
		t.Fatalf("%s failed: expected bo's desc to be %#v but received %#v", name, _desc, desc)
	}
	if isPublished := app.IsPublished(); isPublished != _isPublished {
		t.Fatalf("%s failed: expected bo's is-published to be %#v but received %#v", name, _isPublished, isPublished)
	}
	if numTopics := app.GetNumTopics(); numTopics != _numTopics {
		t.Fatalf("%s failed: expected num-topics to be %#v but received %#v", name, _numTopics, numTopics)
	}
}

func TestNewAppFromUbo(t *testing.T) {
	name := "TestNewAppFromUbo"

	if NewAppFromUbo(nil) != nil {
		t.Fatalf("%s failed: NewAppFromUbo(nil) should return nil", name)
	}
	_tagVersion := uint64(1337)
	_id := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_numTopics := rand.Intn(1024)
	ubo := henge.NewUniversalBo(_id, _tagVersion)
	ubo.SetDataAttr(AppAttrName, _name)
	ubo.SetDataAttr(AppAttrDesc, _desc)
	ubo.SetDataAttr(AppAttrIsPublished, _isPublished)
	ubo.SetDataAttr(AppAttrNumTopics, _numTopics)

	app := NewAppFromUbo(ubo)
	if app == nil {
		t.Fatalf("%s failed: nil", name)
	}
	if tagVersion := app.GetTagVersion(); tagVersion != _tagVersion {
		t.Fatalf("%s failed: expected tag-version to be %#v but received %#v", name, _tagVersion, tagVersion)
	}
	if id := app.GetId(); id != _id {
		t.Fatalf("%s failed: expected bo's id to be %#v but received %#v", name, _id, id)
	}
	if appName := app.GetName(); appName != _name {
		t.Fatalf("%s failed: expected bo's name to be %#v but received %#v", name, _name, appName)
	}
	if desc := app.GetDescription(); desc != _desc {
		t.Fatalf("%s failed: expected bo's desc to be %#v but received %#v", name, _desc, desc)
	}
	if isPublished := app.IsPublished(); isPublished != _isPublished {
		t.Fatalf("%s failed: expected bo's is-published to be %#v but received %#v", name, _isPublished, isPublished)
	}
	if numTopics := app.GetNumTopics(); numTopics != _numTopics {
		t.Fatalf("%s failed: expected num-topics to be %#v but received %#v", name, _numTopics, numTopics)
	}
}

func TestApp_ToMap(t *testing.T) {
	name := "TestApp_ToMap"
	_tagVersion := uint64(1337)
	_id := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_numTopics := rand.Intn(1024)
	app := NewApp(_tagVersion, _id, _name, _desc, _isPublished)
	if app == nil {
		t.Fatalf("%s failed: nil", name)
	}
	app.SetNumTopics(_numTopics)

	m := app.ToMap(nil)
	expected := map[string]interface{}{
		henge.FieldId: _id,
		bo.SerKeyAttrs: map[string]interface{}{
			AppAttrName:        _name,
			AppAttrDesc:        _desc,
			AppAttrIsPublished: _isPublished,
			AppAttrNumTopics:   _numTopics,
		},
	}
	if !reflect.DeepEqual(m, expected) {
		t.Fatalf("%s failed: expected %#v but received %#v", name, expected, m)
	}

	m = app.ToMap(func(input map[string]interface{}) map[string]interface{} {
		return map[string]interface{}{
			"FieldId":     input[henge.FieldId],
			"SerKeyAttrs": input[bo.SerKeyAttrs],
		}
	})
	expected = map[string]interface{}{
		"FieldId": _id,
		"SerKeyAttrs": map[string]interface{}{
			AppAttrName:        _name,
			AppAttrDesc:        _desc,
			AppAttrIsPublished: _isPublished,
			AppAttrNumTopics:   _numTopics,
		},
	}
	if !reflect.DeepEqual(m, expected) {
		t.Fatalf("%s failed: expected %#v but received %#v", name, expected, m)
	}
}

func TestApp_json(t *testing.T) {
	name := "TestApp_json"
	_tagVersion := uint64(1337)
	_id := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_numTopics := rand.Intn(1024)
	app1 := NewApp(_tagVersion, _id, _name, _desc, _isPublished)
	if app1 == nil {
		t.Fatalf("%s failed: nil", name)
	}
	app1.SetNumTopics(_numTopics)
	js1, _ := json.Marshal(app1)

	var app2 *App
	err := json.Unmarshal(js1, &app2)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	if app1.GetTagVersion() != app2.GetTagVersion() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, app1.GetTagVersion(), app2.GetTagVersion())
	}
	if app1.GetId() != app2.GetId() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, app1.GetId(), app2.GetId())
	}
	if app1.GetName() != app2.GetName() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, app1.GetName(), app2.GetName())
	}
	if app1.GetDescription() != app2.GetDescription() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, app1.GetDescription(), app2.GetDescription())
	}
	if app1.IsPublished() != app2.IsPublished() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, app1.IsPublished(), app2.IsPublished())
	}
	if app1.GetNumTopics() != app2.GetNumTopics() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, app1.GetNumTopics(), app2.GetNumTopics())
	}
	if app1.GetChecksum() != app2.GetChecksum() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, app1.GetChecksum(), app2.GetChecksum())
	}
}
