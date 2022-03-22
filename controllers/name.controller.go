package controllers

import (
	"github.com/aimenhamed/go-ms/interfaces"
	"github.com/aimenhamed/go-ms/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type NamesController struct {
	service interfaces.Service
}

func (c NamesController) ServiceHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("NamesController called.")
	c.service.HandleRequest(w, r)
}

func (c NamesController) RegisterRoutes(r *mux.Router) {
	c.service = services.NamesService{}
	r.StrictSlash(true)
	r.HandleFunc("/names", c.ServiceHandler)
	log.Printf("NamesRoutes registered.")
}
