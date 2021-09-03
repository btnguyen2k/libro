package doc

import (
	"os"
	"strings"
	"testing"

	"github.com/btnguyen2k/prom"
)

const (
	testMongoCollection = "test_topic"
)

func initTopicDaoMongo(mc *prom.MongoConnect) TopicDao {
	return NewTopicDaoMongo(mc, testMongoCollection, strings.Index(mc.GetUrl(), "replicaSet=") >= 0)
}

/*----------------------------------------------------------------------*/

func TestNewTopicDaoMongo(t *testing.T) {
	name := "TestNewTopicDaoMongo"
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
	dao := initTopicDaoMongo(mc)
	if dao == nil {
		t.Fatalf("%s failed: nil", name+"/initTopicDaoMongo")
	}
	mc.Close(nil)
}

func TestTopicDaoMongo_CreateGet(t *testing.T) {
	name := "TestTopicDaoMongo_CreateGet"
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
	dao := initTopicDaoMongo(mc)
	doTestTopicDaoCreateGet(t, name, dao)
	mc.Close(nil)
}

func TestTopicDaoMongo_CreateUpdateGet(t *testing.T) {
	name := "TestTopicDaoMongo_CreateGet"
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
	dao := initTopicDaoMongo(mc)
	doTestTopicDaoCreateUpdateGet(t, name, dao)
	mc.Close(nil)
}

func TestTopicDaoMongo_CreateDelete(t *testing.T) {
	name := "TestTopicDaoMongo_CreateDelete"
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
	dao := initTopicDaoMongo(mc)
	doTestTopicDaoCreateDelete(t, name, dao)
	mc.Close(nil)
}

func TestTopicDaoMongo_GetAll(t *testing.T) {
	name := "TestTopicDaoMongo_GetAll"
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
	dao := initTopicDaoMongo(mc)
	doTestTopicDaoGetAll(t, name, dao)
	mc.Close(nil)
}

func TestTopicDaoMongo_GetN(t *testing.T) {
	name := "TestTopicDaoMongo_GetN"
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
	dao := initTopicDaoMongo(mc)
	doTestTopicDaoGetN(t, name, dao)
	mc.Close(nil)
}