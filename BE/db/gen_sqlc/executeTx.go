package db

import (
	"context"
	"fmt"
)

// ExecTx executes a function within a database transaction
func (SQLStore *SQLStorage) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := SQLStore.conn.Begin(ctx)

	if err != nil {
		return err
	}
	newQueriesWithTx := SQLStore.WithTx(tx)
	err = fn(newQueriesWithTx)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(ctx)
}
