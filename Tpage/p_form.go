package Tpage

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/language"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"html/template"
)

func GetForm1Content(ctx *context.Context) (types.Panel, error) {

	components := template2.Get(context.NewContext(ctx.Request), config.GetTheme())

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
	panel.AddField("name", "name", db.Varchar, form.Text).
		FieldFoot(seeCodeHTML(`formList.AddField("name", "name", db.Varchar, form.Text)`))
	panel.AddField("age", "age", db.Int, form.Number).
		FieldFoot(seeCodeHTML(`formList.AddField("age", "age", db.Int, form.Number)`))
	panel.AddField("homePage", "homepage", db.Varchar, form.Url).FieldDefault("http://google.com").
		FieldFoot(seeCodeHTML(`formList.AddField("homePage", "homepage", db.Varchar, form.Url).FieldDefault("http://google.com")`))
	panel.AddField("mailbox", "email", db.Varchar, form.Email).FieldDefault("xxxx@xxx.com").
		FieldFoot(seeCodeHTML(`formList.AddField("mailbox", "email", db.Varchar, form.Email).FieldDefault("xxxx@xxx.com")`))
	panel.AddField("birthday", "birthday", db.Varchar, form.Date).FieldDefault("2010-09-03 18:09:05").
		FieldFoot(seeCodeHTML(`formList.AddField("birthday", "birthday", db.Varchar, form.Date).FieldDefault("2010-09-03 18:09:05")`))
	panel.AddField("time", "time", db.Varchar, form.Datetime).FieldDefault("2010-09-05").
		FieldFoot(seeCodeHTML(`formList.AddField("time", "time", db.Varchar, form.Datetime).FieldDefault("2010-09-05")`))
	panel.AddField("timeLimit", "time_range", db.Varchar, form.DatetimeRange).
		FieldFoot(seeCodeHTML(`formList.AddField("timeLimit", "time_range", db.Varchar, form.DatetimeRange)`))
	panel.AddField("dateRange", "date_range", db.Varchar, form.DateRange).
		FieldFoot(seeCodeHTML(`formList.AddField("dateRange", "date_range", db.Varchar, form.DateRange)`))
	panel.AddField("password", "password", db.Varchar, form.Password).FieldDivider("iAmTheDividingLine").
		FieldFoot(seeCodeHTML(`formList.AddField("密码", "password", db.Varchar, form.Password).FieldDivider("iAmTheDividingLine")`, true))
	panel.AddField("IP", "ip", db.Varchar, form.Ip).
		FieldFoot(seeCodeHTML(`formList.AddField("IP", "ip", db.Varchar, form.Ip)`))
	panel.AddField("credentials", "certificate", db.Varchar, form.Multifile).FieldOptionExt(map[string]interface{}{
		"maxFileCount": 10,
	}).
		FieldFoot(seeCodeHTML(`formList.AddField("证件", "certificate", db.Varchar, form.Multifile).FieldOptionExt(map[string]interface{}{
		"maxFileCount": 10,
	})`))
	panel.AddField("sum", "currency", db.Int, form.Currency).
		FieldFoot(seeCodeHTML(`formList.AddField("金额", "currency", db.Int, form.Currency)`))
	panel.AddField("proportion", "rate", db.Int, form.Rate).
		FieldFoot(seeCodeHTML(`formList.AddField("比例", "rate", db.Int, form.Rate)`))
	panel.AddField("bonus", "reward", db.Int, form.Slider).FieldOptionExt(map[string]interface{}{
		"max":     1000,
		"min":     1,
		"step":    1,
		"postfix": "dollar",
	}).
		FieldFoot(seeCodeHTML(`formList.AddField("奖金", "reward", db.Int, form.Slider).FieldOptionExt(map[string]interface{}{
		"max":     1000,
		"min":     1,
		"step":    1,
		"postfix": "dollar",
	})`))
	panel.AddField("content", "content", db.Text, form.RichText).
		FieldDefault(`<h1>343434</h1><p>34344433434</p><ol><li>23234</li><li>2342342342</li><li>asdfads</li></ol><ul><li>3434334</li><li>34343343434</li><li>44455</li></ul><p><span style="color: rgb(194, 79, 74);">343434</span></p><p><span style="background-color: rgb(194, 79, 74); color: rgb(0, 0, 0);">434434433434</span></p><table border="0" width="100%" cellpadding="0" cellspacing="0"><tbody><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr></tbody></table><p><br></p><p><span style="color: rgb(194, 79, 74);"><br></span></p>`).
		FieldDivider("secondDividingLine").
		FieldFoot(seeCodeHTML(`formList.AddField("内容", "content", db.Text, form.RichText).
		FieldDefault(`+"`"+`<h1>343434</h1><p>34344433434</p><ol><li>23234</li><li>2342342342</li><li>asdfads</li></ol><ul><li>3434334</li><li>34343343434</li><li>44455</li></ul><p><span style="color: rgb(194, 79, 74);">343434</span></p><p><span style="background-color: rgb(194, 79, 74); color: rgb(0, 0, 0);">434434433434</span></p><table border="0" width="100%" cellpadding="0" cellspacing="0"><tbody><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr></tbody></table><p><br></p><p><span style="color: rgb(194, 79, 74);"><br></span></p>`+"`"+`).
		FieldDivider("secondDividingLine")`, true))
	panel.AddField("code", "code", db.Text, form.Code).FieldDefault(`package main

import "fmt"

func main() {
	fmt.Println("hello GoAdmin!")
}
`).
		FieldFoot(seeCodeHTML(`formList.AddField("code", "code", db.Text, form.Code).FieldDefault(` + "`" + `package main

import "fmt"

func main() {
	fmt.Println("hello GoAdmin!")
}` + "`)"))

	panel.AddField("siteSwitch", "website", db.Tinyint, form.Switch).
		FieldHelpMsg("You will no longer be able to access the site after it is closed，youCanLogInNormallyInTheBackground").
		FieldOptions(types.FieldOptions{
			{Value: "0"},
			{Value: "1"},
		}).
		FieldFoot(seeCodeHTML(`formList.AddField("站点开关", "website", db.Tinyint, form.Switch).
		FieldHelpMsg("You will no longer be able to access the site after it is closed，youCanLogInNormallyInTheBackground").
		FieldOptions(types.FieldOptions{
			{Value: "0"},
			{Value: "1"},
		})`))
	panel.AddField("fruit", "fruit", db.Varchar, form.SelectBox).
		FieldOptions(types.FieldOptions{
			{Text: "apple", Value: "apple"},
			{Text: "banana", Value: "banana"},
			{Text: "watermelon", Value: "watermelon"},
			{Text: "pear", Value: "pear"},
		}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return []string{"pear"}
		}).
		FieldFoot(seeCodeHTML(`formList.AddField("fruit", "fruit", db.Varchar, form.SelectBox).
		FieldOptions(types.FieldOptions{
			{Text: "apple", Value: "apple"},
			{Text: "banana", Value: "banana"},
			{Text: "watermelon", Value: "watermelon"},
			{Text: "pear", Value: "pear"},
		}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return []string{"pear"}
		})`))
	panel.AddField("gender", "gender", db.Tinyint, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "boy", Value: "0"},
			{Text: "girl", Value: "1"},
		}).
		FieldFoot(seeCodeHTML(`formList.AddField("性别", "gender", db.Tinyint, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "boy", Value: "0"},
			{Text: "girl", Value: "1"},
		})`))
	panel.AddField("饮料", "drink", db.Tinyint, form.Select).
		FieldOptions(types.FieldOptions{
			{Text: "beer", Value: "beer"},
			{Text: "juice", Value: "juice"},
			{Text: "plainBoiledWater", Value: "water"},
			{Text: "redBull", Value: "red bull"},
		}).FieldDefault("beer").
		FieldFoot(seeCodeHTML(`formList.AddField("beverage", "drink", db.Tinyint, form.Select).
		FieldOptions(types.FieldOptions{
			{Text: "beer", Value: "beer"},
			{Text: "juice", Value: "juice"},
			{Text: "Water", Value: "water"},
			{Text: "redBull", Value: "red bull"},
		}).FieldDefault("beer")`))
	panel.AddField("workExperience", "experience", db.Tinyint, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: "twoYears", Value: "0"},
			{Text: "threeYears", Value: "1"},
			{Text: "fourYears", Value: "2"},
			{Text: "fiveYears", Value: "3"},
		}).FieldDefault("beer").
		FieldFoot(seeCodeHTML(`formList.AddField("workExperience", "experience", db.Tinyint, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: "twoYears", Value: "0"},
			{Text: "threeYears", Value: "1"},
			{Text: "fourYears", Value: "2"},
			{Text: "fiveYears", Value: "3"},
		}).FieldDefault("beer")`))
	panel.AddField("snack", "snacks", db.Varchar, form.Checkbox).
		FieldOptions(types.FieldOptions{
			{Text: "oatmeal", Value: "0"},
			{Text: "fries", Value: "1"},
			{Text: "spicyStrips", Value: "2"},
			{Text: "iceCream", Value: "3"},
		}).
		FieldFoot(seeCodeHTML(`formList.AddField("snack", "snacks", db.Varchar, form.Checkbox).
		FieldOptions(types.FieldOptions{
			{Text: "oatmeal", Value: "0"},
			{Text: "fries", Value: "1"},
			{Text: "spicyStrips", Value: "2"},
			{Text: "iceCream", Value: "3"},
		})`))
	panel.AddField("cat", "cat", db.Varchar, form.CheckboxStacked).
		FieldOptions(types.FieldOptions{
			{Text: "kaffeyCat", Value: "0"},
			{Text: "britishShort", Value: "1"},
			{Text: "beautifulAndShort", Value: "2"},
		}).
		FieldFoot(seeCodeHTML(`formList.AddField("cat", "cat", db.Varchar, form.CheckboxStacked).
		FieldOptions(types.FieldOptions{
			{Text: "kaffeyCat", Value: "0"},
			{Text: "britishShort", Value: "1"},
			{Text: "beautifulAndShort", Value: "2"},
		})`))
	panel.AddRow(func(pa *types.FormPanel) {
		panel.AddField("province", "province", db.Tinyint, form.SelectSingle).
			FieldOptions(types.FieldOptions{
				{Text: "tehran", Value: "0"},
				{Text: "Esfahan", Value: "1"},
				{Text: "mashhad", Value: "2"},
				{Text: "oromiee", Value: "3"},
			}).FieldRowWidth(2)
		panel.AddField("city", "city", db.Tinyint, form.SelectSingle).
			FieldOptions(types.FieldOptions{
				{Text: "tehran", Value: "0"},
				{Text: "esfahan", Value: "1"},
				{Text: "shiraz", Value: "2"},
				{Text: "ahvaz", Value: "3"},
			}).FieldRowWidth(3).FieldHeadWidth(2).FieldInputWidth(10)
		panel.AddField("area", "district", db.Tinyint, form.SelectSingle).
			FieldOptions(types.FieldOptions{
				{Text: "chaoyangDistrict", Value: "0"},
				{Text: "haizhuDistrict", Value: "1"},
				{Text: "pudongNewArea", Value: "2"},
				{Text: "baoAnDistrict", Value: "3"},
			}).FieldRowWidth(3).FieldHeadWidth(2).FieldInputWidth(9)
	}).FieldFoot(seeCodeHTML(`panel.AddRow(func(pa *types.FormPanel) {
		panel.AddField("province", "province", db.Tinyint, form.SelectSingle).
			FieldOptions(types.FieldOptions{
				{Text: "tehran", Value: "0"},
				{Text: "esfahan", Value: "1"},
				{Text: "shiraz", Value: "2"},
				{Text: "ahvaz", Value: "3"},
			}).FieldRowWidth(2)
		panel.AddField("城市", "city", db.Tinyint, form.SelectSingle).
			FieldOptions(types.FieldOptions{
				{Text: "chaoyangDistrict", Value: "0"},
				{Text: "haizhuDistrict", Value: "1"},
				{Text: "pudongNewArea", Value: "2"},
				{Text: "baoAnDistrict", Value: "3"},
			}).FieldRowWidth(3).FieldHeadWidth(2).FieldInputWidth(10)
		panel.AddField("area", "district", db.Tinyint, form.SelectSingle).
			FieldOptions(types.FieldOptions{
				{Text: "chaoyangDistrict", Value: "0"},
				{Text: "haizhuDistrict", Value: "1"},
				{Text: "pudongNewArea", Value: "2"},
				{Text: "baoAnDistrict", Value: "3"},
			}).FieldRowWidth(3).FieldHeadWidth(2).FieldInputWidth(9)
	})`))
	panel.AddField("employee", "employee", db.Varchar, form.Array).
		FieldFoot(seeCodeHTML(`formList.AddField("employee", "employee", db.Varchar, form.Array)`))
	panel.AddTable("setUp", "setting", func(panel *types.FormPanel) {
		panel.AddField("Key", "key", db.Varchar, form.Text).FieldHideLabel()
		panel.AddField("Value", "value", db.Varchar, form.Text).FieldHideLabel()
	}).
		FieldFoot(seeCodeHTML(`formList.AddTable("setUp", "setting", func(panel *types.FormPanel) {
		panel.AddField("Key", "key", db.Varchar, form.Text).FieldHideLabel()
		panel.AddField("Value", "value", db.Varchar, form.Text).FieldHideLabel()
	})`))

	panel.AddField("shape", "shape", db.Tinyint, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: "round", Value: "0"},
			{Text: "square", Value: "1"},
			{Text: "rectangle", Value: "2"},
		}).
		FieldOnChooseHide("0", "side", "length", "width").
		FieldOnChooseShow("0", "radius").
		FieldOnChooseShow("1", "side").
		FieldOnChooseShow("2", "length", "width").
		FieldDefault("0").
		FieldFoot(seeCodeHTML(`formList.AddField("shape", "shape", db.Tinyint, form.SelectSingle).
	FieldOptions(types.FieldOptions{
			{Text: "round", Value: "0"},
			{Text: "square", Value: "1"},
			{Text: "rectangle", Value: "2"},
	}).
	FieldOnChooseHide("0", "side", "length", "width").
	FieldOnChooseShow("0", "radius").
	FieldOnChooseShow("1", "side").
	FieldOnChooseShow("2", "length", "width").
	FieldDefault("0")`))

	panel.AddField("radius", "radius", db.Int, form.Number).
		FieldDefault("3").
		FieldFoot(seeCodeHTML(`formList.AddField("radius", "radius", db.Int, form.Number).
		FieldDefault("3")`))

	panel.AddField("sideLength", "side", db.Int, form.Number).
		FieldDefault("5").
		FieldFoot(seeCodeHTML(`formList.AddField("边长", "side", db.Int, form.Number).
		FieldDefault("5")`))

	panel.AddField("长", "length", db.Int, form.Number).
		FieldDefault("5").
		FieldFoot(seeCodeHTML(`formList.AddField("长", "length", db.Int, form.Number).
		FieldDefault("5")`))

	panel.AddField("宽", "width", db.Int, form.Number).
		FieldDefault("5").
		FieldFoot(seeCodeHTML(`formList.AddField("宽", "width", db.Int, form.Number).
		FieldDefault("5")`))

	panel.SetTabGroups(types.TabGroups{
		{"name", "age", "homepage", "email", "birthday", "time", "time_range", "date_range", "password", "ip",
			"certificate", "currency", "rate", "reward", "content", "code"},
		{"website", "snacks", "fruit", "gender", "cat", "drink", "province", "city", "district", "experience"},
		{"employee", "setting"},
		{"shape", "radius", "side", "length", "width"},
	})
	panel.SetTabHeaders("enter", "select", "multiple", "linkage")

	fields, headers := panel.GroupField()

	aform := components.Form().
		SetTabHeaders(headers).
		SetTabContents(fields).
		SetPrefix(config.PrefixFixSlash()).
		SetUrl("/admin/form/update").
		SetTitle("formExample").
		SetHiddenFields(map[string]string{
			form2.PreviousKey: "/admin",
		}).
		SetOperationFooter(col1 + col2)

	popup := components.Popup().SetID("code_modal").
		SetHideFooter().
		SetTitle("code").
		SetHeight("300px").
		SetBody(template.HTML("")).
		GetContent()

	return types.Panel{
		Content: `<script src="//cdnjs.cloudflare.com/ajax/libs/ace/1.32.9/mode-golang.min.js" integrity="sha512-mIobi3b3I41C0/RBjxbfyY0IurALSs+BEJS+sUImhu8bc9Cs6UZS76eIWrACdj1Xu7CSsk7++wEO7FR12xNzPA==" crossorigin="anonymous" referrerpolicy="no-referrer"></script>` +
			components.Box().
				SetHeader(aform.GetDefaultBoxHeader(true)).
				WithHeadBorder().
				SetBody(aform.GetContent()+panel.FooterHtml).
				GetContent() + popup,
		Title:       "form",
		Description: "formExample",
		CSS:         `.modal.fade.in{z-index:10002}`,
		JS: `
$(".see-code").on("click", function(){
	$('#code_modal .modal-body').html('<div id="pop_code_editor" style="width: 100%;height: 100%;" class="ace_editor"></div>');
	editor = ace.edit("pop_code_editor");
	editor.setTheme("ace/theme/monokai");
	editor.session.setMode("ace/mode/golang");
	editor.setFontSize(14);
	editor.setValue($(this).parent().next().text());
	$("#code_modal").modal();
})
`,
	}, nil
}

func seeCodeHTML(data string, divide ...bool) template.HTML {
	if len(divide) > 0 && divide[0] {
		return template.HTML(fmt.Sprintf(`<div style="margin-top: 24px;"><a class="see-code" href="javascript:;">viewCode</a></div><div style="display:none;">%s</div>`, data))
	}
	return template.HTML(fmt.Sprintf(`<div><a class="see-code" href="javascript:;">viewCode</a></div><div style="display:none;">%s</div>`, data))
}
