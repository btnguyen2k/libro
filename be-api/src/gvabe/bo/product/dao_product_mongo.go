package product

import (
	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

// NewProductDaoMongo is helper method to create MongoDB-implementation of ProductDao
func NewProductDaoMongo(mc *prom.MongoConnect, collectionName string, txModeOnWrite bool) ProductDao {
	dao := &BaseProductDaoImpl{}
	dao.UniversalDao = henge.NewUniversalDaoMongo(mc, collectionName, txModeOnWrite)
	return dao
}
