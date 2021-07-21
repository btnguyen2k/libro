package app

import (
	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

// NewAppDaoSql is helper method to create SQL-implementation of AppDao
func NewAppDaoSql(sqlc *prom.SqlConnect, tableName string, txModeOnWrite bool) AppDao {
	dao := &BaseAppDaoImpl{}
	dao.UniversalDao = henge.NewUniversalDaoSql(sqlc, tableName, txModeOnWrite, nil)
	return dao
}
