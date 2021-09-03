package doc

import (
	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

// NewTopicDaoMongo is helper method to create MongoDB-implementation of TopicDao
func NewTopicDaoMongo(mc *prom.MongoConnect, collectionName string, txModeOnWrite bool) TopicDao {
	dao := &BaseTopicDaoImpl{}
	dao.UniversalDao = henge.NewUniversalDaoMongo(mc, collectionName, txModeOnWrite)
	return dao
}
