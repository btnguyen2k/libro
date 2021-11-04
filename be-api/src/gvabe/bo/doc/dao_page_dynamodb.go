package doc

import (
	"strings"

	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

const (
	awsDynamodbGSIPagePos = "gsi_{tableName}_pos"
)

// NewPageDaoDynamodb is helper method to create AWS DynamoDB-implementation of PageDao
func NewPageDaoDynamodb(adc *prom.AwsDynamodbConnect, tableName string) PageDao {
	dao := &BasePageDaoImpl{}
	spec := &henge.DynamodbDaoSpec{}
	innerDao := henge.NewUniversalDaoDynamodb(adc, tableName, spec)
	innerDao.MapGsi(strings.ReplaceAll(awsDynamodbGSIPagePos, "{tableName}", tableName), PageFieldPosition)
	dao.UniversalDao = innerDao
	return dao
}
