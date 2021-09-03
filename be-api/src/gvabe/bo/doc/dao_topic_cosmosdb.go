package doc

import (
	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

// NewTopicDaoCosmosdb is helper method to create Azure Cosmos DB-implementation of TopicDao
//
// Note: txModeOnWrite is not currently used!
func NewTopicDaoCosmosdb(sqlc *prom.SqlConnect, tableName string, txModeOnWrite bool) TopicDao {
	dao := &BaseTopicDaoImpl{}
	spec := &henge.CosmosdbDaoSpec{
		PkName:        henge.CosmosdbColId,
		TxModeOnWrite: txModeOnWrite,
	}
	dao.UniversalDao = henge.NewUniversalDaoCosmosdbSql(sqlc, tableName, spec)
	return dao
}
