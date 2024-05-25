package main

import (
	"github.com/baselrabia/echo-task/handlers"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/route", handlers.ReconstructRoute)

	e.Logger.Fatal(e.Start(":8080"))

}
