package libro

import (
	"fmt"
	"log"

	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

// NewTopicDaoSql is helper method to create SQL-implementation of TopicDao.
func NewTopicDaoSql(sqlc *prom.SqlConnect, tableName string, txModeOnWrite bool) TopicDao {
	dao := &BaseTopicDaoImpl{}
	dao.UniversalDao = henge.NewUniversalDaoSql(sqlc, tableName, txModeOnWrite,
		map[string]string{TopicColProductId: TopicFieldProductId, TopicColPosition: TopicFieldPosition})
	return dao
}

// InitTopicTableSql is helper function to initialize SQL table to store topics.
// This function also creates necessary indexes.
//
// Note: Application may need to create database before calling this function.
func InitTopicTableSql(sqlc *prom.SqlConnect, tableName string) error {
	var err error
	switch sqlc.GetDbFlavor() {
	case prom.FlavorCosmosDb:
		// spec := &henge.CosmosdbCollectionSpec{Pk: TopicFieldProductId}
		spec := &henge.CosmosdbCollectionSpec{Pk: henge.CosmosdbColId}
		err = henge.InitCosmosdbCollection(sqlc, tableName, spec)
	case prom.FlavorMySql:
		err = henge.InitMysqlTable(sqlc, tableName, map[string]string{TopicColProductId: "VARCHAR(32)", TopicColPosition: "INT"})
	case prom.FlavorPgSql:
		err = henge.InitPgsqlTable(sqlc, tableName, map[string]string{TopicColProductId: "VARCHAR(32)", TopicColPosition: "INT"})
	case prom.FlavorSqlite:
		err = henge.InitSqliteTable(sqlc, tableName, map[string]string{TopicColProductId: "VARCHAR(32)", TopicColPosition: "INT"})
	default:
		err = fmt.Errorf("unsupported database type %#v", sqlc.GetDbFlavor())
	}
	if err == nil {
		switch sqlc.GetDbFlavor() {
		case prom.FlavorMySql, prom.FlavorPgSql, prom.FlavorSqlite:
			idxName := "idx_" + tableName + "_" + TopicColProductId + "_" + TopicColPosition
			idxCols := TopicColProductId + "," + TopicColPosition
			sql := fmt.Sprintf("CREATE INDEX %s ON %s(%s)", idxName, tableName, idxCols)
			_, err = sqlc.GetDB().Exec(sql)
		}
	}
	if err != nil {
		log.Printf("[WARN] Creating table %s (%s) returns the following message: %s\n", tableName, _sqlFlavorStr(sqlc.GetDbFlavor()), err)
	}
	return err
}
