package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/UmerIftikhar/goSqlAzure/common"
	"github.com/UmerIftikhar/goSqlAzure/models"
	"github.com/gorilla/mux"
)

//https://golang.org/pkg/net/http/
//GetResources
//var todos repository.Todos

/*
type TodoInput struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

func (a TodoInput) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required, validation.Length(1, 50), is.Alpha),
		validation.Field(&a.Completed),
		validation.Field(&a.Due, validation.Date("2018-05-05")),
	)
}
*/
type TodoCtrl struct {
	toDoModel models.Todo
}

//var toDoModel models.Todo

func (tdCtrl *TodoCtrl) Index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Welcome!\n")
}

func (tdCtrl *TodoCtrl) TodoIndex(w http.ResponseWriter, r *http.Request) {
	toDoModel := tdCtrl.toDoModel
	toDoResult, err := toDoModel.GetTodos()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(&toDoResult)

}

func (tdCtrl *TodoCtrl) TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var todoId int
	toDoModel := tdCtrl.toDoModel
	var err error
	if todoId, err = strconv.Atoi(vars["todoId"]); err != nil {
		//panic(err)
		errMess := common.ErrorMessage{Name: err.Error(), Code: http.StatusBadRequest}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&errMess)
		return
	}

	toDoResult, err := toDoModel.GetTodo(todoId)
	w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	if err != nil {
		errMess := common.ErrorMessage{Name: "Given ID does not exist in the Data Base", Code: http.StatusNotFound}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&errMess)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&toDoResult)

}

/*
Test with this curl command:

curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos

*/
func (tdCtrl *TodoCtrl) TodoCreate(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	toDoModel := tdCtrl.toDoModel
	if err != nil {
		errMess := common.ErrorMessage{Name: err.Error(), Code: http.StatusBadRequest}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&errMess)
		return
	}

	toDoResult, err := toDoModel.CreateTodo(body)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		errMess := common.ErrorMessage{Name: "Server Encountered error while creating new item", Code: http.StatusInternalServerError}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errMess)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&toDoResult)

}
