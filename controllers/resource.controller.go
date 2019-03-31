package controllers

import (
	"encoding/json"
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
//var resourceModel models.Resource

type ResourceCtrl struct {
	resourceModel models.Resource
}

func (resCtrl *ResourceCtrl) ResourceIndex(w http.ResponseWriter, r *http.Request) {
	resourceModel := resCtrl.resourceModel
	resResult, err := resourceModel.GetResources()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(&resResult)

}

func (resCtrl *ResourceCtrl) ResourceShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var resourceId int
	resourceModel := resCtrl.resourceModel
	var err error
	if resourceId, err = strconv.Atoi(vars["resourceId"]); err != nil {
		//panic(err)
		errMess := common.ErrorMessage{Name: err.Error(), Code: http.StatusBadRequest}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&errMess)
		return
	}

	resResult, err := resourceModel.GetResource(resourceId)
	w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	if err != nil {
		errMess := common.ErrorMessage{Name: "Given ID does not exist in the Data Base", Code: http.StatusNotFound}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&errMess)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&resResult)

}

/*
Test with this curl command:

curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos

*/
func (resCtrl *ResourceCtrl) ResourceCreate(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	resourceModel := resCtrl.resourceModel
	if err != nil {
		errMess := common.ErrorMessage{Name: err.Error(), Code: http.StatusBadRequest}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&errMess)
		return
	}

	resResult, err := resourceModel.CreateResource(body)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		errMess := common.ErrorMessage{Name: "Server Encountered error while creating new item", Code: http.StatusInternalServerError}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errMess)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&resResult)

}
