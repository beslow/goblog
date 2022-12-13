package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetTestUsersTable(ctx *context.Context) table.Table {

	testUsers := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := testUsers.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Age", "age", db.Int)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Updated_at", "updated_at", db.Timestamp)

	info.SetTable("test_users").SetTitle("TestUsers").SetDescription("TestUsers")

	formList := testUsers.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Age", "age", db.Int, form.Number)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime).FieldNotAllowAdd()
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime).FieldNotAllowAdd()

	formList.SetTable("test_users").SetTitle("TestUsers").SetDescription("TestUsers")

	return testUsers
}
