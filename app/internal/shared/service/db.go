package service

import (
	"context"
	"database/sql"
)

func WithTransaction(
	ctx context.Context,
	db *sql.DB,
	fn func(tx *sql.Tx) error,
) (err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		// パニック時はロールバックして再パニック
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
		if err != nil {
			_ = tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	return fn(tx)
}
