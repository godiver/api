package router

import (
	"fmt"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"app/controller"
	"app/graph"
	"app/graph/generated"
)

func Router() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/", controller.Hello)
	e.GET("/api/v1/books", controller.Books)
	e.GET("/api/v1/books/search/:title", controller.SearchBooks)
	e.GET("/api/v1/books/videos/:title", controller.GetVideos)
	e.GET("/api/v1/watch/:videoId", controller.GetVideo)

	// 環境変数からPORTを取得する
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	graphqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	playgroundHandler := playground.Handler("GraphQL playground", "/query")
	e.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})
	e.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
