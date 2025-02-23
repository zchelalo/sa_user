package userInfrastructure

import (
	"context"

	userApplication "github.com/zchelalo/sa_user/internal/modules/user/application"
	Errors "github.com/zchelalo/sa_user/internal/modules/user/errors"
	"github.com/zchelalo/sa_user/pkg/proto"
	userDb "github.com/zchelalo/sa_user/pkg/sqlc/user/db"
	"github.com/zchelalo/sa_user/pkg/util"
	"google.golang.org/grpc/codes"
)

type UserRouter struct {
	ctx     context.Context
	useCase *userApplication.UserUseCases
	proto.UnimplementedUserServiceServer
}

func NewUserRouter(store userDb.Store, ctx context.Context) *UserRouter {
	userRepository := NewUserRepository(ctx, store)
	userUseCase := userApplication.NewUserUseCases(ctx, userRepository)

	return &UserRouter{
		ctx:     ctx,
		useCase: userUseCase,
	}
}

func (userRouter *UserRouter) GetUsers(ctx context.Context, req *proto.GetUsersRequest) (*proto.GetUsersResponse, error) {
	// Create a new context with the request context
	// ctx = context.WithValue(userRouter.ctx, "requestCtx", ctx)

	usersObtained, meta, err := userRouter.useCase.GetAll(req.GetPage(), req.GetLimit())
	if err != nil {
		Error := &proto.Error{}

		if err == Errors.ErrUsersNotFound {
			Error.Code = int32(codes.NotFound)
			Error.Message = err.Error()
		}

		responseError := &proto.GetUsersResponse_Error{
			Error: Error,
		}

		responseProto := &proto.GetUsersResponse{
			Result: responseError,
		}

		return responseProto, nil
	}

	metaProto := &proto.Meta{
		Page:       meta.Page,
		PerPage:    meta.PerPage,
		Count:      meta.PageCount,
		TotalCount: meta.TotalCount,
	}

	usersProto := []*proto.UserData{}
	for _, user := range usersObtained {
		protoUser := &proto.UserData{
			Id:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Verified: user.Verified,
		}

		usersProto = append(usersProto, protoUser)
	}

	responseDataProto := &proto.GetUsersResponse_Data{
		Data: &proto.UsersWithMeta{
			Users: usersProto,
			Meta:  metaProto,
		},
	}

	responseProto := &proto.GetUsersResponse{
		Result: responseDataProto,
	}

	return responseProto, nil
}

func (userRouter *UserRouter) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	userObtained, err := userRouter.useCase.Get(req.GetId())
	if err != nil {
		Error := &proto.Error{}

		if err == Errors.ErrUserNotFound {
			Error.Code = int32(codes.NotFound)
			Error.Message = err.Error()
		} else if err == Errors.ErrIdRequired || err == Errors.ErrIdInvalid {
			Error.Code = int32(codes.InvalidArgument)
			Error.Message = err.Error()
		} else {
			Error.Code = int32(codes.Internal)
			Error.Message = err.Error()
		}

		responseError := &proto.GetUserResponse_Error{
			Error: Error,
		}

		responseProto := &proto.GetUserResponse{
			Result: responseError,
		}

		return responseProto, nil
	}

	protoUser := &proto.UserData{
		Id:       userObtained.ID,
		Name:     userObtained.Name,
		Email:    userObtained.Email,
		Verified: userObtained.Verified,
	}

	responseDataProto := &proto.GetUserResponse_User{
		User: protoUser,
	}

	responseProto := &proto.GetUserResponse{
		Result: responseDataProto,
	}

	return responseProto, nil
}

func (userRouter *UserRouter) GetUserToAuth(ctx context.Context, req *proto.GetUserToAuthRequest) (*proto.GetUserToAuthResponse, error) {
	userObtained, err := userRouter.useCase.GetToAuth(req.GetEmail())
	if err != nil {
		Error := &proto.Error{}

		if err == Errors.ErrUserNotFound {
			Error.Code = int32(codes.NotFound)
			Error.Message = err.Error()
		} else if err == Errors.ErrEmailRequired || err == Errors.ErrEmailInvalid {
			Error.Code = int32(codes.InvalidArgument)
			Error.Message = err.Error()
		} else {
			Error.Code = int32(codes.Internal)
			Error.Message = err.Error()
		}

		responseError := &proto.GetUserToAuthResponse_Error{
			Error: Error,
		}

		responseProto := &proto.GetUserToAuthResponse{
			Result: responseError,
		}

		return responseProto, nil
	}

	protoUser := &proto.UserWithPassword{
		Id:       userObtained.ID,
		Name:     userObtained.Name,
		Email:    userObtained.Email,
		Password: userObtained.Password,
		Verified: userObtained.Verified,
	}

	responseDataProto := &proto.GetUserToAuthResponse_User{
		User: protoUser,
	}

	responseProto := &proto.GetUserToAuthResponse{
		Result: responseDataProto,
	}

	return responseProto, nil
}

