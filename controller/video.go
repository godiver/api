package video

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Index() echo.HandlerFunc {
	return func(c echo.Context) error {     //c をいじって Request, Responseを色々する
    return c.String(http.StatusOK, "Hello World")
  }
}
