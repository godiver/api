package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"app/controller"
)

func Router() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", controller.Hello)
	e.GET("/api/v1/books", controller.Books)
	e.GET("/api/v1/books/search", controller.SearchBooks)
	e.GET("/api/v1/books/videos", controller.GetVideos)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