func (userRouter *UserRouter) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	userCreated, err := userRouter.useCase.Create(req.GetName(), req.GetEmail(), req.GetPassword())
	if err != nil {
		Error := &proto.Error{}

		errorInvalidArgument := []error{
			Errors.ErrNameRequired,
			Errors.ErrNameInvalid,
			Errors.ErrEmailRequired,
			Errors.ErrEmailInvalid,
			Errors.ErrPasswordRequired,
			Errors.ErrPasswordInvalid,
		}

		isInvalidArgumentError := util.IsErrorType(err, errorInvalidArgument)

		if isInvalidArgumentError {
			Error.Code = int32(codes.InvalidArgument)
			Error.Message = err.Error()
		} else if err == Errors.ErrEmailAlreadyExists {
			Error.Code = int32(codes.AlreadyExists)
			Error.Message = err.Error()
		} else {
			Error.Code = int32(codes.Internal)
			Error.Message = err.Error()
		}

		responseError := &proto.CreateUserResponse_Error{
			Error: Error,
		}

		responseProto := &proto.CreateUserResponse{
			Result: responseError,
		}

		return responseProto, nil
	}

	protoUser := &proto.UserData{
		Id:       userCreated.ID,
		Name:     userCreated.Name,
		Email:    userCreated.Email,
		Verified: userCreated.Verified,
	}

	responseDataProto := &proto.CreateUserResponse_User{
		User: protoUser,
	}

	responseProto := &proto.CreateUserResponse{
		Result: responseDataProto,
	}

	return responseProto, nil
}

func (userRouter *UserRouter) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {

	userUpdated, err := userRouter.useCase.Update(req.Id, req.Name, req.Email, req.Password, req.Verified)
	if err != nil {
		Error := &proto.Error{}

		errorInvalidArgument := []error{
			Errors.ErrIdRequired,
			Errors.ErrIdInvalid,
			Errors.ErrNameInvalid,
			Errors.ErrEmailInvalid,
			Errors.ErrPasswordInvalid,
			Errors.ErrVerifiedInvalid,
		}

		isInvalidArgumentError := util.IsErrorType(err, errorInvalidArgument)

		if isInvalidArgumentError {
			Error.Code = int32(codes.InvalidArgument)
			Error.Message = err.Error()
		} else if err == Errors.ErrUserNotFound {
			Error.Code = int32(codes.NotFound)
			Error.Message = err.Error()
		} else if err == Errors.ErrEmailAlreadyExists {
			Error.Code = int32(codes.AlreadyExists)
			Error.Message = err.Error()
		} else {
			Error.Code = int32(codes.Internal)
			Error.Message = err.Error()
		}

		responseError := &proto.UpdateUserResponse_Error{
			Error: Error,
		}

		responseProto := &proto.UpdateUserResponse{
			Result: responseError,
		}

		return responseProto, nil
	}

	protoUser := &proto.UserData{
		Id:       userUpdated.ID,
		Name:     userUpdated.Name,
		Email:    userUpdated.Email,
		Verified: userUpdated.Verified,
	}

	responseDataProto := &proto.UpdateUserResponse_User{
		User: protoUser,
	}

	responseProto := &proto.UpdateUserResponse{
		Result: responseDataProto,
	}

	return responseProto, nil
}

func (userRouter *UserRouter) DeleteUser(ctx context.Context, req *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	err := userRouter.useCase.Delete(req.GetId())
	if err != nil {
		Error := &proto.Error{}

		if err == Errors.ErrUserNotFound {
			Error.Code = int32(codes.NotFound)
			Error.Message = err.Error()
		} else if err == Errors.ErrIdRequired || err == Errors.ErrIdInvalid {
			Error.Code = int32(codes.InvalidArgument)
			Error.Message = err.Error()
		} else {
			Error.Code = int32(codes.Internal)
			Error.Message = err.Error()
		}

		responseError := &proto.DeleteUserResponse_Error{
			Error: Error,
		}

		responseProto := &proto.DeleteUserResponse{
			Result: responseError,
		}

		return responseProto, nil
	}

	responseProto := &proto.DeleteUserResponse{
		Result: &proto.DeleteUserResponse_Success{
			Success: true,
		},
	}

	return responseProto, nil
}
