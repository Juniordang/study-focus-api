package schema

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name             string `gorm:"column:nome;not null"`
	Email            string `gorm:"column:email;not null;unique"`
	Password         string `gorm:"column:senha;not nul"`
	DefaultTimeFocus int    `gorm:"column:tempo_foco_padrao;default:25"`
	DefaultBreakTime int    `gorm:"column:tempo_pausa_padrao;default:5"`
}

type Discipline struct {
	gorm.Model
	Name          string         `gorm:"column:nome;not null"`
	UserID        uint           `gorm:"column:usuario_id;not null"`
	Flashcards    []Flashcard    `gorm:"constraint:OnDelete:CASCADE;"`
	SessoesEstudo []StudySession `gorm:"constraint:OnDelete:CASCADE;"`
}

type Flashcard struct {
	gorm.Model
	Question       string    `gorm:"column:pergunta;not null"`
	Response       string    `gorm:"column:resposta;not null"`
	Dificulty      string    `gorm:"column:nivel_dificuldade"`
	NextReviewDate time.Time `gorm:"column:datetiddata_proxima_revisao"`
	DisciplineID   uint      `gorm:"column:disciplina_id;not null"`
}

type StudySession struct {
	gorm.Model
	durationMinutes int  `gorm:"column:duracao_minutos;not null"`
	CompletedCycles int  `gorm:"column:ciclos_concluidos;default:0"`
	UserID          uint `gorm:"column:usuario_id;not null"`
	DisciplineID    uint `gorm:"column:disciplina_id;not null"`
}
