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
	testSqlTablePage = "test_page"
)

func sqlInitTablePage(sqlc *prom.SqlConnect, tableName string) error {
	rand.Seed(time.Now().UnixNano())
	if sqlc.GetDbFlavor() == prom.FlavorCosmosDb {
		_, err := sqlc.GetDB().Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s WITH MAXRU=10000", cosmosdbDbName))
		if err != nil {
			return err
		}
	}
	sqlc.GetDB().Exec(fmt.Sprintf("DROP TABLE %s", tableName))
	return CreateSqlTableForPages(sqlc, tableName)
}

func initPageDaoSql(sqlc *prom.SqlConnect) PageDao {
	if sqlc.GetDbFlavor() == prom.FlavorCosmosDb {
		return NewPageDaoCosmosdb(sqlc, testSqlTablePage, true)
	}
	return NewPageDaoSql(sqlc, testSqlTablePage, true)
}

/*----------------------------------------------------------------------*/

func TestNewPageDaoSql(t *testing.T) {
	name := "TestNewPageDaoSql"
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
		err = sqlInitTablePage(sqlc, testSqlTablePage)
		if err != nil {
			t.Fatalf("%s failed: error [%s]", name+"/sqlInitTablePage/"+dbtype, err)
		}
		dao := initPageDaoSql(sqlc)
		if dao == nil {
			t.Fatalf("%s failed: nil", name+"/initPageDaoSql")
		}
		sqlc.Close()
	}
}

func TestPageDaoSql_CreateGet(t *testing.T) {
	name := "TestPageDaoSql_CreateGet"
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
		err = sqlInitTablePage(sqlc, testSqlTablePage)
		if err != nil {
			t.Fatalf("%s failed: error [%s]", name+"/sqlInitTablePage/"+dbtype, err)
		}
		dao := initPageDaoSql(sqlc)
		if dao == nil {
			t.Fatalf("%s failed: nil", name)
		}
		if sqlc.GetDbFlavor() == prom.FlavorSqlite {
			henge.TimeLayout = "2006-01-02 15:04:05Z07:00"
		} else {
			henge.TimeLayout = time.RFC3339
		}
		doTestPageDaoCreateGet(t, name, dao)
		sqlc.Close()
	}
}

func TestPageDaoSql_CreateUpdateGet(t *testing.T) {
	name := "TestPageDaoSql_CreateUpdateGet"
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
		err = sqlInitTablePage(sqlc, testSqlTablePage)
		if err != nil {
			t.Fatalf("%s failed: error [%s]", name+"/sqlInitTablePage/"+dbtype, err)
		}
		dao := initPageDaoSql(sqlc)
		if dao == nil {
			t.Fatalf("%s failed: nil", name)
		}
		doTestPageDaoCreateUpdateGet(t, name, dao)
		sqlc.Close()
	}
}

func TestPageDaoSql_CreateDelete(t *testing.T) {
	name := "TestPageDaoSql_CreateDelete"
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
		err = sqlInitTablePage(sqlc, testSqlTablePage)
		if err != nil {
			t.Fatalf("%s failed: error [%s]", name+"/sqlInitTablePage/"+dbtype, err)
		}
		dao := initPageDaoSql(sqlc)
		if dao == nil {
			t.Fatalf("%s failed: nil", name)
		}
		doTestPageDaoCreateDelete(t, name, dao)
		sqlc.Close()
	}
}

func TestPageDaoSql_GetAll(t *testing.T) {
	name := "TestPageDaoSql_GetAll"
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
		err = sqlInitTablePage(sqlc, testSqlTablePage)
		if err != nil {
			t.Fatalf("%s failed: error [%s]", name+"/sqlInitTablePage/"+dbtype, err)
		}
		dao := initPageDaoSql(sqlc)
		if dao == nil {
			t.Fatalf("%s failed: nil", name)
		}
		doTestPageDaoGetAll(t, name, dao)
		sqlc.Close()
	}
}

func TestPageDaoSql_GetN(t *testing.T) {
	name := "TestPageDaoSql_GetN"
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
		err = sqlInitTablePage(sqlc, testSqlTablePage)
		if err != nil {
			t.Fatalf("%s failed: error [%s]", name+"/sqlInitTablePage/"+dbtype, err)
		}
		dao := initPageDaoSql(sqlc)
		if dao == nil {
			t.Fatalf("%s failed: nil", name)
		}
		doTestPageDaoGetN(t, name, dao)
		sqlc.Close()
	}
}
