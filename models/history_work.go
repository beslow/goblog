package models

import (
	"time"

	"github.com/beslow/goblog/initialize"
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
	initialize.DB.Find(&historyWorks)

	return historyWorks
}
