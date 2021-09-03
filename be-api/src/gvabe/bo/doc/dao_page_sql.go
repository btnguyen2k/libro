package doc

import (
	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

// NewPageDaoSql is helper method to create SQL-implementation of PageDao
func NewPageDaoSql(sqlc *prom.SqlConnect, tableName string, txModeOnWrite bool) PageDao {
	dao := &BasePageDaoImpl{}
	dao.UniversalDao = henge.NewUniversalDaoSql(sqlc, tableName, txModeOnWrite,
		map[string]string{PageColAppId: PageFieldAppId, PageColTopicId: PageFieldTopicId})
	return dao
}
