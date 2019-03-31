package main

import (
	"fmt"
	"log"

	"github.com/UmerIftikhar/goSqlAzure/models"
	"github.com/UmerIftikhar/goSqlAzure/routers"
	"github.com/codegangsta/negroni"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := models.OpenConnection()

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	if db != nil {
		fmt.Println("----- Connected to the DB -----")
	}

	//AutoMigrate all the functionality within this function
	models.AutoMigrate()
	//db.AutoMigrate(&Resource{})

	defer models.CloseConnection()
	//fmt.Println(os.Getenv("DBNAME"))
	router := routers.NewRouter()

	n := negroni.New()
	n.Use(negroni.HandlerFunc(routers.Logger))
	n.UseHandler(router)
	n.Run(":4040")
	//log.Fatal(http.ListenAndServe(":4040", router))
}
