package entities

type AudioBase struct {
	Text string `json:"text"`
}

type Audio struct {
	AudioBase
	ID int `json:"id"`
}
