package entities

type CourseBase struct {
	Name        string    `json:"name"`
	Level       int       `json:"level"`
	Description string    `json:"description"`
	Tests       []TestGet `json:"Tests"`
	//CountQ    int    `json:"count_questions"`
}

type Course struct {
	CourseBase
	ID int `json:"id"`
}
