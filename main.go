package main

import (
	"database/sql"
	"fmt"
	"log"
	"main/controllers"
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
	var err error
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection success")
	log.Println(db)
	// database.Migrate(&entity.Technique{})
}

func handleRequests() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8787"
	}
	fmt.Printf("Port is: %s", port)

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
		port = "8787"
	}
	fmt.Println(port)

	fmt.Fprintf(w, "Hi judoka!\nWelcome to the HomePage of judo tecniques! ")
	fmt.Fprintf(w, "\nFor get all techniques visit this endpoint:\n http://18.221.140.18:"+port+"/techniques")

	fmt.Println("Endpoint Hit: homePage")
}
