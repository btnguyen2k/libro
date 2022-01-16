package libro

import (
	"fmt"

	"github.com/btnguyen2k/prom"
)

func _sqlFlavorStr(flavor prom.DbFlavor) string {
	switch flavor {
	case prom.FlavorCosmosDb:
		return "Azure Cosmos DB"
	case prom.FlavorMySql:
		return "MySQL"
	case prom.FlavorPgSql:
		return "PostgreSQL"
	case prom.FlavorSqlite:
		return "SQLite"
	}
	return fmt.Sprintf("%d", flavor)
}
