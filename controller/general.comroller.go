package controller

import (
	"fmt"
	"net/http"

	"github.com/earthrockey/Golang-Structure/config"
	"github.com/earthrockey/Golang-Structure/model"
	"github.com/labstack/echo/v4"
)

func HomePage(e echo.Context) error {
	type HomePage struct {
		Message string `json:"message"`
	}
	return e.JSON(http.StatusOK, HomePage{Message: "API SERVER IS WORK!"})
}

func TestCreateUser(e echo.Context) error {
	db, err := config.ConnectDB()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return echo.ErrInternalServerError
	}
	var newuser model.User
	db.First(&newuser, 5)
	// newuser.Username = "usernamenew"
	// newuser.Password = "passwordnew"
	newuser.Book = append(newuser.Book, model.Book{Name: "booknew"})
	var achievement model.Achievement
	db.First(&achievement, 1)
	newuser.Achievement = append(newuser.Achievement, achievement)
	db.Save(&newuser)
	return e.JSON(http.StatusCreated, newuser)
}
