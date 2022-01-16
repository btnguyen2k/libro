package libro

import (
	"encoding/json"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/btnguyen2k/henge"
	"main/src/gvabe/bo"
	"main/src/utils"
)

func TestNewTopic(t *testing.T) {
	testName := "TestNewTopic"
	teardownTest := setupTest(t, testName, func(t *testing.T, testName string) {
		bo.UboTimestampRouding = henge.TimestampRoundSettingNone
	}, nil)
	defer teardownTest(t)

	tstart := time.Now()
	_tagVersion := uint64(1337)
	_prodId := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_prod := NewProduct(_tagVersion, _prodId, _name, _desc, _isPublished)

	_title := "Quick start"
	_icon := "default"
	_summary := "topic one"
	_pos := rand.Intn(10242048)
	_numPages := _pos%10 + 1
	topic := NewTopic(_tagVersion, _prod, _title, _icon, _summary)
	if topic == nil {
		t.Fatalf("%s failed: nil", testName)
	}
	topic.SetPosition(_pos).SetNumPages(_numPages)

	_id := topic.GetId()
	if tagVersion := topic.GetTagVersion(); tagVersion != _tagVersion {
		t.Fatalf("%s failed: expected tag-version to be %#v but received %#v", testName, _tagVersion, tagVersion)
	}
	if id := topic.GetId(); id != _id {
		t.Fatalf("%s failed: expected id to be %#v but received %#v", testName, _id, id)
	}
	if prodId := topic.GetProductId(); prodId != _prod.GetId() {
		t.Fatalf("%s failed: expected product-id to be %#v but received %#v", testName, _prod.GetId(), prodId)
	}
	if title := topic.GetTitle(); title != _title {
		t.Fatalf("%s failed: expected title to be %#v but received %#v", testName, _title, title)
	}
	if icon := topic.GetIcon(); icon != _icon {
		t.Fatalf("%s failed: expected icon to be %#v but received %#v", testName, _icon, icon)
	}
	if summary := topic.GetSummary(); summary != _summary {
		t.Fatalf("%s failed: expected summary to be %#v but received %#v", testName, _summary, summary)
	}
	if pos := topic.GetPosition(); pos != _pos {
		t.Fatalf("%s failed: expected position to be %#v but received %#v", testName, _pos, pos)
	}
	if numPages := topic.GetNumPages(); numPages != _numPages {
		t.Fatalf("%s failed: expected num-pages to be %#v but received %#v", testName, _numPages, numPages)
	}

	tend := time.Now()
	if topic.GetTimeCreated().Before(tstart) || topic.GetTimeCreated().After(tend) {
		t.Fatalf("%s failed: timestamp-created is invalid\nStart: %s / Created: %s / End: %s", testName, tstart, topic.GetTimeCreated(), tend)
	}
	if topic.GetTimeUpdated().Before(tstart) || topic.GetTimeUpdated().After(tend) || topic.GetTimeUpdated().Before(topic.GetTimeCreated()) {
		t.Fatalf("%s failed: timestamp-updated is invalid\nStart: %s / Updated: %s / End: %s", testName, tstart, topic.GetTimeUpdated(), tend)
	}
}

func TestNewTopicFromUbo(t *testing.T) {
	testName := "TestNewTopicFromUbo"
	teardownTest := setupTest(t, testName, func(t *testing.T, testName string) {
		bo.UboTimestampRouding = henge.TimestampRoundSettingNone
	}, nil)
	defer teardownTest(t)

	if NewTopicFromUbo(nil) != nil {
		t.Fatalf("%s failed: NewTopicFromUbo(nil) should return nil", testName)
	}

	tstart := time.Now()
	_tagVersion := uint64(1337)
	_id := utils.UniqueId()
	_prodId := "libro"
	_title := "Quick start"
	_icon := "default"
	_summary := "topic one"
	_pos := rand.Intn(10242048)
	_numPages := _pos%10 + 1
	ubo := henge.NewUniversalBo(_id, _tagVersion, henge.UboOpt{TimeLayout: bo.UboTimeLayout, TimestampRounding: bo.UboTimestampRouding})
	ubo.SetExtraAttr(TopicFieldProductId, _prodId)
	ubo.SetExtraAttr(TopicFieldPosition, _pos)
	ubo.SetDataAttr(TopicAttrTitle, _title)
	ubo.SetDataAttr(TopicAttrIcon, _icon)
	ubo.SetDataAttr(TopicAttrSummary, _summary)
	ubo.SetDataAttr(TopicAttrNumPages, _numPages)

	topic := NewTopicFromUbo(ubo)
	if topic == nil {
		t.Fatalf("%s failed: nil", testName)
	}
	if tagVersion := topic.GetTagVersion(); tagVersion != _tagVersion {
		t.Fatalf("%s failed: expected tag-version to be %#v but received %#v", testName, _tagVersion, tagVersion)
	}
	if id := topic.GetId(); id != _id {
		t.Fatalf("%s failed: expected id to be %#v but received %#v", testName, _id, id)
	}
	if prodId := topic.GetProductId(); prodId != _prodId {
		t.Fatalf("%s failed: expected product-id to be %#v but received %#v", testName, _prodId, prodId)
	}
	if title := topic.GetTitle(); title != _title {
		t.Fatalf("%s failed: expected title to be %#v but received %#v", testName, _title, title)
	}
	if icon := topic.GetIcon(); icon != _icon {
		t.Fatalf("%s failed: expected icon to be %#v but received %#v", testName, _icon, icon)
	}
	if summary := topic.GetSummary(); summary != _summary {
		t.Fatalf("%s failed: expected summary to be %#v but received %#v", testName, _summary, summary)
	}
	if pos := topic.GetPosition(); pos != _pos {
		t.Fatalf("%s failed: expected position to be %#v but received %#v", testName, _pos, pos)
	}
	if numPages := topic.GetNumPages(); numPages != _numPages {
		t.Fatalf("%s failed: expected num-pages to be %#v but received %#v", testName, _numPages, numPages)
	}

	tend := time.Now()
	if topic.GetTimeCreated().Before(tstart) || topic.GetTimeCreated().After(tend) {
		t.Fatalf("%s failed: timestamp-created is invalid\nStart: %s / Created: %s / End: %s", testName, tstart, topic.GetTimeCreated(), tend)
	}
	if topic.GetTimeUpdated().Before(tstart) || topic.GetTimeUpdated().After(tend) || topic.GetTimeUpdated().Before(topic.GetTimeCreated()) {
		t.Fatalf("%s failed: timestamp-updated is invalid\nStart: %s / Updated: %s / End: %s", testName, tstart, topic.GetTimeUpdated(), tend)
	}
}

