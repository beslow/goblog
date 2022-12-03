package migrations

import (
	"fmt"

	"github.com/beslow/goblog/db"
	"github.com/beslow/goblog/models"
)

func createConsts() {
	tableName := "consts"
	timestamp := "202212031243"

	var migration models.Migration
	db.DB.Where("timestamp = ?", timestamp).Find(&migration)

	if migration.Timestamp == "" && !existTable(tableName) {
		db.DB.Exec(`CREATE TABLE ` + tableName + ` (
			id int(11) unsigned NOT NULL AUTO_INCREMENT,
			name varchar(100) NOT NULL DEFAULT '' COMMENT '常量名称',
			description varchar(500) DEFAULT NULL COMMENT '常量描述',
			value varchar(1000) DEFAULT NULL COMMENT '常量值',
			PRIMARY KEY (id)
		  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`)

		db.DB.Create(&models.Migration{Timestamp: timestamp})

		fmt.Printf("create table %s\n", tableName)
	}
}
