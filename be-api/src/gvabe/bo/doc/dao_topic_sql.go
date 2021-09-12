package doc

import (
	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

// NewTopicDaoSql is helper method to create SQL-implementation of TopicDao
func NewTopicDaoSql(sqlc *prom.SqlConnect, tableName string, txModeOnWrite bool) TopicDao {
	dao := &BaseTopicDaoImpl{}
	dao.UniversalDao = henge.NewUniversalDaoSql(sqlc, tableName, txModeOnWrite, map[string]string{TopicColProductId: TopicFieldProductId})
	return dao
}
