package routers

import (
	"github.com/UmerIftikhar/goSqlAzure/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

/*
router.PathPrefix("/tasks").Handler(negroni.New(
	negroni.HandlerFunc(common.Authorize),
	negroni.Wrap(taskRouter),
))
*/

func TodoRouter(router *mux.Router) *mux.Router {

	TodoCtrl := controllers.TodoCtrl{}
	todoRouter := mux.NewRouter()
	todoRouter.HandleFunc("/", TodoCtrl.Index).Methods("GET")
	todoRouter.HandleFunc("/todos", TodoCtrl.TodoIndex).Methods("GET")
	todoRouter.HandleFunc("/todos", TodoCtrl.TodoCreate).Methods("POST")
	todoRouter.HandleFunc("/todos/{todoId}", TodoCtrl.TodoShow).Methods("GET")
	router.PathPrefix("/todos").Handler(negroni.New(
		//negroni.HandlerFunc(logger),
		negroni.HandlerFunc(middlewareFirst),
		negroni.HandlerFunc(middlewareSecond),
		negroni.Wrap(todoRouter),
	))

	return router
}

func init() {
	//middleArray := []func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc){middlewareFirst, middlewareSecond}

	/*
		routes := Routes{
			Route{
				"Index",
				"GET",
				"/",
				TodoCtrl.Index,
			},
			Route{
				"TodoIndex",
				"GET",
				"/todos",
				TodoCtrl.TodoIndex,
			},
			Route{
				"TodoCreate",
				"POST",
				"/todos",
				TodoCtrl.TodoCreate,
			},
			Route{
				"TodoShow",
				"GET",
				"/todos/{todoId}",
				TodoCtrl.TodoShow,
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
