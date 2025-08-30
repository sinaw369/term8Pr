package Tpage

import (
	"fmt"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	selection "github.com/GoAdminGroup/go-admin/template/types/form/select"
	editType "github.com/GoAdminGroup/go-admin/template/types/table"
)

func GetUserTable(ctx *context.Context) (userTable table.Table) {

	userTable = table.NewDefaultTable(ctx, table.Config{
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

	info := userTable.GetInfo().SetFilterFormLayout(form.LayoutThreeCol).ExportValue()
	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("Name", "name", db.Varchar).FieldEditAble(editType.Text).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Gender", "gender", db.Tinyint).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "0" {
			return "men"
		}
		if model.Value == "1" {
			return "women"
		}
		return "unknown"
	}).FieldEditAble(editType.Switch).FieldEditOptions(types.FieldOptions{
		{Value: "0", Text: "men"},
		{Value: "1", Text: "women"},
	}).FieldFilterable(types.FilterType{FormType: form.SelectSingle}).FieldFilterOptions(types.FieldOptions{
		{Value: "0", Text: "men"},
		{Value: "1", Text: "women"},
	})
	info.AddColumn("personality", func(value types.FieldModel) interface{} {
		return "handsome"
	})
	info.AddColumnButtons(ctx, "seeMore", types.GetColumnButton("more", icon.Info,
		action.PopUp("/see/more/example", "more", func(ctx *context.Context) (success bool, msg string, data interface{}) {
			return true, "ok", "<h1>detail</h1><p>balabala</p><p>thisFeatureV1.2._7Open</p>"
		})))
	info.AddField("Phone", "phone", db.Varchar).FieldFilterable()
	info.AddField("City", "city", db.Varchar).FieldFilterable().
		FieldEditAble(editType.Select).FieldEditOptions(types.FieldOptions{
		{Value: "guangzhou", Text: "canton"},
		{Value: "shanghai", Text: "shanghai"},
		{Value: "beijing", Text: "beijing"},
		{Value: "shenzhen", Text: "shenzhen"},
	})
	//info.AddField("Avatar", "avatar", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
	//	return template.Default().Image().
	//		SetSrc(`../uploads/test1.png`).
	//		SetHeight("120").SetWidth("120").WithModal().GetContent()
	//})
	info.AddField("CreatedAt", "created_at", db.Timestamp).
		FieldFilterable(types.FilterType{FormType: form.DatetimeRange})
	info.AddField("UpdatedAt", "updated_at", db.Timestamp).FieldEditAble(editType.Datetime)

	info.AddActionButton(ctx, "google", action.Jump("https://google.com"))
	info.AddActionButton(ctx, "approval", action.Ajax("/admin/audit",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			return true, "approvalSuccessful，awesome", ""
		}))
	info.AddActionButton(ctx, "preview", action.PopUp("/admin/", "preview",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			return true, "", "<h2>previewContent</h2>"
		}))
	info.AddButton(ctx, "jump", icon.User, action.JumpInNewTab("/admin/info/SearchLogs", "SearchLogs"))
	info.AddButton(ctx, "popup", icon.Terminal, action.PopUp("/admin/popup", "Popup Example",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			return true, "", "<h2>hello world</h2>"
		}))
	info.AddButton(ctx, "iframe", icon.Tv, action.PopUpWithIframe("/admin/iframe", "Iframe Example",
		action.IframeData{Src: "/admin/info/ProfileTable"}, "900px", "600px"))
	info.AddButton(ctx, "ajax", icon.Android, action.Ajax("/admin/ajax",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			return true, "requestSuccessful，awesome", ""
		}))

	info.AddSelectBox(ctx, "gender", types.FieldOptions{
		{Value: "0", Text: "male"},
		{Value: "1", Text: "female"},
	}, action.FieldFilter("gender"))

	info.SetTable("users").SetTitle("Users").SetDescription("Users")

	formList := userTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldDisplayButCanNotEditWhenUpdate().FieldNotAllowAdd()
	formList.AddField("IP", "ip", db.Varchar, form.Text)
	formList.AddField("fullName", "name", db.Varchar, form.Text)
	formList.AddField("gender", "gender", db.Tinyint, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "male", Value: "0"},
			{Text: "female", Value: "1"},
		}).FieldDefault("0")
	formList.AddField("phone", "phone", db.Varchar, form.Text)
	formList.AddField("country", "country", db.Tinyint, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: "china", Value: "0"},
			{Text: "us", Value: "1"},
			{Text: "britain", Value: "2"},
			{Text: "canada", Value: "3"},
		}).FieldDefault("0").FieldOnChooseAjax("city", "/choose/country",
		func(ctx *context.Context) (bool, string, interface{}) {
			country := ctx.FormValue("value")
			var data = make(selection.Options, 0)
			switch country {
			case "0":
				data = selection.Options{
					{Text: "beijing", ID: "beijing"},
					{Text: "shanghai", ID: "shangHai"},
					{Text: "canton", ID: "guangZhou"},
					{Text: "shenzhen", ID: "shenZhen"},
				}
			case "1":
				data = selection.Options{
					{Text: "losAngeles", ID: "los angeles"},
					{Text: "washington", ID: "washington, dc"},
					{Text: "newYork", ID: "new york"},
					{Text: "lasVegas", ID: "las vegas"},
				}
			case "2":
				data = selection.Options{
					{Text: "london", ID: "london"},
					{Text: "cambridge", ID: "cambridge"},
					{Text: "manchester", ID: "manchester"},
					{Text: "liverpool", ID: "liverpool"},
				}
			case "3":
				data = selection.Options{
					{Text: "vancouver", ID: "vancouver"},
					{Text: "toronto", ID: "toronto"},
				}
			default:
				data = selection.Options{
					{Text: "beijing", ID: "beijing"},
					{Text: "shanghai", ID: "shangHai"},
					{Text: "canton", ID: "guangZhou"},
					{Text: "shenzhen", ID: "shenZhen"},
				}
			}
			return true, "ok", data
		})
	formList.AddField("city", "city", db.Varchar, form.SelectSingle).
		FieldOptionInitFn(func(val types.FieldModel) types.FieldOptions {

			if val.Value == "" {
				return types.FieldOptions{
					{Text: "beijing", Value: "beijing"},
					{Text: "shanghai", Value: "shangHai"},
					{Text: "canton", Value: "guangZhou"},
					{Text: "shenzhen", Value: "shenZhen"},
				}
			}

			return types.FieldOptions{
				{Value: val.Value, Text: val.Value, Selected: true},
			}
		}).FieldOnChooseAjax("district", "/choose/city",
		func(ctx *context.Context) (bool, string, interface{}) {
			country := ctx.FormValue("value")
			var data = make(selection.Options, 0)
			switch country {
			case "beijing":
				data = selection.Options{
					{Text: "haveASunnyExposure", ID: "chaoyang"},
					{Text: "haidian", ID: "haidian"},
				}
			case "shangHai":
				data = selection.Options{
					{Text: "yangpu", ID: "yangpu"},
					{Text: "pudong", ID: "pudong"},
				}
			default:
				data = selection.Options{
					{Text: "southArea", ID: "southern"},
					{Text: "northDistrict", ID: "north"},
				}
			}
			return true, "ok", data
		})
	formList.AddField("region", "district", db.Varchar, form.SelectSingle).
		FieldOptionInitFn(func(val types.FieldModel) types.FieldOptions {

			if val.Value == "" {
				return types.FieldOptions{
					{Text: "southArea", Value: "southern"},
					{Text: "northDistrict", Value: "north"},
				}
			}

			return types.FieldOptions{
				{Value: val.Value, Text: val.Value, Selected: true},
			}
		})
	formList.AddField("customFields", "role", db.Varchar, form.Text).
		FieldPostFilterFn(func(value types.PostFieldModel) interface{} {
			fmt.Println("user custom field", value)
			return ""
		})

	formList.AddField("UpdatedAt", "updated_at", db.Timestamp, form.Default).FieldNotAllowAdd()
	formList.AddField("CreatedAt", "created_at", db.Timestamp, form.Default).FieldNotAllowAdd()

	userTable.GetForm().SetTabGroups(types.
		NewTabGroups("id", "ip", "name", "gender", "country", "city", "district").
		AddGroup("phone", "role", "created_at", "updated_at")).
		SetTabHeaders("profile1", "profile2")

	formList.SetTable("users").SetTitle("Users").SetDescription("Users")

	formList.SetPostHook(func(values form2.Values) error {
		fmt.Println("userTable.GetForm().PostHook", values)
		return nil
	})

	return
}
