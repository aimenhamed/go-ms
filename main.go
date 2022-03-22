package main

import (
	"github.com/aimenhamed/go-ms/interfaces"
	"log"
	"net/http"

	"github.com/aimenhamed/go-ms/controllers"
	"github.com/gorilla/mux"
)

func serve(port string, r http.Handler) {
	err := http.ListenAndServe(":"+port, r)

	if err != nil {
		log.Fatal("Serving error.")
	}
}

func initialiseController(c interfaces.Controller, path *mux.Router) {
	c.RegisterRoutes(path)
}

func main() {
	r := mux.NewRouter()
	v1 := r.PathPrefix("/api/v1").Subrouter()
	initialiseController(controllers.NamesController{}, v1)

	serve("8000", r)
}
