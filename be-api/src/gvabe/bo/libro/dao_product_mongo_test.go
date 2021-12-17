package libro

import (
	"strings"
	"testing"

	"github.com/btnguyen2k/prom"
)

const (
	testMongoCollectionProduct = "test_product"
)

func initProductDaoMongo(mc *prom.MongoConnect) ProductDao {
	return NewProductDaoMongo(mc, testMongoCollectionProduct, strings.Index(mc.GetUrl(), "replicaSet=") >= 0)
}

var setupTestProductDaoMongo = func(t *testing.T, testName string) {
	var err error
	testMc, err = newMongoConnect(t, testName, "", "")
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName, err)
	} else if testMc == nil {
		t.Fatalf("%s failed: nil", testName)
	}

	err = mongoInitCollection(testMc, testMongoCollectionProduct)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName+"/mongoInitCollection", err)
	}
}

var teardownTestProductDaoMongo = func(t *testing.T, testName string) {
	if testMc != nil {
		testMc.Close(nil)
		defer func() { testMc = nil }()
	}
}

/*----------------------------------------------------------------------*/

func TestNewProductDaoMongo(t *testing.T) {
	testName := "TestNewProductDaoMongo"
	teardownTest := setupTest(t, testName, setupTestProductDaoMongo, teardownTestProductDaoMongo)
	defer teardownTest(t)

	dao := initProductDaoMongo(testMc)
	if dao == nil {
		t.Fatalf("%s failed: nil", testName+"/initProductDaoMongo")
	}
}

func TestProductDaoMongo_CreateGet(t *testing.T) {
	testName := "TestProductDaoMongo_CreateGet"
	teardownTest := setupTest(t, testName, setupTestProductDaoMongo, teardownTestProductDaoMongo)
	defer teardownTest(t)

	dao := initProductDaoMongo(testMc)
	doTestProductDaoCreateGet(t, testName, dao)
}

func TestProductDaoMongo_CreateUpdateGet(t *testing.T) {
	testName := "TestProductDaoMongo_CreateGet"
	teardownTest := setupTest(t, testName, setupTestProductDaoMongo, teardownTestProductDaoMongo)
	defer teardownTest(t)

	dao := initProductDaoMongo(testMc)
	doTestProductDaoCreateUpdateGet(t, testName, dao)
}

func TestProductDaoMongo_CreateDelete(t *testing.T) {
	testName := "TestProductDaoMongo_CreateDelete"
	teardownTest := setupTest(t, testName, setupTestProductDaoMongo, teardownTestProductDaoMongo)
	defer teardownTest(t)

	dao := initProductDaoMongo(testMc)
	doTestProductDaoCreateDelete(t, testName, dao)
}

func TestProductDaoMongo_GetAll(t *testing.T) {
	testName := "TestProductDaoMongo_GetAll"
	teardownTest := setupTest(t, testName, setupTestProductDaoMongo, teardownTestProductDaoMongo)
	defer teardownTest(t)

	dao := initProductDaoMongo(testMc)
	doTestProductDaoGetAll(t, testName, dao)
}

func TestProductDaoMongo_GetN(t *testing.T) {
	testName := "TestProductDaoMongo_GetN"
	teardownTest := setupTest(t, testName, setupTestProductDaoMongo, teardownTestProductDaoMongo)
	defer teardownTest(t)

	dao := initProductDaoMongo(testMc)
	doTestProductDaoGetN(t, testName, dao)
}
