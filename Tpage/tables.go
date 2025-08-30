package Tpage

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

var Generators = map[string]table.Generator{
	"SearchLogs":   SearchLogs,
	"UserTable":    GetUserTable,
	"ProfileTable": GetProfileTable,
}
