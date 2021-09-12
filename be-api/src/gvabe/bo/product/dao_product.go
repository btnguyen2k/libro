package product

import (
	"github.com/btnguyen2k/godal"
	"github.com/btnguyen2k/henge"
)

const (
	// TableProduct is name of the database table to store products.
	TableProduct = "libro_product"

	// custom db table fields
)

// ProductDao defines API to access Product storage
type ProductDao interface {
	// Delete removes the specified business object from storage
	Delete(bo *Product) (bool, error)

	// Create persists a new business object to storage
	Create(bo *Product) (bool, error)

	// Get retrieves a business object from storage
	Get(id string) (*Product, error)

	// GetN retrieves N business objects from storage
	GetN(fromOffset, maxNumRows int, filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*Product, error)

	// GetAll retrieves all available business objects from storage
	GetAll(filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*Product, error)

	// Update modifies an existing business object
	Update(bo *Product) (bool, error)
}

// BaseProductDaoImpl is a generic implementation of ProductDao
type BaseProductDaoImpl struct {
	henge.UniversalDao
}

// Delete implements ProductDao.Delete
func (dao *BaseProductDaoImpl) Delete(bo *Product) (bool, error) {
	if bo == nil {
		return false, nil
	}
	return dao.UniversalDao.Delete(bo.sync().UniversalBo)
}

// Create implements ProductDao.Create
func (dao *BaseProductDaoImpl) Create(bo *Product) (bool, error) {
	if bo == nil {
		return false, nil
	}
	return dao.UniversalDao.Create(bo.sync().UniversalBo)
}

// Get implements ProductDao.Get
func (dao *BaseProductDaoImpl) Get(id string) (*Product, error) {
	ubo, err := dao.UniversalDao.Get(id)
	if err != nil {
		return nil, err
	}
	return NewProductFromUbo(ubo), nil
}

// GetN implements ProductDao.GetN
func (dao *BaseProductDaoImpl) GetN(fromOffset, maxNumRows int, filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*Product, error) {
	uboList, err := dao.UniversalDao.GetN(fromOffset, maxNumRows, filter, sorting)
	if err != nil {
		return nil, err
	}
	result := make([]*Product, 0)
	for _, ubo := range uboList {
		bo := NewProductFromUbo(ubo)
		result = append(result, bo)
	}
	return result, nil
}

// GetAll implements ProductDao.GetAll
func (dao *BaseProductDaoImpl) GetAll(filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*Product, error) {
	return dao.GetN(0, 0, filter, sorting)
}

// Update implements ProductDao.Update
func (dao *BaseProductDaoImpl) Update(bo *Product) (bool, error) {
	if bo == nil {
		return false, nil
	}
	return dao.UniversalDao.Update(bo.sync().UniversalBo)
}
