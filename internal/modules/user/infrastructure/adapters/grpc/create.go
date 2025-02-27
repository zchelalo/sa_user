package userGRPC

import (
	"context"

	Errors "github.com/zchelalo/sa_user/internal/modules/user/errors"
	"github.com/zchelalo/sa_user/pkg/proto"
	"github.com/zchelalo/sa_user/pkg/util"
	"google.golang.org/grpc/codes"
)

func (userRouter *UserRouter) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	userCreated, err := userRouter.useCase.Create(ctx, req.GetName(), req.GetEmail(), req.GetPassword())
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
