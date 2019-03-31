package routers

import (
	"github.com/UmerIftikhar/goSqlAzure/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func ResourceRouter(router *mux.Router) *mux.Router {

	ResourceCtrl := controllers.ResourceCtrl{}
	resourceRouter := mux.NewRouter()
	resourceRouter.HandleFunc("/resources", ResourceCtrl.ResourceIndex).Methods("GET")
	resourceRouter.HandleFunc("/resources", ResourceCtrl.ResourceShow).Methods("POST")
	resourceRouter.HandleFunc("/resources/{resourceId}", ResourceCtrl.ResourceCreate).Methods("GET")
	router.PathPrefix("/resources").Handler(negroni.New(
		//negroni.HandlerFunc(logger),
		negroni.Wrap(resourceRouter),
	))

	return router
}

func init() {
	/*
		ResourceCtrl := controllers.ResourceCtrl{}
		routes := Routes{
			Route{
				"ResourceIndex",
				"GET",
				"/resources",
				ResourceCtrl.ResourceIndex,
			},
			Route{
				"ResourceCreate",
				"POST",
				"/resources",
				ResourceCtrl.ResourceCreate,
			},
			Route{
				"ResourceShow",
				"GET",
				"/resources/{resourceId}",
				ResourceCtrl.ResourceShow,
			},
		}
		AppendRoutes(routes)
	*/
	//AppendRoutes(Route{"Index", "GET", "/", controllers.Index})
}

/*
routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controllers.Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		controllers.TodoIndex,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		controllers.TodoCreate,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		controllers.TodoShow,
	},
}
*/
