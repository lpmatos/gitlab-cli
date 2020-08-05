package json

// Metrics struct in SpeedioJSON.
type Metrics struct {
	Name        string `json:"name"`
	Value       int    `json:"value"`
	DesiredSize string `json:"desiredSize"`
}

// SpeedioJSON struct - JSON object returned in Speedio performance file.
type SpeedioJSON struct {
	Subject string     `json:"subject"`
	Metrics []*Metrics `json:"metrics"`
}
