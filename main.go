package main

import (
	"fmt"
	"log"
	"main/controllers"
	"main/database"
	"main/entity"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Technique struct {
    Id      string    `json:"id"`
    Name   string `json:"name"`
    Belt string `json:"belt"`
}

func main() {
    initDB()
    handleRequests()
}

func initDB() {
    config :=
        database.Config{
            ServerName: "godockerDB",
            User: "root",
            Password: "judo-test-password",
            DB: "techniques",
        }
    connectionString := database.GetConnectionString(config)
    err := database.Connect(connectionString)
    if err != nil {
        panic(err.Error())
    }
    database.Migrate(&entity.Technique{})
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/techniques", controllers.GetAllTechniques)
    myRouter.HandleFunc("/technique/{id}", controllers.GetTechniqueById)
    myRouter.HandleFunc("/technique", controllers.CreateTechnique).Methods("POST")
    myRouter.HandleFunc("/technique/{id}", controllers.DeleteTechniqueById).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8787", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi judoka!\nWelcome to the HomePage of judo tecniques! ")
    fmt.Fprintf(w, "\nFor get all techniques visit this endpoint:\n http://18.219.167.56:8787/techniques")

    fmt.Println("Endpoint Hit: homePage")
}
