package main

import (
	"github.com/C3LOUD/the-heart-stack/db"
	"github.com/C3LOUD/the-heart-stack/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	db.Conn()
	defer db.Close()
	e.Static("/static", "static")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	indexHandler := handler.IndexHandler{}
	e.GET("/", indexHandler.HandleIndex)

	e.Start(":3000")
}
