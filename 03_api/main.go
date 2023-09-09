package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/valeriatisch/golang-web-development/03_api/api"
)

// "github.com/gorilla/mux"
// "github.com/joho/godotenv"

func main() {
	// Load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// Create router
	router := mux.NewRouter()

	// Add API & Page Routes mit path prefixes
	api.UserAPIRoutes(router.PathPrefix("/api/v1").Subrouter())
	api.UserPageRoutes(router.PathPrefix("/").Subrouter())

	// Listen
	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", router)

}
