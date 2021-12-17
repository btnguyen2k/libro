package libro

import (
	"testing"

	"github.com/btnguyen2k/prom"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/godror/godror"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/mattn/go-sqlite3"
)

const (
	testSqlTableProduct = "test_product"
)

func initProductDaoSql(sqlc *prom.SqlConnect) ProductDao {
	if sqlc.GetDbFlavor() == prom.FlavorCosmosDb {
		return NewProductDaoCosmosdb(sqlc, testSqlTableProduct, true)
	}
	return NewProductDaoSql(sqlc, testSqlTableProduct, true)
}

var setupTestProductDaoSql = func(t *testing.T, testName string) {
	var err error
	testSqlc, err = initSqlConnect(t, testName, testDbType, testDbInfo)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName+"/"+testDbType+"/initSqlConnect", err)
	} else if testSqlc == nil {
		t.Fatalf("%s failed: nil", testName+"/"+testDbType+"/initSqlConnect")
	}
	err = sqlInitTable(testSqlc, testSqlTableProduct)
	if err != nil {
		t.Fatalf("%s failed: error [%s]", testName+"/"+testDbType+"/sqlInitTable", err)
	}
}

var teardownTestProductDaoSql = func(t *testing.T, testName string) {
	if testSqlc != nil {
		testSqlc.Close()
		defer func() { testSqlc = nil }()
	}
}

/*----------------------------------------------------------------------*/

func TestNewProductDaoSql(t *testing.T) {
	testName := "TestNewProductDaoSql"
	urlMap := sqlGetUrlFromEnv()
	if len(urlMap) == 0 {
		t.Skipf("%s skipped", testName)
	}
	for testDbType, testDbInfo = range urlMap {
		t.Run(testDbType, func(t *testing.T) {
			teardownTest := setupTest(t, testName, setupTestProductDaoSql, teardownTestProductDaoSql)
			defer teardownTest(t)

			dao := initProductDaoSql(testSqlc)
			if dao == nil {
				t.Fatalf("%s failed: nil", testName+"/"+testDbType+"/initProductDaoSql")
			}
		})
	}
}

func TestProductDaoSql_CreateGet(t *testing.T) {
	testName := "TestProductDaoSql_CreateGet"
	urlMap := sqlGetUrlFromEnv()
	if len(urlMap) == 0 {
		t.Skipf("%s skipped", testName)
	}
	for testDbType, testDbInfo = range urlMap {
		t.Run(testDbType, func(t *testing.T) {
			teardownTest := setupTest(t, testName, setupTestProductDaoSql, teardownTestProductDaoSql)
			defer teardownTest(t)

			dao := initProductDaoSql(testSqlc)
			// if sqlc.GetDbFlavor() == prom.FlavorSqlite {
			// 	henge.TimeLayout = "2006-01-02 15:04:05Z07:00"
			// } else {
			// 	henge.TimeLayout = time.RFC3339
			// }
			doTestProductDaoCreateGet(t, testName+"/"+testDbType, dao)
		})
	}
}

func TestProductDaoSql_CreateUpdateGet(t *testing.T) {
	testName := "TestProductDaoSql_CreateUpdateGet"
	urlMap := sqlGetUrlFromEnv()
	if len(urlMap) == 0 {
		t.Skipf("%s skipped", testName)
	}
	for testDbType, testDbInfo = range urlMap {
		t.Run(testDbType, func(t *testing.T) {
			teardownTest := setupTest(t, testName, setupTestProductDaoSql, teardownTestProductDaoSql)
			defer teardownTest(t)

			dao := initProductDaoSql(testSqlc)
			doTestProductDaoCreateUpdateGet(t, testName+"/"+testDbType, dao)
		})
	}
}

func TestProductDaoSql_CreateDelete(t *testing.T) {
	testName := "TestProductDaoSql_CreateDelete"
	urlMap := sqlGetUrlFromEnv()
	if len(urlMap) == 0 {
		t.Skipf("%s skipped", testName)
	}
	for testDbType, testDbInfo = range urlMap {
		t.Run(testDbType, func(t *testing.T) {
			teardownTest := setupTest(t, testName, setupTestProductDaoSql, teardownTestProductDaoSql)
			defer teardownTest(t)

			dao := initProductDaoSql(testSqlc)
			doTestProductDaoCreateDelete(t, testName+"/"+testDbType, dao)
		})
	}
}

func TestProductDaoSql_GetAll(t *testing.T) {
	testName := "TestProductDaoSql_GetAll"
	urlMap := sqlGetUrlFromEnv()
	if len(urlMap) == 0 {
		t.Skipf("%s skipped", testName)
	}
	for testDbType, testDbInfo = range urlMap {
		t.Run(testDbType, func(t *testing.T) {
			teardownTest := setupTest(t, testName, setupTestProductDaoSql, teardownTestProductDaoSql)
			defer teardownTest(t)

			dao := initProductDaoSql(testSqlc)
			doTestProductDaoGetAll(t, testName+"/"+testDbType, dao)
		})
	}
}

func TestProductDaoSql_GetN(t *testing.T) {
	testName := "TestProductDaoSql_GetN"
	urlMap := sqlGetUrlFromEnv()
	if len(urlMap) == 0 {
		t.Skipf("%s skipped", testName)
	}
	for testDbType, testDbInfo = range urlMap {
		t.Run(testDbType, func(t *testing.T) {
			teardownTest := setupTest(t, testName, setupTestProductDaoSql, teardownTestProductDaoSql)
			defer teardownTest(t)

			dao := initProductDaoSql(testSqlc)
			doTestProductDaoGetN(t, testName+"/"+testDbType, dao)
		})
	}
}
