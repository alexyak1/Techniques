package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

var Techniques []Technique


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

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the HomePage of judo tecniques!")
    fmt.Fprintf(w, "---All techniques here: http://3.15.169.17:8787/techniques")

    fmt.Println("Endpoint Hit: homePage")
}

func createNewTechnique(w http.ResponseWriter, r *http.Request) {
    reqBody, _ := ioutil.ReadAll(r.Body)
    var technique Technique
    json.Unmarshal(reqBody, &technique)

    // fake update
    Techniques = append(Techniques, technique)

    json.NewEncoder(w).Encode(technique)
}

func deleteTechnique(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    for index, technique := range Techniques {
        if technique.Id == id {
            Techniques = append(Techniques[:index], Techniques[index+1:]...)
        }
    }

}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/techniques", controllers.GetAllTechniques)
    myRouter.HandleFunc("/technique/{id}", controllers.GetTechniqueById)
    myRouter.HandleFunc("/technique", createNewTechnique).Methods("POST")
    myRouter.HandleFunc("/technique/{id}", deleteTechnique).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8787", myRouter))
}

func main() {
    initDB()
    handleRequests()
}