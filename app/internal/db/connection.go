package db

import (
	"cloud.google.com/go/cloudsqlconn"
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/sky0621/my-test-project/app/internal/config"
	"net"
)

func NewDBConnection(ctx context.Context, cfg *config.DBConfig) (*sql.DB, error) {
	if cfg.IsCloudSQL {
		return connectToCloudSQL(ctx, cfg)
	}
	return connectToLocalMySQL(cfg)
}

func connectToCloudSQL(ctx context.Context, cfg *config.DBConfig) (*sql.DB, error) {
	mysql.RegisterDialContext("cloudsql-mysql",
		func(ctx context.Context, addr string) (net.Conn, error) {
			dialer, err := cloudsqlconn.NewDialer(ctx)
			if err != nil {
				return nil, err
			}

			instanceConnName := fmt.Sprintf("%s:%s:%s", cfg.ProjectID, cfg.Region, cfg.Instance)
			return dialer.Dial(ctx, instanceConnName)
		})

	dsn := fmt.Sprintf("%s:%s@cloudsql-mysql(localhost:3306)/%s?parseTime=true",
		cfg.User, cfg.Password, cfg.DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, err
	}

	return sql.Open("mysql", dsn)
}

func connectToLocalMySQL(cfg *config.DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, err
	}

	return db, nil
}
