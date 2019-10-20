package viewmodels

type Sentiment struct {
	Label       string      `json:"label"`
	Probability Probability `json:"probability"`
}

type Probability struct {
	Negative float64 `json:"neg"`
	Neutral  float64 `json:"neutral"`
	Positive float64 `json:"pos"`
}
