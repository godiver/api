package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"src/github.com/godiver/api/config/env"

	"github.com/labstack/echo/v4"
)

func Books(c echo.Context) error {
	id := env.Env.API.Rakuten.ID
	url := fmt.Sprintf("https://app.rakuten.co.jp/services/api/BooksBook/Search/20170404?applicationId=%s&booksGenreId=001006&sort=sales", id)
	res, err := http.Get(url)
	if err != nil {
		return c.String(http.StatusBadRequest, "時間をおいて再度実行してください。")
	}
	defer res.Body.Close()
	var r io.Reader = res.Body
	r = io.TeeReader(r, os.Stderr)
	var foo interface{}
	err = json.NewDecoder(r).Decode(&foo)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, res)
}
