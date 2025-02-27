package userApplication

import (
	"context"

	userDomain "github.com/zchelalo/sa_user/internal/modules/user/domain"
	"golang.org/x/crypto/bcrypt"
)

func (useCase *UserUseCases) Create(ctx context.Context, name, email, password string) (*userDomain.UserEntity, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user, err := userDomain.NewUserEntity(name, email, string(hashedPassword))
	if err != nil {
		return nil, err
	}

	return useCase.userRepository.Create(ctx, user)
}
