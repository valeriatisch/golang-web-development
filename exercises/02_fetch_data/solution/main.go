package main

import (
	"fmt"

	"github.com/joho/godotenv"

	api "github.com/valeriatisch/golang-web-development/exercises/02_fetch_data/solution/api"
	db "github.com/valeriatisch/golang-web-development/exercises/02_fetch_data/solution/db"
	utils "github.com/valeriatisch/golang-web-development/exercises/02_fetch_data/solution/utils"
)

const numActivities = 10

func main() {
	// Load .env file
	err := godotenv.Load("../../../.env")
	utils.CheckError("Error loading .env file: ", err)

	// Connect to the database
	conn, err := db.Connect()
	utils.CheckError("Error connecting to the database: ", err)
	defer conn.Close()

	// Ensure the table exists
	err = db.EnsureTableExists(conn)
	utils.CheckError("Error creating table: ", err)

	// Fetch the activities
	activities, errors := api.FetchMultipleActivities(numActivities)
	for _, err := range errors {
		utils.CheckError("Error fetching activity: ", err)
	}

	// Insert the activities into the database
	for _, activity := range activities {
		err := db.InsertActivity(conn, activity)
		if err != nil {
			fmt.Printf("Error inserting activity into database: %v \n", err)
		}
	}
}