func TestTopic_ToMap(t *testing.T) {
	testName := "TestTopic_ToMap"
	teardownTest := setupTest(t, testName, nil, nil)
	defer teardownTest(t)

	_tagVersion := uint64(1337)
	_prodId := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_prod := NewProduct(_tagVersion, _prodId, _name, _desc, _isPublished)

	_title := "Quick start"
	_icon := "default"
	_summary := "topic one"
	_pos := rand.Intn(10242048)
	_numPages := _pos%10 + 1
	topic := NewTopic(_tagVersion, _prod, _title, _icon, _summary)
	if topic == nil {
		t.Fatalf("%s failed: nil", testName)
	}
	topic.SetPosition(_pos).SetNumPages(_numPages)
	_id := topic.GetId()

	m := topic.ToMap(nil)
	expected := map[string]interface{}{
		henge.FieldId:          _id,
		henge.FieldTimeCreated: topic.GetTimeCreated(),
		henge.FieldTimeUpdated: topic.GetTimeUpdated(),
		bo.SerKeyFields: map[string]interface{}{
			TopicFieldProductId: _prod.GetId(),
			TopicFieldPosition:  _pos,
		},
		bo.SerKeyAttrs: map[string]interface{}{
			TopicAttrTitle:    _title,
			TopicAttrIcon:     _icon,
			TopicAttrSummary:  _summary,
			TopicAttrNumPages: _numPages,
		},
	}
	if !reflect.DeepEqual(m, expected) {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, expected, m)
	}

	m = topic.ToMap(func(input map[string]interface{}) map[string]interface{} {
		return map[string]interface{}{
			"FieldId":          input[henge.FieldId],
			"FieldTimeCreated": input[henge.FieldTimeCreated],
			"FieldTimeUpdated": input[henge.FieldTimeUpdated],
			"SerKeyFields":     input[bo.SerKeyFields],
			"SerKeyAttrs":      input[bo.SerKeyAttrs],
		}
	})
	expected = map[string]interface{}{
		"FieldId":          _id,
		"FieldTimeCreated": topic.GetTimeCreated(),
		"FieldTimeUpdated": topic.GetTimeUpdated(),
		"SerKeyFields": map[string]interface{}{
			TopicFieldProductId: _prod.GetId(),
			TopicFieldPosition:  _pos,
		},
		"SerKeyAttrs": map[string]interface{}{
			TopicAttrTitle:    _title,
			TopicAttrIcon:     _icon,
			TopicAttrSummary:  _summary,
			TopicAttrNumPages: _numPages,
		},
	}
	if !reflect.DeepEqual(m, expected) {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, expected, m)
	}
}

func TestTopic_json(t *testing.T) {
	testName := "TestTopic_json"
	teardownTest := setupTest(t, testName, nil, nil)
	defer teardownTest(t)

	_tagVersion := uint64(1337)
	_prodId := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_prod := NewProduct(_tagVersion, _prodId, _name, _desc, _isPublished)

	_title := "Quick start"
	_icon := "default"
	_summary := "topic one"
	_pos := rand.Intn(10242048)
	_numPages := _pos%10 + 1
	topic1 := NewTopic(_tagVersion, _prod, _title, _icon, _summary)
	if topic1 == nil {
		t.Fatalf("%s failed: nil", testName)
	}
	topic1.SetPosition(_pos).SetNumPages(_numPages)
	js1, _ := json.Marshal(topic1)

	var topic2 *Topic
	err := json.Unmarshal(js1, &topic2)
	if err != nil {
		t.Fatalf("%s failed: %e", testName, err)
	}
	if topic1.GetTagVersion() != topic2.GetTagVersion() {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, topic1.GetTagVersion(), topic2.GetTagVersion())
	}
	if topic1.GetId() != topic2.GetId() {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, topic1.GetId(), topic2.GetId())
	}
	if topic1.GetProductId() != topic2.GetProductId() {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, topic1.GetProductId(), topic2.GetProductId())
	}
	if topic1.GetTitle() != topic2.GetTitle() {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, topic1.GetTitle(), topic2.GetTitle())
	}
	if topic1.GetIcon() != topic2.GetIcon() {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, topic1.GetIcon(), topic2.GetIcon())
	}
	if topic1.GetSummary() != topic2.GetSummary() {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, topic1.GetSummary(), topic2.GetSummary())
	}
	if topic1.GetPosition() != topic2.GetPosition() {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, topic1.GetPosition(), topic2.GetPosition())
	}
	if topic1.GetNumPages() != topic2.GetNumPages() {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, topic1.GetNumPages(), topic2.GetNumPages())
	}
	if topic1.GetChecksum() != topic2.GetChecksum() {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, topic1.GetChecksum(), topic2.GetChecksum())
	}
}
