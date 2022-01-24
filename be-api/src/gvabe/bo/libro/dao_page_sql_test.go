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
	testSqlTablePage = "test_page"
)

// func sqlInitTablePage(sqlc *prom.SqlConnect, tableName string) error {
// 	rand.Seed(time.Now().UnixNano())
// 	if sqlc.GetDbFlavor() == prom.FlavorCosmosDb {
// 		_, err := sqlc.GetDB().Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s WITH MAXRU=10000", cosmosdbDbName))
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	sqlc.GetDB().Exec(fmt.Sprintf("DROP TABLE %s", tableName))
// 	return InitPageTableSql(sqlc, tableName)
// }

func initPageDaoSql(sqlc *prom.SqlConnect) PageDao {
	if sqlc.GetDbFlavor() == prom.FlavorCosmosDb {
		return NewPageDaoCosmosdb(sqlc, testSqlTablePage, true)
	}
	return NewPageDaoSql(sqlc, testSqlTablePage, true)
}

var setupTestPageDaoSql = func(t *testing.T, testName string) {
	var err error
	testSqlc, err = initSqlConnect(t, testName, testDbType, testDbInfo)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName+"/"+testDbType+"/initSqlConnect", err)
	} else if testSqlc == nil {
		t.Fatalf("%s failed: nil", testName+"/"+testDbType+"/initSqlConnect")
	}
	err = sqlInitTable(testSqlc, testSqlTablePage)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName+"/"+testDbType+"/sqlInitTable", err)
	}
	err = InitPageTableSql(testSqlc, testSqlTablePage)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName+"/"+testDbType+"/InitPageTableSql", err)
	}
}

var teardownTestPageDaoSql = func(t *testing.T, testName string) {
	if testSqlc != nil {
		testSqlc.GetDB().Exec(fmt.Sprintf("DROP TABLE %s", testSqlTablePage))
		testSqlc.Close()
		defer func() { testSqlc = nil }()
	}
}

/*----------------------------------------------------------------------*/

func TestNewPageDaoSql(t *testing.T) {
	testName := "TestNewPageDaoSql"
	urlMap := sqlGetUrlFromEnv()
	if len(urlMap) == 0 {
		t.Skipf("%s skipped", testName)
	}
	for testDbType, testDbInfo = range urlMap {
		t.Run(testDbType, func(t *testing.T) {
			teardownTest := setupTest(t, testName, setupTestPageDaoSql, teardownTestPageDaoSql)
			defer teardownTest(t)
			dao := initPageDaoSql(testSqlc)
			if dao == nil {
				t.Fatalf("%s failed: nil", testName+"/initPageDaoSql")
			}
		})
	}
}

func TestPageDaoSql_CreateGet(t *testing.T) {
	testName := "TestPageDaoSql_CreateGet"
	urlMap := sqlGetUrlFromEnv()
	if len(urlMap) == 0 {
		t.Skipf("%s skipped", testName)
	}
	for testDbType, testDbInfo = range urlMap {
		t.Run(testDbType, func(t *testing.T) {
			teardownTest := setupTest(t, testName, setupTestPageDaoSql, teardownTestPageDaoSql)
			defer teardownTest(t)
			dao := initPageDaoSql(testSqlc)
			doTestPageDaoCreateGet(t, testName, dao)
		})
	}
}

func TestPageDaoSql_CreateUpdateGet(t *testing.T) {
	testName := "TestPageDaoSql_CreateUpdateGet"
	urlMap := sqlGetUrlFromEnv()
	if len(urlMap) == 0 {
		t.Skipf("%s skipped", testName)
	}
	for testDbType, testDbInfo = range urlMap {
		t.Run(testDbType, func(t *testing.T) {
			teardownTest := setupTest(t, testName, setupTestPageDaoSql, teardownTestPageDaoSql)
			defer teardownTest(t)
			dao := initPageDaoSql(testSqlc)
			doTestPageDaoCreateUpdateGet(t, testName, dao)
		})
	}
}

func TestPageDaoSql_CreateDelete(t *testing.T) {
	testName := "TestPageDaoSql_CreateDelete"
	urlMap := sqlGetUrlFromEnv()
	if len(urlMap) == 0 {
		t.Skipf("%s skipped", testName)
	}
	for testDbType, testDbInfo = range urlMap {
		t.Run(testDbType, func(t *testing.T) {
			teardownTest := setupTest(t, testName, setupTestPageDaoSql, teardownTestPageDaoSql)
			defer teardownTest(t)
			dao := initPageDaoSql(testSqlc)
			doTestPageDaoCreateDelete(t, testName, dao)
		})
	}
}

func TestPageDaoSql_GetAll(t *testing.T) {
	testName := "TestPageDaoSql_GetAll"
	urlMap := sqlGetUrlFromEnv()
	if len(urlMap) == 0 {
		t.Skipf("%s skipped", testName)
	}
	for testDbType, testDbInfo = range urlMap {
		t.Run(testDbType, func(t *testing.T) {
			teardownTest := setupTest(t, testName, setupTestPageDaoSql, teardownTestPageDaoSql)
			defer teardownTest(t)
			dao := initPageDaoSql(testSqlc)
			doTestPageDaoGetAll(t, testName, dao)
		})
	}
}

func TestPageDaoSql_GetN(t *testing.T) {
	testName := "TestPageDaoSql_GetN"
	urlMap := sqlGetUrlFromEnv()
	if len(urlMap) == 0 {
		t.Skipf("%s skipped", testName)
	}
	for testDbType, testDbInfo = range urlMap {
		t.Run(testDbType, func(t *testing.T) {
			teardownTest := setupTest(t, testName, setupTestPageDaoSql, teardownTestPageDaoSql)
			defer teardownTest(t)
			dao := initPageDaoSql(testSqlc)
			doTestPageDaoGetN(t, testName, dao)
		})
	}
}
