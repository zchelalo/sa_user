package userPostgresRepo

import (
	userDomain "github.com/zchelalo/sa_user/internal/modules/user/domain"
	"github.com/zchelalo/sa_user/pkg/sqlc/db"
)

type PostgresRepository struct {
	store *db.SQLStore
}

func New(store *db.SQLStore) userDomain.UserRepository {
	return &PostgresRepository{
		store: store,
	}
}
