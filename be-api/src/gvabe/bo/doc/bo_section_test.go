package doc

import (
	"encoding/json"
	"math/rand"
	"reflect"
	"testing"

	"github.com/btnguyen2k/henge"
	"main/src/gvabe/bo/app"
	"main/src/utils"
)

func TestNewSection(t *testing.T) {
	name := "TestNewSection"
	_tagVersion := uint64(1337)
	_appId := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isVisible := true
	_app := app.NewApp(_tagVersion, _appId, _name, _desc, _isVisible)

	_title := "Quick start"
	_icon := "default"
	_summary := "section one"
	_pos := rand.Int()
	section := NewSection(_tagVersion, _app, _title, _icon, _summary)
	if section == nil {
		t.Fatalf("%s failed: nil", name)
	}
	section.SetPosition(_pos)

	_id := section.GetId()
	if tagVersion := section.GetTagVersion(); tagVersion != _tagVersion {
		t.Fatalf("%s failed: expected tag-version to be %#v but received %#v", name, _tagVersion, tagVersion)
	}
	if id := section.GetId(); id != _id {
		t.Fatalf("%s failed: expected id to be %#v but received %#v", name, _id, id)
	}
	if appId := section.GetAppId(); appId != _app.GetId() {
		t.Fatalf("%s failed: expected app-id to be %#v but received %#v", name, _app.GetId(), appId)
	}
	if title := section.GetTitle(); title != _title {
		t.Fatalf("%s failed: expected title to be %#v but received %#v", name, _title, title)
	}
	if icon := section.GetIcon(); icon != _icon {
		t.Fatalf("%s failed: expected icon to be %#v but received %#v", name, _icon, icon)
	}
	if summary := section.GetSummary(); summary != _summary {
		t.Fatalf("%s failed: expected summary to be %#v but received %#v", name, _summary, summary)
	}
	if pos := section.GetPosition(); pos != _pos {
		t.Fatalf("%s failed: expected position to be %#v but received %#v", name, _pos, pos)
	}
}

func TestNewSectionFromUbo(t *testing.T) {
	name := "TestNewSectionFromUbo"

	if NewSectionFromUbo(nil) != nil {
		t.Fatalf("%s failed: NewSectionFromUbo(nil) should return nil", name)
	}
	_tagVersion := uint64(1337)
	_id := utils.UniqueId()
	_appId := "libro"
	_title := "Quick start"
	_icon := "default"
	_summary := "section one"
	_pos := rand.Intn(10242048)
	ubo := henge.NewUniversalBo(_id, _tagVersion)
	ubo.SetExtraAttr(SectionFieldAppId, _appId)
	ubo.SetDataAttr(SectionAttrTitle, _title)
	ubo.SetDataAttr(SectionAttrIcon, _icon)
	ubo.SetDataAttr(SectionAttrSummary, _summary)
	ubo.SetDataAttr(SectionAttrPosition, _pos)

	section := NewSectionFromUbo(ubo)
	if section == nil {
		t.Fatalf("%s failed: nil", name)
	}
	if tagVersion := section.GetTagVersion(); tagVersion != _tagVersion {
		t.Fatalf("%s failed: expected tag-version to be %#v but received %#v", name, _tagVersion, tagVersion)
	}
	if id := section.GetId(); id != _id {
		t.Fatalf("%s failed: expected id to be %#v but received %#v", name, _id, id)
	}
	if appId := section.GetAppId(); appId != _appId {
		t.Fatalf("%s failed: expected app-id to be %#v but received %#v", name, _appId, appId)
	}
	if title := section.GetTitle(); title != _title {
		t.Fatalf("%s failed: expected title to be %#v but received %#v", name, _title, title)
	}
	if icon := section.GetIcon(); icon != _icon {
		t.Fatalf("%s failed: expected icon to be %#v but received %#v", name, _icon, icon)
	}
	if summary := section.GetSummary(); summary != _summary {
		t.Fatalf("%s failed: expected summary to be %#v but received %#v", name, _summary, summary)
	}
	if pos := section.GetPosition(); pos != _pos {
		t.Fatalf("%s failed: expected position to be %#v but received %#v", name, _pos, pos)
	}
}

