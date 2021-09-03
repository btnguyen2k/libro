package doc

import (
	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

// NewPageDaoMongo is helper method to create MongoDB-implementation of PageDao
func NewPageDaoMongo(mc *prom.MongoConnect, collectionName string, txModeOnWrite bool) PageDao {
	dao := &BasePageDaoImpl{}
	dao.UniversalDao = henge.NewUniversalDaoMongo(mc, collectionName, txModeOnWrite)
	return dao
}
