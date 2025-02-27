package userPostgresRepo

import (
	"context"

	"github.com/lib/pq"
	userDomain "github.com/zchelalo/sa_user/internal/modules/user/domain"
	userErrors "github.com/zchelalo/sa_user/internal/modules/user/errors"
	userData "github.com/zchelalo/sa_user/pkg/sqlc/data/user/db"
)

func (repo *PostgresRepository) Create(ctx context.Context, user *userDomain.UserEntity) (*userDomain.UserEntity, error) {
	arg := userData.CreateUserParams{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Verified: user.Verified,
	}
	userCreated, err := repo.store.UserQueries.CreateUser(ctx, arg)
	if err != nil {
		if pqError, ok := err.(*pq.Error); ok {
			switch pqError.Code.Name() {
			case "unique_violation":
				return nil, userErrors.ErrEmailAlreadyExists
			}
		}

		return nil, err
	}

	return &userDomain.UserEntity{
		ID:       userCreated.ID,
		Name:     userCreated.Name,
		Email:    userCreated.Email,
		Verified: userCreated.Verified,
	}, nil
}