func TestSection_ToMap(t *testing.T) {
	name := "TestSection_ToMap"
	_tagVersion := uint64(1337)
	_appId := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isVisible := true
	_app := app.NewApp(_tagVersion, _appId, _name, _desc, _isVisible)

	_title := "Quick start"
	_icon := "default"
	_summary := "section one"
	_pos := rand.Intn(10242048)
	section := NewSection(_tagVersion, _app, _title, _icon, _summary)
	if section == nil {
		t.Fatalf("%s failed: nil", name)
	}
	section.SetPosition(_pos)
	_id := section.GetId()

	m := section.ToMap(nil)
	expected := map[string]interface{}{
		henge.FieldId: _id,
		SerKeyFields: map[string]interface{}{
			SectionFieldAppId: _app.GetId(),
		},
		SerKeyAttrs: map[string]interface{}{
			SectionAttrTitle:    _title,
			SectionAttrIcon:     _icon,
			SectionAttrSummary:  _summary,
			SectionAttrPosition: _pos,
		},
	}
	if !reflect.DeepEqual(m, expected) {
		t.Fatalf("%s failed: expected %#v but received %#v", name, expected, m)
	}

	m = section.ToMap(func(input map[string]interface{}) map[string]interface{} {
		return map[string]interface{}{
			"FieldId":      input[henge.FieldId],
			"SerKeyFields": input[SerKeyFields],
			"SerKeyAttrs":  input[SerKeyAttrs],
		}
	})
	expected = map[string]interface{}{
		"FieldId": _id,
		"SerKeyFields": map[string]interface{}{
			SectionFieldAppId: _app.GetId(),
		},
		"SerKeyAttrs": map[string]interface{}{
			SectionAttrTitle:    _title,
			SectionAttrIcon:     _icon,
			SectionAttrSummary:  _summary,
			SectionAttrPosition: _pos,
		},
	}
	if !reflect.DeepEqual(m, expected) {
		t.Fatalf("%s failed: expected %#v but received %#v", name, expected, m)
	}
}

func TestSection_json(t *testing.T) {
	name := "TestSection_json"
	_tagVersion := uint64(1337)
	_appId := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isVisible := true
	_app := app.NewApp(_tagVersion, _appId, _name, _desc, _isVisible)

	_title := "Quick start"
	_icon := "default"
	_summary := "section one"
	_pos := rand.Intn(10242048)
	section1 := NewSection(_tagVersion, _app, _title, _icon, _summary)
	if section1 == nil {
		t.Fatalf("%s failed: nil", name)
	}
	section1.SetPosition(_pos)
	js1, _ := json.Marshal(section1)

	var section2 *Section
	err := json.Unmarshal(js1, &section2)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	if section1.GetTagVersion() != section2.GetTagVersion() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, section1.GetTagVersion(), section2.GetTagVersion())
	}
	if section1.GetId() != section2.GetId() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, section1.GetId(), section2.GetId())
	}
	if section1.GetAppId() != section2.GetAppId() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, section1.GetAppId(), section2.GetAppId())
	}
	if section1.GetTitle() != section2.GetTitle() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, section1.GetTitle(), section2.GetTitle())
	}
	if section1.GetIcon() != section2.GetIcon() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, section1.GetIcon(), section2.GetIcon())
	}
	if section1.GetSummary() != section2.GetSummary() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, section1.GetSummary(), section2.GetSummary())
	}
	if section1.GetPosition() != section2.GetPosition() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, section1.GetPosition(), section2.GetPosition())
	}
	if section1.GetChecksum() != section2.GetChecksum() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, section1.GetChecksum(), section2.GetChecksum())
	}
}
