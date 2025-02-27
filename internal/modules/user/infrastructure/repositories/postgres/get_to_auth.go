package userPostgresRepo

import (
	"context"
	"database/sql"

	userDomain "github.com/zchelalo/sa_user/internal/modules/user/domain"
	userErrors "github.com/zchelalo/sa_user/internal/modules/user/errors"
)

func (repo *PostgresRepository) GetToAuth(ctx context.Context, email string) (*userDomain.UserEntity, error) {
	userObtained, err := repo.store.UserQueries.GetUserToAuth(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, userErrors.ErrUserNotFound
		}

		return nil, err
	}

	result := &userDomain.UserEntity{
		ID:       userObtained.ID,
		Name:     userObtained.Name,
		Email:    userObtained.Email,
		Password: userObtained.Password,
		Verified: userObtained.Verified,
	}

	return result, nil
}
