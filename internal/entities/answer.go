package entities

type AnswerBase struct {
	Name      string `json:"name"`
	IsCorrect string `json:"isCorrect"`
}

type Answer struct {
	AnswerBase
	ID int `json:"id"`
}
