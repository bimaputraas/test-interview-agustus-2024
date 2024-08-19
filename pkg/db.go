package pkg

import (
	"database/sql"
	"fmt"
)

func NewPostgresDb(user, password, dbname, host, port string) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname))

	if err != nil {
		return nil, err
	}

	return db, nil
}
