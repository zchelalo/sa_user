package userApplication

import (
	"context"

	userDomain "github.com/zchelalo/sa_user/internal/modules/user/domain"
)

func (useCase *UserUseCases) GetToAuth(ctx context.Context, email string) (*userDomain.UserEntity, error) {
	err := userDomain.IsEmailValid(email)
	if err != nil {
		return nil, err
	}

	return useCase.userRepository.GetToAuth(ctx, email)
}
