package libro

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

// InitProductTableDynamodb creates AWS DynamoDB table to store products.
//   - Necessary table and index (GSI) are created.
//   - If table/index will be created, RCU=1 and WCU=1 are used.
//   - If table/index already exist, they will be intact.
func InitProductTableDynamodb(adc *prom.AwsDynamodbConnect, tableName string) error {
	spec := &henge.DynamodbTablesSpec{MainTableRcu: 1, MainTableWcu: 1}
	if err := henge.InitDynamodbTables(adc, tableName, spec); err != nil {
		return err
	}
	return nil
}
