package libro

import (
	"fmt"
	"log"

	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

// NewProductDaoSql is helper method to create SQL-implementation of ProductDao
func NewProductDaoSql(sqlc *prom.SqlConnect, tableName string, txModeOnWrite bool) ProductDao {
	dao := &BaseProductDaoImpl{}
	dao.UniversalDao = henge.NewUniversalDaoSql(sqlc, tableName, txModeOnWrite, nil)
	return dao
}

// InitProductTableSql is helper function to initialize SQL table to store products.
// This function also creates necessary indexes.
//
// Note: Application may need to create database before calling this function.
func InitProductTableSql(sqlc *prom.SqlConnect, tableName string) error {
	var err error
	switch sqlc.GetDbFlavor() {
	case prom.FlavorCosmosDb:
		spec := &henge.CosmosdbCollectionSpec{Pk: henge.CosmosdbColId}
		err = henge.InitCosmosdbCollection(sqlc, tableName, spec)
	case prom.FlavorMySql:
		err = henge.InitMysqlTable(sqlc, tableName, nil)
	case prom.FlavorPgSql:
		err = henge.InitPgsqlTable(sqlc, tableName, nil)
	case prom.FlavorSqlite:
		err = henge.InitSqliteTable(sqlc, tableName, nil)
	default:
		err = fmt.Errorf("unsupported database type %#v", sqlc.GetDbFlavor())
	}
	if err != nil {
		log.Printf("[WARN] Creating table %s (%s) returns the following message: %s\n", tableName, _sqlFlavorStr(sqlc.GetDbFlavor()), err)
	}
	return err
}
