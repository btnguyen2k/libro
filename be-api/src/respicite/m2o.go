package respicite

import (
	"time"

	"github.com/btnguyen2k/consu/reddo"
	"github.com/btnguyen2k/godal"
)

// M2oDao defines many-to-one mapping API
type M2oDao interface {
	// Get fetches a mapping {Src->Dest} from storage.
	// If the mapping does not exist, this function returns (nil, ErrNotFound).
	Get(src string) (*Mapping, error)

	// Rget fetches all mappings {Src->Dest} from storage.
	Rget(dest string) ([]*Mapping, error)

	// Set creates a mapping {Src->Dest}.
	// If the mapping already existed, this function returns (true, ErrDuplicated).
	Set(src string, dest string) (bool, error)

	// Remove deletes a mapping {Src->Dest}.
	Remove(src string, dest string) (bool, error)
}

// BaseM2oDao is a generic implementation of M2oDao
type BaseM2oDao struct {
	godal.IGenericDao
	storageId string
}

// ToMapping converts godal.IGenericBo to Mapping instance.
func (dao *BaseM2oDao) ToMapping(gbo godal.IGenericBo) *Mapping {
	if gbo == nil {
		return nil
	}
	result := Mapping{}
	if src, err := gbo.GboGetAttr(MappingFieldSrc, reddo.TypeString); err == nil && src != nil {
		result.Src = src.(string)
	}
	if dest, err := gbo.GboGetAttr(MappingFieldDest, reddo.TypeString); err == nil && dest != nil {
		result.Dest = dest.(string)
		// t := Dest.(string)
		// result.Dest = &t
	}
	if t, err := gbo.GboGetTimeWithLayout(MappingFieldCreatedOn, time.RFC3339); err == nil {
		result.CreatedOn = t
	}
	return &result
}

// ToGbo converts a Mapping instance to godal.IGenericBo.
func (dao *BaseM2oDao) ToGbo(m *Mapping) godal.IGenericBo {
	gbo := godal.NewGenericBo()
	gbo.GboSetAttr(MappingFieldSrc, m.Src)
	gbo.GboSetAttr(MappingFieldDest, m.Dest)
	gbo.GboSetAttr(MappingFieldCreatedOn, m.CreatedOn)
	return gbo
}

// Get implements M2oDao.Get
func (dao *BaseM2oDao) Get(src string) (*Mapping, error) {
	filterBo := godal.NewGenericBo()
	filterBo.GboSetAttr(MappingFieldSrc, src)
	gbo, err := dao.GdaoFetchOne(dao.storageId, dao.GdaoCreateFilter(dao.storageId, filterBo))
	if err != nil {
		return nil, err
	}
	mapping, err := dao.ToMapping(gbo), nil
	if mapping == nil {
		err = ErrNotFound
	}
	return mapping, err
}

// Rget implements M2oDao.Rget
func (dao *BaseM2oDao) Rget(dest string) ([]*Mapping, error) {
	filterBo := godal.NewGenericBo()
	filterBo.GboSetAttr(MappingFieldDest, dest)
	gboList, err := dao.GdaoFetchMany(dao.storageId, dao.GdaoCreateFilter(dao.storageId, filterBo),
		(&godal.SortingField{FieldName: MappingFieldSrc}).ToSortingOpt(), 0, 0)
	result := make([]*Mapping, 0)
	if err != nil {
		return result, err
	}
	for _, gbo := range gboList {
		result = append(result, dao.ToMapping(gbo))
	}
	return result, nil
}

// Set implements M2oDao.Set
func (dao *BaseM2oDao) Set(src string, dest string) (bool, error) {
	mapping := &Mapping{Src: src, Dest: dest, CreatedOn: time.Now()}
	affectedRows, err := dao.GdaoCreate(dao.storageId, dao.ToGbo(mapping))
	if err == godal.ErrGdaoDuplicatedEntry {
		return true, ErrDuplicated
	}
	return affectedRows > 0, err
}

// Remove implements M2oDao.Remove
func (dao *BaseM2oDao) Remove(src string, dest string) (bool, error) {
	mapping := &Mapping{Src: src, Dest: dest}
	affectedRows, err := dao.GdaoDelete(dao.storageId, dao.ToGbo(mapping))
	return affectedRows > 0, err
}
