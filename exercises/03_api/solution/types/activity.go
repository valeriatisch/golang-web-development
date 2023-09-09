package types

type Activity struct {
	ID            int     `json:"id"`
	ActivityName  string  `json:"activity"`
	Type          string  `json:"type"`
	Participants  int     `json:"participants"`
	Price         float64 `json:"price"`
	Link          string  `json:"link"`
	Accessibility float64 `json:"accessibility"`
}
