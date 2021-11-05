package doc

import (
	"fmt"

	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

// NewPageDaoSql is helper method to create SQL-implementation of PageDao
func NewPageDaoSql(sqlc *prom.SqlConnect, tableName string, txModeOnWrite bool) PageDao {
	dao := &BasePageDaoImpl{}
	dao.UniversalDao = henge.NewUniversalDaoSql(sqlc, tableName, txModeOnWrite,
		map[string]string{PageColProductId: PageFieldProductId, PageColTopicId: PageFieldTopicId, PageColPosition: PageFieldPosition})
	return dao
}

// CreateSqlTableForPages creates SQL database table to store document pages.
//   - Necessary table and index are created.
//   - Application may need to create database before calling this function.
func CreateSqlTableForPages(sqlc *prom.SqlConnect, tableName string) error {
	var err error
	switch sqlc.GetDbFlavor() {
	case prom.FlavorCosmosDb:
		spec := &henge.CosmosdbCollectionSpec{Pk: henge.CosmosdbColId}
		err = henge.InitCosmosdbCollection(sqlc, tableName, spec)
	case prom.FlavorMySql:
		err = henge.InitMysqlTable(sqlc, tableName, map[string]string{PageColProductId: "VARCHAR(32)", PageColTopicId: "VARCHAR(32)", PageColPosition: "INT"})
	case prom.FlavorPgSql:
		err = henge.InitPgsqlTable(sqlc, tableName, map[string]string{PageColProductId: "VARCHAR(32)", PageColTopicId: "VARCHAR(32)", PageColPosition: "INT"})
	case prom.FlavorSqlite:
		err = henge.InitSqliteTable(sqlc, tableName, map[string]string{PageColProductId: "VARCHAR(32)", PageColTopicId: "VARCHAR(32)", PageColPosition: "INT"})
	default:
		err = fmt.Errorf("unsupported database type %#v", sqlc.GetDB())
	}
	if err == nil {
		switch sqlc.GetDbFlavor() {
		case prom.FlavorCosmosDb, prom.FlavorMySql, prom.FlavorPgSql, prom.FlavorSqlite:
			idxName := "idx_" + tableName + "_" + PageColTopicId + "_" + PageColPosition
			idxCols := PageColTopicId + "," + PageColPosition
			sql := fmt.Sprintf("CREATE INDEX %s ON %s(%s)", idxName, tableName, idxCols)
			_, err = sqlc.GetDB().Exec(sql)
		}
	}
	return err
}
