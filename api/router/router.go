package router

import (
	"github.com/dung13890/microservices/api/handler"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	e.GET("/users", handler.ListUser)
}
