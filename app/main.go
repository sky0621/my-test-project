package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ヘルスチェック用エンドポイント
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	e.GET("/hello", func(c echo.Context) error {
		name := c.QueryParam("name")
		if name == "" {
			name = "world"
		}
		return c.String(http.StatusOK, "Hello, "+name+"!")
	})

	// ポート 8080 で起動
	e.Logger.Fatal(e.Start(":8080"))
}
