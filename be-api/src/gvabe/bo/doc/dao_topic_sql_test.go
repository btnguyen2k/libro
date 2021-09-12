package doc

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/godror/godror"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/mattn/go-sqlite3"
)

const (
	testSqlTableTopic = "test_topic"
)

func sqlInitTableTopic(sqlc *prom.SqlConnect, table string) error {
	rand.Seed(time.Now().UnixNano())
	var err error
	if sqlc.GetDbFlavor() == prom.FlavorCosmosDb {
		_, err = sqlc.GetDB().Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s WITH MAXRU=10000", cosmosdbDbName))
		if err != nil {
			return err
		}
	}
	sqlc.GetDB().Exec(fmt.Sprintf("DROP TABLE %s", table))
	switch sqlc.GetDbFlavor() {
	case prom.FlavorCosmosDb:
		spec := &henge.CosmosdbCollectionSpec{Pk: henge.CosmosdbColId}
		err = henge.InitCosmosdbCollection(sqlc, table, spec)
	case prom.FlavorMySql:
		err = henge.InitMysqlTable(sqlc, table, map[string]string{TopicColProductId: "VARCHAR(32)"})
	case prom.FlavorPgSql:
		err = henge.InitPgsqlTable(sqlc, table, map[string]string{TopicColProductId: "VARCHAR(32)"})
	case prom.FlavorSqlite:
		err = henge.InitSqliteTable(sqlc, table, map[string]string{TopicColProductId: "VARCHAR(32)"})
	}
	return err
}

func initTopicDaoSql(sqlc *prom.SqlConnect) TopicDao {
	if sqlc.GetDbFlavor() == prom.FlavorCosmosDb {
		return NewTopicDaoCosmosdb(sqlc, testSqlTableTopic, true)
	}
	return NewTopicDaoSql(sqlc, testSqlTableTopic, true)
}

/*----------------------------------------------------------------------*/

func TestNewTopicDaoSql(t *testing.T) {
	name := "TestNewTopicDaoSql"
	urlMap := sqlGetUrlFromEnv()
	if len(urlMap) == 0 {
		t.Skipf("%s skipped", name)
	}
	for dbtype, info := range urlMap {
		sqlc, err := initSqlConnect(t, name, dbtype, info)
		if err != nil {
			t.Fatalf("%s failed: error [%s]", name+"/"+dbtype, err)
		} else if sqlc == nil {
			t.Fatalf("%s failed: nil", name+"/"+dbtype)
		}
		err = sqlInitTableTopic(sqlc, testSqlTableTopic)
		if err != nil {
			t.Fatalf("%s failed: error [%s]", name+"/sqlInitTableTopic/"+dbtype, err)
		}
		dao := initTopicDaoSql(sqlc)
		if dao == nil {
			t.Fatalf("%s failed: nil", name+"/initTopicDaoSql")
		}
		sqlc.Close()
	}
}

func TestTopicDaoSql_CreateGet(t *testing.T) {
	name := "TestTopicDaoSql_CreateGet"
	urlMap := sqlGetUrlFromEnv()
	if len(urlMap) == 0 {
		t.Skipf("%s skipped", name)
	}
	for dbtype, info := range urlMap {
		sqlc, err := initSqlConnect(t, name, dbtype, info)
		if err != nil {
			t.Fatalf("%s failed: error [%s]", name+"/"+dbtype, err)
		} else if sqlc == nil {
			t.Fatalf("%s failed: nil", name+"/"+dbtype)
		}
		err = sqlInitTableTopic(sqlc, testSqlTableTopic)
		if err != nil {
			t.Fatalf("%s failed: error [%s]", name+"/sqlInitTableTopic/"+dbtype, err)
		}
		dao := initTopicDaoSql(sqlc)
		if dao == nil {
			t.Fatalf("%s failed: nil", name)
		}
		if sqlc.GetDbFlavor() == prom.FlavorSqlite {
			henge.TimeLayout = "2006-01-02 15:04:05Z07:00"
		} else {
			henge.TimeLayout = time.RFC3339
		}
		doTestTopicDaoCreateGet(t, name, dao)
		sqlc.Close()
	}
}

