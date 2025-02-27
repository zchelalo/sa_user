package userPostgresRepo

import (
	"context"

	userDomain "github.com/zchelalo/sa_user/internal/modules/user/domain"
	userData "github.com/zchelalo/sa_user/pkg/sqlc/data/user/db"
)

func (repo *PostgresRepository) Update(ctx context.Context, user *userDomain.UserEntity) (*userDomain.UserEntity, error) {
	arg := userData.UpdateUserParams{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Verified: user.Verified,
	}
	userUpdated, err := repo.store.UserQueries.UpdateUser(ctx, arg)
	if err != nil {
		return nil, err
	}

	return &userDomain.UserEntity{
		ID:       userUpdated.ID,
		Name:     userUpdated.Name,
		Email:    userUpdated.Email,
		Verified: userUpdated.Verified,
	}, nil
}
