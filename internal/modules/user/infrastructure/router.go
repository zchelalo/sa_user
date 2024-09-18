package userInfrastructure

import (
	"context"

	userApplication "github.com/zchelalo/sa_user/internal/modules/user/application"
	userErrors "github.com/zchelalo/sa_user/internal/modules/user/errors"
	userProto "github.com/zchelalo/sa_user/pkg/proto/user"
	userDb "github.com/zchelalo/sa_user/pkg/sqlc/user/db"
	"github.com/zchelalo/sa_user/pkg/util"
	"google.golang.org/grpc/codes"
)

type UserRouter struct {
	ctx     context.Context
	useCase *userApplication.UserUseCases
	userProto.UnimplementedUserServiceServer
}

func NewUserRouter(store userDb.Store, ctx context.Context) *UserRouter {
	userRepository := NewUserRepository(ctx, store)
	userUseCase := userApplication.NewUserUseCases(ctx, userRepository)

	return &UserRouter{
		ctx:     ctx,
		useCase: userUseCase,
	}
}

func (userRouter *UserRouter) GetUsers(ctx context.Context, req *userProto.GetUsersRequest) (*userProto.GetUsersResponse, error) {
	// Create a new context with the request context
	// ctx = context.WithValue(userRouter.ctx, "requestCtx", ctx)

	usersObtained, meta, err := userRouter.useCase.GetAll(req.GetPage(), req.GetLimit())
	if err != nil {
		userError := &userProto.UserError{}

		if err == userErrors.ErrUsersNotFound {
			userError.Code = int32(codes.NotFound)
			userError.Message = err.Error()
		}

		responseError := &userProto.GetUsersResponse_Error{
			Error: userError,
		}

		responseProto := &userProto.GetUsersResponse{
			Result: responseError,
		}

		return responseProto, nil
	}

	metaProto := &userProto.Meta{
		Page:       meta.Page,
		PerPage:    meta.PerPage,
		Count:      meta.PageCount,
		TotalCount: meta.TotalCount,
	}

	usersProto := []*userProto.UserData{}
	for _, user := range usersObtained {
		protoUser := &userProto.UserData{
			Id:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Verified: user.Verified,
		}

		usersProto = append(usersProto, protoUser)
	}

	responseDataProto := &userProto.GetUsersResponse_Data{
		Data: &userProto.UsersWithMeta{
			Users: usersProto,
			Meta:  metaProto,
		},
	}

	responseProto := &userProto.GetUsersResponse{
		Result: responseDataProto,
	}

	return responseProto, nil
}

func (userRouter *UserRouter) GetUser(ctx context.Context, req *userProto.GetUserRequest) (*userProto.GetUserResponse, error) {
	userObtained, err := userRouter.useCase.Get(req.GetId())
	if err != nil {
		userError := &userProto.UserError{}

		if err == userErrors.ErrUserNotFound {
			userError.Code = int32(codes.NotFound)
			userError.Message = err.Error()
		} else if err == userErrors.ErrIdRequired || err == userErrors.ErrIdInvalid {
			userError.Code = int32(codes.InvalidArgument)
			userError.Message = err.Error()
		} else {
			userError.Code = int32(codes.Internal)
			userError.Message = err.Error()
		}

		responseError := &userProto.GetUserResponse_Error{
			Error: userError,
		}

		responseProto := &userProto.GetUserResponse{
			Result: responseError,
		}

		return responseProto, nil
	}

	protoUser := &userProto.UserData{
		Id:       userObtained.ID,
		Name:     userObtained.Name,
		Email:    userObtained.Email,
		Verified: userObtained.Verified,
	}

	responseDataProto := &userProto.GetUserResponse_User{
		User: protoUser,
	}

	responseProto := &userProto.GetUserResponse{
		Result: responseDataProto,
	}

	return responseProto, nil
}

func (userRouter *UserRouter) GetUserToAuth(ctx context.Context, req *userProto.GetUserToAuthRequest) (*userProto.GetUserToAuthResponse, error) {
	userObtained, err := userRouter.useCase.GetToAuth(req.GetEmail())
	if err != nil {
		userError := &userProto.UserError{}

		if err == userErrors.ErrUserNotFound {
			userError.Code = int32(codes.NotFound)
			userError.Message = err.Error()
		} else if err == userErrors.ErrEmailRequired || err == userErrors.ErrEmailInvalid {
			userError.Code = int32(codes.InvalidArgument)
			userError.Message = err.Error()
		} else {
			userError.Code = int32(codes.Internal)
			userError.Message = err.Error()
		}

		responseError := &userProto.GetUserToAuthResponse_Error{
			Error: userError,
		}

		responseProto := &userProto.GetUserToAuthResponse{
			Result: responseError,
		}

		return responseProto, nil
	}

	protoUser := &userProto.UserWithPassword{
		Id:       userObtained.ID,
		Name:     userObtained.Name,
		Email:    userObtained.Email,
		Password: userObtained.Password,
		Verified: userObtained.Verified,
	}

	responseDataProto := &userProto.GetUserToAuthResponse_User{
		User: protoUser,
	}

	responseProto := &userProto.GetUserToAuthResponse{
		Result: responseDataProto,
	}

	return responseProto, nil
}

