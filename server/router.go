package server

import (
	"github.com/gorilla/mux"
)

//NewRouter - Returns a pointer to a mux.Router used as a handler.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
