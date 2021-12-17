package libro

import (
	"testing"

	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

const (
	testDynamodbTableProduct = "test_product"
)

func initProductDaoDynamodb(adc *prom.AwsDynamodbConnect) ProductDao {
	return NewProductDaoDynamodb(adc, testDynamodbTableProduct)
}

var testDynamodbTableSpecProduct = &henge.DynamodbTablesSpec{MainTableRcu: 2, MainTableWcu: 1, CreateUidxTable: true, UidxTableRcu: 2, UidxTableWcu: 1}
var setupTestProductDaoDynamodb = func(t *testing.T, testName string) {
	var err error
	testAdc, err = newDynamodbConnect(t, testName)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName, err)
	} else if testAdc == nil {
		t.Fatalf("%s failed: nil", testName)
	}

	err = dynamodbInitTable(testAdc, testDynamodbTableProduct, testDynamodbTableSpecProduct)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName+"/dynamodbInitTable", err)
	}
}

var teardownTestProductDaoDynamodb = func(t *testing.T, testName string) {
	if testAdc != nil {
		testAdc.Close()
		defer func() { testAdc = nil }()
	}
}

/*----------------------------------------------------------------------*/

func TestNewProductDaoDynamodb(t *testing.T) {
	testName := "TestNewProductDaoDynamodb"
	teardownTest := setupTest(t, testName, setupTestProductDaoDynamodb, teardownTestProductDaoDynamodb)
	defer teardownTest(t)

	dao := initProductDaoDynamodb(testAdc)
	if dao == nil {
		t.Fatalf("%s failed: nil", testName+"/initProductDaoDynamodb")
	}
}

func TestProductDaoDynamodb_CreateGet(t *testing.T) {
	testName := "TestProductDaoDynamodb_CreateGet"
	teardownTest := setupTest(t, testName, setupTestProductDaoDynamodb, teardownTestProductDaoDynamodb)
	defer teardownTest(t)

	dao := initProductDaoDynamodb(testAdc)
	doTestProductDaoCreateGet(t, testName, dao)
}

func TestProductDaoDynamodb_CreateUpdateGet(t *testing.T) {
	testName := "TestProductDaoDynamodb_CreateUpdateGet"
	teardownTest := setupTest(t, testName, setupTestProductDaoDynamodb, teardownTestProductDaoDynamodb)
	defer teardownTest(t)

	dao := initProductDaoDynamodb(testAdc)
	doTestProductDaoCreateUpdateGet(t, testName, dao)
}

func TestProductDaoDynamodb_CreateDelete(t *testing.T) {
	testName := "TestProductDaoDynamodb_CreateDelete"
	teardownTest := setupTest(t, testName, setupTestProductDaoDynamodb, teardownTestProductDaoDynamodb)
	defer teardownTest(t)

	dao := initProductDaoDynamodb(testAdc)
	doTestProductDaoCreateDelete(t, testName, dao)
}

func TestProductDaoDynamodb_GetAll(t *testing.T) {
	testName := "TestProductDaoDynamodb_GetAll"
	teardownTest := setupTest(t, testName, setupTestProductDaoDynamodb, teardownTestProductDaoDynamodb)
	defer teardownTest(t)

	dao := initProductDaoDynamodb(testAdc)
	doTestProductDaoGetAll(t, testName, dao)
}

func TestProductDaoDynamodb_GetN(t *testing.T) {
	testName := "TestProductDaoDynamodb_GetN"
	teardownTest := setupTest(t, testName, setupTestProductDaoDynamodb, teardownTestProductDaoDynamodb)
	defer teardownTest(t)

	dao := initProductDaoDynamodb(testAdc)
	doTestProductDaoGetN(t, testName, dao)
}
