package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/earthrockey/Golang-Structure/config"
	"github.com/earthrockey/Golang-Structure/model"
	"github.com/gorilla/securecookie"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(e echo.Context) error {
	var req model.LoginRequest
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
	var user model.User
	db.Where("username = ?", req.Username).First(&user)
	if user.Username == "" {
		return echo.NewHTTPError(400, "username not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(400, "password not correct")
	}

	var hashKey = []byte("very-secret")
	var blockKey = []byte("a-lot-secret1234")
	var s = securecookie.New(hashKey, blockKey)
	// value := map[string]string{
	// 	"username": req.Username,
	// }
	encoded, err := s.Encode("golang-structure", req)
	if err != nil {
		fmt.Println(err)
	}
	cookie := new(http.Cookie)
	cookie.Name = "golang-structure"
	cookie.Value = encoded
	cookie.Expires = time.Now().Add(24 * time.Hour)
	e.SetCookie(cookie)
	return e.JSON(http.StatusOK, "Login!")
}

func CheckAuthentication(e echo.Context) error {
	cookie, err := e.Cookie("golang-structure")
	if err != nil {
		return err
	}
	var hashKey = []byte("very-secret")
	var blockKey = []byte("a-lot-secret1234")
	var s = securecookie.New(hashKey, blockKey)
	var req model.LoginRequest
	err = s.Decode("golang-structure", cookie.Value, &req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(req.Username)
	return e.JSON(http.StatusOK, "Check Authentication!")
}
