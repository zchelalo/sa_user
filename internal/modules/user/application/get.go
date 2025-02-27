package userApplication

import (
	"context"

	userDomain "github.com/zchelalo/sa_user/internal/modules/user/domain"
)

func (useCase *UserUseCases) Get(ctx context.Context, id string) (*userDomain.UserEntity, error) {
	err := userDomain.IsIdValid(id)
	if err != nil {
		return nil, err
	}

	return useCase.userRepository.Get(ctx, id)
}
