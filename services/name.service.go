package services

import (
	"fmt"
	"log"
	"net/http"
)

func NamesHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("NamseHandler called.")
	fmt.Fprint(w, "list names here")
}
