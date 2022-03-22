package services

import (
	"fmt"
	"log"
	"net/http"
)

type NamesService struct {
}

func (s NamesService) HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("NamesService called.")
	fmt.Fprint(w, "list names here")
}
