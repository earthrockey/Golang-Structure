package controller

import (
	"fmt"
	"net/http"

	"github.com/earthrockey/Golang-Structure/config"
	"github.com/earthrockey/Golang-Structure/model"
	"github.com/labstack/echo/v4"
)

func GetAllBook(e echo.Context) error {
	db, err := config.ConnectDB()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return echo.ErrInternalServerError
	}
	var book []model.Book
	db.Find(&book)
	return e.JSON(http.StatusCreated, book)
}

func GetIDBook(e echo.Context) error {
	db, err := config.ConnectDB()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return echo.ErrInternalServerError
	}
	var book model.Book
	db.First(&book, e.Param("id"))
	if book.Name == "" {
		return echo.NewHTTPError(400, "id not found")
	}
	return e.JSON(http.StatusCreated, book)
}

func CreateBook(e echo.Context) error {
	var req model.BookRequest
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
	newBook := model.Book{Name: req.Name, UserID: req.UserID}
	db.Create(&newBook)
	return e.JSON(http.StatusCreated, newBook)
}

func EditBook(e echo.Context) error {
	var req model.BookRequest
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
	var editBook model.Book
	db.First(&editBook, e.Param("id"))
	if editBook.Name == "" {
		return echo.NewHTTPError(400, "id not found")
	}
	editBook.Name = req.Name
	editBook.UserID = req.UserID
	db.Save(&editBook)
	return e.JSON(http.StatusCreated, editBook)
}

func DeleteBook(e echo.Context) error {
	db, err := config.ConnectDB()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return echo.ErrInternalServerError
	}
	var deleteBook model.Book
	db.First(&deleteBook, e.Param("id"))
	if deleteBook.Name != "" {
		db.Delete(&deleteBook)
	} else {
		return echo.NewHTTPError(400, "id not found")
	}
	return e.JSON(http.StatusCreated, deleteBook)
}
