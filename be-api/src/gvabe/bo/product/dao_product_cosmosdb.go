package product

import (
	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

// NewProductDaoCosmosdb is helper method to create Azure Cosmos DB-implementation of ProductDao
//
// Note: txModeOnWrite is not currently used!
func NewProductDaoCosmosdb(sqlc *prom.SqlConnect, tableName string, txModeOnWrite bool) ProductDao {
	dao := &BaseProductDaoImpl{}
	spec := &henge.CosmosdbDaoSpec{
		PkName:        henge.CosmosdbColId,
		TxModeOnWrite: txModeOnWrite,
	}
	dao.UniversalDao = henge.NewUniversalDaoCosmosdbSql(sqlc, tableName, spec)
	return dao
}
