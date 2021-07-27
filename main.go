package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    for _, technique := range Techniques {
        if technique.Id == key {
            json.NewEncoder(w).Encode(technique)
        }
    }
}


func createNewArticle(w http.ResponseWriter, r *http.Request) {
    reqBody, _ := ioutil.ReadAll(r.Body)
    var technique Technique
    json.Unmarshal(reqBody, &technique)
    // update our global Articles array to include
    // our new Article
    Techniques = append(Techniques, technique)

    json.NewEncoder(w).Encode(technique)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    for index, article := range Techniques {
        if article.Id == id {
            Techniques = append(Techniques[:index], Techniques[index+1:]...)
        }
    }

}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/techniques", returnAllTechnique)
    myRouter.HandleFunc("/technique", createNewArticle).Methods("POST")
    myRouter.HandleFunc("/technique/{id}", deleteArticle).Methods("DELETE")
    myRouter.HandleFunc("/technique/{id}", returnSingleArticle)
    log.Fatal(http.ListenAndServe(":8787", myRouter))
}

func main() {
    Techniques = []Technique{
        Technique{Id: "1", Name: "O-soto-otoshi", Belt: "yellow"},
        Technique{Id: "2", Name: "O-goshi", Belt: "yellow"},
        Technique{Id: "3", Name: "Uchi-mata", Belt: "Orange"},
    }
    handleRequests()
}