package db

import (
	"database/sql"

	api "github.com/valeriatisch/golang-web-development/exercises/02_fetch_data/solution/api"
)

const createTableQuery = `
CREATE TABLE IF NOT EXISTS activities (
    id SERIAL PRIMARY KEY,
    activity TEXT NOT NULL,
    type TEXT,
    participants INT,
    price FLOAT,
    link TEXT,
    accessibility FLOAT
);
`

func EnsureTableExists(conn *sql.DB) error {
	_, err := conn.Exec(createTableQuery)
	return err
}

func InsertActivity(conn *sql.DB, activity *api.Activity) error {
	_, err := conn.Exec(`
	INSERT INTO activities (activity, type, participants, price, link, accessibility)
	VALUES ($1, $2, $3, $4, $5, $6)
	`,
		activity.Activity,
		activity.Type,
		activity.Participants,
		activity.Price,
		activity.Link,
		activity.Accessibility)

	return err
}
