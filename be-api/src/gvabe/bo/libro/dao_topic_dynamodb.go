package libro

import (
	"strings"

	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

const (
	awsDynamodbGSITopicPos = "gsi_{tableName}_pos"
)

// NewTopicDaoDynamodb is helper method to create AWS DynamoDB-implementation of TopicDao
func NewTopicDaoDynamodb(adc *prom.AwsDynamodbConnect, tableName string) TopicDao {
	dao := &BaseTopicDaoImpl{}
	spec := &henge.DynamodbDaoSpec{}
	innerDao := henge.NewUniversalDaoDynamodb(adc, tableName, spec)
	innerDao.MapGsi(strings.ReplaceAll(awsDynamodbGSITopicPos, "{tableName}", tableName), TopicFieldPosition)
	dao.UniversalDao = innerDao
	return dao
}

// CreateDynamoTableForTopics creates AWS DynamoDB table to store document topics.
//   - Necessary table and index (GSI) are created.
//   - If table/index will be created, RCU=1 and WCU=1 are used.
//   - If table/index already exist, they will be intact.
func CreateDynamoTableForTopics(adc *prom.AwsDynamodbConnect, tableName string) error {
	spec := &henge.DynamodbTablesSpec{MainTableRcu: 1, MainTableWcu: 1}
	if err := henge.InitDynamodbTables(adc, tableName, spec); err != nil {
		return err
	}

	gsiName := strings.ReplaceAll(awsDynamodbGSITopicPos, "{tableName}", tableName)
	if status, err := adc.GetGlobalSecondaryIndexStatus(nil, tableName, gsiName); err != nil {
		return err
	} else if status == "" {
		attrDef := []prom.AwsDynamodbNameAndType{
			{Name: TopicFieldProductId, Type: prom.AwsAttrTypeString},
			{Name: TopicFieldPosition, Type: prom.AwsAttrTypeNumber},
		}
		keyAttrs := []prom.AwsDynamodbNameAndType{
			{Name: TopicFieldProductId, Type: prom.AwsKeyTypePartition},
			{Name: TopicFieldPosition, Type: prom.AwsKeyTypeSort},
		}
		return adc.CreateGlobalSecondaryIndex(nil, tableName, gsiName, 1, 1, attrDef, keyAttrs)
	}

	return nil
}
