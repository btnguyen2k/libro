package app

import (
	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

// NewAppDaoMongo is helper method to create MongoDB-implementation of AppDao
func NewAppDaoMongo(mc *prom.MongoConnect, collectionName string, txModeOnWrite bool) AppDao {
	dao := &BaseAppDaoImpl{}
	dao.UniversalDao = henge.NewUniversalDaoMongo(mc, collectionName, txModeOnWrite)
	return dao
}
