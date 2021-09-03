package doc

import (
	"testing"

	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

const (
	testDynamodbTable = "test_topic"
)

func initTopicDaoDynamodb(adc *prom.AwsDynamodbConnect) TopicDao {
	return NewTopicDaoDynamodb(adc, testDynamodbTable)
}

/*----------------------------------------------------------------------*/

func TestNewTopicDaoDynamodb(t *testing.T) {
	name := "TestNewTopicDaoDynamodb"
	adc, err := newDynamodbConnect(t, name)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name, err)
	} else if adc == nil {
		t.Fatalf("%s failed: nil", name)
	}
	spec := &henge.DynamodbTablesSpec{MainTableRcu: 2, MainTableWcu: 1, CreateUidxTable: true, UidxTableRcu: 2, UidxTableWcu: 1}
	err = dynamodbInitTable(adc, testDynamodbTable, spec)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name+"/dynamodbInitTable", err)
	}
	dao := initTopicDaoDynamodb(adc)
	if dao == nil {
		t.Fatalf("%s failed: nil", name+"/initTopicDaoDynamodb")
	}
	defer adc.Close()
}

func TestTopicDaoDynamodb_CreateGet(t *testing.T) {
	name := "TestTopicDaoDynamodb_CreateGet"
	adc, err := newDynamodbConnect(t, name)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name, err)
	} else if adc == nil {
		t.Fatalf("%s failed: nil", name)
	}
	spec := &henge.DynamodbTablesSpec{MainTableRcu: 2, MainTableWcu: 1, CreateUidxTable: true, UidxTableRcu: 2, UidxTableWcu: 1}
	err = dynamodbInitTable(adc, testDynamodbTable, spec)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name+"/dynamodbInitTable", err)
	}
	dao := initTopicDaoDynamodb(adc)
	if dao == nil {
		t.Fatalf("%s failed: nil", name+"/initTopicDaoDynamodb")
	}
	defer adc.Close()
	doTestTopicDaoCreateGet(t, name, dao)
}

func TestTopicDaoDynamodb_CreateUpdateGet(t *testing.T) {
	name := "TestTopicDaoDynamodb_CreateUpdateGet"
	adc, err := newDynamodbConnect(t, name)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name, err)
	} else if adc == nil {
		t.Fatalf("%s failed: nil", name)
	}
	spec := &henge.DynamodbTablesSpec{MainTableRcu: 2, MainTableWcu: 1, CreateUidxTable: true, UidxTableRcu: 2, UidxTableWcu: 1}
	err = dynamodbInitTable(adc, testDynamodbTable, spec)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name+"/dynamodbInitTable", err)
	}
	dao := initTopicDaoDynamodb(adc)
	if dao == nil {
		t.Fatalf("%s failed: nil", name+"/initTopicDaoDynamodb")
	}
	defer adc.Close()
	doTestTopicDaoCreateUpdateGet(t, name, dao)
}

func TestTopicDaoDynamodb_CreateDelete(t *testing.T) {
	name := "TestTopicDaoDynamodb_CreateDelete"
	adc, err := newDynamodbConnect(t, name)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name, err)
	} else if adc == nil {
		t.Fatalf("%s failed: nil", name)
	}
	spec := &henge.DynamodbTablesSpec{MainTableRcu: 2, MainTableWcu: 1, CreateUidxTable: true, UidxTableRcu: 2, UidxTableWcu: 1}
	err = dynamodbInitTable(adc, testDynamodbTable, spec)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name+"/dynamodbInitTable", err)
	}
	dao := initTopicDaoDynamodb(adc)
	if dao == nil {
		t.Fatalf("%s failed: nil", name+"/initTopicDaoDynamodb")
	}
	defer adc.Close()
	doTestTopicDaoCreateDelete(t, name, dao)
}

func TestTopicDaoDynamodb_GetAll(t *testing.T) {
	name := "TestTopicDaoDynamodb_GetAll"
	adc, err := newDynamodbConnect(t, name)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name, err)
	} else if adc == nil {
		t.Fatalf("%s failed: nil", name)
	}
	spec := &henge.DynamodbTablesSpec{MainTableRcu: 2, MainTableWcu: 1, CreateUidxTable: true, UidxTableRcu: 2, UidxTableWcu: 1}
	err = dynamodbInitTable(adc, testDynamodbTable, spec)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name+"/dynamodbInitTable", err)
	}
	dao := initTopicDaoDynamodb(adc)
	if dao == nil {
		t.Fatalf("%s failed: nil", name+"/initTopicDaoDynamodb")
	}
	defer adc.Close()
	doTestTopicDaoGetAll(t, name, dao)
}

func TestTopicDaoDynamodb_GetN(t *testing.T) {
	name := "TestTopicDaoDynamodb_GetN"
	adc, err := newDynamodbConnect(t, name)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name, err)
	} else if adc == nil {
		t.Fatalf("%s failed: nil", name)
	}
	spec := &henge.DynamodbTablesSpec{MainTableRcu: 2, MainTableWcu: 1, CreateUidxTable: true, UidxTableRcu: 2, UidxTableWcu: 1}
	err = dynamodbInitTable(adc, testDynamodbTable, spec)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name+"/dynamodbInitTable", err)
	}
	dao := initTopicDaoDynamodb(adc)
	if dao == nil {
		t.Fatalf("%s failed: nil", name+"/initTopicDaoDynamodb")
	}
	defer adc.Close()
	doTestTopicDaoGetN(t, name, dao)
}
