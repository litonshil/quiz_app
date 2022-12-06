package domain

import (
	"quiz_app/app/models"
)

type IQuizRepo interface {
	GetQuiz() ([]models.Question, error)
}

type IQuizSvc interface {
	GetQuiz() ([]models.Question, error)
}
