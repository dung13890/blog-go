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
	Data    []model.User `json:"data,omitempty"`
}

type User struct{}

func (u *User) Index(c echo.Context) error {
	query := c.QueryParams()
	fmt.Println(query)
	context := config.NewContext()
	user := context.DbCollection("users")
	defer context.Close()
	repo := &model.UserRepo{user}
	users := repo.Get(query)

	return c.JSON(http.StatusOK, &response{
		Message: "ok",
		Data:    users,
	})
}

func (u *User) Create(c echo.Context) error {
	params := model.User{}
	c.Bind(&params)
	fmt.Println(params)
	context := config.NewContext()
	user := context.DbCollection("users")
	defer context.Close()
	repo := &model.UserRepo{user}
	_ = repo.Create(&params)
	return c.String(http.StatusCreated, "ok")
}
