package userGRPC

import (
	"context"

	Errors "github.com/zchelalo/sa_user/internal/modules/user/errors"
	"github.com/zchelalo/sa_user/pkg/proto"
	"github.com/zchelalo/sa_user/pkg/util"
	"google.golang.org/grpc/codes"
)

func (userRouter *UserRouter) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {

	userUpdated, err := userRouter.useCase.Update(ctx, req.Id, req.Name, req.Email, req.Password, req.Verified)
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
