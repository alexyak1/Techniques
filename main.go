package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
// Article - Our struct for all articles
type Technique struct {
    Id      string    `json:"Id"`
    Name   string `json:"name"`
    Belt string `json:"belt"`
}

var Techniques []Technique

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the HomePage of judo tecniques!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnAllTechnique(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Techniques)
}

func returnSingleTechnique(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    for _, technique := range Techniques {
        if technique.Id == key {
            json.NewEncoder(w).Encode(technique)
        }
    }
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
    myRouter.HandleFunc("/techniques", returnAllTechnique)
    myRouter.HandleFunc("/technique", createNewTechnique).Methods("POST")
    myRouter.HandleFunc("/technique/{id}", deleteTechnique).Methods("DELETE")
    myRouter.HandleFunc("/technique/{id}", returnSingleTechnique)
    log.Fatal(http.ListenAndServe(":8787", myRouter))
}

func main() {

        // Capture connection properties.
        cfg := mysql.Config{
            User:   os.Getenv("root"),
            Passwd: os.Getenv("judo-test-password"),
            Net:    "tcp",
            Addr:   "127.0.0.1:3306",
            DBName: "techniques",
        }
        // Get a database handle.
        var err error
        db, err = sql.Open("mysql", cfg.FormatDSN())
        if err != nil {
            log.Fatal(err)
        }

        pingErr := db.Ping()
        if pingErr != nil {
            log.Fatal(pingErr)
        }
        fmt.Println("Connected!")
    Techniques = []Technique{
        Technique{Id: "1", Name: "O-soto-otoshi", Belt: "yellow"},
        Technique{Id: "2", Name: "O-goshi", Belt: "yellow"},
        Technique{Id: "3", Name: "Uchi-mata", Belt: "Orange"},
    }
    handleRequests()
}