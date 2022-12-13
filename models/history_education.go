package models

import (
	"time"

	"github.com/beslow/goblog/db"
)

type HistoryEducation struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Course      string
	City        string
	Description string
	FromTo      string
	Sort        int
}

func GetAllHistoryEducations() []HistoryEducation {
	var historyEducations []HistoryEducation
	db.DB.Find(&historyEducations)

	return historyEducations
}
