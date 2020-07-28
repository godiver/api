package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"src/github.com/godiver/api/controller"
)

func Router() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", controller.Hello)
	e.GET("/api/v1/books/:book_id/videos", video.Index())

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}