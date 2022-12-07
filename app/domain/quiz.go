package domain

import (
	"quiz_app/app/models"
	"quiz_app/app/serializers"
)

type IQuizRepo interface {
	GetQuiz(totalQuestion int) ([]models.Question, error)
	SubmitQuiz(req *models.Quiz, response *serializers.SubmitQuizResponse) error
}

type IQuizSvc interface {
	GetQuiz(totalQuestion int) ([]models.Question, error)
	SubmitQuiz(payload *serializers.QuizPayload, response *serializers.SubmitQuizResponse) error
}
