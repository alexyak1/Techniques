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
	config := &database.Config{}
	if db_password != "" {
		*config =
			database.Config{
				ServerName: "sql.freedb.tech",
				User:       "freedb_alexyak1",
				Hash:       db_password,
				DB:         "freedb_techniques",
			}
	} else {
		*config =
			database.Config{
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
	certFile := "/etc/letsencrypt/live/judoquiz.com/fullchain.pem"
	keyFile := "/etc/letsencrypt/live/judoquiz.com/privkey.pem"

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

	log.Fatal(http.ListenAndServeTLS(":"+port, certFile, keyFile, myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8787"
	}

	fmt.Fprintf(w, "Hi judoka!\nWelcome to the HomePage of judo techniques! ")
	fmt.Fprintf(w, "\nFor get all techniques visit this endpoint:\n IP:"+port+"/techniques")

	fmt.Println("Endpoint Hit: homePage")
}
