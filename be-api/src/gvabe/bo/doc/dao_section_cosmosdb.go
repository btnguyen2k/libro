package doc

import (
	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

// NewSectionDaoCosmosdb is helper method to create Azure Cosmos DB-implementation of SectionDao
//
// Note: txModeOnWrite is not currently used!
func NewSectionDaoCosmosdb(sqlc *prom.SqlConnect, tableName string, txModeOnWrite bool) SectionDao {
	dao := &BaseSectionDaoImpl{}
	spec := &henge.CosmosdbDaoSpec{
		PkName:        henge.CosmosdbColId,
		TxModeOnWrite: txModeOnWrite,
	}
	dao.UniversalDao = henge.NewUniversalDaoCosmosdbSql(sqlc, tableName, spec)
	return dao
}
