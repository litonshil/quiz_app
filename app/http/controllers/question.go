package controllers

import (
	"net/http"
	"quiz_app/app/domain"
	"quiz_app/app/serializers"
	"quiz_app/app/utils/msgutil"
	logger "quiz_app/infra/logger"

	"github.com/labstack/echo/v4"
)

type QuestionCnt struct {
	questionSvc domain.IQuestionSvc
}

func NewQuestionController(questionSvc domain.IQuestionSvc) *QuestionCnt {
	questionc := &QuestionCnt{
		questionSvc: questionSvc,
	}
	return questionc
}

func (bc *QuestionCnt) InsertQuestion(c echo.Context) error {
	var req serializers.QuestionPayload
	var err error

	if err = c.Bind(&req); err != nil {
		logger.Error(err)
		return c.JSON(http.StatusBadRequest, msgutil.RequestBodyParseErrorResponseMsg())
	}

	err = bc.questionSvc.InsertQuestion(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, msgutil.EntityCreationFailedMsg("Question"))
	}

	return c.NoContent(http.StatusCreated)
}

func (bc *QuestionCnt) GetQuestion(c echo.Context) error {
	res, err := bc.questionSvc.GetQuestion()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, msgutil.EntityCreationFailedMsg("Question"))
	}

	return c.JSON(http.StatusOK, res)
}
