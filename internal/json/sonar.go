package json

// Conditions struct - SonarRequestJSON.
type Conditions struct {
	Status         string `json:"status"`
	MetricKey      string `json:"metricKey"`
	Comparator     string `json:"comparator"`
	ErrorThreshold string `json:"errorThreshold"`
	ActualValue    string `json:"actualValue"`
}

// Periods struct - SonarRequestJSON.
type Periods struct {
	Index     int    `json:"index"`
	Mode      string `json:"mode"`
	Date      string `json:"date"`
	Parameter string `json:"parameter"`
}

// ProjectStatus struct - SonarRequestJSON.
type ProjectStatus struct {
	Status            string        `json:"status"`
	Conditions        []*Conditions `json:"conditions"`
	Periods           []*Periods    `json:"periods"`
	IgnoredConditions bool          `json:"ignoredConditions"`
}

// SonarRequestJSON struct - JSON object returned in Sonar GET endpoint Quality Gate.
type SonarRequestJSON struct {
	ProjectStatus *ProjectStatus `json:"projectStatus"`
}
