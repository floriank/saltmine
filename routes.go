package main

import (
	"net/http"
)

// Route unifies route generation by moving everything into a struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is just a specific collection of Route
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/projects",
		ProjectIndex,
	},
	Route{
		"Create Projects",
		"POST",
		"/projects",
		ProjectCreate,
	},
	Route{
		"Update Projects",
		"PATCH",
		"/projects/{identifier}",
		ProjectUpdate,
	},
}
