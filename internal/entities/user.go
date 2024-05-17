package entities

type UserBase struct {
	Nick  int64  `json:"nick"`
	Email string `json:"email"`
	Goal  string `json:"goal"`
	Sex   string `json:"sex"`
}

type UserCreate struct {
	UserBase
	Password string `json:"password"`
}

type UserLogin struct {
	Email    int64  `json:"email"`
	Password string `json:"password"`
}

type UserChangePassword struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
}

type UserChangeName struct {
	ID   int    `json:"id"`
	Nick string `json:"nick"`
}

type User struct {
	UserBase
	ID          int `json:"id"`
	Rating      int `json:"rating"`
	Grammar     int `json:"grammar"`
	Vocabulary  int `json:"vocabulary"`
	Speaking    int `json:"speaking"`
	Level       int `json:"level"`
	Days        int `json:"days"`
	Achievement int `json:"achievement"`
}
