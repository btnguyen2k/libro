package doc

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
	testDynamodbTable = "test_section"
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

func initDaoDynamodb(adc *prom.AwsDynamodbConnect) SectionDao {
	return NewSectionDaoDynamodb(adc, testDynamodbTable)
}

/*----------------------------------------------------------------------*/

func TestNewSectionDaoDynamodb(t *testing.T) {
	name := "TestNewSectionDaoDynamodb"
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

func TestSectionDaoDynamodb_CreateGet(t *testing.T) {
	name := "TestSectionDaoDynamodb_CreateGet"
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
	doTestSectionDaoCreateGet(t, name, dao)
}

func TestSectionDaoDynamodb_CreateUpdateGet(t *testing.T) {
	name := "TestSectionDaoDynamodb_CreateUpdateGet"
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
	doTestSectionDaoCreateUpdateGet(t, name, dao)
}

func TestSectionDaoDynamodb_CreateDelete(t *testing.T) {
	name := "TestSectionDaoDynamodb_CreateDelete"
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
	doTestSectionDaoCreateDelete(t, name, dao)
}

func TestSectionDaoDynamodb_GetAll(t *testing.T) {
	name := "TestSectionDaoDynamodb_GetAll"
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
	doTestSectionDaoGetAll(t, name, dao)
}

func TestSectionDaoDynamodb_GetN(t *testing.T) {
	name := "TestSectionDaoDynamodb_GetN"
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
	doTestSectionDaoGetN(t, name, dao)
}