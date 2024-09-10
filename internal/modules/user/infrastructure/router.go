package userInfrastructure

import (
	"context"
	"reflect"

	userApplication "github.com/zchelalo/sa_user/internal/modules/user/application"
	userErrors "github.com/zchelalo/sa_user/internal/modules/user/errors"
	userProto "github.com/zchelalo/sa_user/pkg/proto/user"
	userDb "github.com/zchelalo/sa_user/pkg/sqlc/user/db"
	"github.com/zchelalo/sa_user/pkg/utils"
)

type UserRouter struct {
	useCase *userApplication.UserUseCases
	userProto.UnimplementedUserServiceServer
}

func NewUserRouter(store userDb.Store, ctx context.Context) *UserRouter {
	userRepository := NewUserRepository(ctx, store)
	userUseCase := userApplication.NewUserUseCases(ctx, userRepository)

	return &UserRouter{useCase: userUseCase}
}

func (userRouter *UserRouter) GetUsers(ctx context.Context, req *userProto.GetUsersRequest) (*userProto.GetUsersResponse, error) {
	usersObtained, meta, err := userRouter.useCase.GetAll(req.GetPage(), req.GetLimit())
	if err != nil {
		responseProto := formatError[userProto.GetUsersResponse_Error, userProto.GetUsersResponse](err)

		return responseProto, nil
	}

	metaProto := &userProto.Meta{
		Page:       meta.Page,
		PerPage:    meta.PerPage,
		Count:      meta.PageCount,
		TotalCount: meta.TotalCount,
	}

	usersProto := []*userProto.User{}
	for _, user := range usersObtained {
		protoUser := &userProto.User{
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
		responseProto := formatError[userProto.GetUserResponse_Error, userProto.GetUserResponse](err)

		return responseProto, nil
	}

	protoUser := &userProto.User{
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
		responseProto := formatError[userProto.GetUserToAuthResponse_Error, userProto.GetUserToAuthResponse](err)

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
		responseProto := formatError[userProto.CreateUserResponse_Error, userProto.CreateUserResponse](err)

		return responseProto, nil
	}

	protoUser := &userProto.User{
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

	userUpdated, err := userRouter.useCase.Update(req.GetId(), req.Name, req.Email, req.Password, req.Verified)
	if err != nil {
		responseProto := formatError[userProto.UpdateUserResponse_Error, userProto.UpdateUserResponse](err)

		return responseProto, nil
	}

	protoUser := &userProto.User{
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
		responseProto := formatError[userProto.DeleteUserResponse_Error, userProto.DeleteUserResponse](err)

		return responseProto, nil
	}

	responseProto := &userProto.DeleteUserResponse{
		Result: &userProto.DeleteUserResponse_Success{
			Success: true,
		},
	}

	return responseProto, nil
}

func formatError[TResponeError any, TResponse any](err error) *TResponse {
	if err == nil {
		return nil
	}

	errorProto := convertUserErrorToProtoError(err)

	responseErrorProto := new(TResponeError)
	reflect.ValueOf(responseErrorProto).Elem().FieldByName("Error").Set(reflect.ValueOf(errorProto))

	responseProto := new(TResponse)
	reflect.ValueOf(responseProto).Elem().FieldByName("Result").Set(reflect.ValueOf(responseErrorProto))

	return responseProto
}

func convertUserErrorToProtoError(err error) *userProto.Error {
	if err == nil {
		return nil
	}

	var code int32
	var message string

	error400 := []error{
		userErrors.ErrEmailRequired,
		userErrors.ErrEmailInvalid,

		userErrors.ErrIdRequired,
		userErrors.ErrIdInvalid,

		userErrors.ErrNameRequired,
		userErrors.ErrNameInvalid,

		userErrors.ErrPasswordRequired,
		userErrors.ErrPasswordInvalid,
	}

	error404 := []error{
		userErrors.ErrIdRequired,
	}

	error409 := []error{
		userErrors.ErrEmailAlreadyExists,
	}

	if utils.IsErrorType(err, error400) {
		code = 400
		message = err.Error()
	} else if utils.IsErrorType(err, error404) {
		code = 404
		message = err.Error()
	} else if utils.IsErrorType(err, error409) {
		code = 409
		message = err.Error()
	} else {
		code = 500
		message = err.Error()
	}

	errorProto := &userProto.Error{
		Code:    int32(utils.ConvertStatusCodeToProtoCode(code)),
		Message: message,
	}

	return errorProto
}
