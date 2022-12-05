package repository

import (
	"context"
	"log"
	"quiz_app/app/domain"
	"quiz_app/app/models"
	"strings"
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
func MongoPipeline(str string) mongo.Pipeline {
	var pipeline = []bson.D{}
	str = strings.TrimSpace(str)
	if strings.Index(str, "[") != 0 {
		var doc bson.D
		bson.UnmarshalExtJSON([]byte(str), false, &doc)
		pipeline = append(pipeline, doc)
	} else {
		bson.UnmarshalExtJSON([]byte(str), false, &pipeline)
	}
	return pipeline
}

func (cr *questions) GetQuiz() ([]models.Question, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection := cr.DB.Database("test").Collection("question")

	questions := []models.Question{}

	groupStage := []bson.D{bson.D{{"$sample", bson.D{{"size", 2}}}}}

	res, err := collection.Aggregate(ctx, mongo.Pipeline(groupStage))

	if err != nil {
		return nil, err
	}

	if err = res.All(ctx, &questions); err != nil {
		log.Fatal(err)
	}
	return questions, nil
}
