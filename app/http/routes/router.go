package routes

import (
	"quiz_app/app/http/controllers"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, bc *controllers.QuestionCnt) {
	// Question
	e.POST("/question", bc.InsertQuestion)
}
