package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetConstsTable(ctx *context.Context) table.Table {

	consts := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := consts.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Description", "description", db.Varchar)
	info.AddField("Value", "value", db.Varchar)

	info.SetTable("consts").SetTitle("Consts").SetDescription("Consts")

	formList := consts.GetForm()
	// formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Description", "description", db.Varchar, form.Text)
	formList.AddField("Value", "value", db.Varchar, form.Text)

	formList.SetTable("consts").SetTitle("Consts").SetDescription("Consts")

	return consts
}
