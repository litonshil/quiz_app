package controllers

import (
	"net/http"
	"quiz_app/app/domain"
	"quiz_app/app/utils/msgutil"

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
	res, err := qc.quizSvc.GetQuiz()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, msgutil.EntityCreationFailedMsg("Quiz"))
	}

	return c.JSON(http.StatusOK, res)
}
