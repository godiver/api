package router

import (
	"fmt"
	"os"

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

	// 環境変数からPORTを取得する
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
