package respicite

import (
	"github.com/btnguyen2k/godal"
	"github.com/btnguyen2k/godal/mongo"
	"github.com/btnguyen2k/prom"
	"go.mongodb.org/mongo-driver/bson"
)

// M2oDaoMongoInitCollection is helper function to initialize MongoDB collection to store many-to-many mappings.
func M2oDaoMongoInitCollection(mc *prom.MongoConnect, collectionName string) error {
	_, err := mc.CreateCollection(collectionName)
	if err != nil {
		return err
	}
	indexes := []interface{}{
		map[string]interface{}{
			"key":    bson.D{{MappingFieldDest, 1}},
			"name":   "idx_dest",
			"unique": false,
		},
	}
	_, err = mc.CreateCollectionIndexes(collectionName, indexes)
	return err
}

// NewM2oDaoMongo is helper function to construct a new M2oDaoMongo instance.
func NewM2oDaoMongo(mc *prom.MongoConnect, collectionName string, txMode bool) *M2oDaoMongo {
	dao := &M2oDaoMongo{}
	gdao := mongo.NewGenericDaoMongo(mc, godal.NewAbstractGenericDao(dao)).SetTxModeOnWrite(txMode)
	gdao.SetRowMapper(&myM2oDaoMongoRowMapper{&mongo.GenericRowMapperMongo{}})
	dao.BaseM2oDao = &BaseM2oDao{
		IGenericDao: gdao,
		storageId:   collectionName,
	}
	dao.mc = mc
	return dao
}

type myM2oDaoMongoRowMapper struct {
	*mongo.GenericRowMapperMongo
}

// ToRow implements godal.IRowMapper.ToRow.
func (m *myM2oDaoMongoRowMapper) ToRow(_ string, bo godal.IGenericBo) (interface{}, error) {
	if bo == nil {
		return nil, nil
	}
	result := make(map[string]interface{})
	err := bo.GboTransferViaJson(&result)
	result["_id"] = result[MappingFieldSrc]
	return result, err
}

// M2oDaoMongo is MongoDB-implementation of M2oDao
type M2oDaoMongo struct {
	*BaseM2oDao
	mc *prom.MongoConnect
}

// Destroy is called to close the underlying prom.MongoConnect
func (dao *M2oDaoMongo) Destroy() {
	dao.mc.Close(nil)
}

// GdaoCreateFilter implements godal.IGenericDao.GdaoCreateFilter.
func (dao *M2oDaoMongo) GdaoCreateFilter(collectionName string, bo godal.IGenericBo) godal.FilterOpt {
	if collectionName == dao.storageId {
		/*
			"Universal" filter: filter on "Src", or "Dest" or "Src" AND "Dest
		*/

		vSrc, _ := bo.GboGetAttr(MappingFieldSrc, nil)
		vDest, _ := bo.GboGetAttr(MappingFieldDest, nil)
		f1 := &godal.FilterOptFieldOpValue{FieldName: MappingFieldSrc, Operator: godal.FilterOpEqual, Value: vSrc}
		f2 := &godal.FilterOptFieldOpValue{FieldName: MappingFieldDest, Operator: godal.FilterOpEqual, Value: vDest}

		if vSrc != nil && vDest != nil {
			return (&godal.FilterOptAnd{}).Add(f1).Add(f2)
		}

		if vSrc != nil {
			return f1
		}

		if vDest != nil {
			return f2
		}
	}
	return nil
}
