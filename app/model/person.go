package model

// Person describes a person
type Person struct {
	Name      string  `json:"name"`
	Height    float64 `json:"height"`
	BirthYear string  `json:"birthYear"`
	Gender    string  `json:"gender"`
	Ships     []Ship  `json:"ships,omitempty"`
}

// AddShip adds a ship to a person
func (p *Person) AddShip(s Ship) {
	p.Ships = append(p.Ships, s)
}
