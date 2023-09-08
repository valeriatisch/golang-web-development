package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// A template is a string or file containing HTML & placeholders or loops enclosed in double braces {{ }}
// A template allows us to generate HTML dynamically.
// Template files typically end with .gohtml or .tmpl
var tmpl *template.Template

func init() {
	// Let's ensure that the template parsing and error handling are performed once in the beginning
	// Must() simplifies error handling as it panics when an error occurs.
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

type Person struct {
	Name string
	Age int
}

type Book struct {
	Title string
}

func main() {
	// Create a file called index.html
	file, err := os.Create("index.html")
	checkError(err)
	defer file.Close()

	// Placeholder {{ . }}
	err = tmpl.ExecuteTemplate(file, "placeholder.gohtml", "Jane")
	checkError(err)

	// Empty file
	file.Truncate(0)

	// Variables {{ $name := . }}
	err = tmpl.ExecuteTemplate(file, "variable.gohtml", "John")
	checkError(err)
	file.Truncate(0)

	// Structs {{ .Field }}
	// Loops {{ range . }}
	// Struct of slices
	p1 := Person{
		Name: "Jane",
		Age:  20,
	}
	b1 := Book{
		Title: "Harry Potter",
	}
	b2 := Book{
		Title: "Lord of the Rings",
	}
	data := struct {
		People []Person
		Books []Book
	}{
		People: []Person{p1},
		Books: []Book{b1, b2},
	}
	err = tmpl.ExecuteTemplate(file, "loops.gohtml", data)
	checkError(err)

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
