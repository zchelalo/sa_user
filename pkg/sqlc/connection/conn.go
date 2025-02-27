package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func New(driver, source string) (*sql.DB, error) {
	conn, err := sql.Open(driver, source)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to db: %w", err)
	}

	return conn, nil
}
