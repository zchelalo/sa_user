package userApplication

import (
	"context"
	"log"

	userDomain "github.com/zchelalo/sa_user/internal/modules/user/domain"
	"github.com/zchelalo/sa_user/pkg/config"
	"github.com/zchelalo/sa_user/pkg/meta"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCases struct {
	ctx            context.Context
	UserRepository userDomain.UserRepository
}

func NewUserUseCases(ctx context.Context, userRepository userDomain.UserRepository) *UserUseCases {
	return &UserUseCases{
		ctx:            ctx,
		UserRepository: userRepository,
	}
}

func (useCase *UserUseCases) Get(id string) (*userDomain.UserEntity, error) {
	err := userDomain.IsIdValid(id)
	if err != nil {
		return nil, err
	}

	return useCase.UserRepository.Get(id)
}

func (useCase *UserUseCases) GetPasswordHashAndID(email string) (*userDomain.HashedPasswordAndID, error) {
	err := userDomain.IsEmailValid(email)
	if err != nil {
		return nil, err
	}

	return useCase.UserRepository.GetPasswordHashAndID(email)
}

func (useCase *UserUseCases) GetAll(page, limit int32) ([]*userDomain.UserEntity, *meta.Meta, error) {
	usersCount, err := useCase.Count()
	if err != nil {
		return nil, nil, err
	}

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	meta, err := meta.New(page, limit, int32(usersCount), config.PaginatorLimitDefault)
	if err != nil {
		return nil, nil, err
	}

	usersObtained, err := useCase.UserRepository.GetAll(int32(meta.Offset()), int32(meta.Limit()))
	if err != nil {
		return nil, nil, err
	}

	return usersObtained, meta, nil
}

func (useCase *UserUseCases) Create(name, email, password string) (*userDomain.UserEntity, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user, err := userDomain.NewUserEntity(name, email, string(hashedPassword))
	if err != nil {
		return nil, err
	}

	return useCase.UserRepository.Create(user)
}

func (useCase *UserUseCases) Update(id string, name, email, password *string, verified *bool) (*userDomain.UserEntity, error) {
	userObtained, err := useCase.Get(id)
	if err != nil {
		return nil, err
	}

	userToUpdate := &userDomain.UserEntity{}
	userToUpdate.ID = userObtained.ID

	if name != nil {
		err = userDomain.IsNameValid(*name)
		if err != nil {
			return nil, err
		}
		userToUpdate.Name = *name
	} else {
		userToUpdate.Name = userObtained.Name
	}

	if email != nil {
		err = userDomain.IsEmailValid(*email)
		if err != nil {
			return nil, err
		}
		userToUpdate.Email = *email
	} else {
		userToUpdate.Email = userObtained.Email
	}

	if password != nil {
		err = userDomain.IsPasswordValid(*password)
		if err != nil {
			return nil, err
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		userToUpdate.Password = string(hashedPassword)
	} else {
		userToUpdate.Password = userObtained.Password
	}

	if verified != nil {
		userToUpdate.Verified = *verified
	} else {
		userToUpdate.Verified = userObtained.Verified
	}

	if userToUpdate.Name == userObtained.Name &&
		userToUpdate.Email == userObtained.Email &&
		userToUpdate.Password == userObtained.Password &&
		userToUpdate.Verified == userObtained.Verified {
		return userObtained, nil
	}

	return useCase.UserRepository.Update(userToUpdate)
}

func (useCase *UserUseCases) Delete(id string) error {
	_, err := useCase.Get(id)
	if err != nil {
		return err
	}

	return useCase.UserRepository.Delete(id)
}

func (useCase *UserUseCases) Count() (int32, error) {
	return useCase.UserRepository.Count()
}
