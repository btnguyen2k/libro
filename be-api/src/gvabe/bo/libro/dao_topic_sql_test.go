package libro

import (
	"fmt"
	"testing"

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

func initTopicDaoSql(sqlc *prom.SqlConnect) TopicDao {
	if sqlc.GetDbFlavor() == prom.FlavorCosmosDb {
		return NewTopicDaoCosmosdb(sqlc, testSqlTableTopic, true)
	}
	return NewTopicDaoSql(sqlc, testSqlTableTopic, true)
}

var setupTestTopicDaoSql = func(t *testing.T, testName string) {
	var err error
	testSqlc, err = initSqlConnect(t, testName, testDbType, testDbInfo)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName+"/"+testDbType+"/initSqlConnect", err)
	} else if testSqlc == nil {
		t.Fatalf("%s failed: nil", testName+"/"+testDbType+"/initSqlConnect")
	}
	err = sqlInitTable(testSqlc, testSqlTableTopic)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName+"/"+testDbType+"/sqlInitTable", err)
	}
	err = InitTopicTableSql(testSqlc, testSqlTableTopic)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName+"/"+testDbType+"/InitTopicTableSql", err)
	}
}

var teardownTestTopicDaoSql = func(t *testing.T, testName string) {
	if testSqlc != nil {
		testSqlc.GetDB().Exec(fmt.Sprintf("DROP TABLE %s", testSqlTableTopic))
		testSqlc.Close()
		defer func() { testSqlc = nil }()
	}
}

/*----------------------------------------------------------------------*/

func TestNewTopicDaoSql(t *testing.T) {
	testName := "TestNewTopicDaoSql"
	urlMap := sqlGetUrlFromEnv()
	if len(urlMap) == 0 {
		t.Skipf("%s skipped", testName)
	}
	for testDbType, testDbInfo = range urlMap {
		t.Run(testDbType, func(t *testing.T) {
			teardownTest := setupTest(t, testName, setupTestTopicDaoSql, teardownTestTopicDaoSql)
			defer teardownTest(t)
			dao := initTopicDaoSql(testSqlc)
			if dao == nil {
				t.Fatalf("%s failed: nil", testName+"/initTopicDaoSql")
			}
		})
	}
}

func TestTopicDaoSql_CreateGet(t *testing.T) {
	testName := "TestTopicDaoSql_CreateGet"
	urlMap := sqlGetUrlFromEnv()
	if len(urlMap) == 0 {
		t.Skipf("%s skipped", testName)
	}
	for testDbType, testDbInfo = range urlMap {
		t.Run(testDbType, func(t *testing.T) {
			teardownTest := setupTest(t, testName, setupTestTopicDaoSql, teardownTestTopicDaoSql)
			defer teardownTest(t)
			dao := initTopicDaoSql(testSqlc)
			doTestTopicDaoCreateGet(t, testName, dao)
		})
	}
}

func TestTopicDaoSql_CreateUpdateGet(t *testing.T) {
	testName := "TestTopicDaoSql_CreateUpdateGet"
	urlMap := sqlGetUrlFromEnv()
	if len(urlMap) == 0 {
		t.Skipf("%s skipped", testName)
	}
	for testDbType, testDbInfo = range urlMap {
		t.Run(testDbType, func(t *testing.T) {
			teardownTest := setupTest(t, testName, setupTestTopicDaoSql, teardownTestTopicDaoSql)
			defer teardownTest(t)
			dao := initTopicDaoSql(testSqlc)
			doTestTopicDaoCreateUpdateGet(t, testName, dao)
		})
	}
}

func TestTopicDaoSql_CreateDelete(t *testing.T) {
	testName := "TestTopicDaoSql_CreateDelete"
	urlMap := sqlGetUrlFromEnv()
	if len(urlMap) == 0 {
		t.Skipf("%s skipped", testName)
	}
	for testDbType, testDbInfo = range urlMap {
		t.Run(testDbType, func(t *testing.T) {
			teardownTest := setupTest(t, testName, setupTestTopicDaoSql, teardownTestTopicDaoSql)
			defer teardownTest(t)
			dao := initTopicDaoSql(testSqlc)
			doTestTopicDaoCreateDelete(t, testName, dao)
		})
	}
}

func TestTopicDaoSql_GetAll(t *testing.T) {
	testName := "TestTopicDaoSql_GetAll"
	urlMap := sqlGetUrlFromEnv()
	if len(urlMap) == 0 {
		t.Skipf("%s skipped", testName)
	}
	for testDbType, testDbInfo = range urlMap {
		t.Run(testDbType, func(t *testing.T) {
			teardownTest := setupTest(t, testName, setupTestTopicDaoSql, teardownTestTopicDaoSql)
			defer teardownTest(t)
			dao := initTopicDaoSql(testSqlc)
			doTestTopicDaoGetAll(t, testName, dao)
		})
	}
}

func TestTopicDaoSql_GetN(t *testing.T) {
	testName := "TestTopicDaoSql_GetN"
	urlMap := sqlGetUrlFromEnv()
	if len(urlMap) == 0 {
		t.Skipf("%s skipped", testName)
	}
	for testDbType, testDbInfo = range urlMap {
		t.Run(testDbType, func(t *testing.T) {
			teardownTest := setupTest(t, testName, setupTestTopicDaoSql, teardownTestTopicDaoSql)
			defer teardownTest(t)
			dao := initTopicDaoSql(testSqlc)
			doTestTopicDaoGetN(t, testName, dao)
		})
	}
}
