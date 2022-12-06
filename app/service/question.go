package svc

import (
	"quiz_app/app/domain"
	"quiz_app/app/models"
	"quiz_app/app/serializers"
	"quiz_app/app/utils/errutil"
	methods "quiz_app/app/utils/methodutil"
	logger "quiz_app/infra/logger"
)

type questions struct {
	questionRepo domain.IQuestionRepo
}

func NewQuestionService(questionRepo domain.IQuestionRepo) domain.IQuestionSvc {
	return &questions{
		questionRepo: questionRepo,
	}
}

func (cs *questions) InsertQuestion(req *serializers.QuestionPayload) error {
	question := &models.Question{}
	respErr := methods.CopyStruct(req, &question)
	if respErr != nil {
		return respErr
	}
	if err := cs.questionRepo.InsertQuestion(question); err != nil {
		logger.Error(err)
		return errutil.ErrQuestionCreate
	}

	return nil
}

func (cs *questions) GetQuestion() ([]models.Question, error) {
	var res []models.Question
	var err error
	if res, err = cs.questionRepo.GetQuestion(); err != nil {
		logger.Error(err)
		return nil, errutil.ErrQuestionCreate
	}

	return res, nil
}
