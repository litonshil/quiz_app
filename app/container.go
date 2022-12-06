package container

import (
	"quiz_app/app/http/controllers"
	"quiz_app/app/http/routes"
	repo "quiz_app/app/repository"
	svc "quiz_app/app/service"
	"quiz_app/infra/conn"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {

	db := conn.Db()

	questionRepo := repo.NewQuestionRepository(db)
	quizRepo := repo.NewQuizRepository(db)

	questionSvc := svc.NewQuestionService(questionRepo)
	quizSvc := svc.NewQuizService(quizRepo)

	questionCr := controllers.NewQuestionController(questionSvc)
	quizCr := controllers.NewQuizController(quizSvc)

	routes.Init(e, questionCr, quizCr)

}
