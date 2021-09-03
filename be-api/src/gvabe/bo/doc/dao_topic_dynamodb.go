package doc

import (
	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

// NewTopicDaoDynamodb is helper method to create AWS DynamoDB-implementation of TopicDao
func NewTopicDaoDynamodb(adc *prom.AwsDynamodbConnect, tableName string) TopicDao {
	dao := &BaseTopicDaoImpl{}
	spec := &henge.DynamodbDaoSpec{}
	dao.UniversalDao = henge.NewUniversalDaoDynamodb(adc, tableName, spec)
	return dao
}
