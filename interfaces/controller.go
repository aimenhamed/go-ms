package interfaces

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Controller interface {
	RegisterRoutes(path *mux.Router)
	ServiceHandler(w http.ResponseWriter, r *http.Request)
}
