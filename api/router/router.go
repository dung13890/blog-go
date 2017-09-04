package router

import (
	"github.com/dung13890/blog-go/api/handler"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	user := &handler.User{}
	e.GET("/users", user.Index)
}
