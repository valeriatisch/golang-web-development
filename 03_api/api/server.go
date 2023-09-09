package api

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/valeriatisch/golang-web-development/03_api/db"
	"github.com/valeriatisch/golang-web-development/03_api/types"
)

// Handlers
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Get all users from db & send back to requestor
	conn, err := db.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// Add first_name as potential query parameter
	first_name := r.URL.Query().Get("first_name")

	users, err := db.GetAllUsers(conn, first_name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Handler to display all users
func AllUsersPageHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := db.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	users, err := db.GetAllUsers(conn, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/users.html"))
	tmpl.Execute(w, users)
}

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	// Create POST method to add users to db
	conn, err := db.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	var newUser types.User
	err = json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = db.AddUser(conn, newUser.FirstName, newUser.LastName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func AddUserPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	conn, err := db.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")

	if firstName == "" || lastName == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	err = db.AddUser(conn, firstName, lastName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

// Routes
func UserAPIRoutes(router *mux.Router) {
	router.HandleFunc("/users", GetAllUsersHandler).Methods(http.MethodGet)
	router.HandleFunc("/users", AddUserHandler).Methods(http.MethodPost)
}

func UserPageRoutes(router *mux.Router) {
	router.HandleFunc("/users", AllUsersPageHandler).Methods(http.MethodGet)
	router.HandleFunc("/users", AddUserPageHandler).Methods(http.MethodPost)
}
