package main

import (
	"github.com/earthrockey/Golang-Structure/config"
	"github.com/earthrockey/Golang-Structure/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.CreateTable()
	e := echo.New()
	setMiddlewere(e)

	// Homepage
	e.GET("/", controller.HomePage)

	// User
	e.GET("/api/user/:id", controller.GetIDUser)
	e.GET("/api/user", controller.GetAllUser)
	e.POST("/api/user", controller.CreateUser)
	e.PUT("/api/user/:id", controller.EditUser)
	e.DELETE("/api/user/:id", controller.DeleteUser)
	e.Logger.Fatal(e.Start(":8080"))
}

func setMiddlewere(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:          middleware.DefaultSkipper,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.POST, echo.DELETE},
		AllowCredentials: true,
	}))
}
