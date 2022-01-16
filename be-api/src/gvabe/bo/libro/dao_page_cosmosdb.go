package libro

import (
	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

// NewPageDaoCosmosdb is helper method to create Azure Cosmos DB-implementation of PageDao
//
// Note: txModeOnWrite is not currently used!
func NewPageDaoCosmosdb(sqlc *prom.SqlConnect, tableName string, txModeOnWrite bool) PageDao {
	dao := &BasePageDaoImpl{}
	spec := &henge.CosmosdbDaoSpec{
		// PkName:        PageColProductId,
		PkName:        henge.CosmosdbColId,
		TxModeOnWrite: txModeOnWrite,
	}
	dao.UniversalDao = henge.NewUniversalDaoCosmosdbSql(sqlc, tableName, spec)
	return dao
}

/* There is no function CreateCosmosdbTableForPages, use InitPageTableSql instead. */
