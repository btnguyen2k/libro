package user

import (
	"os"
	"strings"
	"testing"

	"github.com/btnguyen2k/prom"
)

const (
	testMongoCollection = "test_user"
)

func mongoInitCollection(mc *prom.MongoConnect, collection string) error {
	mc.GetCollection(collection).Drop(nil)
	return nil
}

func newMongoConnect(t *testing.T, testName string) (*prom.MongoConnect, error) {
	db := strings.Trim(os.Getenv("MONGO_DB"), "\"")
	url := strings.Trim(os.Getenv("MONGO_URL"), "\"")
	if db == "" || url == "" {
		t.Skipf("%s skipped", testName)
	}
	return prom.NewMongoConnect(url, db, 10000)
}

func initDaoMongo(mc *prom.MongoConnect) UserDao {
	return NewUserDaoMongo(mc, testMongoCollection, strings.Index(mc.GetUrl(), "replicaSet=") >= 0)
}

var setupTestDaoMongo = func(t *testing.T, testName string) {
	var err error
	testMc, err = newMongoConnect(t, testName)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName, err)
	} else if testMc == nil {
		t.Fatalf("%s failed: nil", testName)
	}

	err = mongoInitCollection(testMc, testMongoCollection)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName+"/mongoInitCollection", err)
	}
	err = InitUserTableMongo(testMc, testMongoCollection)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName+"/"+testDbType+"/InitUserTableMongo", err)
	}
}

var teardownTestDaoMongo = func(t *testing.T, testName string) {
	if testMc != nil {
		testMc.Close(nil)
		defer func() { testMc = nil }()
	}
}

/*----------------------------------------------------------------------*/

func TestNewUserDaoMongo(t *testing.T) {
	testName := "TestNewUserDaoMongo"
	teardownTest := setupTest(t, testName, setupTestDaoMongo, teardownTestDaoMongo)
	defer teardownTest(t)
	dao := initDaoMongo(testMc)
	if dao == nil {
		t.Fatalf("%s failed: nil", testName+"/initDaoMongo")
	}
}

func TestUserDaoMongo_CreateGet(t *testing.T) {
	testName := "TestUserDaoMongo_CreateGet"
	teardownTest := setupTest(t, testName, setupTestDaoMongo, teardownTestDaoMongo)
	defer teardownTest(t)
	dao := initDaoMongo(testMc)
	doTestUserDaoCreateGet(t, testName, dao)
}

func TestUserDaoMongo_CreateUpdateGet(t *testing.T) {
	testName := "TestUserDaoMongo_CreateGet"
	teardownTest := setupTest(t, testName, setupTestDaoMongo, teardownTestDaoMongo)
	defer teardownTest(t)
	dao := initDaoMongo(testMc)
	doTestUserDaoCreateUpdateGet(t, testName, dao)
}

func TestUserDaoMongo_CreateDelete(t *testing.T) {
	testName := "TestUserDaoMongo_CreateDelete"
	teardownTest := setupTest(t, testName, setupTestDaoMongo, teardownTestDaoMongo)
	defer teardownTest(t)
	dao := initDaoMongo(testMc)
	doTestUserDaoCreateDelete(t, testName, dao)
}

func TestUserDaoMongo_GetAll(t *testing.T) {
	testName := "TestUserDaoMongo_GetAll"
	teardownTest := setupTest(t, testName, setupTestDaoMongo, teardownTestDaoMongo)
	defer teardownTest(t)
	dao := initDaoMongo(testMc)
	doTestUserDaoGetAll(t, testName, dao)
}

func TestUserDaoMongo_GetN(t *testing.T) {
	testName := "TestUserDaoMongo_GetN"
	teardownTest := setupTest(t, testName, setupTestDaoMongo, teardownTestDaoMongo)
	defer teardownTest(t)
	dao := initDaoMongo(testMc)
	doTestUserDaoGetN(t, testName, dao)
}
