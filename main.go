package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"main/controllers"
	"main/database"
	"main/entity"
	"main/middleware"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/cors"
	"golang.org/x/crypto/bcrypt"
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
		if r.Method == http.MethodGet {
			// Don't cache authenticated/user-specific responses
			if strings.HasPrefix(r.URL.Path, "/auth/") ||
				strings.HasPrefix(r.URL.Path, "/user/") ||
				strings.HasPrefix(r.URL.Path, "/coach/") ||
				strings.HasPrefix(r.URL.Path, "/admin/") {
				w.Header().Set("Cache-Control", "no-store")
			} else {
				w.Header().Set("Cache-Control", "public, max-age=300, stale-while-revalidate=60")
			}
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	initDB()
	handleRequests()
}

func initDB() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	if dbHost == "" {
		dbHost = "godockerDB"
	}
	if dbUser == "" {
		dbUser = "root"
	}
	if dbName == "" {
		dbName = "techniques"
	}

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

	// Migrate new auth tables
	database.Connector.AutoMigrate(
		&entity.Club{},
		&entity.User{},
		&entity.Belt{},
		&entity.Competition{},
		&entity.QuizResult{},
		&entity.CoachStudent{},
		&entity.License{},
		&entity.VerificationToken{},
		&entity.MergeLog{},
	)

	// Seed default admin if none exists
	var count int
	database.Connector.Model(&entity.User{}).Where("role = ?", "admin").Count(&count)
	if count == 0 {
		adminEmail := os.Getenv("ADMIN_EMAIL")
		adminPass := os.Getenv("ADMIN_PASSWORD")
		if adminEmail == "" || adminPass == "" {
			log.Println("ADMIN_EMAIL and ADMIN_PASSWORD env vars required to create admin")
			return
		}
		hash, _ := bcrypt.GenerateFromPassword([]byte(adminPass), bcrypt.DefaultCost)
		admin := entity.User{
			Email:         adminEmail,
			PasswordHash:  string(hash),
			Name:          "Admin",
			Role:          "admin",
			EmailVerified: true,
		}
		database.Connector.Create(&admin)
		log.Printf("Default admin created: %s", adminEmail)
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

	// Existing public routes
	myRouter.HandleFunc("/technique", controllers.CreateTechnique).Methods("POST")
	myRouter.HandleFunc("/technique/{id}", controllers.DeleteTechniqueById).Methods("DELETE")
	myRouter.HandleFunc("/technique/{id}", controllers.UpdateTechniqueById).Methods("PUT")
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/techniques", controllers.GetAllTechniques)
	myRouter.HandleFunc("/technique/{id}", controllers.GetTechniqueById)

	// Kata techniques routes
	myRouter.HandleFunc("/kata", controllers.CreateKataTechnique).Methods("POST")
	myRouter.HandleFunc("/kata", controllers.GetAllKataTechniques)

	// Public auth routes
	myRouter.HandleFunc("/auth/register", controllers.Register).Methods("POST")
	myRouter.HandleFunc("/auth/login", controllers.Login).Methods("POST")
	myRouter.HandleFunc("/auth/verify-email", controllers.VerifyEmail).Methods("GET")
	myRouter.HandleFunc("/auth/resend-verification", controllers.ResendVerification).Methods("POST")
	myRouter.HandleFunc("/auth/confirm-password", controllers.ConfirmPasswordChange).Methods("GET")
	myRouter.HandleFunc("/auth/forgot-password", controllers.ForgotPassword).Methods("POST")
	myRouter.HandleFunc("/auth/reset-password", controllers.ResetPassword).Methods("POST")
	myRouter.HandleFunc("/auth/accept-invite", controllers.AcceptInvite).Methods("POST")
	myRouter.HandleFunc("/auth/accept-club-invite", controllers.AcceptClubInvite).Methods("GET")

	// Protected routes - any authenticated user
	protected := myRouter.PathPrefix("").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.HandleFunc("/auth/me", controllers.GetMe).Methods("GET")
	protected.HandleFunc("/user/clubs", controllers.GetAllClubsForUser).Methods("GET")
	protected.HandleFunc("/user/join-club", controllers.UserJoinClub).Methods("POST")
	protected.HandleFunc("/user/leave-club", controllers.UserLeaveClub).Methods("POST")
	protected.HandleFunc("/user/club-coaches", controllers.GetMyClubCoaches).Methods("GET")
	protected.HandleFunc("/user/club-competitions", controllers.GetMyClubCompetitions).Methods("GET")
	protected.HandleFunc("/user/club-stats", controllers.GetMyClubStats).Methods("GET")
	protected.HandleFunc("/user/club-competitions-full", controllers.GetMyClubCompetitionsFull).Methods("GET")
	protected.HandleFunc("/user/club-member/{id}", controllers.GetClubMemberProfile).Methods("GET")
	protected.HandleFunc("/user/licenses", controllers.AddLicense).Methods("POST")
	protected.HandleFunc("/user/licenses/{id}", controllers.DeleteLicense).Methods("DELETE")
	protected.HandleFunc("/user/profile", controllers.UpdateProfile).Methods("PUT")
	protected.HandleFunc("/user/change-password", controllers.RequestPasswordChange).Methods("PUT")
	protected.HandleFunc("/user/upload-photo", controllers.UploadPhoto).Methods("POST")
	protected.HandleFunc("/user/belts", controllers.AddBelt).Methods("POST")
	protected.HandleFunc("/user/belts/{id}", controllers.DeleteBelt).Methods("DELETE")
	protected.HandleFunc("/user/competitions", controllers.AddCompetition).Methods("POST")
	protected.HandleFunc("/user/competitions/{id}", controllers.DeleteCompetition).Methods("DELETE")
	protected.HandleFunc("/user/quiz-results", controllers.SaveQuizResult).Methods("POST")
	protected.HandleFunc("/user/quiz-results", controllers.GetUserQuizResults).Methods("GET")

	// Coach + Admin routes
	coachRoutes := myRouter.PathPrefix("/coach").Subrouter()
	coachRoutes.Use(middleware.AuthMiddleware)
	coachRoutes.Use(middleware.RequireRole("coach", "admin"))
	coachRoutes.HandleFunc("/students", controllers.GetStudents).Methods("GET")
	coachRoutes.HandleFunc("/students/{id}", controllers.GetStudentProfile).Methods("GET")
	coachRoutes.HandleFunc("/competitions", controllers.GetClubCompetitions).Methods("GET")
	coachRoutes.HandleFunc("/club-stats", controllers.GetClubStats).Methods("GET")
	coachRoutes.HandleFunc("/competitions", controllers.CreateCoachCompetition).Methods("POST")
	coachRoutes.HandleFunc("/competitions/{id}/result", controllers.UpdateCompetitionResult).Methods("PUT")
	coachRoutes.HandleFunc("/competitions/{id}/category", controllers.UpdateCompetitionCategory).Methods("PUT")
	coachRoutes.HandleFunc("/competitions/{id}/weight-class", controllers.UpdateCompetitionWeightClass).Methods("PUT")
	coachRoutes.HandleFunc("/competitions/{id}", controllers.DeleteCompetitionEntry).Methods("DELETE")
	coachRoutes.HandleFunc("/competitions/update-event", controllers.UpdateCompetitionEvent).Methods("PUT")
	coachRoutes.HandleFunc("/competitions/delete-event", controllers.DeleteCompetitionEvent).Methods("POST")
	coachRoutes.HandleFunc("/competitions/restore-event", controllers.RestoreCompetitionEvent).Methods("POST")
	coachRoutes.HandleFunc("/students/{id}/belts", controllers.CoachAddBelt).Methods("POST")
	coachRoutes.HandleFunc("/students/{id}/belts/{beltId}", controllers.CoachDeleteBelt).Methods("DELETE")
	coachRoutes.HandleFunc("/students/{id}/competitions", controllers.CoachAddCompetition).Methods("POST")
	coachRoutes.HandleFunc("/students/{id}/competitions/{compId}", controllers.CoachDeleteCompetition).Methods("DELETE")
	coachRoutes.HandleFunc("/students/{id}/profile", controllers.CoachUpdateStudentProfile).Methods("PUT")
	coachRoutes.HandleFunc("/students/{id}/invite", controllers.InviteStudent).Methods("POST")
	coachRoutes.HandleFunc("/coaches/{id}/licenses", controllers.CoachAddLicense).Methods("POST")
	coachRoutes.HandleFunc("/coaches/{id}/licenses/{licId}", controllers.CoachDeleteLicense).Methods("DELETE")
	coachRoutes.HandleFunc("/available-students", controllers.GetAvailableStudents).Methods("GET")
	coachRoutes.HandleFunc("/add-student", controllers.CoachAddStudent).Methods("POST")
	coachRoutes.HandleFunc("/create-student", controllers.CoachCreateStudent).Methods("POST")
	coachRoutes.HandleFunc("/remove-student/{id}", controllers.CoachRemoveStudent).Methods("DELETE")
	coachRoutes.HandleFunc("/club-coaches", controllers.GetClubCoaches).Methods("GET")
	coachRoutes.HandleFunc("/club-coaches/{id}", controllers.GetCoachProfile).Methods("GET")
	coachRoutes.HandleFunc("/clubs", controllers.GetAllClubsPublic).Methods("GET")
	coachRoutes.HandleFunc("/create-club", controllers.CoachCreateClub).Methods("POST")
	coachRoutes.HandleFunc("/join-club", controllers.CoachJoinClub).Methods("POST")
	coachRoutes.HandleFunc("/club-requests", controllers.GetClubRequests).Methods("GET")
	coachRoutes.HandleFunc("/approve-coach/{id}", controllers.ApproveCoach).Methods("PUT")
	coachRoutes.HandleFunc("/reject-coach/{id}", controllers.RejectCoach).Methods("PUT")

	// Admin-only routes
	adminRoutes := myRouter.PathPrefix("/admin").Subrouter()
	adminRoutes.Use(middleware.AuthMiddleware)
	adminRoutes.Use(middleware.RequireRole("admin"))
	adminRoutes.HandleFunc("/dashboard", controllers.AdminDashboard).Methods("GET")
	adminRoutes.HandleFunc("/create-coach", controllers.AdminCreateCoach).Methods("POST")
	adminRoutes.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	adminRoutes.HandleFunc("/users/{id}/role", controllers.UpdateUserRole).Methods("PUT")
	adminRoutes.HandleFunc("/users/{id}/email", controllers.UpdateUserEmail).Methods("PUT")
	adminRoutes.HandleFunc("/users/{id}/club", controllers.UpdateUserClub).Methods("PUT")
	adminRoutes.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
	adminRoutes.HandleFunc("/coach-students", controllers.AssignStudentToCoach).Methods("POST")
	adminRoutes.HandleFunc("/coach-students/{id}", controllers.RemoveStudentFromCoach).Methods("DELETE")
	adminRoutes.HandleFunc("/invite-to-club", controllers.AdminInviteToClub).Methods("POST")
	adminRoutes.HandleFunc("/merge-preview", controllers.MergePreview).Methods("POST")
	adminRoutes.HandleFunc("/merge-users", controllers.MergeUsers).Methods("POST")
	adminRoutes.HandleFunc("/clubs", controllers.GetAllClubs).Methods("GET")
	adminRoutes.HandleFunc("/clubs", controllers.CreateClub).Methods("POST")
	adminRoutes.HandleFunc("/clubs/{id}", controllers.DeleteClub).Methods("DELETE")

	// Serve uploaded photos
	myRouter.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("/app/uploads"))))

	// CORS configuration
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// Apply middleware: CORS + gzip compression + cache headers
	handler := c.Handler(gzipMiddleware(cacheMiddleware(myRouter)))

	// Start the server
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi judoka!\nWelcome to the HomePage of judo techniques! ")
	fmt.Fprintf(w, "\nTo get all techniques, visit this endpoint: /techniques")
}
