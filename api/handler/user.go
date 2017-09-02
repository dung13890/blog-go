package handler

import (
	_ "log"
	"net/http"

	"github.com/dung13890/blog-go/api/config"
	"github.com/dung13890/blog-go/api/model"
	"github.com/labstack/echo"
)

type response struct {
	Message string       `json:"message,omitempty"`
	Items   []model.User `json:"items,omitempty"`
}

func ListUser(c echo.Context) error {
	context := config.NewContext()
	u := context.DbCollection("users")
	defer context.Close()
	repo := &model.UserRepo{u}
	users := repo.Get()

	return c.JSON(http.StatusOK, users)
}
