package doc

import (
	"github.com/btnguyen2k/godal"
	"github.com/btnguyen2k/henge"
	"main/src/gvabe/bo/app"
)

const (
	// TableSection is name of the database table to store document section records.
	TableSection = "libro_section"

	// SectionColAppId is name of database column for document section's app-id.
	SectionColAppId = "zaid"
)

// SectionDao defines API to access Section storage
type SectionDao interface {
	// Delete removes the specified business object from storage
	Delete(bo *Section) (bool, error)

	// Create persists a new business object to storage
	Create(bo *Section) (bool, error)

	// Get retrieves a business object from storage
	Get(id string) (*Section, error)

	// GetN retrieves N business objects from storage
	GetN(app *app.App, fromOffset, maxNumRows int, filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*Section, error)

	// GetAll retrieves all available business objects from storage
	GetAll(app *app.App, filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*Section, error)

	// Update modifies an existing business object
	Update(bo *Section) (bool, error)
}

// BaseSectionDaoImpl is a generic implementation of SectionDao
type BaseSectionDaoImpl struct {
	henge.UniversalDao
}

// Delete implements SectionDao.Delete
func (dao *BaseSectionDaoImpl) Delete(bo *Section) (bool, error) {
	if bo == nil {
		return false, nil
	}
	return dao.UniversalDao.Delete(bo.sync().UniversalBo)
}

// Create implements SectionDao.Create
func (dao *BaseSectionDaoImpl) Create(bo *Section) (bool, error) {
	if bo == nil {
		return false, nil
	}
	return dao.UniversalDao.Create(bo.sync().UniversalBo)
}

// Get implements SectionDao.Get
func (dao *BaseSectionDaoImpl) Get(id string) (*Section, error) {
	ubo, err := dao.UniversalDao.Get(id)
	if err != nil {
		return nil, err
	}
	return NewSectionFromUbo(ubo), nil
}

// GetN implements SectionDao.GetN
func (dao *BaseSectionDaoImpl) GetN(app *app.App, fromOffset, maxNumRows int, filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*Section, error) {
	if app == nil {
		return make([]*Section, 0), nil
	}
	myFilter := (&godal.FilterOptAnd{}).Add(filter).
		Add(godal.FilterOptFieldOpValue{FieldName: SectionFieldAppId, Operator: godal.FilterOpEqual, Value: app.GetId()})
	uboList, err := dao.UniversalDao.GetN(fromOffset, maxNumRows, myFilter, sorting)
	if err != nil {
		return nil, err
	}
	result := make([]*Section, 0)
	for _, ubo := range uboList {
		bo := NewSectionFromUbo(ubo)
		result = append(result, bo)
	}
	return result, nil
}

// GetAll implements SectionDao.GetAll
func (dao *BaseSectionDaoImpl) GetAll(app *app.App, filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*Section, error) {
	return dao.GetN(app, 0, 0, filter, sorting)
}

// Update implements SectionDao.Update
func (dao *BaseSectionDaoImpl) Update(bo *Section) (bool, error) {
	if bo == nil {
		return false, nil
	}
	return dao.UniversalDao.Update(bo.sync().UniversalBo)
}
