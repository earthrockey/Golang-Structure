package controller

import (
	"fmt"
	"net/http"

	"github.com/earthrockey/Golang-Structure/config"
	"github.com/earthrockey/Golang-Structure/model"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUser(e echo.Context) error {
	db, err := config.ConnectDB()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return echo.ErrInternalServerError
	}
	var user []model.User
	db.Find(&user)
	for i := range user {
		db.Model(&user[i]).Related(&user[i].Book)
		db.Model(&user[i]).Association("Achievement").Find(&user[i].Achievement)
	}
	return e.JSON(http.StatusCreated, user)
}

func GetIDUser(e echo.Context) error {
	db, err := config.ConnectDB()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return echo.ErrInternalServerError
	}
	var user model.User
	db.First(&user, e.Param("id"))
	if user.Username == "" {
		return echo.NewHTTPError(400, "id not found")
	}
	db.Model(&user).Related(&user.Book)
	db.Model(&user).Association("Achievement").Find(&user.Achievement)
	return e.JSON(http.StatusCreated, user)
}

func CreateUser(e echo.Context) error {
	var req model.UserRequest
	if err := e.Bind(&req); err != nil {
		fmt.Println(err)
		return echo.ErrBadRequest
	}
	db, err := config.ConnectDB()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return echo.ErrInternalServerError
	}
	hashpass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	newUser := model.User{Username: req.Username, Password: string(hashpass)}
	db.Create(&newUser)
	return e.JSON(http.StatusCreated, newUser)
}

func EditUser(e echo.Context) error {
	var req model.UserRequest
	if err := e.Bind(&req); err != nil {
		fmt.Println(err)
		return echo.ErrBadRequest
	}
	db, err := config.ConnectDB()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return echo.ErrInternalServerError
	}
	var editUser model.User
	db.First(&editUser, e.Param("id"))
	if editUser.Username == "" {
		return echo.NewHTTPError(400, "id not found")
	}
	editUser.Username = req.Username
	hashpass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	editUser.Password = string(hashpass)
	for _, value := range req.AchievementID {
		var achievement model.Achievement
		db.First(&achievement, value)
		editUser.Achievement = append(editUser.Achievement, achievement)
	}
	db.Save(&editUser)
	return e.JSON(http.StatusCreated, editUser)
}

func DeleteUser(e echo.Context) error {
	db, err := config.ConnectDB()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return echo.ErrInternalServerError
	}
	var deleteUser model.User
	db.First(&deleteUser, e.Param("id"))
	if deleteUser.Username != "" {
		db.Delete(&deleteUser)
	} else {
		return echo.NewHTTPError(400, "id not found")
	}
	return e.JSON(http.StatusCreated, deleteUser)
}
