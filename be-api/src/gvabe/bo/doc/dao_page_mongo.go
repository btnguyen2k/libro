package doc

import (
	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
	"go.mongodb.org/mongo-driver/bson"
)

// NewPageDaoMongo is helper method to create MongoDB-implementation of PageDao
func NewPageDaoMongo(mc *prom.MongoConnect, collectionName string, txModeOnWrite bool) PageDao {
	dao := &BasePageDaoImpl{}
	dao.UniversalDao = henge.NewUniversalDaoMongo(mc, collectionName, txModeOnWrite)
	return dao
}

// CreateMongoCollectionForPages creates MongoDB collection to store document pages.
//   - Necessary collection and index are created.
//   - Application may need to create database before calling this function.
func CreateMongoCollectionForPages(mc *prom.MongoConnect, collectionName string) error {
	err := henge.InitMongoCollection(mc, collectionName)
	if err != nil {
		return err
	}
	indexes := []interface{}{
		map[string]interface{}{
			"key":    bson.D{{PageFieldTopicId, 1}, {PageFieldPosition, 1}},
			"name":   "idx_" + PageFieldTopicId + "_" + PageFieldPosition,
			"unique": false,
		},
	}
	_, err = mc.CreateCollectionIndexes(collectionName, indexes)
	return err
}
