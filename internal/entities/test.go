package entities

type TestBase struct {
	Name      string     `json:"name"`
	Type      string     `json:"type"`
	Level     int        `json:"level"`
	Speed     string     `json:"speed"`
	Questions []Question `json:"questions" `
	//CountQ    int    `json:"count_questions"`
}

type Test struct {
	TestBase
	ID int `json:"id"`
}

type TestGet struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
