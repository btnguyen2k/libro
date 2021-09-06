package respicite

import (
	"strings"
	"testing"

	"github.com/btnguyen2k/prom"
)

const (
	envMongoDb  = "MONGO_DB"
	envMongoUrl = "MONGO_URL"
)

func newMongoConnect(t *testing.T, testName string, db, url string) (*prom.MongoConnect, error) {
	db = strings.Trim(db, "\"")
	url = strings.Trim(url, "\"")
	if db == "" || url == "" {
		t.Skipf("%s skipped", testName)
	}
	return prom.NewMongoConnect(url, db, 10000)
}
