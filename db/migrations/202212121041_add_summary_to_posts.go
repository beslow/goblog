package migrations

import (
	"fmt"
	"os"

	"github.com/beslow/goblog/initialize"
	"github.com/beslow/goblog/models"
)

func addSummaryToPosts() {
	timestamp := "202212121041"

	var migration models.Migration
	initialize.DB.Where("timestamp = ?", timestamp).Find(&migration)

	if migration.Timestamp == "" {
		if err := initialize.DB.Exec(`ALTER TABLE posts ADD COLUMN summary VARCHAR(255) AFTER category_id;`).Error; err != nil {
			fmt.Printf("add category_id to posts fail: %#v\n", err)
			os.Exit(1)
		}

		initialize.DB.Create(&models.Migration{Timestamp: timestamp})
	}
}
