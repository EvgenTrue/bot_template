package storage

import (
	"github.com/jackc/pgx"
)

func Newdbstorage() *pgx.Conn {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		Database: "postgres",
	})
	if err != nil {
		panic(err)

	}
	return conn
}
