//go:build ignore
// +build ignore

package main

//import (
//	"GoAdminTelir/Tpage"
//	"context"
//	_ "github.com/GoAdminGroup/go-admin/adapter/gin" // web framework adapter
//	"github.com/GoAdminGroup/go-admin/modules/config"
//	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite" // sql driver
//	"github.com/GoAdminGroup/go-admin/modules/language"
//	_ "github.com/GoAdminGroup/themes/adminlte" // ui theme
//	"io/ioutil"
//	"log"
//	"net/http"
//	"os"
//	"os/signal"
//	"time"
//
//	"github.com/GoAdminGroup/go-admin/engine"
//	"github.com/GoAdminGroup/go-admin/template"
//	"github.com/GoAdminGroup/go-admin/template/chartjs"
//	"github.com/gin-gonic/gin"
//)
//
//func main() {
//	startServer()
//}
//
//func startServer() {
//	gin.SetMode(gin.ReleaseMode)
//	gin.DefaultWriter = ioutil.Discard
//
//	r := gin.Default()
//
//	eng := engine.Default()
//
//	template.AddComp(chartjs.NewChart())
//
//	cfg := config.Config{
//		Databases: config.DatabaseList{
//			"default": {
//				Driver: config.DriverSqlite,
//				File:   "./data/admin.db",
//			},
//		},
//		UrlPrefix: "admin",
//		IndexUrl:  "/",
//		Debug:     true,
//		Logo:      "<b>Go</b>AdminPanel", // change logo
//		Language:  language.EN,
//	}
//
//	if err := eng.AddConfig(&cfg).
//		Use(r); err != nil {
//		panic(err)
//	}
//
//	r.Static("/uploads", "./uploads")
//
//	eng.HTML("GET", "/admin", Tpage.DashboardPage)
//	eng.HTMLFile("GET", "/admin/hello", "./html/hello.tmpl", map[string]interface{}{
//		"msg": "Hello world",
//	})
//
//	srv := &http.Server{
//		Addr:    ":1234",
//		Handler: r,
//	}
//
//	go func() {
//		if err := srv.ListenAndServe(); err != nil {
//			log.Printf("listen: %s\n", err)
//		}
//	}()
//
//	quit := make(chan os.Signal)
//	signal.Notify(quit, os.Interrupt)
//	<-quit
//
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	if err := srv.Shutdown(ctx); err != nil {
//		log.Fatal("Server Shutdown:", err)
//	}
//	log.Println("Server exiting")
//}
