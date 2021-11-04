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

func TestNewPage(t *testing.T) {
	name := "TestNewPage"
	_tagVersion := uint64(1337)
	_prodId := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_prod := product.NewProduct(_tagVersion, _prodId, _name, _desc, _isPublished)

	_title := "Quick start"
	_icon := "default"
	_summary := "topic one"
	_pos := rand.Intn(10242048)
	_topic := NewTopic(_tagVersion, _prod, _title, _icon, _summary)
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
	if prodId := page.GetProductId(); prodId != _prod.GetId() {
		t.Fatalf("%s failed: expected product-id to be %#v but received %#v", name, _prod.GetId(), prodId)
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
	_prodId := "libro"
	_topicId := "intro"
	_title := "Quick start"
	_icon := "default"
	_summary := "page one"
	_pos := rand.Intn(10242048)
	_content := "page content"
	ubo := henge.NewUniversalBo(_id, _tagVersion)
	ubo.SetExtraAttr(PageFieldProductId, _prodId)
	ubo.SetExtraAttr(PageFieldTopicId, _topicId)
	ubo.SetExtraAttr(PageFieldPosition, _pos)
	ubo.SetDataAttr(PageAttrTitle, _title)
	ubo.SetDataAttr(PageAttrIcon, _icon)
	ubo.SetDataAttr(PageAttrSummary, _summary)
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
	if prodId := page.GetProductId(); prodId != _prodId {
		t.Fatalf("%s failed: expected product-id to be %#v but received %#v", name, _prodId, prodId)
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

func TestPage_ToMap(t *testing.T) {
	name := "TestPage_ToMap"
	_tagVersion := uint64(1337)
	_prodId := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_prod := product.NewProduct(_tagVersion, _prodId, _name, _desc, _isPublished)

	_title := "Quick start"
	_icon := "default"
	_summary := "topic one"
	_pos := rand.Intn(10242048)
	_topic := NewTopic(_tagVersion, _prod, _title, _icon, _summary)
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
	m := page.ToMap(nil)
	expected := map[string]interface{}{
		henge.FieldId: _id,
		bo.SerKeyFields: map[string]interface{}{
			PageFieldProductId: _prod.GetId(),
			PageFieldTopicId:   _topic.GetId(),
			PageFieldPosition:  _pos + 1,
		},
		bo.SerKeyAttrs: map[string]interface{}{
			PageAttrTitle:   _title + "-page",
			PageAttrIcon:    _icon + "-page",
			PageAttrSummary: _summary + "-page",
			PageAttrContent: _content,
		},
	}
	if !reflect.DeepEqual(m, expected) {
		t.Fatalf("%s failed: expected %#v but received %#v", name, expected, m)
	}

	m = page.ToMap(func(input map[string]interface{}) map[string]interface{} {
		return map[string]interface{}{
			"FieldId":      input[henge.FieldId],
			"SerKeyFields": input[bo.SerKeyFields],
			"SerKeyAttrs":  input[bo.SerKeyAttrs],
		}
	})
	expected = map[string]interface{}{
		"FieldId": _id,
		"SerKeyFields": map[string]interface{}{
			PageFieldProductId: _prod.GetId(),
			PageFieldTopicId:   _topic.GetId(),
			PageFieldPosition:  _pos + 1,
		},
		"SerKeyAttrs": map[string]interface{}{
			PageAttrTitle:   _title + "-page",
			PageAttrIcon:    _icon + "-page",
			PageAttrSummary: _summary + "-page",
			PageAttrContent: _content,
		},
	}
	if !reflect.DeepEqual(m, expected) {
		t.Fatalf("%s failed: expected %#v but received %#v", name, expected, m)
	}
}

func TestPage_json(t *testing.T) {
	name := "TestPage_json"
	_tagVersion := uint64(1337)
	_prodId := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_prod := product.NewProduct(_tagVersion, _prodId, _name, _desc, _isPublished)

	_title := "Quick start"
	_icon := "default"
	_summary := "topic one"
	_pos := rand.Intn(10242048)
	_topic := NewTopic(_tagVersion, _prod, _title, _icon, _summary)
	if _topic == nil {
		t.Fatalf("%s failed: nil", name)
	}
	_topic.SetPosition(_pos)

	_content := "page one"
	page1 := NewPage(_tagVersion, _topic, _title+"-page", _icon+"-page", _summary+"-page", _content)
	if page1 == nil {
		t.Fatalf("%s failed: nil", name)
	}
	page1.SetPosition(_pos + 1)
	js1, _ := json.Marshal(page1)

	var page2 *Page
	err := json.Unmarshal(js1, &page2)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	if page1.GetTagVersion() != page2.GetTagVersion() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, page1.GetTagVersion(), page2.GetTagVersion())
	}
	if page1.GetId() != page2.GetId() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, page1.GetId(), page2.GetId())
	}
	if page1.GetProductId() != page2.GetProductId() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, page1.GetProductId(), page2.GetProductId())
	}
	if page1.GetTopicId() != page2.GetTopicId() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, page1.GetTopicId(), page2.GetTopicId())
	}
	if page1.GetTitle() != page2.GetTitle() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, page1.GetTitle(), page2.GetTitle())
	}
	if page1.GetIcon() != page2.GetIcon() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, page1.GetIcon(), page2.GetIcon())
	}
	if page1.GetSummary() != page2.GetSummary() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, page1.GetSummary(), page2.GetSummary())
	}
	if page1.GetPosition() != page2.GetPosition() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, page1.GetPosition(), page2.GetPosition())
	}
	if page1.GetContent() != page2.GetContent() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, page1.GetContent(), page2.GetContent())
	}
	if page1.GetChecksum() != page2.GetChecksum() {
		t.Fatalf("%s failed: expected %#v but received %#v", name, page1.GetChecksum(), page2.GetChecksum())
	}
}
