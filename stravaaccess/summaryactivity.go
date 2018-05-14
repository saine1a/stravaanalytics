package stravaaccess

// MetaAthlete : Strava athlete summary data
type MetaAthlete struct {
	FirstName string `json:"firstname" dimension:"athlete"`
	LastName  string `json:"lastname" dimension:"athlete"`
}

// SummaryActivity : Strava summary activity object
type SummaryActivity struct {
	ID       int64       `json:"id" dimension:"fact"`
	Name     string      `json:"name" dimension:"fact"`
	Distance float32     `json:"distance" dimension:"fact" type:"SMALLINT"`
	Athlete  MetaAthlete `json:"athlete"`
}
