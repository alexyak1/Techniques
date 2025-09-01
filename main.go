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
	// Always use the Docker database container and credentials
	dbHost := "godockerDB"             // Docker service name
	dbUser := "root"                   // User for MySQL
	dbPassword := "judo-test-password" // Password for MySQL
	dbName := "techniques"             // Database name

	// Set up the database configuration
	config := &database.Config{
		ServerName: dbHost,
		User:       dbUser,
		Hash:       dbPassword,
		DB:         dbName,
	}

	// Get the connection string and connect to the database
	connectionString := database.GetConnectionString(*config)
	err := database.Connect(connectionString)
	if err != nil {
		fmt.Printf("Connection problem to SQL: %v\n", err)
		return
	}

	// Run migration to ensure the DB schema is up to date
	database.Migrate(&entity.Technique{})
}

func handleRequests() {
	port := os.Getenv("PORT")

	// Default port is 8787 if not set
	if port == "" {
		port = "8787"
	}

	// Print the port number
	fmt.Println("Starting server on port:", port)

	// Create a new router
	myRouter := mux.NewRouter().StrictSlash(true)

	// Define routes
	myRouter.HandleFunc("/technique", controllers.CreateTechnique).Methods("POST")
	myRouter.HandleFunc("/technique/{id}", controllers.DeleteTechniqueById).Methods("DELETE")
	myRouter.HandleFunc("/technique/{id}", controllers.UpdateTechniqueById).Methods("PUT")
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/techniques", controllers.GetAllTechniques)
	myRouter.HandleFunc("/technique/{id}", controllers.GetTechniqueById)

	// Kata techniques routes
	myRouter.HandleFunc("/kata", controllers.CreateKataTechnique).Methods("POST")
	myRouter.HandleFunc("/kata", controllers.GetAllKataTechniques)

	// Start the server
	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi judoka!\nWelcome to the HomePage of judo techniques! ")
	fmt.Fprintf(w, "\nTo get all techniques, visit this endpoint: /techniques")
}
