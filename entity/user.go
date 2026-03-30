package entity

import "time"

type Club struct {
	ID   uint   `json:"id" gorm:"primary_key;auto_increment"`
	Name string `json:"name" gorm:"not null;size:255"`
}

type User struct {
	ID           uint      `json:"id" gorm:"primary_key;auto_increment"`
	Email        string    `json:"email,omitempty" gorm:"size:255"`
	PasswordHash string    `json:"-" gorm:"size:255"`
	HasPassword  bool      `json:"has_password" gorm:"-"`
	Name         string    `json:"name" gorm:"not null;size:255"`
	Bio          string    `json:"bio,omitempty"`
	PhotoURL     string    `json:"photo_url,omitempty" gorm:"size:512"`
	BirthDate    string    `json:"birth_date,omitempty" gorm:"size:10"`
	Gender       string    `json:"gender,omitempty" gorm:"size:10"`
	Role          string    `json:"role" gorm:"not null;default:'student';size:20"`
	EmailVerified bool     `json:"email_verified" gorm:"default:false"`
	ClubID       *uint     `json:"club_id,omitempty"`
	ClubStatus   string    `json:"club_status,omitempty" gorm:"size:20"`
	Club         *Club     `json:"club,omitempty" gorm:"foreignkey:ClubID"`
	LastLoginAt  *time.Time `json:"last_login_at,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Belts        []Belt         `json:"belts,omitempty" gorm:"foreignkey:UserID"`
	Competitions []Competition  `json:"competitions,omitempty" gorm:"foreignkey:UserID"`
	QuizResults  []QuizResult   `json:"quiz_results,omitempty" gorm:"foreignkey:UserID"`
	Licenses     []License      `json:"licenses,omitempty" gorm:"foreignkey:UserID"`
}

type Belt struct {
	ID             uint   `json:"id" gorm:"primary_key;auto_increment"`
	UserID         uint   `json:"user_id" gorm:"not null"`
	Color          string `json:"color" gorm:"not null;size:20"`
	GraduationDate string `json:"graduation_date" gorm:"not null;size:10"`
	ExaminerID     *uint  `json:"examiner_id,omitempty"`
	ExaminerName   string `json:"examiner_name,omitempty" gorm:"size:255"`
}

type Competition struct {
	ID       uint   `json:"id" gorm:"primary_key;auto_increment"`
	UserID   uint   `json:"user_id" gorm:"not null"`
	ClubID   *uint  `json:"club_id,omitempty"`
	Name     string `json:"name" gorm:"not null;size:255"`
	Date     string `json:"date" gorm:"not null;size:10"`
	Link     string `json:"link,omitempty" gorm:"size:512"`
	Result   string `json:"result" gorm:"not null;size:20"`
	Category    string `json:"category,omitempty" gorm:"size:20"`
	WeightClass string `json:"weight_class,omitempty" gorm:"size:20"`
	Deleted     bool   `json:"deleted" gorm:"default:false"`
}

type QuizResult struct {
	ID             uint      `json:"id" gorm:"primary_key;auto_increment"`
	UserID         uint      `json:"user_id" gorm:"not null"`
	Belt           string    `json:"belt" gorm:"not null;size:20"`
	TotalQuestions int       `json:"total_questions" gorm:"not null"`
	CorrectAnswers int       `json:"correct_answers" gorm:"not null"`
	CreatedAt      time.Time `json:"created_at"`
}

type VerificationToken struct {
	ID        uint      `json:"id" gorm:"primary_key;auto_increment"`
	Email     string    `json:"email" gorm:"not null;size:255"`
	Token     string    `json:"token" gorm:"not null;size:64;index"`
	Purpose   string    `json:"purpose" gorm:"not null;size:20"`
	Data      string    `json:"data,omitempty" gorm:"type:text"`
	ExpiresAt time.Time `json:"expires_at" gorm:"not null"`
	Used      bool      `json:"used" gorm:"default:false"`
}

type License struct {
	ID        uint   `json:"id" gorm:"primary_key;auto_increment"`
	UserID    uint   `json:"user_id" gorm:"not null"`
	Name      string `json:"name" gorm:"not null;size:255"`
	IssuedAt  string `json:"issued_at" gorm:"not null;size:10"`
	ExpiresAt string `json:"expires_at,omitempty" gorm:"size:10"`
}

type CoachStudent struct {
	ID        uint `json:"id" gorm:"primary_key;auto_increment"`
	CoachID   uint `json:"coach_id" gorm:"not null;unique_index:idx_coach_student"`
	StudentID uint `json:"student_id" gorm:"not null;unique_index:idx_coach_student"`
}
