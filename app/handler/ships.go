package handler

import (
	"net/http"

	"github.com/bit-cmdr/go-mux-example/app/model"
)

// GetAllShips gets all ships
func GetAllShips(w http.ResponseWriter, r *http.Request) {
	ships := defaultShips()
	responseJSON(w, 200, ships)
}

func defaultShips() []model.Ship {
	return []model.Ship{
		model.Ship{
			Name:         "Executor",
			Model:        "Executor-class Star Destroyer",
			Manufacturer: "Kuat Drive Yards",
			Length:       19000.0,
			Crew:         279144,
			Passengers:   38000,
			Class:        "Star Dreadnought",
		},
		model.Ship{
			Name:         "Millenium Falcon",
			Model:        "YT-1300 Light Freighter",
			Manufacturer: "Corellian Engineering Corporation",
			Length:       34.37,
			Crew:         4,
			Passengers:   6,
			Class:        "Light Freighter",
		},
		model.Ship{
			Name:         "X-Wing",
			Model:        "T-64 X-Wing",
			Manufacturer: "Incom Corporation",
			Length:       12.5,
			Crew:         1,
			Passengers:   0,
			Class:        "Assault Starfighter",
		},
		model.Ship{
			Name:         "Chimaera",
			Model:        "Imperial-class Star Destroyer Mk II",
			Manufacturer: "Kuat Drive Yards",
			Length:       1600.0,
			Crew:         37000,
			Passengers:   9785,
			Class:        "Star Destroyer",
		},
	}
}
