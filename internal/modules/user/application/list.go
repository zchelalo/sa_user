package userApplication

import (
	"context"

	userDomain "github.com/zchelalo/sa_user/internal/modules/user/domain"
	"github.com/zchelalo/sa_user/pkg/bootstrap"
	"github.com/zchelalo/sa_user/pkg/meta"
)

func (useCase *UserUseCases) GetAll(ctx context.Context, page, limit int32) ([]*userDomain.UserEntity, *meta.Meta, error) {
	usersCount, err := useCase.Count(ctx)
	if err != nil {
		return nil, nil, err
	}

	config := bootstrap.GetConfig()

	meta, err := meta.New(page, limit, int32(usersCount), config.PaginatorLimitDefault)
	if err != nil {
		return nil, nil, err
	}

	usersObtained, err := useCase.userRepository.GetAll(ctx, int32(meta.Offset()), int32(meta.Limit()))
	if err != nil {
		return nil, nil, err
	}

	return usersObtained, meta, nil
}
