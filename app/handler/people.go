package handler

import (
	"net/http"

	"github.com/bit-cmdr/go-mux-example/app/model"
)

// GetAllPeople gets all people
func GetAllPeople(w http.ResponseWriter, r *http.Request) {
	people := defaultPeople()
	responseJSON(w, 200, people)
}

func defaultPeople() []model.Person {
	return []model.Person{
		model.Person{
			Name:      "Luke Skywalker",
			Height:    172.0,
			BirthYear: "19bby",
			Gender:    "Male",
		},
		model.Person{
			Name:      "Han Solo",
			Height:    184.8,
			BirthYear: "29bby",
			Gender:    "Male",
		},
		model.Person{
			Name:      "Chewbacca",
			Height:    220.0,
			BirthYear: "200bby",
			Gender:    "Male",
		},
		model.Person{
			Name:      "Leia Organa",
			Height:    150.0,
			BirthYear: "19bby",
			Gender:    "Female",
		},
		model.Person{
			Name:      "Darth Vader",
			Height:    202.0,
			BirthYear: "41.9bby",
			Gender:    "Male",
		},
	}
}
