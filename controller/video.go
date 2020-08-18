package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetVideos(c echo.Context) error {
	jsonMap := map[string]string{
		"foo": "bar",
		"hoge": "fuga",
	}
	return c.JSON(http.StatusOK, jsonMap)
}
