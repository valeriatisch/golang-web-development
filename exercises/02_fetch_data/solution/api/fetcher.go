package api

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

const apiURL = "https://www.boredapi.com/api/activity"

var client = &http.Client{
	Timeout: time.Second * 30,
}

type Activity struct {
	Activity      string  `json:"activity"`
	Type          string  `json:"type"`
	Participants  int     `json:"participants"`
	Price         float64 `json:"price"`
	Link          string  `json:"link"`
	Accessibility float64 `json:"accessibility"`
}

func FetchActivity() (*Activity, error) {
	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var activity Activity
	if err := json.NewDecoder(resp.Body).Decode(&activity); err != nil {
		return nil, err
	}

	return &activity, nil
}

func FetchMultipleActivities(num int) ([]*Activity, []error) {
	activitiesCh := make(chan *Activity, num)
	errorsCh := make(chan error, num)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			activity, err := FetchActivity()
			if err != nil {
				errorsCh <- err
				return
			}
			activitiesCh <- activity
		}()
	}

	wg.Wait()
	close(activitiesCh)
	close(errorsCh)

	var activities []*Activity
	var errors []error

	for activity := range activitiesCh {
		activities = append(activities, activity)
	}

	for err := range errorsCh {
		errors = append(errors, err)
	}

	return activities, errors
}
