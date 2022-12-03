package consts

import "github.com/beslow/goblog/db"

type Const struct {
	Name        string
	Description string
	Value       string
}

func GetConst(name string) string {
	var constRecord Const

	db.DB.Where("name = ?", name).Find(&constRecord)

	return constRecord.Value
}
