package handler

import (
	"net/http"

	"github.com/dung13890/blog-go/api/config"
	"github.com/dung13890/blog-go/api/model"
	"github.com/labstack/echo"
)

type response struct {
	Message string       `json:"message,omitempty"`
	Data    []model.User `json:"data,omitempty"`
	Count   int          `json:"count,omitempty"`
}

type User struct{}

func (u *User) Index(c echo.Context) error {
	query := c.QueryParams()
	context := config.NewContext()
	user := context.DbCollection("users")
	defer context.Close()
	repo := &model.UserRepo{user}
	users, count := repo.Get(query)

	return c.JSON(http.StatusOK, &response{
		Message: "ok",
		Data:    users,
		Count:   count,
	})
}

func (u *User) Create(c echo.Context) error {
	params := model.User{}
	c.Bind(&params)
	context := config.NewContext()
	user := context.DbCollection("users")
	defer context.Close()
	repo := &model.UserRepo{user}
	_ = repo.Create(&params)

	return c.String(http.StatusCreated, "success")
}
