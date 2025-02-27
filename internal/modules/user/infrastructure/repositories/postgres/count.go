package userPostgresRepo

import (
	"context"
	"database/sql"

	userErrors "github.com/zchelalo/sa_user/internal/modules/user/errors"
)

func (repo *PostgresRepository) Count(ctx context.Context) (int32, error) {
	count, err := repo.store.UserQueries.CountUsers(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, userErrors.ErrUsersNotFound
		}

		return 0, err
	}

	return int32(count), nil
}
