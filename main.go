package main

import (
	"github.com/dung13890/microservices/api/config"
	"github.com/dung13890/microservices/api/router"
	"github.com/labstack/echo"
)

func main() {
	config.Init()
	e := echo.New()
	router.Init(e)
	e.Logger.Fatal(e.Start(":3000"))
}
