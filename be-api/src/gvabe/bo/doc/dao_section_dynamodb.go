package doc

import (
	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

// NewSectionDaoDynamodb is helper method to create AWS DynamoDB-implementation of SectionDao
func NewSectionDaoDynamodb(adc *prom.AwsDynamodbConnect, tableName string) SectionDao {
	dao := &BaseSectionDaoImpl{}
	spec := &henge.DynamodbDaoSpec{}
	dao.UniversalDao = henge.NewUniversalDaoDynamodb(adc, tableName, spec)
	return dao
}
