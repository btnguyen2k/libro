package respicite

import (
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/btnguyen2k/prom"
)

const (
	testMongoCollectionM2o = "test_m2o"
)

func _mongoInitCollectionM2o(mc *prom.MongoConnect, collectionName string) error {
	rand.Seed(time.Now().UnixNano())
	mc.GetCollection(collectionName).Drop(nil)
	return M2oDaoMongoInitCollection(mc, collectionName)
}

func _createDaoMongoM2o(mc *prom.MongoConnect, collectionName string) *M2oDaoMongo {
	return NewM2oDaoMongo(mc, collectionName, strings.Index(mc.GetUrl(), "replicaSet=") >= 0)
}

func _buildMongoConnect(t *testing.T, testName string) *prom.MongoConnect {
	mongoDb := strings.ReplaceAll(os.Getenv("MONGO_DB"), `"`, "")
	mongoUrl := strings.ReplaceAll(os.Getenv("MONGO_URL"), `"`, "")
	if mongoDb == "" || mongoUrl == "" {
		t.Skipf("%s skipped", testName)
		return nil
	}
	mc, _ := prom.NewMongoConnectWithPoolOptions(mongoUrl, mongoDb, 10000, &prom.MongoPoolOpts{
		ConnectTimeout:         10 * time.Second,
		SocketTimeout:          10 * time.Second,
		ServerSelectionTimeout: 10 * time.Second,
	})
	return mc
}

func _initDaoM2oMongo(t *testing.T, testName, collectionName string) *M2oDaoMongo {
	mc := _buildMongoConnect(t, testName)
	if err := _mongoInitCollectionM2o(mc, collectionName); err != nil {
		t.Fatalf("%s failed: %s", testName, err)
	}
	return _createDaoMongoM2o(mc, collectionName)
}

/*----------------------------------------------------------------------*/

func TestNewM2oDaoMongo(t *testing.T) {
	testName := "TestNewM2oDaoMongo"
	dao := _initDaoM2oMongo(t, testName, testMongoCollectionM2o)
	if dao == nil {
		t.Fatalf("%s failed: nil", testName)
	}
	defer dao.Destroy()
}

func TestM2oDaoMongo_GetNotExist(t *testing.T) {
	testName := "TestM2oDaoMongo_GetNotExist"
	dao := _initDaoM2oMongo(t, testName, testMongoCollectionM2o)
	defer dao.Destroy()
	doTestM2mDao_GetNotExist(t, testName, dao)
}

func TestM2oDaoMongo_SetGet(t *testing.T) {
	testName := "TestM2oDaoMongo_SetGet"
	dao := _initDaoM2oMongo(t, testName, testMongoCollectionM2o)
	defer dao.Destroy()
	doTestM2mDao_SetGet(t, testName, dao)
}

func TestM2oDaoMongo_SetDuplicated(t *testing.T) {
	testName := "TestM2oDaoMongo_SetDuplicated"
	dao := _initDaoM2oMongo(t, testName, testMongoCollectionM2o)
	defer dao.Destroy()
	doTestM2mDao_SetDuplicated(t, testName, dao)
}

func TestM2oDaoMongo_SetRemove(t *testing.T) {
	testName := "TestM2oDaoMongo_SetRemove"
	dao := _initDaoM2oMongo(t, testName, testMongoCollectionM2o)
	defer dao.Destroy()
	doTestM2mDao_SetRemove(t, testName, dao)
}

func TestM2oDaoMongo_RemoveNotExist(t *testing.T) {
	testName := "TestM2oDaoMongo_RemoveNotExist"
	dao := _initDaoM2oMongo(t, testName, testMongoCollectionM2o)
	defer dao.Destroy()
	doTestM2mDao_RemoveNotExist(t, testName, dao)
}

func TestM2oDaoMongo_SetRget(t *testing.T) {
	testName := "TestM2oDaoMongo_SetRget"
	dao := _initDaoM2oMongo(t, testName, testMongoCollectionM2o)
	defer dao.Destroy()
	doTestM2mDao_SetRget(t, testName, dao)
}
