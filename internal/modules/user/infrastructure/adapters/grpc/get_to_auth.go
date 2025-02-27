package userGRPC

import (
	"context"

	Errors "github.com/zchelalo/sa_user/internal/modules/user/errors"
	"github.com/zchelalo/sa_user/pkg/proto"
	"google.golang.org/grpc/codes"
)

func (userRouter *UserRouter) GetUserToAuth(ctx context.Context, req *proto.GetUserToAuthRequest) (*proto.GetUserToAuthResponse, error) {
	userObtained, err := userRouter.useCase.GetToAuth(ctx, req.GetEmail())
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
