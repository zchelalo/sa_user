package userPostgresRepo

import "context"

func (repo *PostgresRepository) Delete(ctx context.Context, id string) error {
	err := repo.store.UserQueries.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
