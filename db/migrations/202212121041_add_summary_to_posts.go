package migrations

import (
	"fmt"
	"os"

	"github.com/beslow/goblog/db"
	"github.com/beslow/goblog/models"
)

func addSummaryToPosts() {
	timestamp := "202212121041"

	var migration models.Migration
	db.DB.Where("timestamp = ?", timestamp).Find(&migration)

	if migration.Timestamp == "" {
		if err := db.DB.Exec(`ALTER TABLE posts ADD COLUMN summary VARCHAR(255) AFTER category_id;`).Error; err != nil {
			fmt.Printf("add category_id to posts fail: %#v\n", err)
			os.Exit(1)
		}

		db.DB.Create(&models.Migration{Timestamp: timestamp})
	}
}
