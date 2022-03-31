package main

import (
	"fmt"
	"log"
	"main/controllers"
	"main/database"
	"main/entity"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	initDB()
	handleRequests()
}

func initDB() {
	config :=
		database.Config{
			// ServerName: "godockerDB",
			// User: "root",
			// Password: "judo-test-password",
			// DB: "techniques",
			ServerName: "sql11.freemysqlhosting.net",
			User:       "sql11482611",
			Password:   "ccFfrgmwy7",
			DB:         "sql11482611",
		}
	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Technique{})
}

func handleRequests() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/technique", controllers.CreateTechnique).Methods("POST")
	myRouter.HandleFunc("/technique/{id}", controllers.DeleteTechniqueById).Methods("DELETE")
	myRouter.HandleFunc("/technique/{id}", controllers.UpdateTechniqueById).Methods("PUT")
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/techniques", controllers.GetAllTechniques)
	myRouter.HandleFunc("/technique/{id}", controllers.GetTechniqueById)
	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	fmt.Fprintf(w, "Hi judoka!\nWelcome to the HomePage of judo tecniques! ")
	fmt.Fprintf(w, "\nFor get all techniques visit this endpoint:\n http://18.221.140.18:"+port+"/techniques")

	fmt.Println("Endpoint Hit: homePage")
}
