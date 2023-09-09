package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Project struct to hold the data for each project
type Project struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Link        string `json:"Link"`
}

// Network struct to hold the data for each network
type Network struct {
	Name string `json:"Name"`
	Link string `json:"Link"`
}

// Portfolio struct to hold the data for the entire portfolio
type Portfolio struct {
	Name     string    `json:"Name"`
	Title    string    `json:"Title"`
	About    string    `json:"About"`
	Projects []Project `json:"Projects"`
	Networks []Network `json:"Networks"`
}

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
	// Create the index.html file
	file, err := os.Create("static/index.html")
	if err != nil {
		log.Fatal("Failed to create index.html file:", err)
	}
	defer file.Close()

	// Parse the template file
	tmpl, err := template.ParseFiles("templates/portfolio.gohtml")
	if err != nil {
		log.Fatal("Failed to parse template file:", err)
	}

	// Execute the template and write the output to the index.html
	err = tmpl.Execute(file, portfolio)
	if err != nil {
		log.Fatal("Failed to execute template:", err)
	}
}

func startServer() {
	// Create a file server to serve the files in the static directory
	fs := http.FileServer(http.Dir("static"))
	// Handle the root path "/"
	http.Handle("/", fs)

	log.Println("Listening on :8080...")
	// Start the server, listen on port 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
