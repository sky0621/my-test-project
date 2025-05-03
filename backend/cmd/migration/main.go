package main

import (
	"context"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sky0621/my-test-project/backend/internal/shared/config"
	"github.com/sky0621/my-test-project/backend/internal/shared/rdb"
)

func main() {
	var up, down, test bool
	flag.BoolVar(&up, "up", false, "全マイグレーション適用")
	flag.BoolVar(&down, "down", false, "全マイグレーションロールバック")
	flag.BoolVar(&test, "test", false, "テスト用DBか否か")
	flag.Parse()

	ctx := context.Background()
	var cfg config.Config
	if test {
		cfg = config.NewTestConfig()
	} else {
		cfg = config.NewConfig()
	}
	sqlDB, err := rdb.NewDB(ctx, cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	driver, err := mysql.WithInstance(sqlDB, &mysql.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	if driver == nil {
		fmt.Println("driver is nil")
		return
	}
	m, _ := migrate.NewWithDatabaseInstance(
		"file://schema/db",
		"mysql",
		driver,
	)

	if m == nil {
		fmt.Println("no migration")
		return
	}
	if up {
		if err := m.Up(); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("migration up success")
		return
	}
	if down {
		if err := m.Down(); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("migration down success")
		return
	}
	fmt.Println("no order")
}
