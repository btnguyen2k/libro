package libro

import (
	"os"
	"strings"
	"testing"

	"github.com/btnguyen2k/prom"
)

const (
	testMongoCollectionPage = "test_page"
)

func initPageDaoMongo(mc *prom.MongoConnect) PageDao {
	return NewPageDaoMongo(mc, testMongoCollectionPage, strings.Index(mc.GetUrl(), "replicaSet=") >= 0)
}

/*----------------------------------------------------------------------*/

func TestNewPageDaoMongo(t *testing.T) {
	name := "TestNewPageDaoMongo"
	db := os.Getenv(envMongoDb)
	url := os.Getenv(envMongoUrl)
	mc, err := newMongoConnect(t, name, db, url)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name, err)
	} else if mc == nil {
		t.Fatalf("%s failed: nil", name)
	}
	err = mongoInitCollection(mc, testMongoCollectionPage)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name+"/mongoInitCollection", err)
	}
	dao := initPageDaoMongo(mc)
	if dao == nil {
		t.Fatalf("%s failed: nil", name+"/initPageDaoMongo")
	}
	mc.Close(nil)
}

func TestPageDaoMongo_CreateGet(t *testing.T) {
	name := "TestPageDaoMongo_CreateGet"
	db := os.Getenv(envMongoDb)
	url := os.Getenv(envMongoUrl)
	mc, err := newMongoConnect(t, name, db, url)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name, err)
	} else if mc == nil {
		t.Fatalf("%s failed: nil", name)
	}
	err = mongoInitCollection(mc, testMongoCollectionPage)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name+"/mongoInitCollection", err)
	}
	dao := initPageDaoMongo(mc)
	doTestPageDaoCreateGet(t, name, dao)
	mc.Close(nil)
}

func TestPageDaoMongo_CreateUpdateGet(t *testing.T) {
	name := "TestPageDaoMongo_CreateGet"
	db := os.Getenv(envMongoDb)
	url := os.Getenv(envMongoUrl)
	mc, err := newMongoConnect(t, name, db, url)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name, err)
	} else if mc == nil {
		t.Fatalf("%s failed: nil", name)
	}
	err = mongoInitCollection(mc, testMongoCollectionPage)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name+"/mongoInitCollection", err)
	}
	dao := initPageDaoMongo(mc)
	doTestPageDaoCreateUpdateGet(t, name, dao)
	mc.Close(nil)
}

func TestPageDaoMongo_CreateDelete(t *testing.T) {
	name := "TestPageDaoMongo_CreateDelete"
	db := os.Getenv(envMongoDb)
	url := os.Getenv(envMongoUrl)
	mc, err := newMongoConnect(t, name, db, url)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name, err)
	} else if mc == nil {
		t.Fatalf("%s failed: nil", name)
	}
	err = mongoInitCollection(mc, testMongoCollectionPage)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name+"/mongoInitCollection", err)
	}
	dao := initPageDaoMongo(mc)
	doTestPageDaoCreateDelete(t, name, dao)
	mc.Close(nil)
}

func TestPageDaoMongo_GetAll(t *testing.T) {
	name := "TestPageDaoMongo_GetAll"
	db := os.Getenv(envMongoDb)
	url := os.Getenv(envMongoUrl)
	mc, err := newMongoConnect(t, name, db, url)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name, err)
	} else if mc == nil {
		t.Fatalf("%s failed: nil", name)
	}
	err = mongoInitCollection(mc, testMongoCollectionPage)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name+"/mongoInitCollection", err)
	}
	dao := initPageDaoMongo(mc)
	doTestPageDaoGetAll(t, name, dao)
	mc.Close(nil)
}

func TestPageDaoMongo_GetN(t *testing.T) {
	name := "TestPageDaoMongo_GetN"
	db := os.Getenv(envMongoDb)
	url := os.Getenv(envMongoUrl)
	mc, err := newMongoConnect(t, name, db, url)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name, err)
	} else if mc == nil {
		t.Fatalf("%s failed: nil", name)
	}
	err = mongoInitCollection(mc, testMongoCollectionPage)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", name+"/mongoInitCollection", err)
	}
	dao := initPageDaoMongo(mc)
	doTestPageDaoGetN(t, name, dao)
	mc.Close(nil)
}
