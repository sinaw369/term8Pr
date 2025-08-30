package Tpage

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/paginator"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
)

func GetTableContent(ctx *context.Context) (types.Panel, error) {

	newCtx := context.NewContext(ctx.Request)

	comp := template2.Get(newCtx, config.GetTheme())

	table := comp.DataTable().
		SetInfoList([]map[string]types.InfoItem{
			{
				"id":     {Content: "0"},
				"name":   {Content: "Jack"},
				"gender": {Content: "men"},
				"age":    {Content: "20"},
			},
			{
				"id":     {Content: "1"},
				"name":   {Content: "Jane"},
				"gender": {Content: "women"},
				"age":    {Content: "23"},
			},
		}).
		SetPrimaryKey("id").
		SetThead(types.Thead{
			{Head: "ID", Field: "id"},
			{Head: "Name", Field: "name"},
			{Head: "Gender", Field: "gender"},
			{Head: "Age", Field: "age"},
		})
	body := table.GetContent()
	return types.Panel{
		Content: comp.Box().
			SetBody(body).
			SetNoPadding().
			SetHeader(table.GetDataTableHeader()).
			WithHeadBorder().
			SetFooter(paginator.Get(newCtx, paginator.Config{
				Size:         50,
				PageSizeList: []string{"10", "20", "30", "50"},
				Param:        parameter.GetParam(ctx.Request.URL, 10),
			}).GetContent()).
			GetContent(),
		Title:       "form",
		Description: "tableExample",
	}, nil
}
