package userGRPC

import (
	"context"

	Errors "github.com/zchelalo/sa_user/internal/modules/user/errors"
	"github.com/zchelalo/sa_user/pkg/proto"
	"google.golang.org/grpc/codes"
)

func (userRouter *UserRouter) GetUsers(ctx context.Context, req *proto.GetUsersRequest) (*proto.GetUsersResponse, error) {
	usersObtained, meta, err := userRouter.useCase.GetAll(ctx, req.GetPage(), req.GetLimit())
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
