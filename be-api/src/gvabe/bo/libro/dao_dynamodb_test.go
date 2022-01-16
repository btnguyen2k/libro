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
	tblStatus, err := adc.GetTableStatus(nil, table)
	for ; err == nil && strings.ToUpper(tblStatus) != status; tblStatus, err = adc.GetTableStatus(nil, table) {
		if time.Now().Sub(t).Milliseconds() > timeout.Milliseconds() {
			err = errors.New("timed-out")
		}
	}
	return err
}

func _dynamodbWaitForGSIStatus(adc *prom.AwsDynamodbConnect, table, gsi, status string, timeout time.Duration) error {
	t := time.Now()
	gsiStatus, err := adc.GetGlobalSecondaryIndexStatus(nil, table, gsi)
	for ; err == nil && strings.ToUpper(gsiStatus) != status; gsiStatus, err = adc.GetGlobalSecondaryIndexStatus(nil, table, gsi) {
		if time.Now().Sub(t).Milliseconds() > timeout.Milliseconds() {
			err = errors.New("timed-out")
		}
	}
	return err
}

func dynamodbInitTable(adc *prom.AwsDynamodbConnect, table string) error {
	for _, tblName := range []string{table, table + henge.AwsDynamodbUidxTableSuffix} {
		adc.DeleteTable(nil, tblName)
		if err := _dynamodbWaitForTableStatus(adc, tblName, "", 10*time.Second); err != nil {
			return err
		}
	}
	return nil
}

func dynamodbInitGsi(adc *prom.AwsDynamodbConnect, tableName, gsiName string, attrsDef []prom.AwsDynamodbNameAndType, keyAttrs []prom.AwsDynamodbNameAndType) error {
	if status, err := adc.GetGlobalSecondaryIndexStatus(nil, tableName, gsiName); err != nil {
		return err
	} else {
		if status == "" {
			adc.CreateGlobalSecondaryIndex(nil, tableName, gsiName, 1, 1, attrsDef, keyAttrs)
		}
		if err := _dynamodbWaitForGSIStatus(adc, tableName, gsiName, "ACTIVE", 10*time.Second); err != nil {
			return err
		}
	}
	return nil
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
