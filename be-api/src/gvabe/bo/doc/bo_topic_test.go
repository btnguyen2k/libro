package doc

import (
	"encoding/json"
	"math/rand"
	"reflect"
	"testing"

	"github.com/btnguyen2k/henge"
	"main/src/gvabe/bo"
	"main/src/gvabe/bo/product"
	"main/src/utils"
)

func TestNewTopic(t *testing.T) {
	name := "TestNewTopic"
	_tagVersion := uint64(1337)
	_appId := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_app := product.NewProduct(_tagVersion, _appId, _name, _desc, _isPublished)

	_title := "Quick start"
	_icon := "default"
	_summary := "topic one"
	_pos := rand.Intn(10242048)
	_numPages := _pos%10 + 1
	topic := NewTopic(_tagVersion, _app, _title, _icon, _summary)
	if topic == nil {
		t.Fatalf("%s failed: nil", name)
	}
	topic.SetPosition(_pos).SetNumPages(_numPages)

	_id := topic.GetId()
	if tagVersion := topic.GetTagVersion(); tagVersion != _tagVersion {
		t.Fatalf("%s failed: expected tag-version to be %#v but received %#v", name, _tagVersion, tagVersion)
	}
	if id := topic.GetId(); id != _id {
		t.Fatalf("%s failed: expected id to be %#v but received %#v", name, _id, id)
	}
	if appId := topic.GetAppId(); appId != _app.GetId() {
		t.Fatalf("%s failed: expected app-id to be %#v but received %#v", name, _app.GetId(), appId)
	}
	if title := topic.GetTitle(); title != _title {
		t.Fatalf("%s failed: expected title to be %#v but received %#v", name, _title, title)
	}
	if icon := topic.GetIcon(); icon != _icon {
		t.Fatalf("%s failed: expected icon to be %#v but received %#v", name, _icon, icon)
	}
	if summary := topic.GetSummary(); summary != _summary {
		t.Fatalf("%s failed: expected summary to be %#v but received %#v", name, _summary, summary)
	}
	if pos := topic.GetPosition(); pos != _pos {
		t.Fatalf("%s failed: expected position to be %#v but received %#v", name, _pos, pos)
	}
	if numPages := topic.GetNumPages(); numPages != _numPages {
		t.Fatalf("%s failed: expected num-pages to be %#v but received %#v", name, _numPages, numPages)
	}
}

func TestNewTopicFromUbo(t *testing.T) {
	name := "TestNewTopicFromUbo"

	if NewTopicFromUbo(nil) != nil {
		t.Fatalf("%s failed: NewTopicFromUbo(nil) should return nil", name)
	}
	_tagVersion := uint64(1337)
	_id := utils.UniqueId()
	_appId := "libro"
	_title := "Quick start"
	_icon := "default"
	_summary := "topic one"
	_pos := rand.Intn(10242048)
	_numPages := _pos%10 + 1
	ubo := henge.NewUniversalBo(_id, _tagVersion)
	ubo.SetExtraAttr(TopicFieldAppId, _appId)
	ubo.SetDataAttr(TopicAttrTitle, _title)
	ubo.SetDataAttr(TopicAttrIcon, _icon)
	ubo.SetDataAttr(TopicAttrSummary, _summary)
	ubo.SetDataAttr(TopicAttrPosition, _pos)
	ubo.SetDataAttr(TopicAttrNumPages, _numPages)

	topic := NewTopicFromUbo(ubo)
	if topic == nil {
		t.Fatalf("%s failed: nil", name)
	}
	if tagVersion := topic.GetTagVersion(); tagVersion != _tagVersion {
		t.Fatalf("%s failed: expected tag-version to be %#v but received %#v", name, _tagVersion, tagVersion)
	}
	if id := topic.GetId(); id != _id {
		t.Fatalf("%s failed: expected id to be %#v but received %#v", name, _id, id)
	}
	if appId := topic.GetAppId(); appId != _appId {
		t.Fatalf("%s failed: expected app-id to be %#v but received %#v", name, _appId, appId)
	}
	if title := topic.GetTitle(); title != _title {
		t.Fatalf("%s failed: expected title to be %#v but received %#v", name, _title, title)
	}
	if icon := topic.GetIcon(); icon != _icon {
		t.Fatalf("%s failed: expected icon to be %#v but received %#v", name, _icon, icon)
	}
	if summary := topic.GetSummary(); summary != _summary {
		t.Fatalf("%s failed: expected summary to be %#v but received %#v", name, _summary, summary)
	}
	if pos := topic.GetPosition(); pos != _pos {
		t.Fatalf("%s failed: expected position to be %#v but received %#v", name, _pos, pos)
	}
	if numPages := topic.GetNumPages(); numPages != _numPages {
		t.Fatalf("%s failed: expected num-pages to be %#v but received %#v", name, _numPages, numPages)
	}
}

