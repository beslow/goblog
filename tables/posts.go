package tables

import (
	"strconv"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/beslow/goblog/initialize"
	"github.com/beslow/goblog/models"
)

func GetPostsTable(ctx *context.Context) table.Table {

	posts := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := posts.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Title", "title", db.Varchar)
	info.AddField("Visit_count", "visit_count", db.Int)
	info.AddField("Updated_at", "updated_at", db.Timestamp)
	info.AddField("Created_at", "created_at", db.Timestamp)

	info.SetTable("posts").SetTitle("Posts").SetDescription("Posts")

	formList := posts.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)

	var categories []models.Category
	initialize.DB.Find(&categories)

	var fieldOptions = make(types.FieldOptions, 0, 10)
	for _, category := range categories {
		fieldOptions = append(fieldOptions, types.FieldOption{
			Text:  category.Name,
			Value: strconv.Itoa(category.ID),
		})
	}

	formList.AddField("CategoryID", "category_id", db.Int, form.SelectSingle).FieldOptions(fieldOptions)

	formList.AddField("Title", "title", db.Varchar, form.Text)
	formList.AddField("Summary", "summary", db.Varchar, form.Text)
	formList.AddField("Body", "body", db.Text, form.RichText)

	formList.SetTable("posts").SetTitle("Posts").SetDescription("Posts")

	return posts
}
