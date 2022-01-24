package user

import (
	"log"

	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewUserDaoMongo is helper function to create MongoDB-implementation of UserDao.
//
// Available since template-v0.3.0
func NewUserDaoMongo(mc *prom.MongoConnect, collectionName string, txModeOnWrite bool) UserDao {
	dao := &BaseUserDaoImpl{}
	dao.UniversalDao = henge.NewUniversalDaoMongo(mc, collectionName, txModeOnWrite)
	return dao
}

// InitUserTableMongo is helper function to initialize MongoDB table (collection) to store user accounts.
// This function also creates necessary indexes.
//
// Note: as MongoDB is schemaless, "field name" is used instead of "column name"! Application may need to create database before calling this function.
func InitUserTableMongo(mc *prom.MongoConnect, collectionName string) error {
	err := henge.InitMongoCollection(mc, collectionName)
	if err != nil {
		log.Printf("[WARN] Creating collection %s (%s) returns the following message: %s\n", collectionName, "MongoDB", err)
		return err
	}

	unique := true
	idxName := "idx_" + UserFieldMaskId
	_, err = mc.CreateCollectionIndexes(collectionName, []interface{}{mongo.IndexModel{
		Keys:    bson.D{{UserFieldMaskId, 1}},
		Options: &options.IndexOptions{Name: &idxName, Unique: &unique},
	}})
	if err != nil {
		log.Printf("[WARN] Creating collection index %s/%s (%s) returns the following message: %s\n", collectionName, idxName, "MongoDB", err)
	}
	return err
}
