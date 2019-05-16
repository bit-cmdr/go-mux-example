package app

import (
	"log"
	"net/http"

	"github.com/bit-cmdr/go-mux-example/app/handler"
	"github.com/gorilla/mux"
)

// App has router and db instances
type App struct {
	Router *mux.Router
}

// Initialize starts the app routes
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Get handles get requests
func (a *App) Get(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, handler).Methods("GET")
}

// Run runs the app
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func (a *App) setRouters() {
	a.Get("/people", handler.GetAllPeople)
	a.Get("/ships", handler.GetAllShips)
}
