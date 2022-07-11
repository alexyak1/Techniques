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
	db_password := os.Getenv("DB_PASSWORD")
	config :=
		database.Config{}
	if db_password != "" {
		// config =
		// 	database.Config{
		// 		ServerName: "remotemysql.com",
		// 		User:       "hzhf7kfMUy",
		// 		Hash:       db_password,
		// 		DB:         "hzhf7kfMUy",
		// 	}
		config =
			database.Config{
				ServerName: "db4free.net/",
				User:       "notrootuser",
				Hash:       db_password,
				DB:         "techniques",
			}
	} else {
		config =
			database.Config{
				ServerName: "godockerDB",
				User:       "root",
				Hash:       "judo-test-password",
				DB:         "techniques",
			}
	}
	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		fmt.Printf("Connection problem to SQL")
	}
	database.Migrate(&entity.Technique{})
}

func handleRequests() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8787"
	}
	fmt.Println(port)

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
