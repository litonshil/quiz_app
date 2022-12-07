package svc

import (
	"quiz_app/app/domain"
	"quiz_app/app/models"
	"quiz_app/app/serializers"
	"quiz_app/app/utils/errutil"
	methods "quiz_app/app/utils/methodutil"
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

func (cs *quizes) GetQuiz(totalQuestion int) ([]models.Question, error) {
	var res []models.Question
	var err error
	if res, err = cs.quizRepo.GetQuiz(totalQuestion); err != nil {
		logger.Error(err)
		return nil, errutil.ErrQuestionCreate
	}

	return res, nil
}

func (cs *quizes) SubmitQuiz(req *serializers.QuizPayload, response *serializers.SubmitQuizResponse) error {
	quiz := &models.Quiz{}
	respErr := methods.CopyStruct(req, &quiz)
	if respErr != nil {
		return respErr
	}
	if err := cs.quizRepo.SubmitQuiz(quiz, response); err != nil {
		logger.Error(err)
		return errutil.ErrQuestionCreate
	}
	return nil
}
