package serializers

type QuizPayload struct {
	QuizID  string      `json:"quiz_id,omitempty"`
	UserID  string      `json:"user_id,omitempty"`
	Answers interface{} `json:"answers,omitempty"`
}

type SubmitQuizResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	IsPass  bool        `json:"is_pass"`
	Data    interface{} `json:"data"`
}

type AnswerObj struct {
	QuestionID  string `json:"question_id"`
	Answer      string `json:"answer"`
	GivenAnswer string `json:"given_answer"`
}

type AnswerData struct {
	Number float64 `json:"number"`
}
