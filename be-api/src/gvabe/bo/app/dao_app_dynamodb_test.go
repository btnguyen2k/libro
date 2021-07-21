package app

import (
	"errors"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

const (
	testDynamodbTable = "test_app"
)

func _dynamodbWaitForTableStatus(adc *prom.AwsDynamodbConnect, table, status string, timeout time.Duration) error {
	t := time.Now()
	for tblStatus, err := adc.GetTableStatus(nil, table); ; {
		if err != nil {
			return err
		}
		if strings.ToUpper(tblStatus) == status {
			return nil
		}
		if time.Now().Sub(t).Milliseconds() > timeout.Milliseconds() {
			return errors.New("")
		}
	}
}

func dynamodbInitTable(adc *prom.AwsDynamodbConnect, table string, spec *henge.DynamodbTablesSpec) error {
	rand.Seed(time.Now().UnixNano())
	adc.DeleteTable(nil, table)
	if err := _dynamodbWaitForTableStatus(adc, table, "", 10*time.Second); err != nil {
		return err
	}
	if spec.CreateUidxTable {
		adc.DeleteTable(nil, table+henge.AwsDynamodbUidxTableSuffix)
		if err := _dynamodbWaitForTableStatus(adc, table+henge.AwsDynamodbUidxTableSuffix, "", 10*time.Second); err != nil {
			return err
		}
	}
	return henge.InitDynamodbTables(adc, table, spec)
}

func newDynamodbConnect(t *testing.T, testName string) (*prom.AwsDynamodbConnect, error) {
	awsRegion := strings.ReplaceAll(os.Getenv("AWS_REGION"), `"`, "")
	awsAccessKeyId := strings.ReplaceAll(os.Getenv("AWS_ACCESS_KEY_ID"), `"`, "")
	awsSecretAccessKey := strings.ReplaceAll(os.Getenv("AWS_SECRET_ACCESS_KEY"), `"`, "")
	if awsRegion == "" || awsAccessKeyId == "" || awsSecretAccessKey == "" {
		t.Skipf("%s skipped", testName)
	}
	cfg := &aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewEnvCredentials(),
	}
	if awsDynamodbEndpoint := strings.ReplaceAll(os.Getenv("AWS_DYNAMODB_ENDPOINT"), `"`, ""); awsDynamodbEndpoint != "" {
		cfg.Endpoint = aws.String(awsDynamodbEndpoint)
		if strings.HasPrefix(awsDynamodbEndpoint, "http://") {
			cfg.DisableSSL = aws.Bool(true)
		}
	}
	return prom.NewAwsDynamodbConnect(cfg, nil, nil, 10000)
}

func initDaoDynamodb(adc *prom.AwsDynamodbConnect) AppDao {
	return NewAppDaoDynamodb(adc, testDynamodbTable)
}

/*----------------------------------------------------------------------*/

func TestNewAppDaoDynamodb(t *testing.T) {
	name := "TestNewAppDaoDynamodb"
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
	dao := initDaoDynamodb(adc)
	if dao == nil {
		t.Fatalf("%s failed: nil", name+"/initDaoDynamodb")
	}
	defer adc.Close()
}

func TestAppDaoDynamodb_CreateGet(t *testing.T) {
	name := "TestAppDaoDynamodb_CreateGet"
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
	dao := initDaoDynamodb(adc)
	if dao == nil {
		t.Fatalf("%s failed: nil", name+"/initDaoDynamodb")
	}
	defer adc.Close()
	doTestAppDaoCreateGet(t, name, dao)
}

func TestAppDaoDynamodb_CreateUpdateGet(t *testing.T) {
	name := "TestAppDaoDynamodb_CreateUpdateGet"
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
	dao := initDaoDynamodb(adc)
	if dao == nil {
		t.Fatalf("%s failed: nil", name+"/initDaoDynamodb")
	}
	defer adc.Close()
	doTestAppDaoCreateUpdateGet(t, name, dao)
}

func TestAppDaoDynamodb_CreateDelete(t *testing.T) {
	name := "TestAppDaoDynamodb_CreateDelete"
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
	dao := initDaoDynamodb(adc)
	if dao == nil {
		t.Fatalf("%s failed: nil", name+"/initDaoDynamodb")
	}
	defer adc.Close()
	doTestAppDaoCreateDelete(t, name, dao)
}

func TestAppDaoDynamodb_GetAll(t *testing.T) {
	name := "TestAppDaoDynamodb_GetAll"
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
	dao := initDaoDynamodb(adc)
	if dao == nil {
		t.Fatalf("%s failed: nil", name+"/initDaoDynamodb")
	}
	defer adc.Close()
	doTestAppDaoGetAll(t, name, dao)
}

func TestAppDaoDynamodb_GetN(t *testing.T) {
	name := "TestAppDaoDynamodb_GetN"
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
	dao := initDaoDynamodb(adc)
	if dao == nil {
		t.Fatalf("%s failed: nil", name+"/initDaoDynamodb")
	}
	defer adc.Close()
	doTestAppDaoGetN(t, name, dao)
}
