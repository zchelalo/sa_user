package userApplication

import (
	"context"

	userDomain "github.com/zchelalo/sa_user/internal/modules/user/domain"
	"golang.org/x/crypto/bcrypt"
)

func (useCase *UserUseCases) Update(ctx context.Context, id string, name, email, password *string, verified *bool) (*userDomain.UserEntity, error) {
	userObtained, err := useCase.Get(ctx, id)
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

	return useCase.userRepository.Update(ctx, userToUpdate)
}
