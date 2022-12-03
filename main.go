package main

import (
	container "quiz_app/app"
	"quiz_app/infra/config"
	"quiz_app/infra/conn"

	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadConfig()
	conn.ConnectDb()
	e := echo.New()
	container.Init(e)
	e.Start(":8080")

}
