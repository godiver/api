package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Books(c echo.Context) error {
	// url := "https://www.googleapis.com/books/v1/volumes"
	url := "https://app.rakuten.co.jp/services/api/BooksBook/Search/20170404?applicationId=1011959466247363309&title=ファクト&booksGenreId=001&sort=sales"
	res, err := http.Get(url)
	if err != nil {
		return c.String(http.StatusBadRequest, "時間をおいて再度実行してください。")
	}
	return c.JSON(http.StatusOK, res)
}
