package server

import "net/http"

//Route - Defines a single route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes - an array of Route elements
type Routes []Route

var routes = Routes{
	Route{
		Name:        "Download_Signature",
		Method:      "GET",
		Pattern:     "/signature",
		HandlerFunc: SignatureFunc,
	},
	Route{
		"HealthCheck",
		"GET",
		"/health",
		HealthCheck,
	},
}
