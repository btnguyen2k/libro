package product

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/btnguyen2k/consu/reddo"
	"github.com/btnguyen2k/henge"
	"main/src/gvabe/bo"
)

// NewProduct is helper function to create new Product bo.
func NewProduct(tagVersion uint64, id, name, desc string, isPublished bool) *Product {
	bo := &Product{
		UniversalBo: henge.NewUniversalBo(id, tagVersion),
	}
	return bo.
		SetName(name).
		SetDescription(desc).
		SetPublished(isPublished).
		SetNumTopics(0).
		sync()
}

var TypContactsMap = reflect.TypeOf(map[string]string{})

// NewProductFromUbo is helper function to create Product bo from a universal bo.
func NewProductFromUbo(ubo *henge.UniversalBo) *Product {
	if ubo == nil {
		return nil
	}
	ubo = ubo.Clone()
	bo := &Product{UniversalBo: ubo}
	if v, err := ubo.GetDataAttrAs(ProdAttrName, reddo.TypeString); err != nil {
		return nil
	} else if v != nil {
		bo.name, _ = v.(string)
	}
	if v, err := ubo.GetDataAttrAs(ProdAttrDesc, reddo.TypeString); err != nil {
		return nil
	} else if v != nil {
		bo.description, _ = v.(string)
	}
	if v, err := ubo.GetDataAttrAs(ProdAttrIsPublished, reddo.TypeBool); err != nil {
		return nil
	} else if v != nil {
		bo.isPublished, _ = v.(bool)
	}
	if v, err := ubo.GetDataAttrAs(ProdAttrNumTopics, reddo.TypeInt); err != nil {
		return nil
	} else if temp, ok := v.(int64); ok {
		bo.numTopics = int(temp)
	}
	if v, err := ubo.GetDataAttrAs(ProdAttrContacts, TypContactsMap); err != nil {
		return nil
	} else if v != nil {
		bo.contacts = v.(map[string]string)
	}
	return bo.sync()
}

const (
	// ProdAttrName is product's name, used for displaying purpose.
	ProdAttrName = "name"

	// ProdAttrDesc is product's description text.
	ProdAttrDesc = "desc"

	// ProdAttrIsPublished is a flag to mark if product is enabled/published.
	ProdAttrIsPublished = "ispub"

	// ProdAttrNumTopics is the number of document topics belong to this product.
	ProdAttrNumTopics = "ntopics"

	// ProdAttrContacts is collection of product's contacts (such as GitHub, Website, LinkedIn, etc).
	ProdAttrContacts = "contacts"

	// prodAttrUbo is for internal use only!
	prodAttrUbo = "_ubo"
)

// Product is the business object.
//   - Product inherits unique id from bo.UniversalBo
//   - Product must be unique
type Product struct {
	*henge.UniversalBo `json:"_ubo"`
	name               string            `json:"name"`
	description        string            `json:"desc"`
	isPublished        bool              `json:"ispub"`
	numTopics          int               `json:"ntopics"`
	contacts           map[string]string `json:"contacts"`
}

// ToMap transforms user's attributes to a map.
//
// The function returns a map with the following structure:
//   {
//     henge.FieldId: p.GetId(),
//     SerKeyFields: map[string]interface{}{
//         // all BO's top-level custom fields go here
//		},
//     SerKeyAttrs: map[string]interface{}{
//         // all BO's custom attributes go here
//     },
//   }
func (p *Product) ToMap(postFunc henge.FuncPostUboToMap) map[string]interface{} {
	result := map[string]interface{}{
		henge.FieldId: p.GetId(),
		bo.SerKeyAttrs: map[string]interface{}{
			ProdAttrName:        p.GetName(),
			ProdAttrDesc:        p.GetDescription(),
			ProdAttrIsPublished: p.IsPublished(),
			ProdAttrNumTopics:   p.GetNumTopics(),
			ProdAttrContacts:    p.GetContacts(),
		},
	}
	if postFunc != nil {
		result = postFunc(result)
	}
	return result
}

