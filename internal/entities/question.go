package entities

type QuestionBase struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Answers     []Ans  `json:"answers"`
	//CountA int `json:"count_answers"`
}

type Question struct {
	QuestionBase
	ID int `json:"id"`
}

type QuestionAdd struct {
	IDQuestion int `json:"id_question"`
	IDAnswer   int `json:"id_answer"`
}
