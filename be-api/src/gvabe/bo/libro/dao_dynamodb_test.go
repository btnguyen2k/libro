package libro

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
	envAwsRegion          = "AWS_REGION"
	envAwsAccessKeyId     = "AWS_ACCESS_KEY_ID"
	envAwsSecretAccessKey = "AWS_SECRET_ACCESS_KEY"
	envDynamodbEndpoint   = "AWS_DYNAMODB_ENDPOINT"
)

var testAdc *prom.AwsDynamodbConnect

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

func dynamodbInitGsi(adc *prom.AwsDynamodbConnect, table, gsi string, attrsDef []prom.AwsDynamodbNameAndType, keyAttrs []prom.AwsDynamodbNameAndType) error {
	return adc.CreateGlobalSecondaryIndex(nil, table, gsi, 1, 1, attrsDef, keyAttrs)
}

func newDynamodbConnect(t *testing.T, testName string) (*prom.AwsDynamodbConnect, error) {
	rand.Seed(time.Now().UnixNano())
	awsRegion := strings.ReplaceAll(os.Getenv(envAwsRegion), `"`, "")
	awsAccessKeyId := strings.ReplaceAll(os.Getenv(envAwsAccessKeyId), `"`, "")
	awsSecretAccessKey := strings.ReplaceAll(os.Getenv(envAwsSecretAccessKey), `"`, "")
	if awsRegion == "" || awsAccessKeyId == "" || awsSecretAccessKey == "" {
		t.Skipf("%s skipped", testName)
	}
	cfg := &aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewEnvCredentials(),
	}
	if awsDynamodbEndpoint := strings.ReplaceAll(os.Getenv(envDynamodbEndpoint), `"`, ""); awsDynamodbEndpoint != "" {
		cfg.Endpoint = aws.String(awsDynamodbEndpoint)
		if strings.HasPrefix(awsDynamodbEndpoint, "http://") {
			cfg.DisableSSL = aws.Bool(true)
		}
	}
	return prom.NewAwsDynamodbConnect(cfg, nil, nil, 10000)
}
