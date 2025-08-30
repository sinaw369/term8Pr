package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/GoAdminGroup/components/echarts"
	_ "github.com/GoAdminGroup/go-admin/adapter/echo"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	_ "github.com/GoAdminGroup/themes/adminlte"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sinaw369/term8Pr/Tdata"
	"github.com/sinaw369/term8Pr/Tmenu"
	"github.com/sinaw369/term8Pr/Tpage"
	"github.com/sinaw369/term8Pr/handler/SearchLogsEng"
)

func main() {
	command := flag.String("BasePath", "", "BasePath")
	flag.Parse()
	fmt.Println("command:", *command)
	// initialize go-admin
	DEL := SearchLogsEng.SetUpSearchLogEngine(*command)

	dbFile := "./Tdata/database.db"
	e := echo.New()

	// Add the panic recovery middleware
	e.Use(middleware.Recover())

	//adminPlugin := admin.NewAdmin(t)
	//adminPlugin := admin.NewAdmin()
	eng := engine.Default()
	// Specify the absolute path to your SQLite database file
	cfg := config.Config{
		Env: config.EnvLocal,
		Databases: config.DatabaseList{
			"default": {
				Driver: config.DriverSqlite,
				File:   dbFile,
			},
		},
		UrlPrefix: "admin",
		IndexUrl:  "/",
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Debug:      true,
		Language:   language.EN,
		Theme:      "adminlte",
		Logo:       template.HTML("sina"),
		MiniLogo:   template.HTML("cno"),
		Title:      "SinaPanel",
		LoginTitle: "SinaCoPanel",
		//	LoginLogo:  template.HTML(`<img src="Timage/telir.png" alt="Telir">`),
		LoginLogo: template.HTML(`<link rel="icon" type="image/png" sizes="32x32" href="//quick.go-admin.cn/official/assets/imgs/icons.ico/favicon-32x32.png">`),
		//SessionLifeTime: 30,
		//CustomHeadHtml: template.HTML(`<img src="Timage/telir.png" alt="Telir">`),
		CustomHeadHtml: template.HTML(`<link rel="icon" type="image/png" sizes="32x32" href="//quick.go-admin.cn/official/assets/imgs/icons.ico/favicon-32x32.png">`),
		// logs
		AccessLogPath: "./GoAdminLogs/access.log",
		ErrorLogPath:  "./GoAdminLogs/error.log",
		InfoLogPath:   "./GoAdminLogs/info.log",
	}

	template.AddComp(chartjs.NewChart())
	template.AddComp(echarts.NewChart())
	// customize a plugin

	//examplePlugin := example.NewExample()

	// load from golang.Plugin
	//
	// examplePlugin := plugins.LoadFromPlugin("../datamodel/example.so")

	// customize the login page
	// example: https://github.com/GoAdminGroup/demo.go-admin.cn/blob/master/main.go#L39
	//
	// template.AddComp("login", datamodel.LoginPage)

	// load config from json file
	//
	// eng.AddConfigFromJSON("../datamodel/config.json")
	// init db
	err := Tdata.InitDB()
	if err != nil {
		panic(err)
	}
	if err := eng.AddConfig(&cfg).
		AddGenerators(Tpage.Generators).
		//AddGenerator("SearchLogs", Tpage.GetExternalTable).
		AddDisplayFilterXssJsFilter().
		// add generator, first parameter is the url prefix of table when visit.
		// example:
		//
		// "user" => http://localhost:9033/admin/info/user
		//
		//AddGenerator("user", datamodel.GetUserTable).
		//AddPlugins(examplePlugin).
		Use(e); err != nil {
		fmt.Println(err)
		panic(err)
	}
	for i := 0; i < len(Tmenu.GetMenu()); i++ {
		_, err := menu.NewMenu(eng.SqliteConnection(), Tmenu.GetMenu()[i])
		if err != nil {
			return
		}
	}

	e.Static("/uploads", "./uploads")
	// you can custom your pages like:
	//eng.HTML("GET", "/admin", Tpage.GetContent)
	eng.HTMLFile("GET", "/admin", "./html/helloDashboard.html", map[string]interface{}{})
	eng.Data("GET", "/admin/healthCheck", healthCheck)
	eng.Data("POST", "/search", DEL.SLHa.AdminHandler.SearchLogHandlerExternal)
	eng.HTML("GET", "/admin/DemoPanel", Tpage.GetDashBoard2Content)
	eng.HTML("GET", "/admin/DemoPanel1", Tpage.GetDashBoardContent)
	eng.HTML("GET", "/admin/DemoPanel2", Tpage.GetDashBoard3Content)
	eng.HTML("GET", "/admin/form", Tpage.GetForm1Content)
	eng.HTML("GET", "/admin/table", Tpage.GetTableContent)
	//plug, _ := plugins.FindByName("filemanager")

	// Start server

	go e.Logger.Fatal(e.Start(":1235"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.MysqlConnection().Close()

}
func healthCheck(c *context.Context) {
	c.JSON(http.StatusOK, echo.Map{
		"message": "everything is good!",
	})
}
