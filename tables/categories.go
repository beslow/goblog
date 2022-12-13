package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetCategoriesTable(ctx *context.Context) table.Table {

	categories := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := categories.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Updated_at", "updated_at", db.Timestamp)

	info.SetTable("categories").SetTitle("Categories").SetDescription("Categories")

	formList := categories.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("Name", "name", db.Varchar, form.Text)

	formList.SetTable("categories").SetTitle("Categories").SetDescription("Categories")

	return categories
}
