package api

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/valeriatisch/golang-web-development/02_fetch_data_write_db/utils"
	"github.com/valeriatisch/golang-web-development/exercises/03_api/solution/db"
	"github.com/valeriatisch/golang-web-development/exercises/03_api/solution/types"
)

// Handlers

func GetAllActivitiesHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := db.Connect()
	utils.CheckError("Error connecting to the database: ", err)
	defer conn.Close()

	activityType := r.URL.Query().Get("type")
	participantsStr := r.URL.Query().Get("participants")
	var participants int
	if participantsStr != "" {
		participants, err = strconv.Atoi(participantsStr)
		if err != nil {
			http.Error(w, "Invalid participants number", http.StatusBadRequest)
			return
		}
	}

	activities, err := db.GetAllActivities(conn, activityType, participants)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activities)
}

// Handler to get one activity by ID
func GetActivityHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := db.Connect()
	utils.CheckError("Error connecting to the database: ", err)
	defer conn.Close()

	activity, err := db.GetActivity(conn, mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activity)
}

// Handler to post own activity
func CreateActivityHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := db.Connect()
	utils.CheckError("Error connecting to the database: ", err)
	defer conn.Close()

	var activity types.Activity
	if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.InsertActivity(conn, &activity); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(activity)
}

// Handler to display all activities
func AllActivitiesPageHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := db.Connect()
	utils.CheckError("Error connecting to the database: ", err)
	defer conn.Close()

	activities, err := db.GetAllActivities(conn, "", 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/activities.html"))
	tmpl.Execute(w, activities)
}

// Handler to display a single activity
func SingleActivityPageHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := db.Connect()
	utils.CheckError("Error connecting to the database: ", err)
	defer conn.Close()

	activity, err := db.GetActivity(conn, mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/activity.html"))
	tmpl.Execute(w, activity)
}

func CreateActivityPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/create_activity.html"))
	tmpl.Execute(w, nil)
}

func HandleActivityCreation(w http.ResponseWriter, r *http.Request) {
    conn, err := db.Connect()
    utils.CheckError("Error connecting to the database: ", err)
    defer conn.Close()

    price, err := strconv.ParseFloat(r.FormValue("price"), 64)
    if err != nil {
        http.Error(w, "Invalid price", http.StatusBadRequest)
        return
    }

    participants, err := strconv.Atoi(r.FormValue("participants"))
    if err != nil {
        http.Error(w, "Invalid participants count", http.StatusBadRequest)
        return
    }

    accessibility, err := strconv.ParseFloat(r.FormValue("accessibility"), 64)
    if err != nil {
        http.Error(w, "Invalid accessibility value", http.StatusBadRequest)
        return
    }

    activity := &types.Activity{
        ActivityName:  r.FormValue("activityName"),
        Type:          r.FormValue("type"),
        Participants:  participants,
        Price:         price,
        Link:          r.FormValue("link"),
        Accessibility: accessibility,
    }

    if err := db.InsertActivity(conn, activity); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/activities", http.StatusSeeOther)
}


// Routes

func ActivityAPIRoutes(router *mux.Router) {
	router.HandleFunc("/activities", GetAllActivitiesHandler).Methods(http.MethodGet)
	router.HandleFunc("/activities/{id}", GetActivityHandler).Methods(http.MethodGet)
	router.HandleFunc("/activity", CreateActivityHandler).Methods(http.MethodPost)
}

func ActivityPageRoutes(router *mux.Router) {
	router.HandleFunc("/activities", AllActivitiesPageHandler).Methods(http.MethodGet)
	router.HandleFunc("/activity/{id}", SingleActivityPageHandler).Methods(http.MethodGet)
	router.HandleFunc("/create-activity", CreateActivityPageHandler).Methods(http.MethodGet)
	router.HandleFunc("/create-activity", HandleActivityCreation).Methods(http.MethodPost)
}
