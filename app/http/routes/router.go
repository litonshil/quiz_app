package routes

import (
	"quiz_app/app/http/controllers"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, bc *controllers.QuestionCnt, quizCr *controllers.QuizCnt) {
	// Question
	e.POST("/question", bc.InsertQuestion)
	e.GET("/question", bc.GetQuestion)

	//Quiz
	e.GET("/quiz", quizCr.GetQuiz)
	e.POST("/quiz", quizCr.SubmitQuiz)
}
