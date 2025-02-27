package db

import (
	"context"
	"database/sql"
	"fmt"

	userData "github.com/zchelalo/sa_user/pkg/sqlc/data/user/db"
)

type Querier interface {
	userData.Querier
}

type SQLStore struct {
	DB          *sql.DB
	UserQueries *userData.Queries
}

func New(db *sql.DB) *SQLStore {
	return &SQLStore{
		DB:          db,
		UserQueries: userData.New(db),
	}
}

func (store *SQLStore) ExecTx(ctx context.Context, fn func(*SQLStore) error) error {
	tx, err := store.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	userQueries := userData.New(tx)

	transactionStore := &SQLStore{
		DB:          store.DB,
		UserQueries: userQueries,
	}

	err = fn(transactionStore)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
