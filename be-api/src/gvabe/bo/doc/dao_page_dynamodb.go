package doc

import (
	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

// NewPageDaoDynamodb is helper method to create AWS DynamoDB-implementation of PageDao
func NewPageDaoDynamodb(adc *prom.AwsDynamodbConnect, tableName string) PageDao {
	dao := &BasePageDaoImpl{}
	spec := &henge.DynamodbDaoSpec{}
	dao.UniversalDao = henge.NewUniversalDaoDynamodb(adc, tableName, spec)
	return dao
}
