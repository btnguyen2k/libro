package app

import (
	"github.com/btnguyen2k/godal"
	"github.com/btnguyen2k/henge"
)

const (
	// TableApp is name of the database table to store app records.
	TableApp = "libro_app"

	// custom db table fields
)

// AppDao defines API to access App storage
type AppDao interface {
	// Delete removes the specified business object from storage
	Delete(bo *App) (bool, error)

	// Create persists a new business object to storage
	Create(bo *App) (bool, error)

	// Get retrieves a business object from storage
	Get(id string) (*App, error)

	// GetN retrieves N business objects from storage
	GetN(fromOffset, maxNumRows int, filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*App, error)

	// GetAll retrieves all available business objects from storage
	GetAll(filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*App, error)

	// Update modifies an existing business object
	Update(bo *App) (bool, error)
}

// BaseAppDaoImpl is a generic implementation of AppDao
type BaseAppDaoImpl struct {
	henge.UniversalDao
}

// Delete implements AppDao.Delete
func (dao *BaseAppDaoImpl) Delete(bo *App) (bool, error) {
	if bo == nil {
		return false, nil
	}
	return dao.UniversalDao.Delete(bo.sync().UniversalBo)
}

// Create implements AppDao.Create
func (dao *BaseAppDaoImpl) Create(bo *App) (bool, error) {
	if bo == nil {
		return false, nil
	}
	return dao.UniversalDao.Create(bo.sync().UniversalBo)
}

// Get implements AppDao.Get
func (dao *BaseAppDaoImpl) Get(id string) (*App, error) {
	ubo, err := dao.UniversalDao.Get(id)
	if err != nil {
		return nil, err
	}
	return NewAppFromUbo(ubo), nil
}

// GetN implements AppDao.GetN
func (dao *BaseAppDaoImpl) GetN(fromOffset, maxNumRows int, filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*App, error) {
	uboList, err := dao.UniversalDao.GetN(fromOffset, maxNumRows, filter, sorting)
	if err != nil {
		return nil, err
	}
	result := make([]*App, 0)
	for _, ubo := range uboList {
		app := NewAppFromUbo(ubo)
		result = append(result, app)
	}
	return result, nil
}

// GetAll implements AppDao.GetAll
func (dao *BaseAppDaoImpl) GetAll(filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*App, error) {
	return dao.GetN(0, 0, filter, sorting)
}

// Update implements AppDao.Update
func (dao *BaseAppDaoImpl) Update(bo *App) (bool, error) {
	if bo == nil {
		return false, nil
	}
	return dao.UniversalDao.Update(bo.sync().UniversalBo)
}
