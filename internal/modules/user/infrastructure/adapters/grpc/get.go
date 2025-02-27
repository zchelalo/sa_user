package userGRPC

import (
	"context"

	Errors "github.com/zchelalo/sa_user/internal/modules/user/errors"
	"github.com/zchelalo/sa_user/pkg/proto"
	"google.golang.org/grpc/codes"
)

func (userRouter *UserRouter) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	userObtained, err := userRouter.useCase.Get(ctx, req.GetId())
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