func TestTopic_ToMap(t *testing.T) {
	name := "TestTopic_ToMap"
	_tagVersion := uint64(1337)
	_appId := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_app := product.NewProduct(_tagVersion, _appId, _name, _desc, _isPublished)

	_title := "Quick start"
	_icon := "default"
	_summary := "topic one"
	_pos := rand.Intn(10242048)
	_numPages := _pos%10 + 1
	topic := NewTopic(_tagVersion, _app, _title, _icon, _summary)
	if topic == nil {
		t.Fatalf("%s failed: nil", name)
	}
	topic.SetPosition(_pos).SetNumPages(_numPages)
	_id := topic.GetId()

	m := topic.ToMap(nil)
	expected := map[string]interface{}{
		henge.FieldId: _id,
		bo.SerKeyFields: map[string]interface{}{
			TopicFieldAppId: _app.GetId(),
		},
		bo.SerKeyAttrs: map[string]interface{}{
			TopicAttrTitle:    _title,
			TopicAttrIcon:     _icon,
			TopicAttrSummary:  _summary,
			TopicAttrPosition: _pos,
			TopicAttrNumPages: _numPages,
		},
	}
	if !reflect.DeepEqual(m, expected) {
		t.Fatalf("%s failed: expected %#v but received %#v", name, expected, m)
	}

	m = topic.ToMap(func(input map[string]interface{}) map[string]interface{} {
		return map[string]interface{}{
			"FieldId":      input[henge.FieldId],
			"SerKeyFields": input[bo.SerKeyFields],
			"SerKeyAttrs":  input[bo.SerKeyAttrs],
		}
	})
	expected = map[string]interface{}{
		"FieldId": _id,
		"SerKeyFields": map[string]interface{}{
			TopicFieldAppId: _app.GetId(),
		},
		"SerKeyAttrs": map[string]interface{}{
			TopicAttrTitle:    _title,
			TopicAttrIcon:     _icon,
			TopicAttrSummary:  _summary,
			TopicAttrPosition: _pos,
			TopicAttrNumPages: _numPages,
		},
	}
	if !reflect.DeepEqual(m, expected) {
		t.Fatalf("%s failed: expected %#v but received %#v", name, expected, m)
	}
}

func TestTopic_json(t *testing.T) {
	name := "TestTopic_json"
	_tagVersion := uint64(1337)
	_appId := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_app := product.NewProduct(_tagVersion, _appId, _name, _desc, _isPublished)

	_title := "Quick start"
	_icon := "default"
	_summary := "topic one"
	_pos := rand.Intn(10242048)
	_numPages := _pos%10 + 1
	topic1 := NewTopic(_tagVersion, _app, _title, _icon, _summary)
	if topic1 == nil {
		t.Fatalf("%s failed: nil", name)
	}
	topic1.SetPosition(_pos).SetNumPages(_numPages)
	js1, _ := json.Marshal(topic1)

	var topic2 *Topic
	err := json.Unmarshal(js1, &topic2)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	if topic1.GetTagVersion() != topic2.GetTagVersion() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, topic1.GetTagVersion(), topic2.GetTagVersion())
	}
	if topic1.GetId() != topic2.GetId() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, topic1.GetId(), topic2.GetId())
	}
	if topic1.GetAppId() != topic2.GetAppId() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, topic1.GetAppId(), topic2.GetAppId())
	}
	if topic1.GetTitle() != topic2.GetTitle() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, topic1.GetTitle(), topic2.GetTitle())
	}
	if topic1.GetIcon() != topic2.GetIcon() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, topic1.GetIcon(), topic2.GetIcon())
	}
	if topic1.GetSummary() != topic2.GetSummary() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, topic1.GetSummary(), topic2.GetSummary())
	}
	if topic1.GetPosition() != topic2.GetPosition() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, topic1.GetPosition(), topic2.GetPosition())
	}
	if topic1.GetNumPages() != topic2.GetNumPages() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, topic1.GetNumPages(), topic2.GetNumPages())
	}
	if topic1.GetChecksum() != topic2.GetChecksum() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, topic1.GetChecksum(), topic2.GetChecksum())
	}
}
