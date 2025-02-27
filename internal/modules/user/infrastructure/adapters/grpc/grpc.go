package userGRPC

import (
	userApplication "github.com/zchelalo/sa_user/internal/modules/user/application"
	"github.com/zchelalo/sa_user/pkg/proto"
)

type UserRouter struct {
	useCase *userApplication.UserUseCases
	proto.UnimplementedUserServiceServer
}

func New(userUseCase *userApplication.UserUseCases) *UserRouter {
	return &UserRouter{
		useCase: userUseCase,
	}
}
