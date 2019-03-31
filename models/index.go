package models

import (
	"fmt"
	"net/url"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var db *gorm.DB
var err error

func OpenConnection() (*gorm.DB, error) {
	dbString := getDBString()
	db, err = gorm.Open("mssql", dbString)
	return db, err
}

func CloseConnection() {
	db.Close()
	fmt.Println(db)
	fmt.Println("------ CLOSING THE DB CONNECTION ------")
}

func AutoMigrate() {
	db.AutoMigrate(&Todo{})
	db.AutoMigrate(&Resource{})
}

func getDBString() string {
	query := url.Values{}
	query.Add("database", os.Getenv("DBNAME"))
	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(os.Getenv("USER"), os.Getenv("PASSWORD")),
		Host:     fmt.Sprintf("%s:%d", os.Getenv("HOST"), 1433),
		RawQuery: query.Encode(),
	}
	return u.String()
}
