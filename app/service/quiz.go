package svc

import (
	"quiz_app/app/domain"
	"quiz_app/app/models"
	"quiz_app/app/utils/errutil"
	logger "quiz_app/infra/logger"
)

type quizes struct {
	quizRepo domain.IQuizRepo
}

func NewQuizService(quizRepo domain.IQuizRepo) domain.IQuizSvc {
	return &quizes{
		quizRepo: quizRepo,
	}
}

func (cs *quizes) GetQuiz() ([]models.Question, error) {
	var res []models.Question
	var err error
	if res, err = cs.quizRepo.GetQuiz(); err != nil {
		logger.Error(err)
		return nil, errutil.ErrQuestionCreate
	}

	return res, nil
}
