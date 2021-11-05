package doc

import (
	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
	"go.mongodb.org/mongo-driver/bson"
)

// NewTopicDaoMongo is helper method to create MongoDB-implementation of TopicDao
func NewTopicDaoMongo(mc *prom.MongoConnect, collectionName string, txModeOnWrite bool) TopicDao {
	dao := &BaseTopicDaoImpl{}
	dao.UniversalDao = henge.NewUniversalDaoMongo(mc, collectionName, txModeOnWrite)
	return dao
}

// CreateMongoCollectionForTopics creates MongoDB collection to store document topics.
//   - Necessary collection and index are created.
//   - Application may need to create database before calling this function.
func CreateMongoCollectionForTopics(mc *prom.MongoConnect, collectionName string) error {
	err := henge.InitMongoCollection(mc, collectionName)
	if err != nil {
		return err
	}
	indexes := []interface{}{
		map[string]interface{}{
			"key":    bson.D{{TopicFieldProductId, 1}, {TopicFieldPosition, 1}},
			"name":   "idx_" + TopicFieldProductId + "_" + TopicFieldPosition,
			"unique": false,
		},
	}
	_, err = mc.CreateCollectionIndexes(collectionName, indexes)
	return err
}
