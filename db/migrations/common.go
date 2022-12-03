package migrations

import (
	"github.com/beslow/goblog/config"
	"github.com/beslow/goblog/db"
)

func Migrate() {
	createMigrations()
	createConsts()
}

func createMigrations() {
	if !existTable("migrations") {
		db.DB.Exec(`CREATE TABLE migrations (
			timestamp varchar(100) NOT NULL DEFAULT ''
		  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`)
	}
}

func existTable(table string) bool {
	var existTable int
	db.DB.Raw(`SELECT EXISTS (
		SELECT 
			TABLE_NAME
		FROM 
		information_schema.TABLES 
		WHERE 
		TABLE_SCHEMA LIKE '` + config.MySQL.DBName + `' AND 
			TABLE_TYPE LIKE 'BASE TABLE' AND
			TABLE_NAME = '` + table + `'
		)`).Scan(&existTable)

	return existTable == 1
}
