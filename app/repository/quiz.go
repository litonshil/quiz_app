package repository

import (
	"context"
	"log"
	"net/http"
	"quiz_app/app/domain"
	"quiz_app/app/models"
	"quiz_app/app/serializers"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	methods "quiz_app/app/utils/methodutil"

	"go.mongodb.org/mongo-driver/mongo"
)

type quizes struct {
	DB *mongo.Client
}

func NewQuizRepository(dbc *mongo.Client) domain.IQuizRepo {
	return &quizes{
		DB: dbc,
	}
}

func (cr *quizes) GetQuiz(totalQuestion int) ([]models.Question, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection := cr.DB.Database("test").Collection("question")

	quizes := []models.Question{}

	groupStage := []bson.D{bson.D{{"$sample", bson.D{{"size", totalQuestion}}}}}

	res, err := collection.Aggregate(ctx, mongo.Pipeline(groupStage))

	if err != nil {
		return nil, err
	}

	if err = res.All(ctx, &quizes); err != nil {
		log.Fatal(err)
	}
	return quizes, nil
}

func (cr *quizes) SubmitQuiz(req *models.Quiz, response *serializers.SubmitQuizResponse) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection := cr.DB.Database("test").Collection("quiz")
	newQuiz := models.Quiz{
		QuizID:  req.QuizID,
		UserID:  req.UserID,
		Answers: req.Answers,
	}
	_, err := collection.InsertOne(ctx, newQuiz)
	if err != nil {
		return err
	}

	obj := []serializers.AnswerObj{}
	respErr := methods.CopyStruct(req.Answers, &obj)
	if respErr != nil {
		return respErr
	}
	number := 0
	for _, val := range obj {
		if val.Answer == val.GivenAnswer {
			number++
		}
	}
	percent := float64(float64(number) * (float64(33) / float64(100)))
	response.IsPass = float64(number) >= percent
	response.Data = serializers.AnswerData{
		Number: float64(number),
	}
	response.Status = http.StatusOK

	return nil
}
