package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"main/database"
	"main/entity"
	"main/middleware"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func GetAllClubsForUser(w http.ResponseWriter, r *http.Request) {
	var clubs []entity.Club
	database.Connector.Find(&clubs)
	if clubs == nil {
		clubs = []entity.Club{}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clubs)
}

func UserJoinClub(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r)

	var req struct {
		ClubID uint `json:"club_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	database.Connector.Model(&entity.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"club_id":     req.ClubID,
		"club_status": "pending",
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "pending"})
}

func UserLeaveClub(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r)
	database.Connector.Model(&entity.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"club_id":     nil,
		"club_status": "",
	})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "left"})
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r)

	var req struct {
		Name     string `json:"name"`
		PhotoURL string `json:"photo_url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	updates := map[string]interface{}{}
	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.PhotoURL != "" {
		updates["photo_url"] = req.PhotoURL
	}

	if len(updates) == 0 {
		http.Error(w, `{"error":"No fields to update"}`, http.StatusBadRequest)
		return
	}

	if err := database.Connector.Model(&entity.User{}).Where("id = ?", userID).Updates(updates).Error; err != nil {
		http.Error(w, `{"error":"Failed to update profile"}`, http.StatusInternalServerError)
		return
	}

	var user entity.User
	database.Connector.Preload("Belts").Preload("Competitions").Preload("QuizResults").First(&user, userID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func UploadPhoto(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r)

	r.ParseMultipartForm(5 << 20) // 5MB max

	file, header, err := r.FormFile("photo")
	if err != nil {
		http.Error(w, `{"error":"No photo file provided"}`, http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Validate file type
	ext := strings.ToLower(filepath.Ext(header.Filename))
	allowed := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".webp": true}
	if !allowed[ext] {
		http.Error(w, `{"error":"Only jpg, png, gif, webp files are allowed"}`, http.StatusBadRequest)
		return
	}

	// Create uploads directory
	uploadDir := "/app/uploads"
	os.MkdirAll(uploadDir, 0755)

	// Generate unique filename
	filename := fmt.Sprintf("user_%d_%d%s", userID, time.Now().UnixNano(), ext)
	filePath := filepath.Join(uploadDir, filename)

	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, `{"error":"Failed to save file"}`, http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, `{"error":"Failed to save file"}`, http.StatusInternalServerError)
		return
	}

	// Update user photo URL
	photoURL := "/uploads/" + filename
	database.Connector.Model(&entity.User{}).Where("id = ?", userID).Update("photo_url", photoURL)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"photo_url": photoURL})
}

