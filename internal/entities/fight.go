package entities

type Fight struct {
	Session int `json:"session"`
	Test    int `json:"test"`
	ID_1    int `json:"id_1"`
	ID_2    int `json:"id_2"`
	Res_1   int `json:"res_1"`
	Res_2   int `json:"res_2"`
}

type FightStart struct {
	AnswerID []bool `json:"answer_id"`
}
