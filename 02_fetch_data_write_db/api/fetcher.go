package api

import (
	"encoding/json"
	"net/http"
	"time"
)

const apiURL = "https://random-data-api.com/api/users/random_user"

// Create client
var client = &http.Client{
	Timeout: time.Second * 30,
}

// Create user struct
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Fetch eser & decode into user struct
func FetchUser() (*User, error) {
	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
