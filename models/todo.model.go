package models

import (
	"encoding/json"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type Todo struct {
	gorm.Model

	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

func (t Todo) GetTodos() (*[]Todo, error) {
	var todos []Todo
	result := db.Find(&todos)
	if result.Error != nil {
		panic(result.Error)
	}
	//fmt.Println(todos)
	return &todos, nil
	//json.NewEncoder(w).Encode(&resources)
}

func (t Todo) GetTodo(todoId int) (*Todo, error) {
	var todo Todo
	result := db.First(&todo, todoId)
	if result.Error != nil {
		return nil, result.Error
		//fmt.Println(result.Error)
		//panic(result.Error)
	}
	return &todo, nil
	//json.NewEncoder(w).Encode(&resource)
}

func (t Todo) CreateTodo(body []byte) (*Todo, error) {
	var todo Todo
	//json.Unmarshal(body, &todo)
	if err := json.Unmarshal(body, &todo); err != nil {
		return nil, err
	}
	result := db.Create(&todo)
	if result.Error != nil {
		return nil, result.Error
	}
	return &todo, nil
}
