package pages

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/language"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetHistoryWorkFormContent(ctx *context.Context) (types.Panel, error) {
	components := template2.Get(config.GetTheme())

	col1 := components.Col().GetContent()
	btn1 := components.Button().SetType("submit").
		SetContent(language.GetFromHtml("Save")).
		SetThemePrimary().
		SetOrientationRight().
		SetLoadingText(icon.Icon("fa-spinner fa-spin", 2) + `Save`).
		GetContent()
	btn2 := components.Button().SetType("reset").
		SetContent(language.GetFromHtml("Reset")).
		SetThemeWarning().
		SetOrientationLeft().
		GetContent()
	col2 := components.Col().SetSize(types.SizeMD(8)).
		SetContent(btn1 + btn2).GetContent()

	var panel = types.NewFormPanel()
	panel.AddField("公司名称", "company", db.Varchar, form.Text).FieldMust()
	panel.AddField("职位名称", "job_title", db.Varchar, form.Text).FieldMust()
	panel.AddField("工作描述", "description", db.Varchar, form.TextArea).FieldMust()
	panel.AddField("时间段", "from_to", db.Varchar, form.Text).FieldMust()
	panel.AddField("排序", "sort", db.Int, form.Text)

	panel.SetTabGroups(types.TabGroups{
		{"company", "job_title", "description", "from_to", "sort"},
	})

	panel.SetTabHeaders("工作经历")

	fields, headers := panel.GroupField()

	aform := components.Form().
		SetTabHeaders(headers).
		SetTabContents(fields).
		SetPrefix(config.PrefixFixSlash()).
		SetUrl("/admin/history_works").
		SetHiddenFields(map[string]string{
			form2.PreviousKey: "/admin",
		}).
		SetOperationFooter(col1 + col2)

	return types.Panel{
		Content: components.Box().
			SetHeader(aform.GetDefaultBoxHeader(true)).
			WithHeadBorder().
			SetBody(aform.GetContent()).
			GetContent(),
		Title:       "创建履历",
		Callbacks:   panel.Callbacks,
		Description: "",
	}, nil
}
