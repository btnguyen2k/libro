package doc

import (
	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

// NewSectionDaoMongo is helper method to create MongoDB-implementation of SectionDao
func NewSectionDaoMongo(mc *prom.MongoConnect, collectionName string, txModeOnWrite bool) SectionDao {
	dao := &BaseSectionDaoImpl{}
	dao.UniversalDao = henge.NewUniversalDaoMongo(mc, collectionName, txModeOnWrite)
	return dao
}
