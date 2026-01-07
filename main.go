package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"main/controllers"
	"main/database"
	"main/entity"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// gzipResponseWriter wraps http.ResponseWriter to provide gzip compression
type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

// gzipMiddleware compresses responses with gzip if the client supports it
func gzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if client accepts gzip
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		// Set gzip headers
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Del("Content-Length") // Content length changes with compression

		gz := gzip.NewWriter(w)
		defer gz.Close()

		gzw := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		next.ServeHTTP(gzw, r)
	})
}

// cacheMiddleware adds Cache-Control headers to GET responses
func cacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Only cache GET requests
		if r.Method == http.MethodGet {
			// Cache for 5 minutes, allow stale content while revalidating
			w.Header().Set("Cache-Control", "public, max-age=300, stale-while-revalidate=60")
		}
		next.ServeHTTP(w, r)
	})
}

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
	if err := database.Migrate(&entity.Technique{}); err != nil {
		fmt.Printf("Migration failed: %v\n", err)
		return
	}
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

	// Apply middleware: gzip compression + cache headers
	handler := gzipMiddleware(cacheMiddleware(myRouter))

	// Start the server
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi judoka!\nWelcome to the HomePage of judo techniques! ")
	fmt.Fprintf(w, "\nTo get all techniques, visit this endpoint: /techniques")
}
