package schema

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name             string `gorm:"column:name;not null"`
	Email            string `gorm:"column:email;unique;not null"`
	Password         string `gorm:"column:password;not null"`
	DefaultFocusTime int    `gorm:"column:default_focus_time;default:25"`
	DefaultRestTime  int    `gorm:"column:default_rest_time;default:5"`

	Subjects       []Subject      `gorm:"constraint:OnDelete:CASCADE;"`
	StudySessions  []StudySession `gorm:"constraint:OnDelete:CASCADE;"`
	ScheduleEvents []Schedule     `gorm:"constraint:OnDelete:CASCADE;"`
	AIHistories    []AIHistory    `gorm:"constraint:OnDelete:CASCADE;"`
}

type Subject struct {
	gorm.Model
	Name   string `gorm:"column:name;not null"`
	UserID uint   `gorm:"column:user_id;not null"`

	Flashcards     []Flashcard    `gorm:"constraint:OnDelete:CASCADE;"`
	StudySessions  []StudySession `gorm:"constraint:OnDelete:CASCADE;"`
	ScheduleEvents []Schedule     `gorm:"constraint:OnDelete:SET NULL;"`
}

type Flashcard struct {
	gorm.Model
	Question        string    `gorm:"column:question;not null"`
	Answer          string    `gorm:"column:answer;not null"`
	DifficultyLevel string    `gorm:"column:difficulty_level"`
	NextReviewDate  time.Time `gorm:"column:next_review_date"`
	SubjectID       uint      `gorm:"column:subject_id;not null"`

	StudyHistories []StudyHistory `gorm:"constraint:OnDelete:CASCADE;"`
}

type StudySession struct {
	gorm.Model
	ExecutionDate   time.Time `gorm:"column:execution_date"`
	DurationMinutes int       `gorm:"column:duration_minutes;not null"`
	CompletedCycles int       `gorm:"column:completed_cycles;default:0"`
	UserID          uint      `gorm:"column:user_id;not null"`
	SubjectID       uint      `gorm:"column:subject_id;not null"`
}

type StudyHistory struct {
	gorm.Model
	ReviewDate  time.Time `gorm:"column:review_date"`
	Performance string    `gorm:"column:performance"` // Ex: "Easy", "Medium", "Hard"
	FlashcardID uint      `gorm:"column:flashcard_id;not null"`
}

type Schedule struct {
	gorm.Model
	Title       string    `gorm:"column:title;not null"`
	Description string    `gorm:"column:description"`
	Date        time.Time `gorm:"column:date;not null"`
	Priority    string    `gorm:"column:priority"`
	UserID      uint      `gorm:"column:user_id;not null"`
	SubjectID   uint      `gorm:"column:subject_id"`
}

type AIHistory struct {
	gorm.Model
	Question  string `gorm:"column:question;not null"`
	Answer    string `gorm:"column:answer;not null"`
	UserID    uint   `gorm:"column:user_id;not null"`
	SubjectID uint   `gorm:"column:subject_id"`
}
