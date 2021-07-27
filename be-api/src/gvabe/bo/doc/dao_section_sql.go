package doc

import (
	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

// NewSectionDaoSql is helper method to create SQL-implementation of SectionDao
func NewSectionDaoSql(sqlc *prom.SqlConnect, tableName string, txModeOnWrite bool) SectionDao {
	dao := &BaseSectionDaoImpl{}
	dao.UniversalDao = henge.NewUniversalDaoSql(sqlc, tableName, txModeOnWrite, map[string]string{SectionColAppId: SectionFieldAppId})
	return dao
}
