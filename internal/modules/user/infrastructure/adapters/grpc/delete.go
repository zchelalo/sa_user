package userGRPC

import (
	"context"

	Errors "github.com/zchelalo/sa_user/internal/modules/user/errors"
	"github.com/zchelalo/sa_user/pkg/proto"
	"google.golang.org/grpc/codes"
)

func (userRouter *UserRouter) DeleteUser(ctx context.Context, req *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	err := userRouter.useCase.Delete(ctx, req.GetId())
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
