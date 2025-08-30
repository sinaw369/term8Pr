package Tpage

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"time"
)

// GetExternalTable return the model from external data source.
func SearchLogs(ctx *context.Context) (externalTable table.Table) {

	externalTable = table.NewDefaultTable(ctx, table.Config{
		Driver:     db.DriverSqlite,
		CanAdd:     true,
		Editable:   false,
		Deletable:  false,
		Exportable: true,
		Connection: table.DefaultConnectionName,
		PrimaryKey: table.PrimaryKey{
			Type: db.Int,
			Name: table.DefaultPrimaryKeyName,
		},
	})

	info := externalTable.GetInfo().SetFilterFormLayout(form.LayoutFourCol)

	currentTime := time.Now().Format("15:04:05")
	info.AddField("Search Time", "SearchTime", db.Varchar).
		FieldFilterable(types.FilterType{
			FormType: form.Text,
			//Placeholder: currentTime,
			HeadWidth: 7,
			NoIcon:    true,
		}).FieldFilterOptions(types.FieldOptions{
		{Value: currentTime, Text: currentTime},
	}).FieldHide().AddJS(`let input1 = document.querySelector(".SearchTime");
	input1.setAttribute("readonly", "readonly");
	var endTime = new Date();
	var formattedResponseTime = endTime.getHours() + ':' + endTime.getMinutes() + ':' + endTime.getSeconds();
	input1.value = formattedResponseTime;`).AddCSS(`
	.SearchTime{
		font-size: 25px !important;
	   text-align: center;
	   border: none;
	   background-color: white  !important;
	}
	`)

	info.AddField("نام فایل", "FileName", db.Varchar).FieldHide().FieldFilterable(types.FilterType{
		FormType:    form.Text,
		Placeholder: "نام فایل را وارد کنید",
		HeadWidth:   6,
	}).AddJS(`
let input2 = document.querySelector(".FileName");
	input2.setAttribute("autocomplete","false");`)

	info.AddField("تاریخ", "Date", db.Timestamp).FieldHide().FieldFilterable(types.FilterType{
		FormType:    form.Date,
		Placeholder: "تاریخ را انتخاب کنید",
		HeadWidth:   4,
	}).AddJS(`
let input3 = document.querySelector(".Date");
input3.setAttribute("autocomplete", "false");
//var CurrentDate = new Date();
//var month = (CurrentDate.getMonth() + 1).toString().padStart(2, '0'); // Adjusting month format
//var day = CurrentDate.getDate().toString().padStart(2, '0'); // Adjusting day format
//input3.value = CurrentDate.getFullYear() + '-' + month + '-' + day;

	`)
	info.AddField("نوع لاگ", "LogType", db.Int).FieldHide().
		FieldFilterable(types.FilterType{
			FormType: form.SelectSingle,
			//Operator:    ,
			Placeholder: "نوع لاگ را انتخاب کنید",
			//Width:       150,
			HeadWidth: 6,
		}).FieldFilterOptions(types.FieldOptions{
		{Value: "errors", Text: "خطا"},
		{Value: "info", Text: "عادی"},
	}).AddJS(`
	// Select the "Error" option in the dropdown
	//$('select.LogType').val('errors').trigger('change');
	
		//var selectElement = $("select.LogType");
	   //Remove all options except "Error" and "info"
	   //selectElement.find("option").each(function() {
	   //   if ($(this).val() !== "errors" && $(this).val() !== "info") {
	   //       $(this).remove();
	   //   }
	   //});
		//
	   //Initialize the select2 plugin with allowClear set to false
	   //selectElement.select2({
	   //   allowClear: false
	   //});
	`)

	info.AddField("شامل این فایل نباشد", "notIncludeFileName", db.Varchar).FieldHide().FieldFilterable(types.FilterType{
		FormType:    form.Text,
		Placeholder: "بدون فایل ...",
		HeadWidth:   6,
	}).AddJS(`
let input4 = document.querySelector(".notIncludeFileName");
	input4.setAttribute("autocomplete","false");`)
	info.AddField("محدودیت جستجو", "Limit", db.Int).FieldHide().FieldFilterable(types.FilterType{
		FormType:    form.Text,
		Placeholder: "شامل چه تعداد نتیجه باشد",
		HeadWidth:   6,
	}).AddJS(`
	let input5 = document.querySelector(".Limit");
    input5.setAttribute("autocomplete","false");
    input5.value=100;
	`)
	info.AddField("کلید جستجو", "searchKey", db.Varchar).FieldHide().FieldFilterable(types.FilterType{
		FormType:    form.Text,
		Placeholder: "کلید جستجو",
		HeadWidth:   6,
	}).AddJS(`
	let input6 = document.querySelector(".searchKey");
input5.setAttribute("autocomplete","false");
	`)
	info.AddField("شامل این کلید نباشد", "notIncludeSearchKey", db.Varchar).FieldHide().FieldFilterable(types.FilterType{
		FormType:    form.Text,
		Placeholder: "بدون کلید ...",
		HeadWidth:   6,
	}).AddJS(`
	let input7 = document.querySelector(".notIncludeSearchKey");
input7.setAttribute("autocomplete","false");
	`)
	//.FieldFilterProcess(func(val string) string {
	//	return val + "ms"
	//})
	info.AddField("فقط شمارش", "justCount", db.Tinyint).FieldHide().FieldFilterable(types.FilterType{
		FormType:  form.Switch,
		HeadWidth: 4,
	}).FieldFilterOptions(types.FieldOptions{
		{Text: "ON", Value: "1"},
		{Text: "OFF", Value: "0", Selected: true},
	})
	//info.DefaultPageSize = 100
	/*! table columns  */
	info.AddField("ID", "id", db.Int).FieldSortable()
	if ctx.FormValue("justCount") == "0" {
		info.AddField("FileName", "filename", db.Varchar).FieldCopyable()
		info.AddField("LogLevel", "level", db.Varchar)
		info.AddField("Body", "message", db.JSON).FieldCopyable()
		info.AddField("Time", "time", db.Varchar)
	} else {
		info.AddField("NumberOfResult", "numberOfResult", db.Int)
	}

	info.HideDeleteButton().HideEditButton()

	Data, JustCount, _ := DeliverTpage().SearchLogHandler(ctx)
	info.SetTable("SearchLogs").
		SetTitle("SearchLogs").
		SetDescription("SearchLogs").
		SetGetDataFn(func(param parameter.Parameters) ([]map[string]interface{}, int) {
			sliceOfMaps, LensliceOfMaps := FixDataForShowInGoAdminSearchLog(Data, JustCount)
			return sliceOfMaps, LensliceOfMaps
		})

	//detail := externalTable.GetDetail()
	//
	//detail.AddField("allinone", "allin", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
	//
	//	return value.Row["filename"].(string) + " - " + value.Row["message"].(string)
	//})
	//detail.SetTable("SearchLogs").SetTitle("SearchLogsDetails").SetDescription("SearchLogsDetails")
	return
}
