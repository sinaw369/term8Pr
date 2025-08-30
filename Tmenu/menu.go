package Tmenu

import (
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/template/icon"
)

// GetMenu gets the menu for the application.
func GetMenu() []menu.NewMenuData {
	return []menu.NewMenuData{

		{
			Header:   "Demo Charts and Panel",
			ParentId: 0,
			Type:     1,
			Order:    4,
			Title:    "Demo Charts and Panel",
			Icon:     icon.AreaChart,
		},
		{
			Header:   "Demo Charts and Panel",
			ParentId: 8,
			Type:     1,
			Order:    4,
			Title:    "demo 1",
			Icon:     icon.AreaChart,
			Uri:      "/DemoPanel",
		},
		{
			ParentId: 8,
			Type:     1,
			Order:    4,
			Title:    "demo 2",
			Icon:     icon.AreaChart,
			Uri:      "/DemoPanel1",
		},
		{
			ParentId: 8,
			Type:     1,
			Order:    4,
			Title:    "demo 3",
			Icon:     icon.AreaChart,
			Uri:      "/DemoPanel2",
		},
		{
			ParentId: 0,
			Type:     1,
			Order:    3,
			Title:    "Search Logs",
			Icon:     icon.Search,
			Uri:      "/info/SearchLogs",
		},
		{
			ParentId: 0,
			Type:     1,
			Order:    5,
			Title:    "sample form",
			Icon:     icon.Wpforms,
			Uri:      "/form",
		},
		{
			ParentId: 0,
			Type:     1,
			Order:    6,
			Title:    "sample table",
			Icon:     icon.Table,
			Uri:      "/table",
		},
		{
			ParentId: 0,
			Header:   "Tables",
			Type:     1,
			Order:    7,
			Title:    "Tables",
			Icon:     icon.Table,
		},
		{
			ParentId: 15,
			Type:     1,
			Order:    8,
			Title:    "UserTable",
			Icon:     icon.Anchor,
			Uri:      "/info/UserTable",
		},
		{
			ParentId: 15,
			Type:     1,
			Order:    9,
			Title:    "ProfileTable",
			Icon:     icon.ProductHunt,
			Uri:      "/info/ProfileTable",
		},
	}

}
