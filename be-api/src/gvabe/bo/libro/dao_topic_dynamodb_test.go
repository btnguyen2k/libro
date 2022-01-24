package libro

import (
	"strings"
	"testing"

	"github.com/btnguyen2k/prom"
)

const (
	testDynamodbTableTopic = "test_topic"
)

func initTopicDaoDynamodb(adc *prom.AwsDynamodbConnect) TopicDao {
	return NewTopicDaoDynamodb(adc, testDynamodbTableTopic)
}

var setupTestTopicDaoDynamodb = func(t *testing.T, testName string) {
	var err error
	testAdc, err = newDynamodbConnect(t, testName)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName, err)
	} else if testAdc == nil {
		t.Fatalf("%s failed: nil", testName)
	}

	err = dynamodbInitTable(testAdc, testDynamodbTableTopic)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName+"/dynamodbInitTable", err)
	}
	err = InitTopicTableDynamodb(testAdc, testDynamodbTableTopic)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName+"/"+testDbType+"/InitTopicTableDynamodb", err)
	}
}

var teardownTestTopicDaoDynamodb = func(t *testing.T, testName string) {
	if testAdc != nil {
		testAdc.Close()
		defer func() { testAdc = nil }()
	}
}

/*----------------------------------------------------------------------*/

func TestNewTopicDaoDynamodb(t *testing.T) {
	testName := "TestNewTopicDaoDynamodb"
	teardownTest := setupTest(t, testName, setupTestTopicDaoDynamodb, teardownTestTopicDaoDynamodb)
	defer teardownTest(t)
	dao := initTopicDaoDynamodb(testAdc)
	if dao == nil {
		t.Fatalf("%s failed: nil", testName+"/initTopicDaoDynamodb")
	}
}

func TestTopicDaoDynamodb_CreateGet(t *testing.T) {
	testName := "TestTopicDaoDynamodb_CreateGet"
	teardownTest := setupTest(t, testName, setupTestTopicDaoDynamodb, teardownTestTopicDaoDynamodb)
	defer teardownTest(t)
	dao := initTopicDaoDynamodb(testAdc)
	doTestTopicDaoCreateGet(t, testName, dao)
}

func TestTopicDaoDynamodb_CreateUpdateGet(t *testing.T) {
	testName := "TestTopicDaoDynamodb_CreateUpdateGet"
	teardownTest := setupTest(t, testName, setupTestTopicDaoDynamodb, teardownTestTopicDaoDynamodb)
	defer teardownTest(t)
	dao := initTopicDaoDynamodb(testAdc)
	doTestTopicDaoCreateUpdateGet(t, testName, dao)
}

func TestTopicDaoDynamodb_CreateDelete(t *testing.T) {
	testName := "TestTopicDaoDynamodb_CreateDelete"
	teardownTest := setupTest(t, testName, setupTestTopicDaoDynamodb, teardownTestTopicDaoDynamodb)
	defer teardownTest(t)
	dao := initTopicDaoDynamodb(testAdc)
	doTestTopicDaoCreateDelete(t, testName, dao)
}

func TestTopicDaoDynamodb_GetAll(t *testing.T) {
	testName := "TestTopicDaoDynamodb_GetAll"
	teardownTest := setupTest(t, testName, setupTestTopicDaoDynamodb, teardownTestTopicDaoDynamodb)
	defer teardownTest(t)
	gsiName := strings.ReplaceAll(awsDynamodbGSITopicPos, "{tableName}", testDynamodbTableTopic)
	err := dynamodbInitGsi(testAdc, testDynamodbTableTopic, gsiName,
		[]prom.AwsDynamodbNameAndType{{Name: TopicFieldProductId, Type: prom.AwsAttrTypeString}, {Name: TopicFieldPosition, Type: prom.AwsAttrTypeNumber}},
		[]prom.AwsDynamodbNameAndType{{Name: TopicFieldProductId, Type: prom.AwsKeyTypePartition}, {Name: TopicFieldPosition, Type: prom.AwsKeyTypeSort}})
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName+"/dynamodbInitGsi", err)
	}
	dao := initTopicDaoDynamodb(testAdc)
	doTestTopicDaoGetAll(t, testName, dao)
}

func TestTopicDaoDynamodb_GetN(t *testing.T) {
	testName := "TestTopicDaoDynamodb_GetN"
	teardownTest := setupTest(t, testName, setupTestTopicDaoDynamodb, teardownTestTopicDaoDynamodb)
	defer teardownTest(t)
	gsiName := strings.ReplaceAll(awsDynamodbGSITopicPos, "{tableName}", testDynamodbTableTopic)
	err := dynamodbInitGsi(testAdc, testDynamodbTableTopic, gsiName,
		[]prom.AwsDynamodbNameAndType{{Name: TopicFieldProductId, Type: prom.AwsAttrTypeString}, {Name: TopicFieldPosition, Type: prom.AwsAttrTypeNumber}},
		[]prom.AwsDynamodbNameAndType{{Name: TopicFieldProductId, Type: prom.AwsKeyTypePartition}, {Name: TopicFieldPosition, Type: prom.AwsKeyTypeSort}})
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName+"/dynamodbInitGsi", err)
	}
	dao := initTopicDaoDynamodb(testAdc)
	doTestTopicDaoGetN(t, testName, dao)
}
