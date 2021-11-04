package doc

import (
	"strings"
	"testing"

	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

const (
	testDynamodbTablePage = "test_page"
)

func initPageDaoDynamodb(adc *prom.AwsDynamodbConnect) PageDao {
	return NewPageDaoDynamodb(adc, testDynamodbTablePage)
}

/*----------------------------------------------------------------------*/

func TestNewPageDaoDynamodb(t *testing.T) {
	name := "TestNewPageDaoDynamodb"
	adc, err := newDynamodbConnect(t, name)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name, err)
	} else if adc == nil {
		t.Fatalf("%s failed: nil", name)
	}
	spec := &henge.DynamodbTablesSpec{MainTableRcu: 2, MainTableWcu: 1, CreateUidxTable: true, UidxTableRcu: 2, UidxTableWcu: 1}
	err = dynamodbInitTable(adc, testDynamodbTablePage, spec)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name+"/dynamodbInitTable", err)
	}
	dao := initPageDaoDynamodb(adc)
	if dao == nil {
		t.Fatalf("%s failed: nil", name+"/initPageDaoDynamodb")
	}
	defer adc.Close()
}

func TestPageDaoDynamodb_CreateGet(t *testing.T) {
	name := "TestPageDaoDynamodb_CreateGet"
	adc, err := newDynamodbConnect(t, name)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name, err)
	} else if adc == nil {
		t.Fatalf("%s failed: nil", name)
	}
	spec := &henge.DynamodbTablesSpec{MainTableRcu: 2, MainTableWcu: 1, CreateUidxTable: true, UidxTableRcu: 2, UidxTableWcu: 1}
	err = dynamodbInitTable(adc, testDynamodbTablePage, spec)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name+"/dynamodbInitTable", err)
	}
	dao := initPageDaoDynamodb(adc)
	if dao == nil {
		t.Fatalf("%s failed: nil", name+"/initPageDaoDynamodb")
	}
	defer adc.Close()
	doTestPageDaoCreateGet(t, name, dao)
}

func TestPageDaoDynamodb_CreateUpdateGet(t *testing.T) {
	name := "TestPageDaoDynamodb_CreateUpdateGet"
	adc, err := newDynamodbConnect(t, name)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name, err)
	} else if adc == nil {
		t.Fatalf("%s failed: nil", name)
	}
	spec := &henge.DynamodbTablesSpec{MainTableRcu: 2, MainTableWcu: 1, CreateUidxTable: true, UidxTableRcu: 2, UidxTableWcu: 1}
	err = dynamodbInitTable(adc, testDynamodbTablePage, spec)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name+"/dynamodbInitTable", err)
	}
	dao := initPageDaoDynamodb(adc)
	if dao == nil {
		t.Fatalf("%s failed: nil", name+"/initPageDaoDynamodb")
	}
	defer adc.Close()
	doTestPageDaoCreateUpdateGet(t, name, dao)
}

func TestPageDaoDynamodb_CreateDelete(t *testing.T) {
	name := "TestPageDaoDynamodb_CreateDelete"
	adc, err := newDynamodbConnect(t, name)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name, err)
	} else if adc == nil {
		t.Fatalf("%s failed: nil", name)
	}
	spec := &henge.DynamodbTablesSpec{MainTableRcu: 2, MainTableWcu: 1, CreateUidxTable: true, UidxTableRcu: 2, UidxTableWcu: 1}
	err = dynamodbInitTable(adc, testDynamodbTablePage, spec)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name+"/dynamodbInitTable", err)
	}
	dao := initPageDaoDynamodb(adc)
	if dao == nil {
		t.Fatalf("%s failed: nil", name+"/initPageDaoDynamodb")
	}
	defer adc.Close()
	doTestPageDaoCreateDelete(t, name, dao)
}

func TestPageDaoDynamodb_GetAll(t *testing.T) {
	name := "TestPageDaoDynamodb_GetAll"
	adc, err := newDynamodbConnect(t, name)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name, err)
	} else if adc == nil {
		t.Fatalf("%s failed: nil", name)
	}
	spec := &henge.DynamodbTablesSpec{MainTableRcu: 2, MainTableWcu: 1, CreateUidxTable: true, UidxTableRcu: 2, UidxTableWcu: 1}
	err = dynamodbInitTable(adc, testDynamodbTablePage, spec)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name+"/dynamodbInitTable", err)
	}
	gsiName := strings.ReplaceAll(awsDynamodbGSIPagePos, "{tableName}", testDynamodbTablePage)
	err = dynamodbInitGsi(adc, testDynamodbTablePage, gsiName,
		[]prom.AwsDynamodbNameAndType{{Name: PageFieldTopicId, Type: prom.AwsAttrTypeString}, {Name: PageFieldPosition, Type: prom.AwsAttrTypeNumber}},
		[]prom.AwsDynamodbNameAndType{{Name: PageFieldTopicId, Type: prom.AwsKeyTypePartition}, {Name: PageFieldPosition, Type: prom.AwsKeyTypeSort}})
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name+"/dynamodbInitGsi", err)
	}
	dao := initPageDaoDynamodb(adc)
	if dao == nil {
		t.Fatalf("%s failed: nil", name+"/initPageDaoDynamodb")
	}
	defer adc.Close()
	doTestPageDaoGetAll(t, name, dao)
}

func TestPageDaoDynamodb_GetN(t *testing.T) {
	name := "TestPageDaoDynamodb_GetN"
	adc, err := newDynamodbConnect(t, name)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name, err)
	} else if adc == nil {
		t.Fatalf("%s failed: nil", name)
	}
	spec := &henge.DynamodbTablesSpec{MainTableRcu: 2, MainTableWcu: 1, CreateUidxTable: true, UidxTableRcu: 2, UidxTableWcu: 1}
	err = dynamodbInitTable(adc, testDynamodbTablePage, spec)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name+"/dynamodbInitTable", err)
	}
	gsiName := strings.ReplaceAll(awsDynamodbGSIPagePos, "{tableName}", testDynamodbTablePage)
	err = dynamodbInitGsi(adc, testDynamodbTablePage, gsiName,
		[]prom.AwsDynamodbNameAndType{{Name: PageFieldTopicId, Type: prom.AwsAttrTypeString}, {Name: PageFieldPosition, Type: prom.AwsAttrTypeNumber}},
		[]prom.AwsDynamodbNameAndType{{Name: PageFieldTopicId, Type: prom.AwsKeyTypePartition}, {Name: PageFieldPosition, Type: prom.AwsKeyTypeSort}})
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name+"/dynamodbInitGsi", err)
	}
	dao := initPageDaoDynamodb(adc)
	if dao == nil {
		t.Fatalf("%s failed: nil", name+"/initPageDaoDynamodb")
	}
	defer adc.Close()
	doTestPageDaoGetN(t, name, dao)
}
