package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Question struct {
	ID          primitive.ObjectID `json:"id,omitempty"`
	Title       string             `json:"title,omitempty"`
	Description *string            `json:"description,omitempty"`
	Type        string             `json:"type,omitempty"`
	Option      interface{}        `json:"option,omitempty"`
	Answer      string             `json:"answer,omitempty"`
	Point       int                `json:"point,omitempty"`
	Image       *string            `json:"image,omitempty"`
	Author      int                `json:"author,omitempty"`
}
