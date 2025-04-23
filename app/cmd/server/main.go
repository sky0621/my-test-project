package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sky0621/my-test-project/app/internal/config"
	"github.com/sky0621/my-test-project/app/internal/db"
	"github.com/sky0621/my-test-project/app/internal/infra"
	"log"
	"net/http"
	"strconv"
)

func main() {
	ctx := context.Background()
	cfg := config.NewDBConfig()
	sqlDB, err := db.NewDBConnection(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	q := infra.New(sqlDB)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ヘルスチェック用エンドポイント
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	e.GET("/users/create", func(c echo.Context) error {
		name := c.QueryParam("name")
		userID, err := q.CreateUser(c.Request().Context(), name)
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, fmt.Sprintf("User ID: %d, Name: %s", userID, name))
	})

	e.GET("/users/get", func(c echo.Context) error {
		id := c.QueryParam("id")
		i64, err := strconv.ParseInt(id, 10, 32)
		if err != nil {
			return err
		}
		user, err := q.GetUser(c.Request().Context(), i64)
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, fmt.Sprintf("User ID: %d, Name: %s, CreatedAt: %v", user.ID, user.Name, user.CreatedAt))
	})

	e.GET("/users/list", func(c echo.Context) error {
		users, err := q.ListUsers(c.Request().Context())
		if err != nil {
			return err
		}
		for _, user := range users {
			fmt.Printf("User ID: %d, Name: %s, CreatedAt: %v\n", user.ID, user.Name, user.CreatedAt)
		}
		return c.String(http.StatusOK, fmt.Sprintf("User length: %d", len(users)))
	})

	// ポート 8080 で起動
	e.Logger.Fatal(e.Start(":8080"))
}
