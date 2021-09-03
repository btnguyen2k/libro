package doc

import (
	"math/rand"
	"testing"

	"github.com/btnguyen2k/henge"
	"main/src/gvabe/bo/app"
	"main/src/utils"
)

func TestNewPage(t *testing.T) {
	name := "TestNewPage"
	_tagVersion := uint64(1337)
	_appId := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isVisible := true
	_app := app.NewApp(_tagVersion, _appId, _name, _desc, _isVisible)

	_title := "Quick start"
	_icon := "default"
	_summary := "topic one"
	_pos := rand.Int()
	_topic := NewTopic(_tagVersion, _app, _title, _icon, _summary)
	if _topic == nil {
		t.Fatalf("%s failed: nil", name)
	}
	_topic.SetPosition(_pos)

	_content := "page one"
	page := NewPage(_tagVersion, _topic, _title+"-page", _icon+"-page", _summary+"-page", _content)
	if page == nil {
		t.Fatalf("%s failed: nil", name)
	}
	page.SetPosition(_pos + 1)

	_id := page.GetId()
	if tagVersion := page.GetTagVersion(); tagVersion != _tagVersion {
		t.Fatalf("%s failed: expected tag-version to be %#v but received %#v", name, _tagVersion, tagVersion)
	}
	if id := page.GetId(); id != _id {
		t.Fatalf("%s failed: expected id to be %#v but received %#v", name, _id, id)
	}
	if appId := page.GetAppId(); appId != _app.GetId() {
		t.Fatalf("%s failed: expected app-id to be %#v but received %#v", name, _app.GetId(), appId)
	}
	if topicId := page.GetTopicId(); topicId != _topic.GetId() {
		t.Fatalf("%s failed: expected topic-id to be %#v but received %#v", name, _topic.GetId(), topicId)
	}
	if title := page.GetTitle(); title != _title+"-page" {
		t.Fatalf("%s failed: expected title to be %#v but received %#v", name, _title+"-page", title)
	}
	if icon := page.GetIcon(); icon != _icon+"-page" {
		t.Fatalf("%s failed: expected icon to be %#v but received %#v", name, _icon+"-page", icon)
	}
	if summary := page.GetSummary(); summary != _summary+"-page" {
		t.Fatalf("%s failed: expected summary to be %#v but received %#v", name, _summary+"-page", summary)
	}
	if pos := page.GetPosition(); pos != _pos+1 {
		t.Fatalf("%s failed: expected position to be %#v but received %#v", name, _pos+1, pos)
	}
	if content := page.GetContent(); content != _content {
		t.Fatalf("%s failed: expected content to be %#v but received %#v", name, _content, content)
	}
}

func TestNewPageFromUbo(t *testing.T) {
	name := "TestNewPageFromUbo"

	if NewPageFromUbo(nil) != nil {
		t.Fatalf("%s failed: NewPageFromUbo(nil) should return nil", name)
	}
	_tagVersion := uint64(1337)
	_id := utils.UniqueId()
	_appId := "libro"
	_topicId := "intro"
	_title := "Quick start"
	_icon := "default"
	_summary := "page one"
	_pos := rand.Intn(10242048)
	_content := "page content"
	ubo := henge.NewUniversalBo(_id, _tagVersion)
	ubo.SetExtraAttr(PageFieldAppId, _appId)
	ubo.SetExtraAttr(PageFieldTopicId, _topicId)
	ubo.SetDataAttr(PageAttrTitle, _title)
	ubo.SetDataAttr(PageAttrIcon, _icon)
	ubo.SetDataAttr(PageAttrSummary, _summary)
	ubo.SetDataAttr(PageAttrPosition, _pos)
	ubo.SetDataAttr(PageAttrContent, _content)

	page := NewPageFromUbo(ubo)
	if page == nil {
		t.Fatalf("%s failed: nil", name)
	}
	if tagVersion := page.GetTagVersion(); tagVersion != _tagVersion {
		t.Fatalf("%s failed: expected tag-version to be %#v but received %#v", name, _tagVersion, tagVersion)
	}
	if id := page.GetId(); id != _id {
		t.Fatalf("%s failed: expected id to be %#v but received %#v", name, _id, id)
	}
	if appId := page.GetAppId(); appId != _appId {
		t.Fatalf("%s failed: expected app-id to be %#v but received %#v", name, _appId, appId)
	}
	if topicId := page.GetTopicId(); topicId != _topicId {
		t.Fatalf("%s failed: expected topic-id to be %#v but received %#v", name, _topicId, topicId)
	}
	if title := page.GetTitle(); title != _title {
		t.Fatalf("%s failed: expected title to be %#v but received %#v", name, _title, title)
	}
	if icon := page.GetIcon(); icon != _icon {
		t.Fatalf("%s failed: expected icon to be %#v but received %#v", name, _icon, icon)
	}
	if summary := page.GetSummary(); summary != _summary {
		t.Fatalf("%s failed: expected summary to be %#v but received %#v", name, _summary, summary)
	}
	if pos := page.GetPosition(); pos != _pos {
		t.Fatalf("%s failed: expected position to be %#v but received %#v", name, _pos, pos)
	}
	if content := page.GetContent(); content != _content {
		t.Fatalf("%s failed: expected content to be %#v but received %#v", name, _content, content)
	}
}

