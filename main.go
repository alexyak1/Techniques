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
	// db_password := os.Getenv("DB_PASSWORD")
	// config := &database.Config{}
	// if db_password != "" {
	// 	*config =
	// 		database.Config{
	// 			ServerName: "34.88.169.215",
	// 			User:       "kano",
	// 			Hash:       "workwork",
	// 			DB:         "techniques",
	// 		}
	// } else {
	// 	*config =
	// 		database.Config{
	// 			ServerName: "godockerDB",
	// 			User:       "root",
	// 			Hash:       "judo-test-password",
	// 			DB:         "techniques",
	// 		}
	// }
	// connectionString := database.GetConnectionString(*config)

	connectionName := "concrete-plasma-424808-a7:europe-north1:kano"
	user := "kano"
	password := "workwork"
	databaseName := "techniques"
	socketDir := "/cloudsql"

	// Construct the connection string
	connectionString := fmt.Sprintf("%s:%s@unix(%s/%s)/%s", user, password, socketDir, connectionName, databaseName)

	// Connect to the database using the connection string
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	// Ping the database to check the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	fmt.Println("Connected to the database!")
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

	// Kata techniques
	myRouter.HandleFunc("/kata", controllers.CreateKataTechnique).Methods("POST")
	myRouter.HandleFunc("/kata", controllers.GetAllKataTechniques)

	myRouter.HandleFunc("/blog", controllers.GetBlogData)

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
