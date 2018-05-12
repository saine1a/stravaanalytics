package stravaaccess

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// StravaAccess : Core object for Strava Access
type StravaAccess struct {
	bearerHeader string
}

// Init : Create the Strava Access
func Init() *StravaAccess {

	token := os.Getenv("STRAVA_ACCESS_TOKEN")

	if token == "" {
		fmt.Println("Need STRAVA_ACCESS_TOKEN to contain access token for Strava")
		os.Exit(-1)
	}

	bearer := fmt.Sprintf("Bearer %s", token)

	obj := &StravaAccess{bearerHeader: bearer}

	return obj
}

// GetClubs : Get clubs for the user
func (obj *StravaAccess) GetClubs() *([]Club) {

	// Prepare get request
	request, _ := http.NewRequest("GET", "https://www.strava.com/api/v3/athlete/clubs", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", obj.bearerHeader)
	client := &http.Client{}

	// Make request
	response, err := client.Do(request)

	defer response.Body.Close()

	if err != nil {
		panic(err.Error())
	}

	// Unmarshal response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}

	clubs := make([]Club, 0)

	err = json.Unmarshal(body, &clubs)

	if err != nil {
		panic(err.Error())
	}

	return &clubs
}

// GetActivities : Get activities for a club
func (obj *StravaAccess) GetActivities(club int64) *([]SummaryActivity) {

	// Prepare get request
	url := fmt.Sprintf("https://www.strava.com/api/v3/clubs/%d/activities&per_page=%d", club, 999)

	fmt.Println(url)

	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", obj.bearerHeader)
	client := &http.Client{}

	// Make request
	response, err := client.Do(request)

	defer response.Body.Close()

	if err != nil {
		panic(err.Error())
	}

	// Unmarshal response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}

	activities := make([]SummaryActivity, 0)

	err = json.Unmarshal(body, &activities)

	if err != nil {
		panic(err.Error())
	}

	// Go through and convert distances to miles

	for a := range activities {
		activities[a].Distance = (activities[a].Distance / 1000) / 1.60934
	}

	return &activities
}
