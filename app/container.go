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

	questionSvc := svc.NewQuestionService(questionRepo)

	questionCr := controllers.NewQuestionController(questionSvc)

	routes.Init(e, questionCr)

}
