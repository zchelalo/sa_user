package connection

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewConnection(driver, source string) (*sql.DB, error) {
	conn, err := sql.Open(driver, source)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	return conn, nil
}
