package stravaaccess

// MetaAthlete : Strava athlete summary data
type MetaAthlete struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// SummaryActivity : Strava summary activity object
type SummaryActivity struct {
	ID       int64       `json:"id"`
	Name     string      `json:"name"`
	Distance float32     `json:"distance"`
	Athlete  MetaAthlete `json:"athlete"`
}
