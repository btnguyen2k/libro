package doc

import (
	"strings"

	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

const (
	awsDynamodbGSIPagePos = "gsi_{tableName}_pos"
)

// NewPageDaoDynamodb is helper method to create AWS DynamoDB-implementation of PageDao
func NewPageDaoDynamodb(adc *prom.AwsDynamodbConnect, tableName string) PageDao {
	dao := &BasePageDaoImpl{}
	spec := &henge.DynamodbDaoSpec{}
	innerDao := henge.NewUniversalDaoDynamodb(adc, tableName, spec)
	innerDao.MapGsi(strings.ReplaceAll(awsDynamodbGSIPagePos, "{tableName}", tableName), PageFieldPosition)
	dao.UniversalDao = innerDao
	return dao
}

// CreateDynamoTableForPages creates AWS DynamoDB table to store document pages.
//   - Necessary table and index (GSI) are created.
//   - If table/index will be created, RCU=1 and WCU=1 are used.
//   - If table/index already exist, they will be intact.
func CreateDynamoTableForPages(adc *prom.AwsDynamodbConnect, tableName string) error {
	spec := &henge.DynamodbTablesSpec{MainTableRcu: 1, MainTableWcu: 1}
	if err := henge.InitDynamodbTables(adc, tableName, spec); err != nil {
		return err
	}

	gsiName := strings.ReplaceAll(awsDynamodbGSIPagePos, "{tableName}", tableName)
	if status, err := adc.GetGlobalSecondaryIndexStatus(nil, tableName, gsiName); err != nil {
		return err
	} else if status == "" {
		attrDef := []prom.AwsDynamodbNameAndType{
			{Name: PageFieldTopicId, Type: prom.AwsAttrTypeString},
			{Name: PageFieldPosition, Type: prom.AwsAttrTypeNumber},
		}
		keyAttrs := []prom.AwsDynamodbNameAndType{
			{Name: PageFieldTopicId, Type: prom.AwsKeyTypePartition},
			{Name: PageFieldPosition, Type: prom.AwsKeyTypeSort},
		}
		return adc.CreateGlobalSecondaryIndex(nil, tableName, gsiName, 1, 1, attrDef, keyAttrs)
	}

	return nil
}
