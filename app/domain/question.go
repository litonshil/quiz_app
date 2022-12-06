package domain

import (
	"quiz_app/app/models"
	"quiz_app/app/serializers"
)

type IQuestionRepo interface {
	InsertQuestion(req *models.Question) error
	GetQuestion() ([]models.Question, error)
}

type IQuestionSvc interface {
	InsertQuestion(payload *serializers.QuestionPayload) error
	GetQuestion() ([]models.Question, error)
}
