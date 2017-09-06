package main

import (
	"github.com/dung13890/blog-go/api/config"
	"github.com/dung13890/blog-go/api/router"
	"github.com/labstack/echo"
)

func main() {
	config.Init()
	e := echo.New()
	router.Init(e)
	e.Logger.Fatal(e.Start(config.AppConfig.Server))
}
