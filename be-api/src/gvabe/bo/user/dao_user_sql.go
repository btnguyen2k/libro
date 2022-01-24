package user

import (
	"fmt"
	"log"

	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

// NewUserDaoSql is helper method to create SQL-implementation of UserDao
//
// Available since template-v0.2.0
func NewUserDaoSql(sqlc *prom.SqlConnect, tableName string, txModeOnWrite bool) UserDao {
	dao := &BaseUserDaoImpl{}
	dao.UniversalDao = henge.NewUniversalDaoSql(
		sqlc, tableName, txModeOnWrite,
		map[string]string{UserColMaskUid: UserFieldMaskId})
	return dao
}

func _sqlFlavorStr(flavor prom.DbFlavor) string {
	switch flavor {
	case prom.FlavorCosmosDb:
		return "Azure Cosmos DB"
	case prom.FlavorMySql:
		return "MySQL"
	case prom.FlavorPgSql:
		return "PostgreSQL"
	case prom.FlavorSqlite:
		return "SQLite"
	}
	return fmt.Sprintf("%d", flavor)
}

// InitUserTableSql is helper function to initialize SQL table to store users.
// This function also creates necessary indexes.
//
// Note: Application may need to create database before calling this function.
func InitUserTableSql(sqlc *prom.SqlConnect, tableName string) error {
	var err error
	switch sqlc.GetDbFlavor() {
	case prom.FlavorCosmosDb:
		spec := &henge.CosmosdbCollectionSpec{Pk: henge.CosmosdbColId, Uk: [][]string{{"/" + UserColMaskUid}}}
		err = henge.InitCosmosdbCollection(sqlc, tableName, spec)
	case prom.FlavorMySql:
		err = henge.InitMysqlTable(sqlc, tableName, map[string]string{UserColMaskUid: "VARCHAR(32)"})
	case prom.FlavorPgSql:
		err = henge.InitPgsqlTable(sqlc, tableName, map[string]string{UserColMaskUid: "VARCHAR(32)"})
	case prom.FlavorSqlite:
		err = henge.InitSqliteTable(sqlc, tableName, map[string]string{UserColMaskUid: "VARCHAR(32)"})
	default:
		err = fmt.Errorf("unsupported database type %#v", sqlc.GetDbFlavor())
	}
	if err == nil {
		switch sqlc.GetDbFlavor() {
		case prom.FlavorMySql, prom.FlavorPgSql, prom.FlavorSqlite:
			idxName := "uidx_" + tableName + "_" + UserColMaskUid
			idxCols := UserColMaskUid
			sql := fmt.Sprintf("CREATE INDEX %s ON %s(%s)", idxName, tableName, idxCols)
			_, err = sqlc.GetDB().Exec(sql)
		}
	}
	if err != nil {
		log.Printf("[WARN] Creating table %s (%s) returns the following message: %s\n", tableName, _sqlFlavorStr(sqlc.GetDbFlavor()), err)
	}
	return err
}
