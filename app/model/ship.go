package model

// Ship describes a starship
type Ship struct {
	Name         string  `json:"name"`
	Model        string  `json:"model"`
	Manufacturer string  `json:"manufacturer"`
	Length       float64 `json:"length"`
	Crew         int     `json:"crew"`
	Passengers   int     `json:"passangers"`
	Class        string  `json:"class"`
}
