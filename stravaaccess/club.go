package stravaaccess

// Club : Strava club object
type Club struct {
	ID   int64  `json:"id" dimension:"Club"`
	Name string `json:"name" dimension:"Club"`
}
