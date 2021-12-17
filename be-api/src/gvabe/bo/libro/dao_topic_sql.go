package libro

import (
	"fmt"

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

// CreateSqlTableForTopics creates SQL database table to store document topics.
//   - Necessary table and index are created.
//   - Application may need to create database before calling this function.
func CreateSqlTableForTopics(sqlc *prom.SqlConnect, tableName string) error {
	var err error
	switch sqlc.GetDbFlavor() {
	case prom.FlavorCosmosDb:
		spec := &henge.CosmosdbCollectionSpec{Pk: henge.CosmosdbColId}
		err = henge.InitCosmosdbCollection(sqlc, tableName, spec)
	case prom.FlavorMySql:
		err = henge.InitMysqlTable(sqlc, tableName, map[string]string{TopicColProductId: "VARCHAR(32)", TopicColPosition: "INT"})
	case prom.FlavorPgSql:
		err = henge.InitPgsqlTable(sqlc, tableName, map[string]string{TopicColProductId: "VARCHAR(32)", TopicColPosition: "INT"})
	case prom.FlavorSqlite:
		err = henge.InitSqliteTable(sqlc, tableName, map[string]string{TopicColProductId: "VARCHAR(32)", TopicColPosition: "INT"})
	default:
		err = fmt.Errorf("unsupported database type %#v", sqlc.GetDB())
	}
	if err == nil {
		switch sqlc.GetDbFlavor() {
		case prom.FlavorCosmosDb, prom.FlavorMySql, prom.FlavorPgSql, prom.FlavorSqlite:
			idxName := "idx_" + tableName + "_" + TopicColProductId + "_" + TopicColPosition
			idxCols := TopicColProductId + "," + TopicColPosition
			sql := fmt.Sprintf("CREATE INDEX %s ON %s(%s)", idxName, tableName, idxCols)
			_, err = sqlc.GetDB().Exec(sql)
		}
	}
	return err
}
