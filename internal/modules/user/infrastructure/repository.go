package userInfrastructure

import (
	"context"

	userDomain "github.com/zchelalo/sa_user/internal/modules/user/domain"
	userDb "github.com/zchelalo/sa_user/pkg/sqlc/user/db"
)

type UserRepository struct {
	ctx   context.Context
	store userDb.Store
}

func NewUserRepository(ctx context.Context, store userDb.Store) userDomain.UserRepository {
	return &UserRepository{
		ctx:   ctx,
		store: store,
	}
}

func (repo *UserRepository) Get(id string) (*userDomain.UserEntity, error) {
	userObtained, err := repo.store.GetUser(repo.ctx, id)
	if err != nil {
		return nil, err
	}

	return &userDomain.UserEntity{
		ID:       userObtained.ID,
		Name:     userObtained.Name,
		Email:    userObtained.Email,
		Password: userObtained.Password,
		Verified: userObtained.Verified,
	}, nil
}

func (repo *UserRepository) GetPasswordHashAndID(email string) (*userDomain.HashedPasswordAndID, error) {
	passwordAndID, err := repo.store.GetUserPasswordHashAndID(repo.ctx, email)
	if err != nil {
		return nil, err
	}

	result := &userDomain.HashedPasswordAndID{
		ID:       passwordAndID.ID,
		Password: passwordAndID.Password,
	}

	return result, nil
}

func (repo *UserRepository) GetAll(offset, limit int32) ([]*userDomain.UserEntity, error) {
	arg := userDb.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	}
	usersObtained, err := repo.store.ListUsers(repo.ctx, arg)
	if err != nil {
		return nil, err
	}

	users := make([]*userDomain.UserEntity, 0)
	for _, userObtained := range usersObtained {
		users = append(users, &userDomain.UserEntity{
			ID:       userObtained.ID,
			Name:     userObtained.Name,
			Email:    userObtained.Email,
			Password: userObtained.Password,
			Verified: userObtained.Verified,
		})
	}

	return users, nil
}

func (repo *UserRepository) Create(user *userDomain.UserEntity) (*userDomain.UserEntity, error) {
	arg := userDb.CreateUserParams{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Verified: user.Verified,
	}
	userCreated, err := repo.store.CreateUser(repo.ctx, arg)
	if err != nil {
		return nil, err
	}

	return &userDomain.UserEntity{
		ID:       userCreated.ID,
		Name:     userCreated.Name,
		Email:    userCreated.Email,
		Password: userCreated.Password,
		Verified: userCreated.Verified,
	}, nil
}

func (repo *UserRepository) Update(user *userDomain.UserEntity) (*userDomain.UserEntity, error) {
	arg := userDb.UpdateUserParams{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Verified: user.Verified,
	}
	userUpdated, err := repo.store.UpdateUser(repo.ctx, arg)
	if err != nil {
		return nil, err
	}

	return &userDomain.UserEntity{
		ID:       userUpdated.ID,
		Name:     userUpdated.Name,
		Email:    userUpdated.Email,
		Password: userUpdated.Password,
		Verified: userUpdated.Verified,
	}, nil
}

func (repo *UserRepository) Delete(id string) error {
	err := repo.store.DeleteUser(repo.ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) Count() (int32, error) {
	count, err := repo.store.CountUsers(repo.ctx)
	if err != nil {
		return 0, err
	}

	return int32(count), nil
}
