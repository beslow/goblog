package migrations

import (
	"fmt"
	"os"

	"github.com/beslow/goblog/initialize"
	"github.com/beslow/goblog/models"
)

func Migrate() {
	createMigrations()
	createConsts()
	createHistoryWorks()
	createHistoryEducations()
	createGoAdminTables()
	createPosts()
	createComments()
	createCategories()
	addCategoryIDToPosts()
	addSummaryToPosts()
}

func createMigrations() {
	if !existTable("migrations") {
		initialize.DB.Exec(`CREATE TABLE migrations (
			timestamp varchar(100) NOT NULL DEFAULT ''
		  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`)
	}
}

func existTable(table string) bool {
	var existTable int
	raw := `SELECT EXISTS (
		SELECT 
			TABLE_NAME
		FROM 
		information_schema.TABLES 
		WHERE 
		TABLE_SCHEMA LIKE '` + initialize.DefaultConfig.Name + `' AND 
			TABLE_TYPE LIKE 'BASE TABLE' AND
			TABLE_NAME = '` + table + `'
		)`
	if err := initialize.DB.Raw(raw).Scan(&existTable).Error; err != nil {
		fmt.Printf("check table's existence fail: %#v\n", err)
		os.Exit(1)
	}

	return existTable == 1
}

// create table and mark the migration's timestamp
func createTableWithTimestamp(table, createSql, timestamp string) {
	var migration models.Migration
	if err := initialize.DB.Where("timestamp = ?", timestamp).Find(&migration).Error; err != nil {
		fmt.Printf("query migration fail: %#v\n", err)
		os.Exit(1)
	}

	if migration.Timestamp == "" {
		createTable(table, createSql)

		if err := initialize.DB.Create(&models.Migration{Timestamp: timestamp}).Error; err != nil {
			fmt.Printf("add migration timestamp fail: %#v\n", err)
			os.Exit(1)
		}
	}
}

func createTable(table, createSql string) {
	if !existTable(table) {
		if err := initialize.DB.Exec(createSql).Error; err != nil {
			fmt.Printf("create %s fail: %#v\n", table, err)
			os.Exit(1)
		}

		fmt.Printf("create table %s successfully\n", table)
	}
}
