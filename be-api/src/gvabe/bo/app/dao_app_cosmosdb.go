package app

import (
	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

// NewAppDaoCosmosdb is helper method to create Azure Cosmos DB-implementation of AppDao
//
// Note: txModeOnWrite is not currently used!
func NewAppDaoCosmosdb(sqlc *prom.SqlConnect, tableName string, txModeOnWrite bool) AppDao {
	dao := &BaseAppDaoImpl{}
	spec := &henge.CosmosdbDaoSpec{
		PkName:        henge.CosmosdbColId,
		TxModeOnWrite: txModeOnWrite,
	}
	dao.UniversalDao = henge.NewUniversalDaoCosmosdbSql(sqlc, tableName, spec)
	return dao
}
