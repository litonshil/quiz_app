package controllers

import (
	"net/http"
	"quiz_app/app/domain"
	"quiz_app/app/serializers"
	"quiz_app/app/utils/msgutil"
	logger "quiz_app/infra/logger"
	"strconv"

	"github.com/labstack/echo/v4"
)

type QuizCnt struct {
	quizSvc domain.IQuizSvc
}

func NewQuizController(quizSvc domain.IQuizSvc) *QuizCnt {
	quizc := &QuizCnt{
		quizSvc: quizSvc,
	}
	return quizc
}

func (qc *QuizCnt) GetQuiz(c echo.Context) error {
	total := c.QueryParam("total_question")
	totalQues, err := strconv.Atoi(total)
	if err != nil {
		return err
	}
	res, err := qc.quizSvc.GetQuiz(totalQues)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, msgutil.EntityCreationFailedMsg("Quiz"))
	}

	return c.JSON(http.StatusOK, res)
}

func (bc *QuizCnt) SubmitQuiz(c echo.Context) error {
	var req serializers.QuizPayload
	var err error

	if err = c.Bind(&req); err != nil {
		logger.Error(err)
		return c.JSON(http.StatusBadRequest, msgutil.RequestBodyParseErrorResponseMsg())
	}
	response := &serializers.SubmitQuizResponse{}
	err = bc.quizSvc.SubmitQuiz(&req, response)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, msgutil.EntityCreationFailedMsg("Question"))
	}

	// return c.NoContent(http.StatusCreated)
	return c.JSON(http.StatusOK, response)
}
