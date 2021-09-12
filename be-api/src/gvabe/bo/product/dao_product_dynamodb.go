package product

import (
	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

// NewProductDaoDynamodb is helper method to create AWS DynamoDB-implementation of ProductDao
func NewProductDaoDynamodb(adc *prom.AwsDynamodbConnect, tableName string) ProductDao {
	dao := &BaseProductDaoImpl{}
	spec := &henge.DynamodbDaoSpec{}
	dao.UniversalDao = henge.NewUniversalDaoDynamodb(adc, tableName, spec)
	return dao
}
