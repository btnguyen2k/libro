package doc

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/btnguyen2k/consu/reddo"
	"github.com/btnguyen2k/henge"
	"main/src/gvabe/bo"
	"main/src/gvabe/bo/app"
	"main/src/utils"
)

// NewTopic is helper function to create new Topic bo.
func NewTopic(appVersion uint64, app *app.App, title, icon, summary string) *Topic {
	id := utils.UniqueId()
	bo := &Topic{
		UniversalBo: henge.NewUniversalBo(id, appVersion),
	}
	position := time.Now().Unix()
	return bo.
		SetAppId(app.GetId()).
		SetTitle(title).
		SetIcon(icon).
		SetSummary(summary).
		SetPosition(int(position)).
		sync()
}

// NewTopicFromUbo is helper function to create Topic bo from a universal bo.
func NewTopicFromUbo(ubo *henge.UniversalBo) *Topic {
	if ubo == nil {
		return nil
	}
	ubo = ubo.Clone()
	bo := &Topic{UniversalBo: ubo}
	if v, err := ubo.GetExtraAttrAs(TopicFieldAppId, reddo.TypeString); err != nil {
		return nil
	} else {
		bo.appId, _ = v.(string)
	}
	if v, err := ubo.GetDataAttrAs(TopicAttrTitle, reddo.TypeString); err != nil {
		return nil
	} else {
		bo.title, _ = v.(string)
	}
	if v, err := ubo.GetDataAttrAs(TopicAttrIcon, reddo.TypeString); err != nil {
		return nil
	} else {
		bo.icon, _ = v.(string)
	}
	if v, err := ubo.GetDataAttrAs(TopicAttrSummary, reddo.TypeString); err != nil {
		return nil
	} else {
		bo.summary, _ = v.(string)
	}
	if v, err := ubo.GetDataAttrAs(TopicAttrPosition, reddo.TypeInt); err != nil {
		return nil
	} else if temp, ok := v.(int64); ok {
		bo.position = int(temp)
	}
	return bo.sync()
}

const (
	// TopicFieldAppId holds id of the application that the document topic belongs to
	TopicFieldAppId = "app"

	// TopicAttrTitle is document topic's title
	TopicAttrTitle = "title"

	// TopicAttrIcon is the id of document topic's icon
	TopicAttrIcon = "icon"

	// TopicAttrSummary is the summary text of document topic
	TopicAttrSummary = "summary"

	// TopicAttrPosition is the relative position of document topic (for ordering purpose)
	TopicAttrPosition = "pos"

	// topicAttr_Ubo is for internal use only!
	topicAttr_Ubo = "_ubo"
)

// Topic is the business object.
//   - Topic inherits unique id from bo.UniversalBo
type Topic struct {
	*henge.UniversalBo `json:"_ubo"`
	appId              string `json:"app"`
	title              string `json:"title"`
	icon               string `json:"icon"`
	summary            string `json:"summary"`
	position           int    `json:"pos"`
}

// ToMap transforms topic's attributes to a map.
//
// The function returns a map with the following structure:
//   {
//     henge.FieldId: t.GetId(),
//     SerKeyFields: map[string]interface{}{
//         // all BO's top-level custom fields go here
//		},
//     SerKeyAttrs: map[string]interface{}{
//         // all BO's custom attributes go here
//     },
//   }
func (t *Topic) ToMap(postFunc henge.FuncPostUboToMap) map[string]interface{} {
	result := map[string]interface{}{
		henge.FieldId: t.GetId(),
		bo.SerKeyFields: map[string]interface{}{
			TopicFieldAppId: t.appId,
		},
		bo.SerKeyAttrs: map[string]interface{}{
			TopicAttrTitle:    t.title,
			TopicAttrIcon:     t.icon,
			TopicAttrSummary:  t.summary,
			TopicAttrPosition: t.position,
		},
	}
	if postFunc != nil {
		result = postFunc(result)
	}
	return result
}

