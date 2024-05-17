package entities

type QuestionBase struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Answers     []Answer `json:"answers"`
	//CountA int `json:"count_answers"`
}

type Question struct {
	QuestionBase
	ID int `json:"id"`
}
