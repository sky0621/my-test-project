package db

import (
	"cloud.google.com/go/cloudsqlconn"
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/sky0621/my-test-project/app/internal/config"
	"github.com/sky0621/my-test-project/app/internal/logger"
	"net"
)

func NewDB(ctx context.Context, cfg config.Config, l logger.AppLogger) (*sql.DB, error) {
	network, addr := "tcp", fmt.Sprintf("%s:%s", cfg.DBHost, cfg.DBPort)

	if cfg.UseCloudSQL {
		dialer, err := cloudsqlconn.NewDialer(ctx)
		if err != nil {
			werr := fmt.Errorf("dialer 作成失敗: %w", err)
			l.Log(werr.Error())
			return nil, werr
		}
		network = "cloudsql-mysql"
		addr = cfg.DBHost
		mysql.RegisterDialContext(network, func(ctx context.Context, addr string) (net.Conn, error) {
			return dialer.Dial(ctx, addr)
		})
	}

	dsn := fmt.Sprintf(
		"%s:%s@%s(%s)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPass, network, addr, cfg.DBName,
	)
	l.Log(fmt.Sprintf("dsn: %v", dsn))

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open エラー: %w", err)
	}
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping エラー: %w", err)
	}
	return db, nil
}
