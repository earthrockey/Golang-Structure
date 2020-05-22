package controller

import (
	"fmt"
	"net/http"

	"github.com/earthrockey/Golang-Structure/config"
	"github.com/earthrockey/Golang-Structure/model"
	"github.com/labstack/echo/v4"
)

func GetAllAchievement(e echo.Context) error {
	db, err := config.ConnectDB()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return echo.ErrInternalServerError
	}
	var achievement []model.Achievement
	db.Find(&achievement)
	return e.JSON(http.StatusCreated, achievement)
}

func GetIDAchievement(e echo.Context) error {
	db, err := config.ConnectDB()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return echo.ErrInternalServerError
	}
	var achievement model.Achievement
	db.First(&achievement, e.Param("id"))
	if achievement.Name == "" {
		return echo.NewHTTPError(400, "id not found")
	}
	return e.JSON(http.StatusCreated, achievement)
}

func CreateAchievement(e echo.Context) error {
	var req model.AchievementRequest
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
	newAchievement := model.Achievement{Name: req.Name}
	db.Create(&newAchievement)
	return e.JSON(http.StatusCreated, newAchievement)
}

func EditAchievement(e echo.Context) error {
	var req model.AchievementRequest
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
	var editAchievement model.Achievement
	db.First(&editAchievement, e.Param("id"))
	if editAchievement.Name == "" {
		return echo.NewHTTPError(400, "id not found")
	}
	editAchievement.Name = req.Name
	db.Save(&editAchievement)
	return e.JSON(http.StatusCreated, editAchievement)
}

func DeleteAchievement(e echo.Context) error {
	db, err := config.ConnectDB()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return echo.ErrInternalServerError
	}
	var deleteAchievement model.Achievement
	db.First(&deleteAchievement, e.Param("id"))
	if deleteAchievement.Name != "" {
		db.Delete(&deleteAchievement)
	} else {
		return echo.NewHTTPError(400, "id not found")
	}
	return e.JSON(http.StatusCreated, deleteAchievement)
}