// MarshalJSON implements json.encode.Marshaler.MarshalJSON
// TODO: lock for read?
func (p *Product) MarshalJSON() ([]byte, error) {
	p.sync()
	m := map[string]interface{}{
		prodAttrUbo: p.UniversalBo.Clone(),
		bo.SerKeyAttrs: map[string]interface{}{
			ProdAttrName:        p.GetName(),
			ProdAttrDesc:        p.GetDescription(),
			ProdAttrIsPublished: p.IsPublished(),
			ProdAttrNumTopics:   p.GetNumTopics(),
			ProdAttrContacts:    p.GetContacts(),
		},
	}
	return json.Marshal(m)
}

// UnmarshalJSON implements json.decode.Unmarshaler.UnmarshalJSON
// TODO: lock for write?
func (p *Product) UnmarshalJSON(data []byte) error {
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	var err error
	if m[prodAttrUbo] != nil {
		js, _ := json.Marshal(m[prodAttrUbo])
		if err = json.Unmarshal(js, &p.UniversalBo); err != nil {
			return err
		}
	}
	if _attrs, ok := m[bo.SerKeyAttrs].(map[string]interface{}); ok {
		if p.name, err = reddo.ToString(_attrs[ProdAttrName]); err != nil {
			return err
		}
		if p.description, err = reddo.ToString(_attrs[ProdAttrDesc]); err != nil {
			return err
		}
		if p.isPublished, err = reddo.ToBool(_attrs[ProdAttrIsPublished]); err != nil {
			return err
		}
		if v, err := reddo.ToInt(_attrs[ProdAttrNumTopics]); err != nil {
			return err
		} else {
			p.numTopics = int(v)
		}
		if v, err := reddo.ToMap(_attrs[ProdAttrContacts], TypContactsMap); err != nil {
			return err
		} else {
			p.contacts = v.(map[string]string)
		}
	}
	p.sync()
	return nil
}

// GetName returns value of product's 'name' attribute.
func (p *Product) GetName() string {
	return p.name
}

// SetName sets value of product's 'name' attribute.
func (p *Product) SetName(v string) *Product {
	p.name = strings.TrimSpace(v)
	return p
}

// GetDescription returns value of product's 'desc' attribute.
func (p *Product) GetDescription() string {
	return p.description
}

// SetDescription sets value of product's 'desc' attribute.
func (p *Product) SetDescription(v string) *Product {
	p.description = strings.TrimSpace(v)
	return p
}

// IsPublished returns value of product's 'is-published' attribute.
func (p *Product) IsPublished() bool {
	return p.isPublished
}

// SetPublished sets value of product's 'is-published' attribute.
func (p *Product) SetPublished(v bool) *Product {
	p.isPublished = v
	return p
}

// GetNumTopics returns value of product's 'num-topics' attribute.
func (p *Product) GetNumTopics() int {
	return p.numTopics
}

// SetNumTopics sets value of product's 'num-topic' attribute.
func (p *Product) SetNumTopics(v int) *Product {
	p.numTopics = v
	return p
}

// GetContacts returns value of product's 'contacts' attribute.
func (p *Product) GetContacts() map[string]string {
	result := make(map[string]string)
	for k, v := range p.contacts {
		result[k] = v
	}
	return result
}

// SetContacts sets value of product's 'contacts' attribute.
func (p *Product) SetContacts(contacts map[string]string) *Product {
	m := make(map[string]string)
	for k, v := range contacts {
		m[k] = v
	}
	p.contacts = m
	return p
}

// AddContact adds a new contact (type and value) to product's 'contacts' attribute.
func (p *Product) AddContact(typ, val string) *Product {
	if p.contacts == nil {
		p.contacts = make(map[string]string)
	}
	p.contacts[typ] = val
	return p
}

// sync is called to synchronize BO's attributes to its UniversalBo.
func (p *Product) sync() *Product {
	p.SetDataAttr(ProdAttrName, p.name)
	p.SetDataAttr(ProdAttrDesc, p.description)
	p.SetDataAttr(ProdAttrIsPublished, p.isPublished)
	p.SetDataAttr(ProdAttrNumTopics, p.numTopics)
	p.SetDataAttr(ProdAttrContacts, p.contacts)
	p.UniversalBo.Sync()
	return p
}
