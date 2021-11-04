package doc

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

// // CreateDynamoTableForTopics creates AWS DynamoDB table to store topics.
// //   - Necessary table and index (GSI) are created.
// func CreateDynamoTableForTopics(adc *prom.AwsDynamodbConnect, tableName string) error {
// 	spec := &henge.DynamodbTablesSpec{MainTableRcu: 1, MainTableWcu: 1, CreateUidxTable: true, UidxTableRcu: 1, UidxTableWcu: 1}
// 	if err := henge.InitDynamodbTables(adc, tableName, spec); err != nil {
// 		return err
// 	}
//
// 	// gsiName := "gsi_" + tableName + "_" + TopicFieldProductId + "_" + TopicFieldPosition
// 	// attrDef := []prom.AwsDynamodbNameAndType{
// 	// 	{Name: TopicFieldProductId, Type: prom.AwsAttrTypeString},
// 	// 	{Name: TopicFieldPosition, Type: prom.AwsAttrTypeNumber},
// 	// }
// 	// keyAttrs := []prom.AwsDynamodbNameAndType{
// 	// 	{Name: TopicFieldProductId, Type: prom.AwsKeyTypePartition},
// 	// 	{Name: TopicFieldPosition, Type: prom.AwsKeyTypeSort},
// 	// }
// 	// if err := adc.CreateGlobalSecondaryIndex(nil, tableName, gsiName, 1, 1, attrDef, keyAttrs); err != nil {
// 	// 	return err
// 	// }
//
// 	// TODO create GSI?
//
// 	return nil
// }

// // DynamodbTopicDao is AWS DynamoDB implementation of TopicDao
// type DynamodbTopicDao struct {
// 	*BaseTopicDaoImpl
// 	tableName            string
// 	gsiProductIdPosition string
// }
//
// // GetN implements TopicDao.GetN
// func (dao *DynamodbTopicDao) GetN(prod *product.Product, fromOffset, maxNumRows int, filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*Topic, error) {
// 	result, err := dao.BaseTopicDaoImpl.GetN(prod, fromOffset, maxNumRows, filter, sorting)
// 	if result != nil {
// 		// TODO this does not work with paging!
// 		sort.Slice(result, func(i, j int) bool {
// 			return result[i].position < result[j].position
// 		})
// 	}
// 	return result, err
// }
//
// // GetAll implements TopicDao.GetAll
// func (dao *DynamodbTopicDao) GetAll(prod *product.Product, filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*Topic, error) {
// 	return dao.GetN(prod, 0, 0, filter, sorting)
// }
