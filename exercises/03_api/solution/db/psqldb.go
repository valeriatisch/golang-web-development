package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/valeriatisch/golang-web-development/exercises/03_api/solution/types"
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

func GetAllActivities(db *sql.DB, activityType string, participants int) ([]types.Activity, error) {
	query := "SELECT * FROM activities WHERE 1=1"
	var args []interface{}
	var argID int = 1

	if activityType != "" {
		query += fmt.Sprintf(" AND type = $%d", argID)
		args = append(args, activityType)
		argID++
	}

	if participants != 0 {
		query += fmt.Sprintf(" AND participants = $%d", argID)
		args = append(args, participants)
		argID++
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []types.Activity
	for rows.Next() {
		var a types.Activity
		if err := rows.Scan(&a.ID, &a.ActivityName, &a.Type, &a.Participants, &a.Price, &a.Link, &a.Accessibility); err != nil {
			return nil, err
		}
		activities = append(activities, a)
	}

	return activities, nil
}

func GetActivity(db *sql.DB, id string) (types.Activity, error) {
	var activity types.Activity
	row := db.QueryRow("SELECT * FROM activities WHERE id=$1", id)
	err := row.Scan(&activity.ID, &activity.ActivityName, &activity.Type, &activity.Participants, &activity.Price, &activity.Link, &activity.Accessibility)
	if err != nil {
		return types.Activity{}, err
	}
	
	return activity, nil
}

func InsertActivity(db *sql.DB, activity *types.Activity) error {
	query := `INSERT INTO activities (activity, type, participants, price, link, accessibility) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(query, activity.ActivityName, activity.Type, activity.Participants, activity.Price, activity.Link, activity.Accessibility)
	return err
}
