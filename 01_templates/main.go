package main

import (
	"fmt"
	"log"
	"net/http"
)

// A template is a string or file containing HTML & placeholders or loops enclosed in double braces {{ }}
// A template allows us to generate HTML dynamically.
// Template files typically end with .gohtml or .tmpl

func init() {
	// Let's ensure that the template parsing and error handling are performed once in the beginning
	// Must() simplifies error handling as it panics when an error occurs.

}

func main() {
	// Create a file called index.html

	// Placeholder {{ . }}

	// Empty file

	// Variables {{ $name := . }}

	// Structs {{ .Field }}

	// Loops {{ range . }}
	// Struct of slices

	// Functions

	startServer()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func startServer() {
	fs := http.FileServer(http.Dir("."))

	http.Handle("/", fs)

	fmt.Println("Server listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	// ListenAndServe starts an HTTP server with a given address and handler.
	// The handler is usually nil, which means to use DefaultServeMux.
	// Handle and HandleFunc add handlers to DefaultServeMux:
	// http.Handle("/", http.HandlerFunc(index))
	// http.HandleFunc("/", index)
	// http.ListenAndServe(":8080", nil)
}
