package entities

type UserBase struct {
	Nick  string `json:"nick"`
	Email string `json:"email"`
	Goal  string `json:"goal"`
	Sex   string `json:"sex"`
}

type UserCreate struct {
	UserBase
	Password string `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email"`
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

type Friend struct {
	ID          int           `json:"id"`
	Nick        string        `json:"nick"`
	Achievement int           `json:"achievement"`
	FriendsList []FriendsList `json:"friendsList"`
	Level       int           `json:"level"`
	Grammar     int           `json:"grammar"`
	Vocabulary  int           `json:"vocabulary"`
	Speaking    int           `json:"speaking"`
	Sex         string        `json:"sex"`
}

type FriendsList struct {
	ID   int    `json:"id"`
	Nick string `json:"nick"`
	Sex  string `json:"sex"`
}

type Man struct {
	ID          int    `json:"id"`
	Nick        string `json:"nick"`
	Level       int    `json:"level"`
	Achievement int    `json:"achievement"`
}

type UserAddFriend struct {
	UserID   int `json:"user_id"`
	FriendID int `json:"friend_id"`
}

type User struct {
	UserBase
	ID          int           `json:"id"`
	Rating      int           `json:"rating"`
	Grammar     int           `json:"grammar"`
	Vocabulary  int           `json:"vocabulary"`
	Speaking    int           `json:"speaking"`
	Level       int           `json:"level"`
	Days        int           `json:"days"`
	FriendsList []FriendsList `json:"friendsList"`
	Achievement int           `json:"achievement"`
}

type UserUpGrammar struct {
	ID      int `json:"id"`
	Grammar int `json:"grammar"`
}

type UserUpVocabulary struct {
	ID         int `json:"id"`
	Vocabulary int `json:"vocabulary"`
}

type UserUpSpeaking struct {
	ID       int `json:"id"`
	Speaking int `json:"speaking"`
}

type UserUpLevel struct {
	ID       int `json:"id"`
	Speaking int `json:"speaking"`
}
