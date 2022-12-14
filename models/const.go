package models

import (
	"github.com/beslow/goblog/initialize"
)

type Const struct {
	ID          int
	Name        string
	Description string
	Value       string
}

func GetConst(name string) string {
	var constRecord Const

	initialize.DB.Where("name = ?", name).Find(&constRecord)

	return constRecord.Value
}
