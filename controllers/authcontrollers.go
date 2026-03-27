package controllers

import (
	"encoding/json"
	"fmt"
	"main/database"
	"main/email"
	"main/entity"
	"main/middleware"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type authResponse struct {
	User  entity.User `json:"user"`
	Token string      `json:"token"`
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func getFrontendURL() string {
	url := os.Getenv("FRONTEND_URL")
	if url == "" {
		return "http://localhost:3000"
	}
	return url
}

// Register creates account (unverified) and sends verification email
func Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	req.Name = strings.TrimSpace(req.Name)

	if req.Name == "" {
		http.Error(w, `{"error":"Name is required"}`, http.StatusBadRequest)
		return
	}
	if !emailRegex.MatchString(req.Email) {
		http.Error(w, `{"error":"Invalid email format"}`, http.StatusBadRequest)
		return
	}
	if len(req.Password) < 6 {
		http.Error(w, `{"error":"Password must be at least 6 characters"}`, http.StatusBadRequest)
		return
	}

	var existing entity.User
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, `{"error":"Failed to hash password"}`, http.StatusInternalServerError)
		return
	}

	var user entity.User
	if database.Connector.Where("email = ?", req.Email).First(&existing).Error == nil {
		if existing.EmailVerified && existing.PasswordHash != "" {
			http.Error(w, `{"error":"Email already registered"}`, http.StatusConflict)
			return
		}
		// Coach-created or unverified account — take it over, keep existing data (belts, competitions, club)
		database.Connector.Model(&existing).Updates(map[string]interface{}{
			"password_hash":  string(hash),
			"name":           req.Name,
			"email_verified": false,
		})
		user = existing
	} else {
		user = entity.User{
			Email:         req.Email,
			PasswordHash:  string(hash),
			Name:          req.Name,
			Role:          "student",
			EmailVerified: false,
		}
		if err := database.Connector.Create(&user).Error; err != nil {
			http.Error(w, `{"error":"Failed to create user"}`, http.StatusInternalServerError)
			return
		}
	}

	// Create verification token
	token := email.GenerateToken()
	vt := entity.VerificationToken{
		Email:     req.Email,
		Token:     token,
		Purpose:   "register",
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	database.Connector.Create(&vt)

	link := fmt.Sprintf("%s/verify?token=%s", getFrontendURL(), token)
	email.SendVerificationLink(req.Email, link, "register")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "Verification email sent. Please check your inbox."})
}

