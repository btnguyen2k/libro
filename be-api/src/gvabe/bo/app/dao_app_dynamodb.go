package app

import (
	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

// NewAppDaoDynamodb is helper method to create AWS DynamoDB-implementation of AppDao
func NewAppDaoDynamodb(adc *prom.AwsDynamodbConnect, tableName string) AppDao {
	dao := &BaseAppDaoImpl{}
	spec := &henge.DynamodbDaoSpec{}
	dao.UniversalDao = henge.NewUniversalDaoDynamodb(adc, tableName, spec)
	return dao
}
