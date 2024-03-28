package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sagoresarker/firecracker-rest-api/handlers"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/create-bridge", handlers.CreateBridge)
	e.Logger.Fatal(e.Start(":8080"))
}
