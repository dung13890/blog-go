package router

import (
	"github.com/dung13890/blog-go/api/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	user := &handler.User{}
	e.GET("/users", user.Index)
	e.POST("/users", user.Create)
	e.DELETE("/users/:id", user.Destroy)
}