func TestTopicDaoSql_CreateUpdateGet(t *testing.T) {
	name := "TestTopicDaoSql_CreateUpdateGet"
	urlMap := sqlGetUrlFromEnv()
	if len(urlMap) == 0 {
		t.Skipf("%s skipped", name)
	}
	for dbtype, info := range urlMap {
		sqlc, err := initSqlConnect(t, name, dbtype, info)
		if err != nil {
			t.Fatalf("%s failed: error [%s]", name+"/"+dbtype, err)
		} else if sqlc == nil {
			t.Fatalf("%s failed: nil", name+"/"+dbtype)
		}
		err = sqlInitTableTopic(sqlc, testSqlTableTopic)
		if err != nil {
			t.Fatalf("%s failed: error [%s]", name+"/sqlInitTableTopic/"+dbtype, err)
		}
		dao := initTopicDaoSql(sqlc)
		if dao == nil {
			t.Fatalf("%s failed: nil", name)
		}
		doTestTopicDaoCreateUpdateGet(t, name, dao)
		sqlc.Close()
	}
}

func TestTopicDaoSql_CreateDelete(t *testing.T) {
	name := "TestTopicDaoSql_CreateDelete"
	urlMap := sqlGetUrlFromEnv()
	if len(urlMap) == 0 {
		t.Skipf("%s skipped", name)
	}
	for dbtype, info := range urlMap {
		sqlc, err := initSqlConnect(t, name, dbtype, info)
		if err != nil {
			t.Fatalf("%s failed: error [%s]", name+"/"+dbtype, err)
		} else if sqlc == nil {
			t.Fatalf("%s failed: nil", name+"/"+dbtype)
		}
		err = sqlInitTableTopic(sqlc, testSqlTableTopic)
		if err != nil {
			t.Fatalf("%s failed: error [%s]", name+"/sqlInitTableTopic/"+dbtype, err)
		}
		dao := initTopicDaoSql(sqlc)
		if dao == nil {
			t.Fatalf("%s failed: nil", name)
		}
		doTestTopicDaoCreateDelete(t, name, dao)
		sqlc.Close()
	}
}

func TestTopicDaoSql_GetAll(t *testing.T) {
	name := "TestTopicDaoSql_GetAll"
	urlMap := sqlGetUrlFromEnv()
	if len(urlMap) == 0 {
		t.Skipf("%s skipped", name)
	}
	for dbtype, info := range urlMap {
		sqlc, err := initSqlConnect(t, name, dbtype, info)
		if err != nil {
			t.Fatalf("%s failed: error [%s]", name+"/"+dbtype, err)
		} else if sqlc == nil {
			t.Fatalf("%s failed: nil", name+"/"+dbtype)
		}
		err = sqlInitTableTopic(sqlc, testSqlTableTopic)
		if err != nil {
			t.Fatalf("%s failed: error [%s]", name+"/sqlInitTableTopic/"+dbtype, err)
		}
		dao := initTopicDaoSql(sqlc)
		if dao == nil {
			t.Fatalf("%s failed: nil", name)
		}
		doTestTopicDaoGetAll(t, name, dao)
		sqlc.Close()
	}
}

func TestTopicDaoSql_GetN(t *testing.T) {
	name := "TestTopicDaoSql_GetN"
	urlMap := sqlGetUrlFromEnv()
	if len(urlMap) == 0 {
		t.Skipf("%s skipped", name)
	}
	for dbtype, info := range urlMap {
		sqlc, err := initSqlConnect(t, name, dbtype, info)
		if err != nil {
			t.Fatalf("%s failed: error [%s]", name+"/"+dbtype, err)
		} else if sqlc == nil {
			t.Fatalf("%s failed: nil", name+"/"+dbtype)
		}
		err = sqlInitTableTopic(sqlc, testSqlTableTopic)
		if err != nil {
			t.Fatalf("%s failed: error [%s]", name+"/sqlInitTableTopic/"+dbtype, err)
		}
		dao := initTopicDaoSql(sqlc)
		if dao == nil {
			t.Fatalf("%s failed: nil", name)
		}
		doTestTopicDaoGetN(t, name, dao)
		sqlc.Close()
	}
}
