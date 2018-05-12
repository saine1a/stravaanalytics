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
	accessToken string
}

// Init : Create the Strava Access
func Init() *StravaAccess {

	token := os.Getenv("STRAVA_ACCESS_TOKEN")

	if token == "" {
		fmt.Println("Need STRAVA_ACCESS_TOKEN to contain access token for Strava")
		os.Exit(-1)
	}

	obj := &StravaAccess{accessToken: token}

	return obj
}

// GetClubs : Get clubs for the user
func (obj *StravaAccess) GetClubs() *([]Club) {

	// Prepare get request
	bearer := fmt.Sprintf("Bearer %s", obj.accessToken)

	request, _ := http.NewRequest("GET", "https://www.strava.com/api/v3/athlete/clubs", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", bearer)
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
func (obj *StravaAccess) GetActivities() {

}