func (userRouter *UserRouter) CreateUser(ctx context.Context, req *userProto.CreateUserRequest) (*userProto.CreateUserResponse, error) {
	userCreated, err := userRouter.useCase.Create(req.GetName(), req.GetEmail(), req.GetPassword())
	if err != nil {
		userError := &userProto.UserError{}

		errorInvalidArgument := []error{
			userErrors.ErrNameRequired,
			userErrors.ErrNameInvalid,
			userErrors.ErrEmailRequired,
			userErrors.ErrEmailInvalid,
			userErrors.ErrPasswordRequired,
			userErrors.ErrPasswordInvalid,
		}

		isInvalidArgumentError := util.IsErrorType(err, errorInvalidArgument)

		if isInvalidArgumentError {
			userError.Code = int32(codes.InvalidArgument)
			userError.Message = err.Error()
		} else if err == userErrors.ErrEmailAlreadyExists {
			userError.Code = int32(codes.AlreadyExists)
			userError.Message = err.Error()
		} else {
			userError.Code = int32(codes.Internal)
			userError.Message = err.Error()
		}

		responseError := &userProto.CreateUserResponse_Error{
			Error: userError,
		}

		responseProto := &userProto.CreateUserResponse{
			Result: responseError,
		}

		return responseProto, nil
	}

	protoUser := &userProto.UserData{
		Id:       userCreated.ID,
		Name:     userCreated.Name,
		Email:    userCreated.Email,
		Verified: userCreated.Verified,
	}

	responseDataProto := &userProto.CreateUserResponse_User{
		User: protoUser,
	}

	responseProto := &userProto.CreateUserResponse{
		Result: responseDataProto,
	}

	return responseProto, nil
}

func (userRouter *UserRouter) UpdateUser(ctx context.Context, req *userProto.UpdateUserRequest) (*userProto.UpdateUserResponse, error) {

	userUpdated, err := userRouter.useCase.Update(req.Id, req.Name, req.Email, req.Password, req.Verified)
	if err != nil {
		userError := &userProto.UserError{}

		errorInvalidArgument := []error{
			userErrors.ErrIdRequired,
			userErrors.ErrIdInvalid,
			userErrors.ErrNameInvalid,
			userErrors.ErrEmailInvalid,
			userErrors.ErrPasswordInvalid,
			userErrors.ErrVerifiedInvalid,
		}

		isInvalidArgumentError := util.IsErrorType(err, errorInvalidArgument)

		if isInvalidArgumentError {
			userError.Code = int32(codes.InvalidArgument)
			userError.Message = err.Error()
		} else if err == userErrors.ErrUserNotFound {
			userError.Code = int32(codes.NotFound)
			userError.Message = err.Error()
		} else if err == userErrors.ErrEmailAlreadyExists {
			userError.Code = int32(codes.AlreadyExists)
			userError.Message = err.Error()
		} else {
			userError.Code = int32(codes.Internal)
			userError.Message = err.Error()
		}

		responseError := &userProto.UpdateUserResponse_Error{
			Error: userError,
		}

		responseProto := &userProto.UpdateUserResponse{
			Result: responseError,
		}

		return responseProto, nil
	}

	protoUser := &userProto.UserData{
		Id:       userUpdated.ID,
		Name:     userUpdated.Name,
		Email:    userUpdated.Email,
		Verified: userUpdated.Verified,
	}

	responseDataProto := &userProto.UpdateUserResponse_User{
		User: protoUser,
	}

	responseProto := &userProto.UpdateUserResponse{
		Result: responseDataProto,
	}

	return responseProto, nil
}

func (userRouter *UserRouter) DeleteUser(ctx context.Context, req *userProto.DeleteUserRequest) (*userProto.DeleteUserResponse, error) {
	err := userRouter.useCase.Delete(req.GetId())
	if err != nil {
		userError := &userProto.UserError{}

		if err == userErrors.ErrUserNotFound {
			userError.Code = int32(codes.NotFound)
			userError.Message = err.Error()
		} else if err == userErrors.ErrIdRequired || err == userErrors.ErrIdInvalid {
			userError.Code = int32(codes.InvalidArgument)
			userError.Message = err.Error()
		} else {
			userError.Code = int32(codes.Internal)
			userError.Message = err.Error()
		}

		responseError := &userProto.DeleteUserResponse_Error{
			Error: userError,
		}

		responseProto := &userProto.DeleteUserResponse{
			Result: responseError,
		}

		return responseProto, nil
	}

	responseProto := &userProto.DeleteUserResponse{
		Result: &userProto.DeleteUserResponse_Success{
			Success: true,
		},
	}

	return responseProto, nil
}
