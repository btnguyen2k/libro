package libro

import (
	"strings"
	"testing"

	"github.com/btnguyen2k/prom"
)

const (
	testMongoCollectionTopic = "test_topic"
)

func initTopicDaoMongo(mc *prom.MongoConnect) TopicDao {
	return NewTopicDaoMongo(mc, testMongoCollectionTopic, strings.Index(mc.GetUrl(), "replicaSet=") >= 0)
}

var setupTestTopicDaoMongo = func(t *testing.T, testName string) {
	var err error
	testMc, err = newMongoConnect(t, testName)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName, err)
	} else if testMc == nil {
		t.Fatalf("%s failed: nil", testName)
	}

	err = mongoInitCollection(testMc, testMongoCollectionTopic)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName+"/mongoInitCollection", err)
	}
	err = InitTopicTableMongo(testMc, testMongoCollectionTopic)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName+"/"+testDbType+"/InitTopicTableMongo", err)
	}
}

var teardownTestTopicDaoMongo = func(t *testing.T, testName string) {
	if testMc != nil {
		testMc.Close(nil)
		defer func() { testMc = nil }()
	}
}

/*----------------------------------------------------------------------*/

func TestNewTopicDaoMongo(t *testing.T) {
	testName := "TestNewTopicDaoMongo"
	teardownTest := setupTest(t, testName, setupTestTopicDaoMongo, teardownTestTopicDaoMongo)
	defer teardownTest(t)
	dao := initTopicDaoMongo(testMc)
	if dao == nil {
		t.Fatalf("%s failed: nil", testName+"/initTopicDaoMongo")
	}
}

func TestTopicDaoMongo_CreateGet(t *testing.T) {
	testName := "TestTopicDaoMongo_CreateGet"
	teardownTest := setupTest(t, testName, setupTestTopicDaoMongo, teardownTestTopicDaoMongo)
	defer teardownTest(t)
	dao := initTopicDaoMongo(testMc)
	doTestTopicDaoCreateGet(t, testName, dao)
}

func TestTopicDaoMongo_CreateUpdateGet(t *testing.T) {
	testName := "TestTopicDaoMongo_CreateGet"
	teardownTest := setupTest(t, testName, setupTestTopicDaoMongo, teardownTestTopicDaoMongo)
	defer teardownTest(t)
	dao := initTopicDaoMongo(testMc)
	doTestTopicDaoCreateUpdateGet(t, testName, dao)
}

func TestTopicDaoMongo_CreateDelete(t *testing.T) {
	testName := "TestTopicDaoMongo_CreateDelete"
	teardownTest := setupTest(t, testName, setupTestTopicDaoMongo, teardownTestTopicDaoMongo)
	defer teardownTest(t)
	dao := initTopicDaoMongo(testMc)
	doTestTopicDaoCreateDelete(t, testName, dao)
}

func TestTopicDaoMongo_GetAll(t *testing.T) {
	testName := "TestTopicDaoMongo_GetAll"
	teardownTest := setupTest(t, testName, setupTestTopicDaoMongo, teardownTestTopicDaoMongo)
	defer teardownTest(t)
	dao := initTopicDaoMongo(testMc)
	doTestTopicDaoGetAll(t, testName, dao)
}

func TestTopicDaoMongo_GetN(t *testing.T) {
	testName := "TestTopicDaoMongo_GetN"
	teardownTest := setupTest(t, testName, setupTestTopicDaoMongo, teardownTestTopicDaoMongo)
	defer teardownTest(t)
	dao := initTopicDaoMongo(testMc)
	doTestTopicDaoGetN(t, testName, dao)
}
