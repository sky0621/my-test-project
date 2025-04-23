package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sky0621/my-test-project/app/internal/config"
	"github.com/sky0621/my-test-project/app/internal/db"
	"log"
)

func main() {
	ctx := context.Background()
	cfg := config.NewDBConfig()
	sqlDB, err := db.NewDBConnection(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	driver, _ := mysql.WithInstance(sqlDB, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)

	if m == nil {
		fmt.Println("no migration")
		return
	}
	if config.IsMigrateUp() {
		if err := m.Up(); err != nil {
			fmt.Println(err)
			return
		}
		return
	}
	if config.IsMigrateDown() {
		if err := m.Down(); err != nil {
			fmt.Println(err)
			return
		}
		return
	}
	fmt.Println("no operation")
}
