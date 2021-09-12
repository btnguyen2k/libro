package doc

import (
	"github.com/btnguyen2k/godal"
	"github.com/btnguyen2k/henge"
	"main/src/gvabe/bo/product"
)

const (
	// TableTopic is name of the database table to store document topic records.
	TableTopic = "libro_topic"

	// TopicColProductId is name of database column for document topic's product-id.
	TopicColProductId = "zpid"
)

// TopicDao defines API to access Topic storage
type TopicDao interface {
	// Delete removes the specified business object from storage
	Delete(bo *Topic) (bool, error)

	// Create persists a new business object to storage
	Create(bo *Topic) (bool, error)

	// Get retrieves a business object from storage
	Get(id string) (*Topic, error)

	// GetN retrieves N business objects from storage
	GetN(prod *product.Product, fromOffset, maxNumRows int, filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*Topic, error)

	// GetAll retrieves all available business objects from storage
	GetAll(prod *product.Product, filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*Topic, error)

	// Update modifies an existing business object
	Update(bo *Topic) (bool, error)
}

// BaseTopicDaoImpl is a generic implementation of TopicDao
type BaseTopicDaoImpl struct {
	henge.UniversalDao
}

// Delete implements TopicDao.Delete
func (dao *BaseTopicDaoImpl) Delete(bo *Topic) (bool, error) {
	if bo == nil {
		return false, nil
	}
	return dao.UniversalDao.Delete(bo.sync().UniversalBo)
}

// Create implements TopicDao.Create
func (dao *BaseTopicDaoImpl) Create(bo *Topic) (bool, error) {
	if bo == nil {
		return false, nil
	}
	return dao.UniversalDao.Create(bo.sync().UniversalBo)
}

// Get implements TopicDao.Get
func (dao *BaseTopicDaoImpl) Get(id string) (*Topic, error) {
	ubo, err := dao.UniversalDao.Get(id)
	if err != nil {
		return nil, err
	}
	return NewTopicFromUbo(ubo), nil
}

// GetN implements TopicDao.GetN
func (dao *BaseTopicDaoImpl) GetN(prod *product.Product, fromOffset, maxNumRows int, filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*Topic, error) {
	if prod == nil {
		return make([]*Topic, 0), nil
	}
	myFilter := (&godal.FilterOptAnd{}).Add(filter).
		Add(godal.FilterOptFieldOpValue{FieldName: TopicFieldProductId, Operator: godal.FilterOpEqual, Value: prod.GetId()})
	uboList, err := dao.UniversalDao.GetN(fromOffset, maxNumRows, myFilter, sorting)
	if err != nil {
		return nil, err
	}
	result := make([]*Topic, 0)
	for _, ubo := range uboList {
		bo := NewTopicFromUbo(ubo)
		result = append(result, bo)
	}
	return result, nil
}

// GetAll implements TopicDao.GetAll
func (dao *BaseTopicDaoImpl) GetAll(prod *product.Product, filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*Topic, error) {
	return dao.GetN(prod, 0, 0, filter, sorting)
}

// Update implements TopicDao.Update
func (dao *BaseTopicDaoImpl) Update(bo *Topic) (bool, error) {
	if bo == nil {
		return false, nil
	}
	return dao.UniversalDao.Update(bo.sync().UniversalBo)
}
