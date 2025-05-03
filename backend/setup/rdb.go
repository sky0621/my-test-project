package setup

import (
	"cloud.google.com/go/cloudsqlconn"
	cloudsqlmysql "cloud.google.com/go/cloudsqlconn/mysql/mysql"
	"context"
	"database/sql"
	"fmt"
	"github.com/sky0621/my-test-project/backend/internal/shared/config"
)

func newDB(ctx context.Context, cfg config.Config) (*sql.DB, error) {
	network, addr := "tcp", fmt.Sprintf("%s:%s", cfg.DBHost, cfg.DBPort)

	if cfg.UseCloudSQL {
		cleanup, err := cloudsqlmysql.RegisterDriver(
			"cloudsql-mysql",
			cloudsqlconn.WithIAMAuthN(),
		)
		if err != nil {
			return nil, fmt.Errorf("RegisterDriver: %w", err)
		}
		defer func() { _ = cleanup() }()
	}

	dsn := fmt.Sprintf(
		"%s:%s@%s(%s)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPass, network, addr, cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open エラー: %w", err)
	}
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping エラー: %w", err)
	}
	return db, nil
}
