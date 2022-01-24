package libro

import (
	"log"

	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

// NewProductDaoMongo is helper method to create MongoDB-implementation of ProductDao
func NewProductDaoMongo(mc *prom.MongoConnect, collectionName string, txModeOnWrite bool) ProductDao {
	dao := &BaseProductDaoImpl{}
	dao.UniversalDao = henge.NewUniversalDaoMongo(mc, collectionName, txModeOnWrite)
	return dao
}

// InitProductTableMongo is helper function to initialize MongoDB table (collection) to store products.
// This function also creates necessary indexes.
//
// Note: as MongoDB is schemaless, "field name" is used instead of "column name"! Application may need to create database before calling this function.
func InitProductTableMongo(mc *prom.MongoConnect, collectionName string) error {
	err := henge.InitMongoCollection(mc, collectionName)
	if err != nil {
		log.Printf("[WARN] Creating collection %s (%s) returns the following message: %s\n", collectionName, "MongoDB", err)
		return err
	}
	return err
}
