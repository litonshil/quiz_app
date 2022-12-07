package models

type Quiz struct {
	QuizID  string      `json:"quiz_id,omitempty"`
	UserID  string      `json:"user_id,omitempty"`
	Answers interface{} `json:"answers,omitempty"`
}
