package product

import (
	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

// NewProductDaoSql is helper method to create SQL-implementation of ProductDao
func NewProductDaoSql(sqlc *prom.SqlConnect, tableName string, txModeOnWrite bool) ProductDao {
	dao := &BaseProductDaoImpl{}
	dao.UniversalDao = henge.NewUniversalDaoSql(sqlc, tableName, txModeOnWrite, nil)
	return dao
}
