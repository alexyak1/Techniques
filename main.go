package main

import (
	"fmt"
	"log"
	"main/controllers"
	"main/database"
	"main/entity"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	initDB()
	handleRequests()
}

func initDB() {
	db_password := os.Getenv("DB_PASSWORD")
	config := &database.Config{}
	if db_password != "" {
		*config = database.Config{
			ServerName: "sql.freedb.tech",
			User:       "freedb_alexyak1",
			Hash:       db_password,
			DB:         "freedb_techniques",
		}
	} else {
		*config = database.Config{
			ServerName: "godockerDB",
			User:       "root",
			Hash:       "judo-test-password",
			DB:         "techniques",
		}
	}
	connectionString := database.GetConnectionString(*config)
	err := database.Connect(connectionString)
	if err != nil {
		fmt.Printf("Connection problem to SQL. ")
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
	myRouter.HandleFunc("/techniques", controllers.GetAllTechniques).Methods("GET")
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

	// CORS middleware configuration
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"https://www.judoquiz.com"}, // Adjust this to your frontend URL
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	// Use the CORS middleware with your router
	handler := corsHandler.Handler(myRouter)

	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi judoka!\nWelcome to the HomePage of judo techniques! ")
	fmt.Fprintf(w, "\nTo get all techniques, visit this endpoint: /techniques")
}
