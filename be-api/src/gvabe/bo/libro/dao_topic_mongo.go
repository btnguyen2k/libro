package libro

import (
	"log"

	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewTopicDaoMongo is helper method to create MongoDB-implementation of TopicDao
func NewTopicDaoMongo(mc *prom.MongoConnect, collectionName string, txModeOnWrite bool) TopicDao {
	dao := &BaseTopicDaoImpl{}
	dao.UniversalDao = henge.NewUniversalDaoMongo(mc, collectionName, txModeOnWrite)
	return dao
}

// InitTopicTableMongo is helper function to initialize MongoDB table (collection) to store topics.
// This function also creates necessary indexes.
//
// Note: as MongoDB is schemaless, "field name" is used instead of "column name"! Application may need to create database before calling this function.
func InitTopicTableMongo(mc *prom.MongoConnect, collectionName string) error {
	err := henge.InitMongoCollection(mc, collectionName)
	if err != nil {
		log.Printf("[WARN] Creating collection %s (%s) returns the following message: %s\n", collectionName, "MongoDB", err)
		return err
	}

	unique := false
	idxName := "idx_" + TopicFieldProductId + "_" + TopicFieldPosition
	_, err = mc.CreateCollectionIndexes(collectionName, []interface{}{mongo.IndexModel{
		Keys:    bson.D{{TopicFieldProductId, 1}, {TopicFieldPosition, 1}},
		Options: &options.IndexOptions{Name: &idxName, Unique: &unique},
	}})
	if err != nil {
		log.Printf("[WARN] Creating collection index %s/%s (%s) returns the following message: %s\n", collectionName, idxName, "MongoDB", err)
	}
	return err
}
