package app

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/btnguyen2k/henge"
)

func TestNewApp(t *testing.T) {
	name := "TestNewApp"
	_tagVersion := uint64(1337)
	_id := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isVisible := true
	app := NewApp(_tagVersion, _id, _name, _desc, _isVisible)
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
	if isVisible := app.IsVisible(); isVisible != _isVisible {
		t.Fatalf("%s failed: expected bo's is-visible to be %#v but received %#v", name, _isVisible, isVisible)
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
	_isVisible := true
	ubo := henge.NewUniversalBo(_id, _tagVersion)
	ubo.SetDataAttr(AppAttrName, _name)
	ubo.SetDataAttr(AppAttrDesc, _desc)
	ubo.SetDataAttr(AppAttrIsVisible, _isVisible)

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
	if isVisible := app.IsVisible(); isVisible != _isVisible {
		t.Fatalf("%s failed: expected bo's is-visible to be %#v but received %#v", name, _isVisible, isVisible)
	}
}

func TestApp_ToMap(t *testing.T) {
	name := "TestApp_ToMap"
	_tagVersion := uint64(1337)
	_id := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isVisible := true
	app := NewApp(_tagVersion, _id, _name, _desc, _isVisible)
	if app == nil {
		t.Fatalf("%s failed: nil", name)
	}

	m := app.ToMap(nil)
	expected := map[string]interface{}{
		henge.FieldId: _id,
		SerKeyAttrs: map[string]interface{}{
			AppAttrName:      _name,
			AppAttrDesc:      _desc,
			AppAttrIsVisible: _isVisible,
		},
	}
	if !reflect.DeepEqual(m, expected) {
		t.Fatalf("%s failed: expected %#v but received %#v", name, expected, m)
	}

	m = app.ToMap(func(input map[string]interface{}) map[string]interface{} {
		return map[string]interface{}{
			"FieldId":     input[henge.FieldId],
			"SerKeyAttrs": input[SerKeyAttrs],
		}
	})
	expected = map[string]interface{}{
		"FieldId": _id,
		"SerKeyAttrs": map[string]interface{}{
			AppAttrName:      _name,
			AppAttrDesc:      _desc,
			AppAttrIsVisible: _isVisible,
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
	_isVisible := true
	app1 := NewApp(_tagVersion, _id, _name, _desc, _isVisible)
	if app1 == nil {
		t.Fatalf("%s failed: nil", name)
	}
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
	if app1.IsVisible() != app2.IsVisible() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, app1.IsVisible(), app2.IsVisible())
	}
	if app1.GetChecksum() != app2.GetChecksum() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, app1.GetChecksum(), app2.GetChecksum())
	}
}
