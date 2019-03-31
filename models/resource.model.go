package models

import (
	"encoding/json"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type Resource struct {
	gorm.Model

	Link        string
	Name        string
	Author      string
	Description string
}

func (t Resource) GetResources() (*[]Resource, error) {
	var resource []Resource
	result := db.Find(&resource)
	if result.Error != nil {
		panic(result.Error)
	}
	//fmt.Println(todos)
	return &resource, nil
	//json.NewEncoder(w).Encode(&resources)
}

func (t Resource) GetResource(resId int) (*Resource, error) {
	var resource Resource
	result := db.First(&resource, resId)
	if result.Error != nil {
		return nil, result.Error
		//fmt.Println(result.Error)
		//panic(result.Error)
	}
	return &resource, nil
	//json.NewEncoder(w).Encode(&resource)
}

func (t Resource) CreateResource(body []byte) (*Resource, error) {
	var resource Resource
	json.Unmarshal(body, &resource)
	if err := json.Unmarshal(body, &resource); err != nil {
		return nil, err
	}
	result := db.Create(&resource)
	if result.Error != nil {
		return nil, result.Error
	}
	return &resource, nil
}
