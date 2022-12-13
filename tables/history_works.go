package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetHistoryWorksTable(ctx *context.Context) table.Table {

	historyWorks := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := historyWorks.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Company", "company", db.Varchar)
	info.AddField("Job_title", "job_title", db.Varchar)
	info.AddField("Description", "description", db.Varchar)
	info.AddField("From_to", "from_to", db.Varchar)
	info.AddField("Sort", "sort", db.Int)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Updated_at", "updated_at", db.Datetime)

	info.SetTable("history_works").SetTitle("HistoryWorks").SetDescription("HistoryWorks")

	formList := historyWorks.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).FieldDisplayButCanNotEditWhenUpdate().FieldNotAllowAdd()
	formList.AddField("Company", "company", db.Varchar, form.Text)
	formList.AddField("Job_title", "job_title", db.Varchar, form.Text)
	formList.AddField("Description", "description", db.Varchar, form.Text)
	formList.AddField("From_to", "from_to", db.Varchar, form.Text)
	formList.AddField("Sort", "sort", db.Int, form.Text)

	formList.SetTable("history_works").SetTitle("HistoryWorks").SetDescription("HistoryWorks")

	return historyWorks
}
