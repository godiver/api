package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"app/config/env"

	"github.com/labstack/echo/v4"
)

func Books(c echo.Context) error {
	id := env.Env.API.Rakuten.ID
	booksGenreID := "001006007" // マネジメント・人材管理
	// booksGenreID := "001006009" // 自己啓発
	url := fmt.Sprintf("https://app.rakuten.co.jp/services/api/BooksBook/Search/20170404?applicationId=%s&booksGenreId=%s&sort=sales", id, booksGenreID)
	res, err := http.Get(url)
	if err != nil {
		return c.String(http.StatusBadRequest, "時間をおいて再度実行してください。")
	}
	defer res.Body.Close()
	var r io.Reader = res.Body
	r = io.TeeReader(r, os.Stderr)
	var books interface{}
	err = json.NewDecoder(r).Decode(&books)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, books)
}

type Parameter struct {
	Title string `json:"title" form:"title" query:"title"`
}

func SearchBooks(c echo.Context) error {
	id := env.Env.API.Rakuten.ID
	parameter := new(Parameter)
	if err := c.Bind(parameter); err != nil {
		return c.String(http.StatusBadRequest, "時間をおいて再度実行してください。")
	}

	url := fmt.Sprintf("https://app.rakuten.co.jp/services/api/BooksBook/Search/20170404?applicationId=%s&sort=sales&title=%s", id, parameter.Title)
	res, err := http.Get(url)
	if err != nil {
		return c.String(http.StatusBadRequest, "時間をおいて再度実行してください。")
	}
	defer res.Body.Close()
	var r io.Reader = res.Body
	r = io.TeeReader(r, os.Stderr)
	var books interface{}
	err = json.NewDecoder(r).Decode(&books)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, books)
}
