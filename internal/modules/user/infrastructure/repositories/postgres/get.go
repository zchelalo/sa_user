package userPostgresRepo

import (
	"context"
	"database/sql"

	userDomain "github.com/zchelalo/sa_user/internal/modules/user/domain"
	userErrors "github.com/zchelalo/sa_user/internal/modules/user/errors"
)

func (repo *PostgresRepository) Get(ctx context.Context, id string) (*userDomain.UserEntity, error) {
	userObtained, err := repo.store.UserQueries.GetUser(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, userErrors.ErrUserNotFound
		}

		return nil, err
	}

	return &userDomain.UserEntity{
		ID:       userObtained.ID,
		Name:     userObtained.Name,
		Email:    userObtained.Email,
		Verified: userObtained.Verified,
	}, nil
}