// func TestPage_ToMap(t *testing.T) {
// 	name := "TestPage_ToMap"
// 	_tagVersion := uint64(1337)
// 	_appId := "libro"
// 	_name := "Libro"
// 	_desc := "Libro description"
// 	_isVisible := true
// 	_app := app.NewApp(_tagVersion, _appId, _name, _desc, _isVisible)
//
// 	_title := "Quick start"
// 	_icon := "default"
// 	_summary := "topic one"
// 	_pos := rand.Intn(10242048)
// 	topic := NewPage(_tagVersion, _app, _title, _icon, _summary)
// 	if topic == nil {
// 		t.Fatalf("%s failed: nil", name)
// 	}
// 	topic.SetPosition(_pos)
// 	_id := topic.GetId()
//
// 	m := topic.ToMap(nil)
// 	expected := map[string]interface{}{
// 		henge.FieldId: _id,
// 		SerKeyFields: map[string]interface{}{
// 			PageFieldAppId: _app.GetId(),
// 		},
// 		SerKeyAttrs: map[string]interface{}{
// 			PageAttrTitle:    _title,
// 			PageAttrIcon:     _icon,
// 			PageAttrSummary:  _summary,
// 			PageAttrPosition: _pos,
// 		},
// 	}
// 	if !reflect.DeepEqual(m, expected) {
// 		t.Fatalf("%s failed: expected %#v but received %#v", name, expected, m)
// 	}
//
// 	m = topic.ToMap(func(input map[string]interface{}) map[string]interface{} {
// 		return map[string]interface{}{
// 			"FieldId":      input[henge.FieldId],
// 			"SerKeyFields": input[SerKeyFields],
// 			"SerKeyAttrs":  input[SerKeyAttrs],
// 		}
// 	})
// 	expected = map[string]interface{}{
// 		"FieldId": _id,
// 		"SerKeyFields": map[string]interface{}{
// 			PageFieldAppId: _app.GetId(),
// 		},
// 		"SerKeyAttrs": map[string]interface{}{
// 			PageAttrTitle:    _title,
// 			PageAttrIcon:     _icon,
// 			PageAttrSummary:  _summary,
// 			PageAttrPosition: _pos,
// 		},
// 	}
// 	if !reflect.DeepEqual(m, expected) {
// 		t.Fatalf("%s failed: expected %#v but received %#v", name, expected, m)
// 	}
// }
//
// func TestPage_json(t *testing.T) {
// 	name := "TestPage_json"
// 	_tagVersion := uint64(1337)
// 	_appId := "libro"
// 	_name := "Libro"
// 	_desc := "Libro description"
// 	_isVisible := true
// 	_app := app.NewApp(_tagVersion, _appId, _name, _desc, _isVisible)
//
// 	_title := "Quick start"
// 	_icon := "default"
// 	_summary := "topic one"
// 	_pos := rand.Intn(10242048)
// 	topic1 := NewPage(_tagVersion, _app, _title, _icon, _summary)
// 	if topic1 == nil {
// 		t.Fatalf("%s failed: nil", name)
// 	}
// 	topic1.SetPosition(_pos)
// 	js1, _ := json.Marshal(topic1)
//
// 	var topic2 *Page
// 	err := json.Unmarshal(js1, &topic2)
// 	if err != nil {
// 		t.Fatalf("%s failed: %e", name, err)
// 	}
// 	if topic1.GetTagVersion() != topic2.GetTagVersion() {
// 		t.Fatalf("%s failed: expected %#v but received %#v", name, topic1.GetTagVersion(), topic2.GetTagVersion())
// 	}
// 	if topic1.GetId() != topic2.GetId() {
// 		t.Fatalf("%s failed: expected %#v but received %#v", name, topic1.GetId(), topic2.GetId())
// 	}
// 	if topic1.GetAppId() != topic2.GetAppId() {
// 		t.Fatalf("%s failed: expected %#v but received %#v", name, topic1.GetAppId(), topic2.GetAppId())
// 	}
// 	if topic1.GetTitle() != topic2.GetTitle() {
// 		t.Fatalf("%s failed: expected %#v but received %#v", name, topic1.GetTitle(), topic2.GetTitle())
// 	}
// 	if topic1.GetIcon() != topic2.GetIcon() {
// 		t.Fatalf("%s failed: expected %#v but received %#v", name, topic1.GetIcon(), topic2.GetIcon())
// 	}
// 	if topic1.GetSummary() != topic2.GetSummary() {
// 		t.Fatalf("%s failed: expected %#v but received %#v", name, topic1.GetSummary(), topic2.GetSummary())
// 	}
// 	if topic1.GetPosition() != topic2.GetPosition() {
// 		t.Fatalf("%s failed: expected %#v but received %#v", name, topic1.GetPosition(), topic2.GetPosition())
// 	}
// 	if topic1.GetChecksum() != topic2.GetChecksum() {
// 		t.Fatalf("%s failed: expected %#v but received %#v", name, topic1.GetChecksum(), topic2.GetChecksum())
// 	}
// }