package repository

import (
	"context"
	"log"
	"quiz_app/app/domain"
	"quiz_app/app/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"

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

func (cr *quizes) GetQuiz() ([]models.Question, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection := cr.DB.Database("test").Collection("question")

	quizes := []models.Question{}

	groupStage := []bson.D{bson.D{{"$sample", bson.D{{"size", 2}}}}}

	res, err := collection.Aggregate(ctx, mongo.Pipeline(groupStage))

	if err != nil {
		return nil, err
	}

	if err = res.All(ctx, &quizes); err != nil {
		log.Fatal(err)
	}
	return quizes, nil
}
