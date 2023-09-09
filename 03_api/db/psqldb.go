package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/valeriatisch/golang-web-development/03_api/types"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	return sql.Open("postgres", psqlInfo)
}

func GetAllUsers(db *sql.DB, firstName string) ([]types.User, error) {
	query := "SELECT * FROM users WHERE 1=1"
	var args []interface{}
	var argID int = 1

	// If first_name given, append to query
	if firstName != "" {
		query += fmt.Sprintf(" AND first_name = $%d", argID)
		args = append(args, firstName)
		argID++
	}
	
	// Query DB
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Scan each result into struct & append to slice
	var users []types.User
	for rows.Next() {
		var a types.User
		if err := rows.Scan(&a.ID, &a.FirstName, &a.LastName); err != nil {
			return nil, err
		}
		users = append(users, a)
	}

	return users, nil
}

func AddUser(db *sql.DB, firstName, lastName string) error {
	_, err := db.Exec("INSERT INTO users(first_name, last_name) VALUES($1, $2)", firstName, lastName)
	return err
}
