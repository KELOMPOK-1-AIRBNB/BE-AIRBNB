package main

import (
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/app/configs"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/app/databases"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/app/migrations"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/app/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := configs.InitConfig()
	dbMySql := databases.InitDBMysql(cfg)
	// dbMysql := databases.InitDBPosgres(cfg)

	// create new instance echo
	e := echo.New()

	migrations.InitialMigration()
	routes.InitRouter(e, dbMySql)
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	// start server and port
	e.Logger.Fatal(e.Start(":8080"))
}

// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
// 	Format: "method=${method}, uri=${uri}, status=${status}\n",
// }))
