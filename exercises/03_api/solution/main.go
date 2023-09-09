package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/valeriatisch/golang-web-development/exercises/03_api/solution/api"
	"github.com/valeriatisch/golang-web-development/exercises/03_api/solution/utils"
)

func main() {
	// Load .env file
	err := godotenv.Load("../../../.env")
	utils.CheckError("Error loading .env file: ", err)

	router := mux.NewRouter()

	api.ActivityAPIRoutes(router.PathPrefix("/api/v1").Subrouter())
	api.ActivityPageRoutes(router.PathPrefix("/").Subrouter())

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", router)
}
