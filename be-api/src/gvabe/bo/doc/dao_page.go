package doc

import (
	"github.com/btnguyen2k/godal"
	"github.com/btnguyen2k/henge"
)

const (
	// TablePage is name of the database table to store document page records.
	TablePage = "libro_page"

	// PageColProductId is name of database column for document page's product-id.
	PageColProductId = "zpid"

	// PageColTopicId is name of database column for document page's topic-id.
	PageColTopicId = "ztid"
)

// PageDao defines API to access Page storage
type PageDao interface {
	// Delete removes the specified business object from storage
	Delete(bo *Page) (bool, error)

	// Create persists a new business object to storage
	Create(bo *Page) (bool, error)

	// Get retrieves a business object from storage
	Get(id string) (*Page, error)

	// GetN retrieves N business objects from storage
	GetN(topic *Topic, fromOffset, maxNumRows int, filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*Page, error)

	// GetAll retrieves all available business objects from storage
	GetAll(topic *Topic, filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*Page, error)

	// Update modifies an existing business object
	Update(bo *Page) (bool, error)
}

// BasePageDaoImpl is a generic implementation of PageDao
type BasePageDaoImpl struct {
	henge.UniversalDao
}

// Delete implements PageDao.Delete
func (dao *BasePageDaoImpl) Delete(bo *Page) (bool, error) {
	if bo == nil {
		return false, nil
	}
	return dao.UniversalDao.Delete(bo.sync().UniversalBo)
}

// Create implements PageDao.Create
func (dao *BasePageDaoImpl) Create(bo *Page) (bool, error) {
	if bo == nil {
		return false, nil
	}
	return dao.UniversalDao.Create(bo.sync().UniversalBo)
}

// Get implements PageDao.Get
func (dao *BasePageDaoImpl) Get(id string) (*Page, error) {
	ubo, err := dao.UniversalDao.Get(id)
	if err != nil {
		return nil, err
	}
	return NewPageFromUbo(ubo), nil
}

// GetN implements PageDao.GetN
func (dao *BasePageDaoImpl) GetN(topic *Topic, fromOffset, maxNumRows int, filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*Page, error) {
	if topic == nil {
		return make([]*Page, 0), nil
	}
	if sorting == nil {
		sorting = (&godal.SortingField{FieldName: PageFieldPosition}).ToSortingOpt()
	}
	myFilter := (&godal.FilterOptAnd{}).Add(filter).
		Add(godal.FilterOptFieldOpValue{FieldName: PageFieldTopicId, Operator: godal.FilterOpEqual, Value: topic.GetId()})
	uboList, err := dao.UniversalDao.GetN(fromOffset, maxNumRows, myFilter, sorting)
	if err != nil {
		return nil, err
	}
	result := make([]*Page, 0)
	for _, ubo := range uboList {
		bo := NewPageFromUbo(ubo)
		result = append(result, bo)
	}
	return result, nil
}

// GetAll implements PageDao.GetAll
func (dao *BasePageDaoImpl) GetAll(topic *Topic, filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*Page, error) {
	return dao.GetN(topic, 0, 0, filter, sorting)
}

// Update implements PageDao.Update
func (dao *BasePageDaoImpl) Update(bo *Page) (bool, error) {
	if bo == nil {
		return false, nil
	}
	return dao.UniversalDao.Update(bo.sync().UniversalBo)
}
