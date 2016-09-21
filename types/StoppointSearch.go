package types

type StoppointMatch struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Latitude   float32  `json:"lat"`
	Longditude float32  `json:"lon"`
	IcsID      string   `json:"icsID"`
	Modes      []string `json:"modes"`
}

type StoppointSearchResponse struct {
	Query        string           `json:"query"`
	TotalMatches int              `json:"total"`
	Matches      []StoppointMatch `json:"matches"`
}
