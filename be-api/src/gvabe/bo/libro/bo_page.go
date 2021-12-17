package libro

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/btnguyen2k/consu/reddo"
	"github.com/btnguyen2k/henge"
	"main/src/gvabe/bo"
	"main/src/utils"
)

// NewPage is helper function to create new Page bo.
func NewPage(tagVersion uint64, topic *Topic, title, icon, summary, content string) *Page {
	id := utils.UniqueId()
	bo := &Page{
		UniversalBo: henge.NewUniversalBo(id, tagVersion),
	}
	position := time.Now().Unix()
	return bo.
		SetProductId(topic.GetProductId()).
		SetTopicId(topic.GetId()).
		SetTitle(title).
		SetIcon(icon).
		SetSummary(summary).
		SetPosition(int(position)).
		SetContent(content).
		sync()
}

// NewPageFromUbo is helper function to create Page bo from a universal bo.
func NewPageFromUbo(ubo *henge.UniversalBo) *Page {
	if ubo == nil {
		return nil
	}
	ubo = ubo.Clone()
	bo := &Page{UniversalBo: ubo}
	if v, err := ubo.GetExtraAttrAs(PageFieldProductId, reddo.TypeString); err != nil {
		return nil
	} else {
		bo.productId, _ = v.(string)
	}
	if v, err := ubo.GetExtraAttrAs(PageFieldTopicId, reddo.TypeString); err != nil {
		return nil
	} else {
		bo.topicId, _ = v.(string)
	}
	if v, err := ubo.GetExtraAttrAs(PageFieldPosition, reddo.TypeInt); err != nil {
		return nil
	} else if temp, ok := v.(int64); ok {
		bo.position = int(temp)
	}
	if v, err := ubo.GetDataAttrAs(PageAttrTitle, reddo.TypeString); err != nil {
		return nil
	} else {
		bo.title, _ = v.(string)
	}
	if v, err := ubo.GetDataAttrAs(PageAttrIcon, reddo.TypeString); err != nil {
		return nil
	} else {
		bo.icon, _ = v.(string)
	}
	if v, err := ubo.GetDataAttrAs(PageAttrSummary, reddo.TypeString); err != nil {
		return nil
	} else {
		bo.summary, _ = v.(string)
	}
	if v, err := ubo.GetDataAttrAs(PageAttrContent, reddo.TypeString); err != nil {
		return nil
	} else {
		bo.content, _ = v.(string)
	}
	return bo.sync()
}

const (
	// PageFieldProductId holds id of the product that the document page belongs to
	PageFieldProductId = "prod"

	// PageFieldTopicId holds id of the topic that the document page belongs to
	PageFieldTopicId = "topic"

	// PageFieldPosition is the relative position of document page (for ordering purpose)
	PageFieldPosition = "pos"

	// PageAttrTitle is document page's title
	PageAttrTitle = "title"

	// PageAttrIcon is the id of document page's icon
	PageAttrIcon = "icon"

	// PageAttrSummary is the summary text of document page
	PageAttrSummary = "summary"

	// PageAttrContent is the content of document page
	PageAttrContent = "content"

	// pageAttrUbo is for internal use only!
	pageAttrUbo = "_ubo"
)

// Page is the business object.
//   - Page inherits unique id from bo.UniversalBo
//   - Page's id must be unique throughout the system
type Page struct {
	*henge.UniversalBo `json:"_ubo"`
	productId          string `json:"prod"`
	topicId            string `json:"topic"`
	position           int    `json:"pos"`
	title              string `json:"title"`
	icon               string `json:"icon"`
	summary            string `json:"summary"`
	content            string `json:"content"`
}

// ToMap transforms page's attributes to a map.
//
// The function returns a map with the following structure:
//   {
//     henge.FieldId: s.GetId(),
//     SerKeyFields: map[string]interface{}{
//         // all BO's top-level custom fields go here
//		},
//     SerKeyAttrs: map[string]interface{}{
//         // all BO's custom attributes go here
//     },
//   }
func (p *Page) ToMap(postFunc henge.FuncPostUboToMap) map[string]interface{} {
	result := map[string]interface{}{
		henge.FieldId: p.GetId(),
		henge.FieldTimeCreated: p.GetTimeCreated(),
		henge.FieldTimeUpdated: p.GetTimeUpdated(),
		bo.SerKeyFields: map[string]interface{}{
			PageFieldProductId: p.productId,
			PageFieldTopicId:   p.topicId,
			PageFieldPosition:  p.position,
		},
		bo.SerKeyAttrs: map[string]interface{}{
			PageAttrTitle:   p.title,
			PageAttrIcon:    p.icon,
			PageAttrSummary: p.summary,
			PageAttrContent: p.content,
		},
	}
	if postFunc != nil {
		result = postFunc(result)
	}
	return result
}

