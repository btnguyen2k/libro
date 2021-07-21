package app

import (
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
)

const (
	testMongoCollection = "test_app"
)

func mongoInitCollection(mc *prom.MongoConnect, collection string) error {
	rand.Seed(time.Now().UnixNano())
	mc.GetCollection(collection).Drop(nil)
	return henge.InitMongoCollection(mc, collection)
}

func newMongoConnect(t *testing.T, testName string, db, url string) (*prom.MongoConnect, error) {
	db = strings.Trim(db, "\"")
	url = strings.Trim(url, "\"")
	if db == "" || url == "" {
		t.Skipf("%s skipped", testName)
	}
	return prom.NewMongoConnect(url, db, 10000)
}

func initDaoMongo(mc *prom.MongoConnect) AppDao {
	return NewAppDaoMongo(mc, testMongoCollection, strings.Index(mc.GetUrl(), "replicaSet=") >= 0)
}

const (
	envMongoDb  = "MONGO_DB"
	envMongoUrl = "MONGO_URL"
)

/*----------------------------------------------------------------------*/

func TestNewAppDaoMongo(t *testing.T) {
	name := "TestNewAppDaoMongo"
	db := os.Getenv(envMongoDb)
	url := os.Getenv(envMongoUrl)
	mc, err := newMongoConnect(t, name, db, url)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name, err)
	} else if mc == nil {
		t.Fatalf("%s failed: nil", name)
	}
	err = mongoInitCollection(mc, testMongoCollection)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name+"/mongoInitCollection", err)
	}
	dao := initDaoMongo(mc)
	if dao == nil {
		t.Fatalf("%s failed: nil", name+"/initDaoMongo")
	}
	mc.Close(nil)
}

func TestAppDaoMongo_CreateGet(t *testing.T) {
	name := "TestAppDaoMongo_CreateGet"
	db := os.Getenv(envMongoDb)
	url := os.Getenv(envMongoUrl)
	mc, err := newMongoConnect(t, name, db, url)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name, err)
	} else if mc == nil {
		t.Fatalf("%s failed: nil", name)
	}
	err = mongoInitCollection(mc, testMongoCollection)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name+"/mongoInitCollection", err)
	}
	dao := initDaoMongo(mc)
	doTestAppDaoCreateGet(t, name, dao)
	mc.Close(nil)
}

func TestAppDaoMongo_CreateUpdateGet(t *testing.T) {
	name := "TestAppDaoMongo_CreateGet"
	db := os.Getenv(envMongoDb)
	url := os.Getenv(envMongoUrl)
	mc, err := newMongoConnect(t, name, db, url)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name, err)
	} else if mc == nil {
		t.Fatalf("%s failed: nil", name)
	}
	err = mongoInitCollection(mc, testMongoCollection)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name+"/mongoInitCollection", err)
	}
	dao := initDaoMongo(mc)
	doTestAppDaoCreateUpdateGet(t, name, dao)
	mc.Close(nil)
}

func TestAppDaoMongo_CreateDelete(t *testing.T) {
	name := "TestAppDaoMongo_CreateDelete"
	db := os.Getenv(envMongoDb)
	url := os.Getenv(envMongoUrl)
	mc, err := newMongoConnect(t, name, db, url)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name, err)
	} else if mc == nil {
		t.Fatalf("%s failed: nil", name)
	}
	err = mongoInitCollection(mc, testMongoCollection)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name+"/mongoInitCollection", err)
	}
	dao := initDaoMongo(mc)
	doTestAppDaoCreateDelete(t, name, dao)
	mc.Close(nil)
}

func TestAppDaoMongo_GetAll(t *testing.T) {
	name := "TestAppDaoMongo_GetAll"
	db := os.Getenv(envMongoDb)
	url := os.Getenv(envMongoUrl)
	mc, err := newMongoConnect(t, name, db, url)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name, err)
	} else if mc == nil {
		t.Fatalf("%s failed: nil", name)
	}
	err = mongoInitCollection(mc, testMongoCollection)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name+"/mongoInitCollection", err)
	}
	dao := initDaoMongo(mc)
	doTestAppDaoGetAll(t, name, dao)
	mc.Close(nil)
}

func TestAppDaoMongo_GetN(t *testing.T) {
	name := "TestAppDaoMongo_GetN"
	db := os.Getenv(envMongoDb)
	url := os.Getenv(envMongoUrl)
	mc, err := newMongoConnect(t, name, db, url)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name, err)
	} else if mc == nil {
		t.Fatalf("%s failed: nil", name)
	}
	err = mongoInitCollection(mc, testMongoCollection)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name+"/mongoInitCollection", err)
	}
	dao := initDaoMongo(mc)
	doTestAppDaoGetN(t, name, dao)
	mc.Close(nil)
}
