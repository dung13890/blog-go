package handler

import (
	"fmt"
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

type User struct{}

func (u *User) Index(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	order := c.QueryParam("order")
	sort := c.QueryParam("sort")
	fmt.Println(sort)
	context := config.NewContext()
	user := context.DbCollection("users")
	defer context.Close()
	repo := &model.UserRepo{user}
	users := repo.Get(limit, offset, order, sort)

	return c.JSON(http.StatusOK, users)
}
