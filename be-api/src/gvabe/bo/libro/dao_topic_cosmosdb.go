package libro

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
		// PkName:        TopicFieldProductId,
		PkName:        henge.CosmosdbColId,
		TxModeOnWrite: txModeOnWrite,
	}
	dao.UniversalDao = henge.NewUniversalDaoCosmosdbSql(sqlc, tableName, spec)
	return dao
}

/* There is no function InitTopicTableCosmosdb, use InitTopicTableSql instead. */
