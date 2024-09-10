package userDomain

type UserRepository interface {
	Get(id string) (*UserEntity, error)
	GetToAuth(email string) (*UserEntity, error)
	GetAll(offset, limit int32) ([]*UserEntity, error)
	Create(user *UserEntity) (*UserEntity, error)
	Update(user *UserEntity) (*UserEntity, error)
	Delete(id string) error
	Count() (int32, error)
}
