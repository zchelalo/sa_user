package userDomain

import "context"

type UserRepository interface {
	Get(ctx context.Context, id string) (*UserEntity, error)
	GetToAuth(ctx context.Context, email string) (*UserEntity, error)
	GetAll(ctx context.Context, offset, limit int32) ([]*UserEntity, error)
	Create(ctx context.Context, user *UserEntity) (*UserEntity, error)
	Update(ctx context.Context, user *UserEntity) (*UserEntity, error)
	Delete(ctx context.Context, id string) error
	Count(ctx context.Context) (int32, error)
}
