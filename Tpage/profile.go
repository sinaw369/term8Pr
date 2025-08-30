package Tpage

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"path/filepath"
)

func GetProfileTable(ctx *context.Context) table.Table {

	profile := table.NewDefaultTable(ctx, table.Config{
		Driver:     db.DriverSqlite,
		CanAdd:     true,
		Editable:   true,
		Deletable:  true,
		Exportable: true,
		Connection: table.DefaultConnectionName,
		PrimaryKey: table.PrimaryKey{
			Type: db.Int,
			Name: table.DefaultPrimaryKeyName,
		},
	})

	info := profile.GetInfo().HideFilterArea()
	info.AddField("ID", "id", db.Int).FieldFilterable()
	info.AddField("UUID", "uuid", db.Varchar).FieldCopyable()
	info.AddField("pass", "pass", db.Tinyint).FieldBool("1", "0")
	//info.AddField("photo", "photos", db.Varchar).FieldCarousel(func(value string) []string {
	//	return strings.Split(value, ",")
	//}, 150, 100)
	info.AddField("finish_state", "finish_state", db.Tinyint).
		FieldDisplay(func(value types.FieldModel) interface{} {
			if value.Value == "0" {
				return "step1"
			}
			if value.Value == "1" {
				return "step2"
			}
			if value.Value == "2" {
				return "step3"
			}
			return "unknown"
		}).
		FieldDot(map[string]types.FieldDotColor{
			"step1": types.FieldDotColorDanger,
			"step2": types.FieldDotColorInfo,
			"step3": types.FieldDotColorPrimary,
		}, types.FieldDotColorDanger)
	info.AddField("finish_progress", "finish_progress", db.Int).FieldProgressBar()
	info.AddField("resume", "resume", db.Varchar).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return filepath.Base(value.Value)
		}).
		FieldDownLoadable("http://yinyanghu.github.io/files/")
	info.AddField("resume_size", "resume_size", db.Int).FieldFileSize()

	info.AddButton(ctx, "more", icon.FolderO, action.PopUpWithForm(action.PopUpData{
		Id:     "/admin/popup/form",
		Title:  "Popup Form Example",
		Width:  "900px",
		Height: "540px",
	}, func(panel *types.FormPanel) *types.FormPanel {
		panel.AddField("name", "name", db.Varchar, form.Text)
		panel.AddField("age", "age", db.Int, form.Number)
		panel.AddField("homePage", "homepage", db.Varchar, form.Url).FieldDefault("http://google.com")
		panel.AddField("Email", "email", db.Varchar, form.Email).FieldDefault("xxxx@xxx.com")
		panel.AddField("birthday", "birthday", db.Varchar, form.Date).FieldDefault("2010-09-03 18:09:05")
		panel.AddField("time", "time", db.Varchar, form.Datetime).FieldDefault("2010-09-05")
		panel.EnableAjax("addedSuccessfully", "failedToAdd")
		return panel
	}, "/admin/popup/form"))

	info.SetTable("profile").SetTitle("profile").SetDescription("profile")

	formList := profile.GetForm()
	formList.AddField("UUID", "uuid", db.Varchar, form.Text)
	//formList.AddField("photos", "photos", db.Varchar, form.Text)
	formList.AddField("resume", "resume", db.Varchar, form.Text)
	formList.AddField("resumeSize", "resume_size", db.Int, form.Number)
	formList.AddField("finishedCondition", "finish_state", db.Tinyint, form.Number)
	formList.AddField("completeSchedule", "finish_progress", db.Int, form.Number)
	formList.AddField("clearance", "pass", db.Tinyint, form.Number)

	formList.SetTable("profile").SetTitle("profile").SetDescription("profile")

	return profile
}
