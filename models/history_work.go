package models

import (
	"time"

	"github.com/beslow/goblog/db"
)

type HistoryWork struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Company     string
	JobTitle    string
	Description string
	FromTo      string
	Sort        int
}

func GetAllHistoryWorks() []HistoryWork {
	var historyWorks []HistoryWork
	db.DB.Find(&historyWorks)

	return historyWorks
}
