package entities

type AnswerBase struct {
	Name      string `json:"name"`
	IsCorrect bool   `json:"isCorrect"`
}

type Answer struct {
	AnswerBase
	ID int `json:"id"`
}

type AnswerChange struct {
	ID        int  `json:"id"`
	IsCorrect bool `json:"value"`
}

type Ans struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	IsCorrect bool   `json:"isCorrect"`
}
