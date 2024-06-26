package entities

type TeachBase struct {
	Nick  string `json:"nick"`
	Email string `json:"email"`
}

type TeachCreate struct {
	TeachBase
	Password string `json:"password"`
}

type TeachLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TeachChangePassword struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
}

type TeachChangeName struct {
	ID   int    `json:"id"`
	Nick string `json:"nick"`
}

type Teach struct {
	TeachBase
	ID int `json:"id"`
}
