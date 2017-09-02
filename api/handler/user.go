package handler

import (
	"net/http"

	"github.com/dung13890/microservices/api/model"
	"github.com/labstack/echo"
)

func ListUser(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
