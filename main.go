package main

import (
	"log"
	"net/http"

	"github.com/aimenhamed/go-ms/controllers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	v1 := r.PathPrefix("/api/v1").Subrouter()
	controllers.RegisterNamesRoutes(v1)

	err := http.ListenAndServe(":8000", r)

	if err != nil {
		log.Fatal("Serving error.")
	}
}
