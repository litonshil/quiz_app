package serializers

import "go.mongodb.org/mongo-driver/bson/primitive"

type QuestionPayload struct {
	ID          primitive.ObjectID `json:"id,omitempty"`
	Title       *string            `json:"title,omitempty"`
	Description string             `json:"description,omitempty"`
	Type        string             `json:"type,omitempty"`
	Option      *interface{}       `json:"option,omitempty"`
	Answer      string             `json:"answer,omitempty"`
	Point       int                `json:"point,omitempty"`
	Image       *string            `json:"image,omitempty"`
	Author      int                `json:"author,omitempty"`
}

type InsertQuestionRespose struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
