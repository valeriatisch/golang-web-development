package db

import (
	"database/sql"

	api "github.com/valeriatisch/golang-web-development/02_fetch_data_write_db/api"
)

const createTableQuery = `
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
	last_name TEXT NOT NULL
);
`

func EnsureTableExists(conn *sql.DB) error {
	_, err := conn.Exec(createTableQuery)
	return err
}

// Insert User into table
func InsertUser(conn *sql.DB, user *api.User) error {
	_, err := conn.Exec(`
	INSERT INTO users (first_name, last_name)
	VALUES ($1, $2)
	`,
		user.FirstName,
		user.LastName,)

	return err
}
