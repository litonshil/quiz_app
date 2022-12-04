package repository

import (
	"context"
	"quiz_app/app/domain"
	"quiz_app/app/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

type questions struct {
	DB *mongo.Client
}

func NewQuestionRepository(dbc *mongo.Client) domain.IQuestionRepo {
	return &questions{
		DB: dbc,
	}
}

func (cr *questions) InsertQuestion(req *models.Question) error {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection := cr.DB.Database("test").Collection("question")
	newQuestion := models.Question{
		ID:          primitive.NewObjectID(),
		Title:       req.Title,
		Author:      req.Author,
		Description: req.Description,
		Type:        req.Type,
		Option:      req.Option,
		Answer:      req.Answer,
		Point:       req.Point,
	}
	_, err := collection.InsertOne(ctx, newQuestion)
	if err != nil {
		return err
	}
	return nil
}

func (cr *questions) GetQuestion() ([]models.Question, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection := cr.DB.Database("test").Collection("question")

	questions := []models.Question{}

	results, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleQuestion models.Question
		if err = results.Decode(&singleQuestion); err != nil {
			return nil, err
		}

		questions = append(questions, singleQuestion)
	}
	return questions, nil
}

func (cr *questions) GetQuiz() ([]models.Question, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection := cr.DB.Database("test").Collection("question")

	questions := []models.Question{}

	// collection.Aggregate()

	// if err != nil {
	// 	return nil, err
	// }

	// //reading from the db in an optimal way
	// defer results.Close(ctx)
	// for results.Next(ctx) {
	// 	var singleQuestion models.Question
	// 	if err = results.Decode(&singleQuestion); err != nil {
	// 		return nil, err
	// 	}

	// 	questions = append(questions, singleQuestion)
	// }
	return questions, nil
}
