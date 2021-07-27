package doc

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/btnguyen2k/consu/reddo"
	"github.com/btnguyen2k/henge"
	"main/src/gvabe/bo/app"
	"main/src/utils"
)

// NewSection is helper function to create new Section bo.
func NewSection(appVersion uint64, app *app.App, title, icon, summary string) *Section {
	id := utils.UniqueId()
	bo := &Section{
		UniversalBo: henge.NewUniversalBo(id, appVersion),
	}
	position := time.Now().Unix()
	return bo.SetAppId(app.GetId()).SetTitle(title).SetIcon(icon).SetSummary(summary).SetPosition(int(position)).sync()
}

// NewSectionFromUbo is helper function to create Section bo from a universal bo.
func NewSectionFromUbo(ubo *henge.UniversalBo) *Section {
	if ubo == nil {
		return nil
	}
	ubo = ubo.Clone()
	bo := &Section{UniversalBo: ubo}
	if v, err := ubo.GetExtraAttrAs(SectionFieldAppId, reddo.TypeString); err != nil {
		return nil
	} else {
		bo.appId, _ = v.(string)
	}
	if v, err := ubo.GetDataAttrAs(SectionAttrTitle, reddo.TypeString); err != nil {
		return nil
	} else {
		bo.title, _ = v.(string)
	}
	if v, err := ubo.GetDataAttrAs(SectionAttrIcon, reddo.TypeString); err != nil {
		return nil
	} else {
		bo.icon, _ = v.(string)
	}
	if v, err := ubo.GetDataAttrAs(SectionAttrSummary, reddo.TypeString); err != nil {
		return nil
	} else {
		bo.summary, _ = v.(string)
	}
	if v, err := ubo.GetDataAttrAs(SectionAttrPosition, reddo.TypeInt); err != nil {
		return nil
	} else if temp, ok := v.(int64); ok {
		bo.position = int(temp)
	}
	return bo.sync()
}

const (
	// SectionFieldAppId holds id of the application that the document section belong to
	SectionFieldAppId = "app"

	// SectionAttrTitle is document section's title
	SectionAttrTitle = "title"

	// SectionAttrIcon is the id of document section's icon
	SectionAttrIcon = "icon"

	// SectionAttrSummary is the summary text of document section
	SectionAttrSummary = "summary"

	// SectionAttrPosition is the relative position of document section (for ordering purpose)
	SectionAttrPosition = "pos"

	// sectionAttr_Ubo is for internal use only!
	sectionAttr_Ubo = "_ubo"

	// SerKeyAttrs is a key used by Section.ToMap and Section.MarshalJSON/Section.UnmarshalJSON to store BO's custom attributes.
	SerKeyAttrs = "_attrs"

	// SerKeyFields is a key used by Section.ToMap and Section.MarshalJSON/Section.UnmarshalJSON to store BO's top-level custom fields.
	SerKeyFields = "_fields"
)

// Section is the business object.
//   - Section inherits unique id from bo.UniversalBo
type Section struct {
	*henge.UniversalBo `json:"_ubo"`
	appId              string `json:"app"`
	title              string `json:"title"`
	icon               string `json:"icon"`
	summary            string `json:"summary"`
	position           int    `json:"pos"`
}

// ToMap transforms user's attributes to a map.
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
func (s *Section) ToMap(postFunc henge.FuncPostUboToMap) map[string]interface{} {
	result := map[string]interface{}{
		henge.FieldId: s.GetId(),
		SerKeyFields: map[string]interface{}{
			SectionFieldAppId: s.appId,
		},
		SerKeyAttrs: map[string]interface{}{
			SectionAttrTitle:    s.title,
			SectionAttrIcon:     s.icon,
			SectionAttrSummary:  s.summary,
			SectionAttrPosition: s.position,
		},
	}
	if postFunc != nil {
		result = postFunc(result)
	}
	return result
}

// MarshalJSON implements json.encode.Marshaler.MarshalJSON
// TODO: lock for read?
func (s *Section) MarshalJSON() ([]byte, error) {
	s.sync()
	m := map[string]interface{}{
		sectionAttr_Ubo: s.UniversalBo.Clone(),
		SerKeyFields: map[string]interface{}{
			SectionFieldAppId: s.appId,
		},
		SerKeyAttrs: map[string]interface{}{
			SectionAttrTitle:    s.title,
			SectionAttrIcon:     s.icon,
			SectionAttrSummary:  s.summary,
			SectionAttrPosition: s.position,
		},
	}
	return json.Marshal(m)
}

// UnmarshalJSON implements json.decode.Unmarshaler.UnmarshalJSON
// TODO: lock for write?
func (s *Section) UnmarshalJSON(data []byte) error {
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	var err error
	if m[sectionAttr_Ubo] != nil {
		js, _ := json.Marshal(m[sectionAttr_Ubo])
		if err = json.Unmarshal(js, &s.UniversalBo); err != nil {
			return err
		}
	}
	if _fields, ok := m[SerKeyFields].(map[string]interface{}); ok {
		if s.appId, err = reddo.ToString(_fields[SectionFieldAppId]); err != nil {
			return err
		}
	}
	if _attrs, ok := m[SerKeyAttrs].(map[string]interface{}); ok {
		if s.title, err = reddo.ToString(_attrs[SectionAttrTitle]); err != nil {
			return err
		}
		if s.icon, err = reddo.ToString(_attrs[SectionAttrIcon]); err != nil {
			return err
		}
		if s.summary, err = reddo.ToString(_attrs[SectionAttrSummary]); err != nil {
			return err
		}
		if v, err := reddo.ToInt(_attrs[SectionAttrPosition]); err != nil {
			return err
		} else {
			s.position = int(v)
		}
	}
	s.sync()
	return nil
}

// GetAppId returns value of section's 'app-id' attribute.
func (s *Section) GetAppId() string {
	return s.appId
}

// SetAppId sets value of app's 'app-id' attribute.
func (s *Section) SetAppId(v string) *Section {
	s.appId = strings.ToLower(strings.TrimSpace(v))
	return s
}

// GetTitle returns value of section's 'title' attribute.
func (s *Section) GetTitle() string {
	return s.title
}

// SetTitle sets value of app's 'title' attribute.
func (s *Section) SetTitle(v string) *Section {
	s.title = strings.TrimSpace(v)
	return s
}

// GetIcon returns value of app's 'icon' attribute.
func (s *Section) GetIcon() string {
	return s.icon
}

// SetIcon sets value of app's 'icon' attribute.
func (s *Section) SetIcon(v string) *Section {
	s.icon = strings.TrimSpace(v)
	return s
}

// GetSummary returns value of app's 'summary' attribute.
func (s *Section) GetSummary() string {
	return s.summary
}

// SetSummary sets value of app's 'summary' attribute.
func (s *Section) SetSummary(v string) *Section {
	s.summary = strings.TrimSpace(v)
	return s
}

// GetPosition returns value of app's 'position' attribute.
func (s *Section) GetPosition() int {
	return s.position
}

// SetPosition sets value of app's 'position' attribute.
func (s *Section) SetPosition(v int) *Section {
	s.position = v
	return s
}

// sync is called to synchronize BO's attributes to its UniversalBo.
func (s *Section) sync() *Section {
	s.SetExtraAttr(SectionFieldAppId, s.appId)
	s.SetDataAttr(SectionAttrTitle, s.title)
	s.SetDataAttr(SectionAttrIcon, s.icon)
	s.SetDataAttr(SectionAttrSummary, s.summary)
	s.SetDataAttr(SectionAttrPosition, s.position)
	s.UniversalBo.Sync()
	return s
}
