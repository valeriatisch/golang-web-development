package main

import (
	"encoding/json"
	"log"
	"os"
)

// 1) TODO: Fill out the personal_data.json file in the data directory

// Project struct to hold the data for each project
type Project struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Link        string `json:"Link"`
}

// 2) TODO: Create a Network struct to hold the data for each network

// 3) TODO: Create a Portfolio struct to hold the data for the entire portfolio
// Hint: Notice that the attributs "Projects" and "Networks" must be slices/lists

// 4) TODO: Go into templates/portfolio.gohtml and call the data from the Portfolio struct everywhere it says "TODO"

func main() {
	portfolio := handleData()
	createHTML(portfolio)
	startServer()
}

func handleData() Portfolio {
	// Read the data from the JSON file
	file, err := os.ReadFile("static/data/personal_data.json")
	if err != nil {
		log.Fatal("Failed to read personal_data.json:", err)
	}

	// Unmarshal the JSON data into a Portfolio struct
	var portfolio Portfolio
	err = json.Unmarshal(file, &portfolio)
	if err != nil {
		log.Fatal("Failed to unmarshal JSON data:", err)
	}

	return portfolio
}

func createHTML(portfolio Portfolio) {
	// TODO: Create the static/index.html file

	// TODO: Parse the template file

	// TODO: Execute the template and write the output to the index.html

}

func startServer() {
	// TODO: Create a file server to serve the files in the static directory

	// TODO: Handle the root path "/"

	// TODO: Start the server, listen on port 8080

}
