package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/valeriatisch/golang-web-development/02_fetch_data_write_db/api"
	"github.com/valeriatisch/golang-web-development/02_fetch_data_write_db/db"
	"github.com/valeriatisch/golang-web-development/02_fetch_data_write_db/utils"
)

func main() {
	// Load .env file
	err := godotenv.Load("../.env")
	utils.CheckError("Error while reading .env file:", err)

	// Connect to the database
	conn, err := db.Connect()
	utils.CheckError("Error while connecting to the database:", err)
	defer conn.Close()
	fmt.Println("Connected to the database!")

	// Ensure the table exists
	err = db.EnsureTableExists(conn)
	utils.CheckError("Error whhile checking the table:", err)
	fmt.Println("Table exists!")

	// Fetch the user
	user, err := api.FetchUser()
	utils.CheckError("Error while fetching the user:", err)
	fmt.Println("Fetched user!")

	// Insert the user into the database
	err = db.InsertUser(conn, user)
	utils.CheckError("Error while inserting the user:", err)
	fmt.Println("Inserted user!")

}
