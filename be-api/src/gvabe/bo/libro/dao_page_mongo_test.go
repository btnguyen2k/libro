package libro

import (
	"strings"
	"testing"

	"github.com/btnguyen2k/prom"
)

const (
	testMongoCollectionPage = "test_page"
)

func initPageDaoMongo(mc *prom.MongoConnect) PageDao {
	return NewPageDaoMongo(mc, testMongoCollectionPage, strings.Index(mc.GetUrl(), "replicaSet=") >= 0)
}

var setupTestPageDaoMongo = func(t *testing.T, testName string) {
	var err error
	testMc, err = newMongoConnect(t, testName)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName, err)
	} else if testMc == nil {
		t.Fatalf("%s failed: nil", testName)
	}

	err = mongoInitCollection(testMc, testMongoCollectionPage)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName+"/mongoInitCollection", err)
	}
	err = InitPageTableMongo(testMc, testMongoCollectionPage)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName+"/"+testDbType+"/InitPageTableMongo", err)
	}
}

var teardownTestPageDaoMongo = func(t *testing.T, testName string) {
	if testMc != nil {
		testMc.Close(nil)
		defer func() { testMc = nil }()
	}
}

/*----------------------------------------------------------------------*/

func TestNewPageDaoMongo(t *testing.T) {
	testName := "TestNewPageDaoMongo"
	teardownTest := setupTest(t, testName, setupTestPageDaoMongo, teardownTestPageDaoMongo)
	defer teardownTest(t)
	dao := initPageDaoMongo(testMc)
	if dao == nil {
		t.Fatalf("%s failed: nil", testName+"/initPageDaoMongo")
	}
}

func TestPageDaoMongo_CreateGet(t *testing.T) {
	testName := "TestPageDaoMongo_CreateGet"
	teardownTest := setupTest(t, testName, setupTestPageDaoMongo, teardownTestPageDaoMongo)
	defer teardownTest(t)
	dao := initPageDaoMongo(testMc)
	doTestPageDaoCreateGet(t, testName, dao)
}

func TestPageDaoMongo_CreateUpdateGet(t *testing.T) {
	testName := "TestPageDaoMongo_CreateGet"
	teardownTest := setupTest(t, testName, setupTestPageDaoMongo, teardownTestPageDaoMongo)
	defer teardownTest(t)
	dao := initPageDaoMongo(testMc)
	doTestPageDaoCreateUpdateGet(t, testName, dao)
}

func TestPageDaoMongo_CreateDelete(t *testing.T) {
	testName := "TestPageDaoMongo_CreateDelete"
	teardownTest := setupTest(t, testName, setupTestPageDaoMongo, teardownTestPageDaoMongo)
	defer teardownTest(t)
	dao := initPageDaoMongo(testMc)
	doTestPageDaoCreateDelete(t, testName, dao)
}

func TestPageDaoMongo_GetAll(t *testing.T) {
	testName := "TestPageDaoMongo_GetAll"
	teardownTest := setupTest(t, testName, setupTestPageDaoMongo, teardownTestPageDaoMongo)
	defer teardownTest(t)
	dao := initPageDaoMongo(testMc)
	doTestPageDaoGetAll(t, testName, dao)
}

func TestPageDaoMongo_GetN(t *testing.T) {
	testName := "TestPageDaoMongo_GetN"
	teardownTest := setupTest(t, testName, setupTestPageDaoMongo, teardownTestPageDaoMongo)
	defer teardownTest(t)
	dao := initPageDaoMongo(testMc)
	doTestPageDaoGetN(t, testName, dao)
}
