package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes Routes

func AppendRoutes(rt Routes) {
	routes = append(routes, rt...)
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router = ResourceRouter(router)
	router = TodoRouter(router)

	/*
		n := negroni.Classic()
		n.Use(negroni.HandlerFunc(middlewareFirst))
		n.Use(negroni.HandlerFunc(middlewareSecond))
		n.UseHandler(router)
	*/

	/*
		for _, route := range routes {
			var handler http.Handler

			handler = route.HandlerFunc
			handler = Logger(handler, route.Name)

			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)

		}
	*/

	return router
}
