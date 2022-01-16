package libro

import (
	"strings"
	"testing"

	"github.com/btnguyen2k/prom"
)

const (
	testDynamodbTablePage = "test_page"
)

func initPageDaoDynamodb(adc *prom.AwsDynamodbConnect) PageDao {
	return NewPageDaoDynamodb(adc, testDynamodbTablePage)
}

var setupTestPageDaoDynamodb = func(t *testing.T, testName string) {
	var err error
	testAdc, err = newDynamodbConnect(t, testName)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName, err)
	} else if testAdc == nil {
		t.Fatalf("%s failed: nil", testName)
	}

	err = dynamodbInitTable(testAdc, testDynamodbTablePage)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName+"/dynamodbInitTable", err)
	}
	err = InitPageTableDynamodb(testAdc, testDynamodbTablePage)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName+"/"+testDbType+"/InitPageTableDynamodb", err)
	}
}

var teardownTestPageDaoDynamodb = func(t *testing.T, testName string) {
	if testAdc != nil {
		testAdc.Close()
		defer func() { testAdc = nil }()
	}
}

/*----------------------------------------------------------------------*/

func TestNewPageDaoDynamodb(t *testing.T) {
	testName := "TestNewPageDaoDynamodb"
	teardownTest := setupTest(t, testName, setupTestPageDaoDynamodb, teardownTestPageDaoDynamodb)
	defer teardownTest(t)
	dao := initPageDaoDynamodb(testAdc)
	if dao == nil {
		t.Fatalf("%s failed: nil", testName+"/initPageDaoDynamodb")
	}
}

func TestPageDaoDynamodb_CreateGet(t *testing.T) {
	testName := "TestPageDaoDynamodb_CreateGet"
	teardownTest := setupTest(t, testName, setupTestPageDaoDynamodb, teardownTestPageDaoDynamodb)
	defer teardownTest(t)
	dao := initPageDaoDynamodb(testAdc)
	doTestPageDaoCreateGet(t, testName, dao)
}

func TestPageDaoDynamodb_CreateUpdateGet(t *testing.T) {
	testName := "TestPageDaoDynamodb_CreateUpdateGet"
	teardownTest := setupTest(t, testName, setupTestPageDaoDynamodb, teardownTestPageDaoDynamodb)
	defer teardownTest(t)
	dao := initPageDaoDynamodb(testAdc)
	doTestPageDaoCreateUpdateGet(t, testName, dao)
}

func TestPageDaoDynamodb_CreateDelete(t *testing.T) {
	testName := "TestPageDaoDynamodb_CreateDelete"
	teardownTest := setupTest(t, testName, setupTestPageDaoDynamodb, teardownTestPageDaoDynamodb)
	defer teardownTest(t)
	dao := initPageDaoDynamodb(testAdc)
	doTestPageDaoCreateDelete(t, testName, dao)
}

func TestPageDaoDynamodb_GetAll(t *testing.T) {
	testName := "TestPageDaoDynamodb_GetAll"
	teardownTest := setupTest(t, testName, setupTestPageDaoDynamodb, teardownTestPageDaoDynamodb)
	defer teardownTest(t)
	dao := initPageDaoDynamodb(testAdc)
	gsiName := strings.ReplaceAll(awsDynamodbGSIPagePos, "{tableName}", testDynamodbTablePage)
	err := dynamodbInitGsi(testAdc, testDynamodbTablePage, gsiName,
		[]prom.AwsDynamodbNameAndType{{Name: PageFieldTopicId, Type: prom.AwsAttrTypeString}, {Name: PageFieldPosition, Type: prom.AwsAttrTypeNumber}},
		[]prom.AwsDynamodbNameAndType{{Name: PageFieldTopicId, Type: prom.AwsKeyTypePartition}, {Name: PageFieldPosition, Type: prom.AwsKeyTypeSort}})
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName+"/dynamodbInitGsi", err)
	}
	doTestPageDaoGetAll(t, testName, dao)
}

func TestPageDaoDynamodb_GetN(t *testing.T) {
	testName := "TestPageDaoDynamodb_GetN"
	teardownTest := setupTest(t, testName, setupTestPageDaoDynamodb, teardownTestPageDaoDynamodb)
	defer teardownTest(t)
	dao := initPageDaoDynamodb(testAdc)
	gsiName := strings.ReplaceAll(awsDynamodbGSIPagePos, "{tableName}", testDynamodbTablePage)
	err := dynamodbInitGsi(testAdc, testDynamodbTablePage, gsiName,
		[]prom.AwsDynamodbNameAndType{{Name: PageFieldTopicId, Type: prom.AwsAttrTypeString}, {Name: PageFieldPosition, Type: prom.AwsAttrTypeNumber}},
		[]prom.AwsDynamodbNameAndType{{Name: PageFieldTopicId, Type: prom.AwsKeyTypePartition}, {Name: PageFieldPosition, Type: prom.AwsKeyTypeSort}})
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName+"/dynamodbInitGsi", err)
	}
	doTestPageDaoGetN(t, testName, dao)
}