// MarshalJSON implements json.encode.Marshaler.MarshalJSON
// TODO: lock for read?
func (p *Page) MarshalJSON() ([]byte, error) {
	p.sync()
	m := map[string]interface{}{
		pageAttrUbo: p.UniversalBo.Clone(),
		bo.SerKeyFields: map[string]interface{}{
			PageFieldProductId: p.productId,
			PageFieldTopicId:   p.topicId,
			PageFieldPosition:  p.position,
		},
		bo.SerKeyAttrs: map[string]interface{}{
			PageAttrTitle:   p.title,
			PageAttrIcon:    p.icon,
			PageAttrSummary: p.summary,
			PageAttrContent: p.content,
		},
	}
	return json.Marshal(m)
}

// UnmarshalJSON implements json.decode.Unmarshaler.UnmarshalJSON
// TODO: lock for write?
func (p *Page) UnmarshalJSON(data []byte) error {
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	var err error
	if m[pageAttrUbo] != nil {
		js, _ := json.Marshal(m[pageAttrUbo])
		if err = json.Unmarshal(js, &p.UniversalBo); err != nil {
			return err
		}
	}
	if _fields, ok := m[bo.SerKeyFields].(map[string]interface{}); ok {
		if p.productId, err = reddo.ToString(_fields[PageFieldProductId]); err != nil {
			return err
		}
		if p.topicId, err = reddo.ToString(_fields[PageFieldTopicId]); err != nil {
			return err
		}
		if v, err := reddo.ToInt(_fields[PageFieldPosition]); err != nil {
			return err
		} else {
			p.position = int(v)
		}
	}
	if _attrs, ok := m[bo.SerKeyAttrs].(map[string]interface{}); ok {
		if p.title, err = reddo.ToString(_attrs[PageAttrTitle]); err != nil {
			return err
		}
		if p.icon, err = reddo.ToString(_attrs[PageAttrIcon]); err != nil {
			return err
		}
		if p.summary, err = reddo.ToString(_attrs[PageAttrSummary]); err != nil {
			return err
		}
		if p.content, err = reddo.ToString(_attrs[PageAttrContent]); err != nil {
			return err
		}
	}
	p.sync()
	return nil
}

// GetProductId returns value of page's 'product-id' attribute.
func (p *Page) GetProductId() string {
	return p.productId
}

// SetProductId sets value of page's 'product-id' attribute.
func (p *Page) SetProductId(v string) *Page {
	p.productId = strings.ToLower(strings.TrimSpace(v))
	return p
}

// GetTopicId returns value of page's 'topic-id' attribute.
func (p *Page) GetTopicId() string {
	return p.topicId
}

// SetTopicId sets value of page's 'topic-id' attribute.
func (p *Page) SetTopicId(v string) *Page {
	p.topicId = strings.ToLower(strings.TrimSpace(v))
	return p
}

// GetTitle returns value of page's 'title' attribute.
func (p *Page) GetTitle() string {
	return p.title
}

// SetTitle sets value of page's 'title' attribute.
func (p *Page) SetTitle(v string) *Page {
	p.title = strings.TrimSpace(v)
	return p
}

// GetIcon returns value of page's 'icon' attribute.
func (p *Page) GetIcon() string {
	return p.icon
}

// SetIcon sets value of page's 'icon' attribute.
func (p *Page) SetIcon(v string) *Page {
	p.icon = strings.TrimSpace(v)
	return p
}

// GetSummary returns value of page's 'summary' attribute.
func (p *Page) GetSummary() string {
	return p.summary
}

// SetSummary sets value of page's 'summary' attribute.
func (p *Page) SetSummary(v string) *Page {
	p.summary = strings.TrimSpace(v)
	return p
}

// GetPosition returns value of page's 'position' attribute.
func (p *Page) GetPosition() int {
	return p.position
}

// SetPosition sets value of page's 'position' attribute.
func (p *Page) SetPosition(v int) *Page {
	p.position = v
	return p
}

// GetContent returns value of page's 'content' attribute.
func (p *Page) GetContent() string {
	return p.content
}

// SetContent sets value of page's 'content' attribute.
func (p *Page) SetContent(v string) *Page {
	p.content = strings.TrimSpace(v)
	return p
}

// sync is called to synchronize BO's attributes to its UniversalBo.
func (p *Page) sync() *Page {
	p.SetExtraAttr(PageFieldProductId, p.productId)
	p.SetExtraAttr(PageFieldTopicId, p.topicId)
	p.SetExtraAttr(PageFieldPosition, p.position)
	p.SetDataAttr(PageAttrTitle, p.title)
	p.SetDataAttr(PageAttrIcon, p.icon)
	p.SetDataAttr(PageAttrSummary, p.summary)
	p.SetDataAttr(PageAttrContent, p.content)
	p.UniversalBo.Sync()
	return p
}
