package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetHistoryEducationsTable(ctx *context.Context) table.Table {

	historyEducations := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := historyEducations.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Course", "course", db.Varchar)
	info.AddField("City", "city", db.Varchar)
	info.AddField("Description", "description", db.Varchar)
	info.AddField("From_to", "from_to", db.Varchar)
	info.AddField("Sort", "sort", db.Int)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Updated_at", "updated_at", db.Timestamp)

	info.SetTable("history_educations").SetTitle("HistoryEducations").SetDescription("HistoryEducations")

	formList := historyEducations.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).FieldDisplayButCanNotEditWhenUpdate().FieldNotAllowAdd()
	formList.AddField("Course", "course", db.Varchar, form.Text)
	formList.AddField("City", "city", db.Varchar, form.Text)
	formList.AddField("Description", "description", db.Varchar, form.Text)
	formList.AddField("From_to", "from_to", db.Varchar, form.Text)
	formList.AddField("Sort", "sort", db.Int, form.Number)

	formList.SetTable("history_educations").SetTitle("HistoryEducations").SetDescription("HistoryEducations")

	return historyEducations
}