func AddBelt(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r)

	var belt entity.Belt
	if err := json.NewDecoder(r.Body).Decode(&belt); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if belt.Color == "" || belt.GraduationDate == "" {
		http.Error(w, `{"error":"Color and graduation_date are required"}`, http.StatusBadRequest)
		return
	}

	belt.UserID = userID
	if err := database.Connector.Create(&belt).Error; err != nil {
		http.Error(w, `{"error":"Failed to add belt"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(belt)
}

func DeleteBelt(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r)
	beltID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		http.Error(w, `{"error":"Invalid belt ID"}`, http.StatusBadRequest)
		return
	}

	result := database.Connector.Where("id = ? AND user_id = ?", beltID, userID).Delete(&entity.Belt{})
	if result.Error != nil {
		http.Error(w, `{"error":"Failed to delete belt"}`, http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(w, `{"error":"Belt not found"}`, http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func AddCompetition(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r)

	var comp entity.Competition
	if err := json.NewDecoder(r.Body).Decode(&comp); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if comp.Name == "" || comp.Date == "" || comp.Result == "" {
		http.Error(w, `{"error":"Name, date, and result are required"}`, http.StatusBadRequest)
		return
	}

	comp.UserID = userID
	if err := database.Connector.Create(&comp).Error; err != nil {
		http.Error(w, `{"error":"Failed to add competition"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comp)
}

func DeleteCompetition(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r)
	compID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		http.Error(w, `{"error":"Invalid competition ID"}`, http.StatusBadRequest)
		return
	}

	result := database.Connector.Where("id = ? AND user_id = ?", compID, userID).Delete(&entity.Competition{})
	if result.Error != nil {
		http.Error(w, `{"error":"Failed to delete competition"}`, http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(w, `{"error":"Competition not found"}`, http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func SaveQuizResult(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r)

	var result entity.QuizResult
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if result.Belt == "" || result.TotalQuestions == 0 {
		http.Error(w, `{"error":"Belt and total_questions are required"}`, http.StatusBadRequest)
		return
	}

	result.UserID = userID
	if err := database.Connector.Create(&result).Error; err != nil {
		http.Error(w, `{"error":"Failed to save quiz result"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func GetUserQuizResults(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r)

	var results []entity.QuizResult
	if err := database.Connector.Where("user_id = ?", userID).Order("created_at desc").Find(&results).Error; err != nil {
		http.Error(w, `{"error":"Failed to retrieve quiz results"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

// Coach endpoints

func CreateCoachCompetition(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name       string `json:"name"`
		Date       string `json:"date"`
		Link       string `json:"link"`
		StudentIDs []uint `json:"student_ids"`
		Category   string `json:"category"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.Name == "" || req.Date == "" || len(req.StudentIDs) == 0 {
		http.Error(w, `{"error":"Name, date, and at least one student are required"}`, http.StatusBadRequest)
		return
	}

	var created []entity.Competition
	for _, sid := range req.StudentIDs {
		comp := entity.Competition{
			UserID:   sid,
			Name:     req.Name,
			Date:     req.Date,
			Link:     req.Link,
			Result:   "participated",
			Category: req.Category,
		}
		if err := database.Connector.Create(&comp).Error; err != nil {
			continue
		}
		created = append(created, comp)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func UpdateCompetitionCategory(w http.ResponseWriter, r *http.Request) {
	compID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		http.Error(w, `{"error":"Invalid competition ID"}`, http.StatusBadRequest)
		return
	}

	var req struct {
		Category string `json:"category"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	database.Connector.Model(&entity.Competition{}).Where("id = ?", compID).Update("category", req.Category)

	var comp entity.Competition
	database.Connector.First(&comp, compID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comp)
}

func UpdateCompetitionResult(w http.ResponseWriter, r *http.Request) {
	compID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		http.Error(w, `{"error":"Invalid competition ID"}`, http.StatusBadRequest)
		return
	}

	var req struct {
		Result string `json:"result"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validResults := map[string]bool{"gold": true, "silver": true, "bronze": true, "participated": true}
	if !validResults[req.Result] {
		http.Error(w, `{"error":"Result must be gold, silver, bronze, or participated"}`, http.StatusBadRequest)
		return
	}

	if err := database.Connector.Model(&entity.Competition{}).Where("id = ?", compID).Update("result", req.Result).Error; err != nil {
		http.Error(w, `{"error":"Failed to update result"}`, http.StatusInternalServerError)
		return
	}

	var comp entity.Competition
	database.Connector.First(&comp, compID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comp)
}

func getCoachClubID(coachID uint) *uint {
	var coach entity.User
	database.Connector.First(&coach, coachID)
	if coach.ClubStatus != "approved" {
		return nil
	}
	return coach.ClubID
}

func GetAvailableStudents(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r)
	clubID := getCoachClubID(userID)

	var students []entity.User
	query := database.Connector.Where("role = ?", "student")
	if clubID != nil {
		query = query.Where("club_id = ?", *clubID)
	}

	// Exclude already assigned
	var assigned []entity.CoachStudent
	database.Connector.Where("coach_id = ?", userID).Find(&assigned)
	assignedIDs := make([]uint, len(assigned))
	for i, cs := range assigned {
		assignedIDs[i] = cs.StudentID
	}
	if len(assignedIDs) > 0 {
		query = query.Where("id NOT IN (?)", assignedIDs)
	}

	query.Find(&students)
	if students == nil {
		students = []entity.User{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func CoachCreateStudent(w http.ResponseWriter, r *http.Request) {
	coachID := middleware.GetUserIDFromContext(r)

	var req struct {
		Name        string `json:"name"`
		ClubID      *uint  `json:"club_id"`
		NewClubName string `json:"new_club_name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.Name == "" {
		http.Error(w, `{"error":"Name is required"}`, http.StatusBadRequest)
		return
	}

	// Determine club: explicit > new club > coach's club
	var clubID *uint
	if req.NewClubName != "" {
		club := entity.Club{Name: req.NewClubName}
		database.Connector.Create(&club)
		clubID = &club.ID
	} else if req.ClubID != nil {
		clubID = req.ClubID
	} else {
		clubID = getCoachClubID(coachID)
	}

	student := entity.User{
		Name:   req.Name,
		Role:   "student",
		ClubID: clubID,
	}
	if err := database.Connector.Create(&student).Error; err != nil {
		http.Error(w, `{"error":"Failed to create student"}`, http.StatusInternalServerError)
		return
	}

	// Auto-assign to this coach
	cs := entity.CoachStudent{CoachID: coachID, StudentID: student.ID}
	database.Connector.Create(&cs)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

func CoachAddStudent(w http.ResponseWriter, r *http.Request) {
	coachID := middleware.GetUserIDFromContext(r)

	var req struct {
		StudentID uint `json:"student_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	cs := entity.CoachStudent{CoachID: coachID, StudentID: req.StudentID}
	if err := database.Connector.Create(&cs).Error; err != nil {
		http.Error(w, `{"error":"Failed to add student"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cs)
}

func CoachRemoveStudent(w http.ResponseWriter, r *http.Request) {
	coachID := middleware.GetUserIDFromContext(r)
	studentID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		http.Error(w, `{"error":"Invalid student ID"}`, http.StatusBadRequest)
		return
	}

	result := database.Connector.Where("coach_id = ? AND student_id = ?", coachID, studentID).Delete(&entity.CoachStudent{})
	if result.Error != nil {
		http.Error(w, `{"error":"Failed to remove student"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r)
	role := middleware.GetUserRoleFromContext(r)

	var students []entity.User

	if role == "admin" {
		database.Connector.Where("role = ?", "student").
			Preload("Belts").Preload("Competitions").Preload("QuizResults").Preload("Club").
			Find(&students)
	} else {
		// Coach: get students in same club
		clubID := getCoachClubID(userID)
		if clubID != nil {
			database.Connector.Where("role = ? AND club_id = ?", "student", *clubID).
				Preload("Belts").Preload("Competitions").Preload("QuizResults").Preload("Club").
				Find(&students)
		}
	}

	if students == nil {
		students = []entity.User{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func GetStudentProfile(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r)
	role := middleware.GetUserRoleFromContext(r)

	studentID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		http.Error(w, `{"error":"Invalid student ID"}`, http.StatusBadRequest)
		return
	}

	// Coach can view students in their club or assigned to them
	if role == "coach" {
		clubID := getCoachClubID(userID)
		var student entity.User
		database.Connector.First(&student, studentID)

		inClub := clubID != nil && student.ClubID != nil && *clubID == *student.ClubID
		var cs entity.CoachStudent
		assigned := database.Connector.Where("coach_id = ? AND student_id = ?", userID, studentID).First(&cs).Error == nil

		if !inClub && !assigned {
			http.Error(w, `{"error":"Not authorized to view this student"}`, http.StatusForbidden)
			return
		}
	}

	var student entity.User
	if err := database.Connector.
		Preload("Belts").Preload("Competitions").Preload("QuizResults").
		First(&student, studentID).Error; err != nil {
		http.Error(w, `{"error":"Student not found"}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func CoachAddBelt(w http.ResponseWriter, r *http.Request) {
	coachID := middleware.GetUserIDFromContext(r)
	studentID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		http.Error(w, `{"error":"Invalid student ID"}`, http.StatusBadRequest)
		return
	}

	var belt entity.Belt
	if err := json.NewDecoder(r.Body).Decode(&belt); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if belt.Color == "" || belt.GraduationDate == "" {
		http.Error(w, `{"error":"Color and graduation_date are required"}`, http.StatusBadRequest)
		return
	}

	belt.UserID = uint(studentID)

	// Default examiner to current coach if not specified
	if belt.ExaminerID == nil {
		belt.ExaminerID = &coachID
	}
	// Resolve examiner name
	if belt.ExaminerID != nil && belt.ExaminerName == "" {
		var examiner entity.User
		if database.Connector.First(&examiner, *belt.ExaminerID).Error == nil {
			belt.ExaminerName = examiner.Name
		}
	}

	if err := database.Connector.Create(&belt).Error; err != nil {
		http.Error(w, `{"error":"Failed to add belt"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(belt)
}

func CoachDeleteBelt(w http.ResponseWriter, r *http.Request) {
	studentID, _ := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	beltID, _ := strconv.ParseUint(mux.Vars(r)["beltId"], 10, 64)

	result := database.Connector.Where("id = ? AND user_id = ?", beltID, studentID).Delete(&entity.Belt{})
	if result.RowsAffected == 0 {
		http.Error(w, `{"error":"Belt not found"}`, http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func CoachAddCompetition(w http.ResponseWriter, r *http.Request) {
	studentID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		http.Error(w, `{"error":"Invalid student ID"}`, http.StatusBadRequest)
		return
	}

	var comp entity.Competition
	if err := json.NewDecoder(r.Body).Decode(&comp); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if comp.Name == "" || comp.Date == "" || comp.Result == "" {
		http.Error(w, `{"error":"Name, date, and result are required"}`, http.StatusBadRequest)
		return
	}

	comp.UserID = uint(studentID)
	if err := database.Connector.Create(&comp).Error; err != nil {
		http.Error(w, `{"error":"Failed to add competition"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comp)
}

func CoachDeleteCompetition(w http.ResponseWriter, r *http.Request) {
	studentID, _ := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	compID, _ := strconv.ParseUint(mux.Vars(r)["compId"], 10, 64)

	result := database.Connector.Where("id = ? AND user_id = ?", compID, studentID).Delete(&entity.Competition{})
	if result.RowsAffected == 0 {
		http.Error(w, `{"error":"Competition not found"}`, http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func CoachUpdateStudentProfile(w http.ResponseWriter, r *http.Request) {
	studentID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		http.Error(w, `{"error":"Invalid student ID"}`, http.StatusBadRequest)
		return
	}

	var req struct {
		Name     string `json:"name"`
		PhotoURL string `json:"photo_url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	updates := map[string]interface{}{}
	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.PhotoURL != "" {
		updates["photo_url"] = req.PhotoURL
	}

	if len(updates) > 0 {
		database.Connector.Model(&entity.User{}).Where("id = ?", studentID).Updates(updates)
	}

	var user entity.User
	database.Connector.Preload("Belts").Preload("Competitions").Preload("QuizResults").First(&user, studentID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Admin endpoints

func AdminCreateCoach(w http.ResponseWriter, r *http.Request) {
	adminID := middleware.GetUserIDFromContext(r)

	var req struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.Name == "" {
		http.Error(w, `{"error":"Name is required"}`, http.StatusBadRequest)
		return
	}

	// Get admin's club
	var admin entity.User
	database.Connector.First(&admin, adminID)

	coach := entity.User{
		Name:          req.Name,
		Role:          "coach",
		ClubID:        admin.ClubID,
		ClubStatus:    "approved",
		EmailVerified: true,
	}
	if err := database.Connector.Create(&coach).Error; err != nil {
		http.Error(w, `{"error":"Failed to create coach"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(coach)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []entity.User
	database.Connector.Find(&users)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func UpdateUserRole(w http.ResponseWriter, r *http.Request) {
	targetID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		http.Error(w, `{"error":"Invalid user ID"}`, http.StatusBadRequest)
		return
	}

	var req struct {
		Role string `json:"role"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validRoles := map[string]bool{"student": true, "coach": true, "admin": true}
	if !validRoles[req.Role] {
		http.Error(w, `{"error":"Role must be student, coach, or admin"}`, http.StatusBadRequest)
		return
	}

	if err := database.Connector.Model(&entity.User{}).Where("id = ?", targetID).Update("role", req.Role).Error; err != nil {
		http.Error(w, `{"error":"Failed to update role"}`, http.StatusInternalServerError)
		return
	}

	var user entity.User
	database.Connector.First(&user, targetID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func AssignStudentToCoach(w http.ResponseWriter, r *http.Request) {
	currentUserID := middleware.GetUserIDFromContext(r)
	role := middleware.GetUserRoleFromContext(r)

	var req struct {
		CoachID   uint `json:"coach_id"`
		StudentID uint `json:"student_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Coach can only assign to themselves
	if role == "coach" && req.CoachID != currentUserID {
		http.Error(w, `{"error":"Coaches can only assign students to themselves"}`, http.StatusForbidden)
		return
	}

	cs := entity.CoachStudent{CoachID: req.CoachID, StudentID: req.StudentID}
	if err := database.Connector.Create(&cs).Error; err != nil {
		http.Error(w, `{"error":"Failed to assign student"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cs)
}

func RemoveStudentFromCoach(w http.ResponseWriter, r *http.Request) {
	currentUserID := middleware.GetUserIDFromContext(r)
	role := middleware.GetUserRoleFromContext(r)

	csID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		http.Error(w, `{"error":"Invalid ID"}`, http.StatusBadRequest)
		return
	}

	// Coach can only remove their own assignments
	query := database.Connector.Where("id = ?", csID)
	if role == "coach" {
		query = query.Where("coach_id = ?", currentUserID)
	}

	result := query.Delete(&entity.CoachStudent{})
	if result.Error != nil {
		http.Error(w, `{"error":"Failed to remove assignment"}`, http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(w, `{"error":"Assignment not found"}`, http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetClubCoaches(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r)
	role := middleware.GetUserRoleFromContext(r)
	clubID := getCoachClubID(userID)

	var coaches []entity.User
	if clubID != nil {
		if role == "admin" {
			database.Connector.Where("club_id = ? AND club_status = ? AND (role = ? OR role = ?)", *clubID, "approved", "coach", "admin").Find(&coaches)
		} else {
			database.Connector.Where("club_id = ? AND club_status = ? AND role = ?", *clubID, "approved", "coach").Find(&coaches)
		}
	}
	if coaches == nil {
		coaches = []entity.User{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coaches)
}

func GetCoachProfile(w http.ResponseWriter, r *http.Request) {
	coachID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		http.Error(w, `{"error":"Invalid coach ID"}`, http.StatusBadRequest)
		return
	}

	var coach entity.User
	if err := database.Connector.
		Preload("Belts").Preload("Competitions").Preload("Club").
		First(&coach, coachID).Error; err != nil {
		http.Error(w, `{"error":"Coach not found"}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coach)
}

// Club management (coach)

func GetAllClubsPublic(w http.ResponseWriter, r *http.Request) {
	var clubs []entity.Club
	database.Connector.Find(&clubs)
	if clubs == nil {
		clubs = []entity.Club{}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clubs)
}

func CoachCreateClub(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r)

	var req struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.Name == "" {
		http.Error(w, `{"error":"Club name is required"}`, http.StatusBadRequest)
		return
	}

	club := entity.Club{Name: req.Name}
	if err := database.Connector.Create(&club).Error; err != nil {
		http.Error(w, `{"error":"Failed to create club"}`, http.StatusInternalServerError)
		return
	}

	// Auto-assign and approve the creator
	database.Connector.Model(&entity.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"club_id":     club.ID,
		"club_status": "approved",
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(club)
}

func CoachJoinClub(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r)

	var req struct {
		ClubID uint `json:"club_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	database.Connector.Model(&entity.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"club_id":     req.ClubID,
		"club_status": "pending",
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "pending"})
}

func GetClubRequests(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r)
	role := middleware.GetUserRoleFromContext(r)

	var pending []entity.User

	if role == "admin" {
		database.Connector.Where("club_status = ?", "pending").Preload("Club").Find(&pending)
	} else {
		clubID := getCoachClubID(userID)
		if clubID != nil {
			database.Connector.Where("club_id = ? AND club_status = ?", *clubID, "pending").Preload("Club").Find(&pending)
		}
	}

	if pending == nil {
		pending = []entity.User{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pending)
}

func ApproveCoach(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r)
	role := middleware.GetUserRoleFromContext(r)

	targetID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		http.Error(w, `{"error":"Invalid user ID"}`, http.StatusBadRequest)
		return
	}

	// Verify permission: admin or approved coach in same club
	var target entity.User
	database.Connector.First(&target, targetID)

	if role != "admin" {
		clubID := getCoachClubID(userID)
		var approver entity.User
		database.Connector.First(&approver, userID)
		if clubID == nil || target.ClubID == nil || *clubID != *target.ClubID || approver.ClubStatus != "approved" {
			http.Error(w, `{"error":"Not authorized to approve"}`, http.StatusForbidden)
			return
		}
	}

	database.Connector.Model(&entity.User{}).Where("id = ?", targetID).Update("club_status", "approved")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "approved"})
}

func RejectCoach(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r)
	role := middleware.GetUserRoleFromContext(r)

	targetID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		http.Error(w, `{"error":"Invalid user ID"}`, http.StatusBadRequest)
		return
	}

	var target entity.User
	database.Connector.First(&target, targetID)

	if role != "admin" {
		clubID := getCoachClubID(userID)
		var approver entity.User
		database.Connector.First(&approver, userID)
		if clubID == nil || target.ClubID == nil || *clubID != *target.ClubID || approver.ClubStatus != "approved" {
			http.Error(w, `{"error":"Not authorized"}`, http.StatusForbidden)
			return
		}
	}

	// Remove club association
	database.Connector.Model(&entity.User{}).Where("id = ?", targetID).Updates(map[string]interface{}{
		"club_id":     nil,
		"club_status": "",
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "rejected"})
}

// Club management (admin)

func GetAllClubs(w http.ResponseWriter, r *http.Request) {
	var clubs []entity.Club
	database.Connector.Find(&clubs)
	if clubs == nil {
		clubs = []entity.Club{}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clubs)
}

func CreateClub(w http.ResponseWriter, r *http.Request) {
	var club entity.Club
	if err := json.NewDecoder(r.Body).Decode(&club); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if club.Name == "" {
		http.Error(w, `{"error":"Club name is required"}`, http.StatusBadRequest)
		return
	}

	if err := database.Connector.Create(&club).Error; err != nil {
		http.Error(w, `{"error":"Failed to create club"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(club)
}

func DeleteClub(w http.ResponseWriter, r *http.Request) {
	clubID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		http.Error(w, `{"error":"Invalid club ID"}`, http.StatusBadRequest)
		return
	}

	database.Connector.Where("id = ?", clubID).Delete(&entity.Club{})
	w.WriteHeader(http.StatusNoContent)
}

func UpdateUserClub(w http.ResponseWriter, r *http.Request) {
	targetID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		http.Error(w, `{"error":"Invalid user ID"}`, http.StatusBadRequest)
		return
	}

	var req struct {
		ClubID *uint `json:"club_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	database.Connector.Model(&entity.User{}).Where("id = ?", targetID).Update("club_id", req.ClubID)

	var user entity.User
	database.Connector.Preload("Club").First(&user, targetID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
