package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

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
    log.Fatal(http.ListenAndServe(":8788", myRouter))
}

func main() {
    // db, err := sql.Open("mysql", "root:judo-test-password@tcp(0.0.0.1:3306)/techniques")

    // // if there is an error opening the connection, handle it
    // if err != nil {
    //     log.Print(err.Error())
    // }
    // defer db.Close()

    // // Execute the query
    // results, err := db.Query("SELECT id, name FROM techniques")
    // if err != nil {
    //     panic(err.Error()) // proper error handling instead of panic in your app
    // }

    // for results.Next() {
    //     var technique Technique
    //     // for each row, scan the result into our tag composite object
    //     err = results.Scan(&technique.Id, &technique.Name)
    //     if err != nil {
    //         panic(err.Error()) // proper error handling instead of panic in your app
    //     }
    //             // and then print out the tag's Name attribute
    //     log.Printf(technique.Name)
    // }


    Techniques = []Technique{
        Technique{Id: "1", Name: "O-soto-otoshi", Belt: "yellow"},
        Technique{Id: "2", Name: "O-goshi", Belt: "yellow"},
        Technique{Id: "3", Name: "Uchi-mata", Belt: "Orange"},
    }
    handleRequests()
}