// MarshalJSON implements json.encode.Marshaler.MarshalJSON
// TODO: lock for read?
func (t *Topic) MarshalJSON() ([]byte, error) {
	t.sync()
	m := map[string]interface{}{
		topicAttr_Ubo: t.UniversalBo.Clone(),
		bo.SerKeyFields: map[string]interface{}{
			TopicFieldAppId: t.appId,
		},
		bo.SerKeyAttrs: map[string]interface{}{
			TopicAttrTitle:    t.title,
			TopicAttrIcon:     t.icon,
			TopicAttrSummary:  t.summary,
			TopicAttrPosition: t.position,
		},
	}
	return json.Marshal(m)
}

// UnmarshalJSON implements json.decode.Unmarshaler.UnmarshalJSON
// TODO: lock for write?
func (t *Topic) UnmarshalJSON(data []byte) error {
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	var err error
	if m[topicAttr_Ubo] != nil {
		js, _ := json.Marshal(m[topicAttr_Ubo])
		if err = json.Unmarshal(js, &t.UniversalBo); err != nil {
			return err
		}
	}
	if _fields, ok := m[bo.SerKeyFields].(map[string]interface{}); ok {
		if t.appId, err = reddo.ToString(_fields[TopicFieldAppId]); err != nil {
			return err
		}
	}
	if _attrs, ok := m[bo.SerKeyAttrs].(map[string]interface{}); ok {
		if t.title, err = reddo.ToString(_attrs[TopicAttrTitle]); err != nil {
			return err
		}
		if t.icon, err = reddo.ToString(_attrs[TopicAttrIcon]); err != nil {
			return err
		}
		if t.summary, err = reddo.ToString(_attrs[TopicAttrSummary]); err != nil {
			return err
		}
		if v, err := reddo.ToInt(_attrs[TopicAttrPosition]); err != nil {
			return err
		} else {
			t.position = int(v)
		}
	}
	t.sync()
	return nil
}

// GetAppId returns value of topic's 'app-id' attribute.
func (t *Topic) GetAppId() string {
	return t.appId
}

// SetAppId sets value of topic's 'app-id' attribute.
func (t *Topic) SetAppId(v string) *Topic {
	t.appId = strings.ToLower(strings.TrimSpace(v))
	return t
}

// GetTitle returns value of topic's 'title' attribute.
func (t *Topic) GetTitle() string {
	return t.title
}

// SetTitle sets value of topic's 'title' attribute.
func (t *Topic) SetTitle(v string) *Topic {
	t.title = strings.TrimSpace(v)
	return t
}

// GetIcon returns value of topic's 'icon' attribute.
func (t *Topic) GetIcon() string {
	return t.icon
}

// SetIcon sets value of topic's 'icon' attribute.
func (t *Topic) SetIcon(v string) *Topic {
	t.icon = strings.TrimSpace(v)
	return t
}

// GetSummary returns value of topic's 'summary' attribute.
func (t *Topic) GetSummary() string {
	return t.summary
}

// SetSummary sets value of topic's 'summary' attribute.
func (t *Topic) SetSummary(v string) *Topic {
	t.summary = strings.TrimSpace(v)
	return t
}

// GetPosition returns value of topic's 'position' attribute.
func (t *Topic) GetPosition() int {
	return t.position
}

// SetPosition sets value of topic's 'position' attribute.
func (t *Topic) SetPosition(v int) *Topic {
	t.position = v
	return t
}

// sync is called to synchronize BO's attributes to its UniversalBo.
func (t *Topic) sync() *Topic {
	t.SetExtraAttr(TopicFieldAppId, t.appId)
	t.SetDataAttr(TopicAttrTitle, t.title)
	t.SetDataAttr(TopicAttrIcon, t.icon)
	t.SetDataAttr(TopicAttrSummary, t.summary)
	t.SetDataAttr(TopicAttrPosition, t.position)
	t.UniversalBo.Sync()
	return t
}