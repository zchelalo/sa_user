package userPostgresRepo

import (
	"context"
	"database/sql"

	userDomain "github.com/zchelalo/sa_user/internal/modules/user/domain"
	userErrors "github.com/zchelalo/sa_user/internal/modules/user/errors"
	userData "github.com/zchelalo/sa_user/pkg/sqlc/data/user/db"
)

func (repo *PostgresRepository) GetAll(ctx context.Context, offset, limit int32) ([]*userDomain.UserEntity, error) {
	arg := userData.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	}
	usersObtained, err := repo.store.UserQueries.ListUsers(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, userErrors.ErrUsersNotFound
		}

		return nil, err
	}

	if len(usersObtained) == 0 {
		return nil, userErrors.ErrUsersNotFound
	}

	users := make([]*userDomain.UserEntity, 0)
	for _, userObtained := range usersObtained {
		users = append(users, &userDomain.UserEntity{
			ID:       userObtained.ID,
			Name:     userObtained.Name,
			Email:    userObtained.Email,
			Verified: userObtained.Verified,
		})
	}

	return users, nil
}
