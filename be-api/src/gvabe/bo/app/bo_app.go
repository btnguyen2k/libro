package app

import (
	"encoding/json"
	"strings"

	"github.com/btnguyen2k/consu/reddo"
	"github.com/btnguyen2k/henge"
)

// NewApp is helper function to create new App bo.
func NewApp(appVersion uint64, id, name, desc string, isPublished bool) *App {
	bo := &App{
		UniversalBo: henge.NewUniversalBo(id, appVersion),
	}
	return bo.SetName(name).SetDescription(desc).SetPublished(isPublished).sync()
}

// NewAppFromUbo is helper function to create App bo from a universal bo.
func NewAppFromUbo(ubo *henge.UniversalBo) *App {
	if ubo == nil {
		return nil
	}
	ubo = ubo.Clone()
	bo := &App{UniversalBo: ubo}
	if v, err := ubo.GetDataAttrAs(AppAttrName, reddo.TypeString); err != nil {
		return nil
	} else {
		bo.name, _ = v.(string)
	}
	if v, err := ubo.GetDataAttrAs(AppAttrDesc, reddo.TypeString); err != nil {
		return nil
	} else {
		bo.description, _ = v.(string)
	}
	if v, err := ubo.GetDataAttrAs(AppAttrIsPublished, reddo.TypeBool); err != nil {
		return nil
	} else {
		bo.isPublished, _ = v.(bool)
	}
	return bo.sync()
}

const (
	// AppAttrName is app's name, used for displaying purpose.
	AppAttrName = "name"

	// AppAttrDesc is app's description text.
	AppAttrDesc = "desc"

	// AppAttrIsPublished is a flag to mark if app is enabled/published.
	AppAttrIsPublished = "ispub"

	// appAttr_Ubo is for internal use only!
	appAttr_Ubo = "_ubo"

	// SerKeyAttrs is a key used by App.ToMap and App.MarshalJSON/App.UnmarshalJSON to store BO's custom attributes.
	SerKeyAttrs = "_attrs"

	// SerKeyFields is a key used by App.ToMap and App.MarshalJSON/App.UnmarshalJSON to store BO's top-level custom fields.
	SerKeyFields = "_fields"
)

// App is the business object.
//   - App inherits unique id from bo.UniversalBo
type App struct {
	*henge.UniversalBo `json:"_ubo"`
	name               string `json:"name"`
	description        string `json:"desc"`
	isPublished        bool   `json:"ispub"`
}

// ToMap transforms user's attributes to a map.
//
// The function returns a map with the following structure:
//   {
//     henge.FieldId: a.GetId(),
//     SerKeyFields: map[string]interface{}{
//         // all BO's top-level custom fields go here
//		},
//     SerKeyAttrs: map[string]interface{}{
//         // all BO's custom attributes go here
//     },
//   }
func (a *App) ToMap(postFunc henge.FuncPostUboToMap) map[string]interface{} {
	result := map[string]interface{}{
		henge.FieldId: a.GetId(),
		SerKeyAttrs: map[string]interface{}{
			AppAttrName:        a.name,
			AppAttrDesc:        a.description,
			AppAttrIsPublished: a.isPublished,
		},
	}
	if postFunc != nil {
		result = postFunc(result)
	}
	return result
}

// MarshalJSON implements json.encode.Marshaler.MarshalJSON
// TODO: lock for read?
func (a *App) MarshalJSON() ([]byte, error) {
	a.sync()
	m := map[string]interface{}{
		appAttr_Ubo: a.UniversalBo.Clone(),
		SerKeyAttrs: map[string]interface{}{
			AppAttrName:        a.name,
			AppAttrDesc:        a.description,
			AppAttrIsPublished: a.isPublished,
		},
	}
	return json.Marshal(m)
}

// UnmarshalJSON implements json.decode.Unmarshaler.UnmarshalJSON
// TODO: lock for write?
func (a *App) UnmarshalJSON(data []byte) error {
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	var err error
	if m[appAttr_Ubo] != nil {
		js, _ := json.Marshal(m[appAttr_Ubo])
		if err = json.Unmarshal(js, &a.UniversalBo); err != nil {
			return err
		}
	}
	if _attrs, ok := m[SerKeyAttrs].(map[string]interface{}); ok {
		if a.name, err = reddo.ToString(_attrs[AppAttrName]); err != nil {
			return err
		}
		if a.description, err = reddo.ToString(_attrs[AppAttrDesc]); err != nil {
			return err
		}
		if a.isPublished, err = reddo.ToBool(_attrs[AppAttrIsPublished]); err != nil {
			return err
		}
	}
	a.sync()
	return nil
}

// GetName returns value of app's 'name' attribute.
func (a *App) GetName() string {
	return a.name
}

// SetName sets value of app's 'name' attribute.
func (a *App) SetName(v string) *App {
	a.name = strings.TrimSpace(v)
	return a
}

// GetDescription returns value of app's 'desc' attribute.
func (a *App) GetDescription() string {
	return a.description
}

// SetDescription sets value of app's 'desc' attribute.
func (a *App) SetDescription(v string) *App {
	a.description = strings.TrimSpace(v)
	return a
}

// IsPublished returns value of app's 'is-published' attribute.
func (a *App) IsPublished() bool {
	return a.isPublished
}

// SetPublished sets value of app's 'is-published' attribute.
func (a *App) SetPublished(v bool) *App {
	a.isPublished = v
	return a
}

// sync is called to synchronize BO's attributes to its UniversalBo.
func (a *App) sync() *App {
	a.SetDataAttr(AppAttrName, a.name)
	a.SetDataAttr(AppAttrDesc, a.description)
	a.SetDataAttr(AppAttrIsPublished, a.isPublished)
	a.UniversalBo.Sync()
	return a
}
