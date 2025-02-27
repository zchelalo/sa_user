package userApplication

import (
	userDomain "github.com/zchelalo/sa_user/internal/modules/user/domain"
)

type UserUseCases struct {
	userRepository userDomain.UserRepository
}

func New(userRepository userDomain.UserRepository) *UserUseCases {
	return &UserUseCases{
		userRepository: userRepository,
	}
}
