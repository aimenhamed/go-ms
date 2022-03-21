package controllers

import (
	"log"
	"net/http"

	"github.com/aimenhamed/go-ms/services"
	"github.com/gorilla/mux"
)

func RegisterNamesRoutes(r *mux.Router) {
	r.StrictSlash(true)
	r.HandleFunc("/names", NamesController)
	log.Printf("NamesRoutes registered.")
}

func NamesController(w http.ResponseWriter, r *http.Request) {
	log.Printf("NamesController called.")

	services.NamesHandler(w, r)
}