// VerifyEmail handles the token from the email link
func VerifyEmail(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, `{"error":"Token is required"}`, http.StatusBadRequest)
		return
	}

	var vt entity.VerificationToken
	if err := database.Connector.
		Where("token = ? AND purpose = ? AND used = ? AND expires_at > ?", token, "register", false, time.Now()).
		First(&vt).Error; err != nil {
		http.Error(w, `{"error":"Invalid or expired link"}`, http.StatusBadRequest)
		return
	}

	// Mark token as used
	database.Connector.Model(&vt).Update("used", true)

	// Verify the user
	database.Connector.Model(&entity.User{}).Where("email = ?", vt.Email).Update("email_verified", true)

	// Load user and generate JWT
	var user entity.User
	database.Connector.Preload("Club").Preload("Belts").Preload("Competitions").Preload("QuizResults").
		Where("email = ?", vt.Email).First(&user)

	jwtToken, err := middleware.GenerateToken(user.ID, user.Role)
	if err != nil {
		http.Error(w, `{"error":"Failed to generate token"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authResponse{User: user, Token: jwtToken})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	req.Email = strings.TrimSpace(strings.ToLower(req.Email))

	var user entity.User
	if err := database.Connector.Preload("Club").Where("email = ?", req.Email).First(&user).Error; err != nil {
		http.Error(w, `{"error":"Invalid email or password"}`, http.StatusUnauthorized)
		return
	}

	if !user.EmailVerified {
		http.Error(w, `{"error":"Please verify your email first. Check your inbox."}`, http.StatusForbidden)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		http.Error(w, `{"error":"Invalid email or password"}`, http.StatusUnauthorized)
		return
	}

	token, err := middleware.GenerateToken(user.ID, user.Role)
	if err != nil {
		http.Error(w, `{"error":"Failed to generate token"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authResponse{User: user, Token: token})
}

// ResendVerification resends the verification email
func ResendVerification(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	req.Email = strings.TrimSpace(strings.ToLower(req.Email))

	var user entity.User
	if err := database.Connector.Where("email = ? AND email_verified = ?", req.Email, false).First(&user).Error; err != nil {
		http.Error(w, `{"error":"Email not found or already verified"}`, http.StatusBadRequest)
		return
	}

	// Invalidate old tokens
	database.Connector.Model(&entity.VerificationToken{}).
		Where("email = ? AND purpose = ? AND used = ?", req.Email, "register", false).
		Update("used", true)

	token := email.GenerateToken()
	vt := entity.VerificationToken{
		Email:     req.Email,
		Token:     token,
		Purpose:   "register",
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	database.Connector.Create(&vt)

	link := fmt.Sprintf("%s/verify?token=%s", getFrontendURL(), token)
	email.SendVerificationLink(req.Email, link, "register")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Verification email sent"})
}

// RequestPasswordChange sends a confirmation link for password change
func RequestPasswordChange(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r)

	var req struct {
		CurrentPassword string `json:"current_password"`
		NewPassword     string `json:"new_password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if len(req.NewPassword) < 6 {
		http.Error(w, `{"error":"New password must be at least 6 characters"}`, http.StatusBadRequest)
		return
	}

	var user entity.User
	if err := database.Connector.First(&user, userID).Error; err != nil {
		http.Error(w, `{"error":"User not found"}`, http.StatusNotFound)
		return
	}

	// Verify current password
	if user.PasswordHash != "" {
		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.CurrentPassword)); err != nil {
			http.Error(w, `{"error":"Current password is incorrect"}`, http.StatusUnauthorized)
			return
		}
	}

	// Hash new password and store in token data
	newHash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, `{"error":"Failed to process password"}`, http.StatusInternalServerError)
		return
	}

	// Invalidate old tokens
	database.Connector.Model(&entity.VerificationToken{}).
		Where("email = ? AND purpose = ? AND used = ?", user.Email, "password", false).
		Update("used", true)

	token := email.GenerateToken()
	vt := entity.VerificationToken{
		Email:     user.Email,
		Token:     token,
		Purpose:   "password",
		Data:      string(newHash),
		ExpiresAt: time.Now().Add(10 * time.Minute),
	}
	database.Connector.Create(&vt)

	link := fmt.Sprintf("%s/verify?token=%s", getFrontendURL(), token)
	email.SendVerificationLink(user.Email, link, "password")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Confirmation email sent. Check your inbox."})
}

// ConfirmPasswordChange handles the token from the password change email
func ConfirmPasswordChange(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, `{"error":"Token is required"}`, http.StatusBadRequest)
		return
	}

	var vt entity.VerificationToken
	if err := database.Connector.
		Where("token = ? AND purpose = ? AND used = ? AND expires_at > ?", token, "password", false, time.Now()).
		First(&vt).Error; err != nil {
		http.Error(w, `{"error":"Invalid or expired link"}`, http.StatusBadRequest)
		return
	}

	// Mark token as used
	database.Connector.Model(&vt).Update("used", true)

	// Apply the new password hash
	database.Connector.Model(&entity.User{}).Where("email = ?", vt.Email).Update("password_hash", vt.Data)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Password changed successfully"})
}

// ForgotPassword sends a reset link to the user's email
func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	req.Email = strings.TrimSpace(strings.ToLower(req.Email))

	// Always return success to not leak whether email exists
	w.Header().Set("Content-Type", "application/json")

	var user entity.User
	if database.Connector.Where("email = ?", req.Email).First(&user).Error != nil {
		json.NewEncoder(w).Encode(map[string]string{"status": "If the email exists, a reset link has been sent."})
		return
	}

	// Invalidate old tokens
	database.Connector.Model(&entity.VerificationToken{}).
		Where("email = ? AND purpose = ? AND used = ?", req.Email, "reset", false).
		Update("used", true)

	token := email.GenerateToken()
	vt := entity.VerificationToken{
		Email:     req.Email,
		Token:     token,
		Purpose:   "reset",
		ExpiresAt: time.Now().Add(30 * time.Minute),
	}
	database.Connector.Create(&vt)

	link := fmt.Sprintf("%s/reset-password?token=%s", getFrontendURL(), token)
	email.SendVerificationLink(req.Email, link, "password")

	json.NewEncoder(w).Encode(map[string]string{"status": "If the email exists, a reset link has been sent."})
}

// ResetPassword sets a new password using the reset token
func ResetPassword(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Token       string `json:"token"`
		NewPassword string `json:"new_password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if len(req.NewPassword) < 6 {
		http.Error(w, `{"error":"Password must be at least 6 characters"}`, http.StatusBadRequest)
		return
	}

	var vt entity.VerificationToken
	if err := database.Connector.
		Where("token = ? AND purpose = ? AND used = ? AND expires_at > ?", req.Token, "reset", false, time.Now()).
		First(&vt).Error; err != nil {
		http.Error(w, `{"error":"Invalid or expired reset link"}`, http.StatusBadRequest)
		return
	}

	database.Connector.Model(&vt).Update("used", true)

	hash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, `{"error":"Failed to process password"}`, http.StatusInternalServerError)
		return
	}

	database.Connector.Model(&entity.User{}).Where("email = ?", vt.Email).Update("password_hash", string(hash))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Password reset successfully"})
}

func AcceptInvite(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Token    string `json:"token"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if len(req.Password) < 6 {
		http.Error(w, `{"error":"Password must be at least 6 characters"}`, http.StatusBadRequest)
		return
	}

	var vt entity.VerificationToken
	if err := database.Connector.
		Where("token = ? AND purpose = ? AND used = ? AND expires_at > ?", req.Token, "invite", false, time.Now()).
		First(&vt).Error; err != nil {
		http.Error(w, `{"error":"Invalid or expired invite link"}`, http.StatusBadRequest)
		return
	}

	database.Connector.Model(&vt).Update("used", true)

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, `{"error":"Failed to process password"}`, http.StatusInternalServerError)
		return
	}

	database.Connector.Model(&entity.User{}).Where("email = ?", vt.Email).Updates(map[string]interface{}{
		"password_hash":  string(hash),
		"email_verified": true,
	})

	var user entity.User
	database.Connector.Preload("Club").Preload("Belts").Preload("Competitions").Preload("QuizResults").
		Where("email = ?", vt.Email).First(&user)

	jwtToken, err := middleware.GenerateToken(user.ID, user.Role)
	if err != nil {
		http.Error(w, `{"error":"Failed to generate token"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authResponse{User: user, Token: jwtToken})
}

func GetMe(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r)

	var user entity.User
	if err := database.Connector.
		Preload("Belts").
		Preload("Competitions").
		Preload("QuizResults").
		Preload("Club").
		First(&user, userID).Error; err != nil {
		http.Error(w, `{"error":"User not found"}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
