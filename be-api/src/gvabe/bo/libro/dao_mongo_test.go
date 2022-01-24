package libro

import (
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/btnguyen2k/prom"
)

const (
	envMongoDb  = "MONGO_DB"
	envMongoUrl = "MONGO_URL"
)

var testMc *prom.MongoConnect

func mongoInitCollection(mc *prom.MongoConnect, collection string) error {
	mc.GetCollection(collection).Drop(nil)
	return nil
}

func newMongoConnect(t *testing.T, testName string) (*prom.MongoConnect, error) {
	rand.Seed(time.Now().UnixNano())
	db := strings.Trim(os.Getenv(envMongoDb), "\"")
	url := strings.Trim(os.Getenv(envMongoUrl), "\"")
	if db == "" || url == "" {
		t.Skipf("%s skipped", testName)
	}
	return prom.NewMongoConnect(url, db, 10000)
}
