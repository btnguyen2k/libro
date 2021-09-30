package doc

import (
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

const (
	envMongoDb  = "MONGO_DB"
	envMongoUrl = "MONGO_URL"
)

func mongoInitCollection(mc *prom.MongoConnect, collection string) error {
	mc.GetCollection(collection).Drop(nil)
	return henge.InitMongoCollection(mc, collection)
}

func newMongoConnect(t *testing.T, testName string, db, url string) (*prom.MongoConnect, error) {
	rand.Seed(time.Now().UnixNano())
	db = strings.Trim(db, "\"")
	url = strings.Trim(url, "\"")
	if db == "" || url == "" {
		t.Skipf("%s skipped", testName)
	}
	return prom.NewMongoConnect(url, db, 10000)
}
