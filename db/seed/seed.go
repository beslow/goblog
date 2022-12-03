package seed

import (
	"fmt"

	"github.com/beslow/goblog/db"
	"github.com/beslow/goblog/models/consts"
)

func Seed() {
	var allConsts = [][3]string{
		{"real_name", "", "your name"},
		{"profession", "", "高级服务器端工程师"},
		{"age", "", "30"},
		{"city", "", "Shanghai"},
		{"hobby", "", "game"},
		{"email", "", "549174542@qq.com"},
	}

	var constRecord consts.Const
	for _, c := range allConsts {
		db.DB.Where(consts.Const{Name: c[0]}).Attrs(consts.Const{Description: c[1], Value: c[2]}).FirstOrCreate(&constRecord)
	}

	fmt.Printf("consts seed finish.\n")
}
