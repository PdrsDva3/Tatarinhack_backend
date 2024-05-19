package repository

import (
	"Tatarinhack_backend/internal/entities"
	"context"
	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	Create(ctx context.Context, user entities.UserCreate) (int, error)         // id, err
	Get(ctx context.Context, id int) (*entities.User, error)                   //
	GetFriendByID(ctx context.Context, id int) (*entities.Friend, error)       //
	GetManByID(ctx context.Context, id int) (*entities.Man, error)             //
	GetHshPwddByEmail(ctx context.Context, email string) (int, string, error)  //id pwd err
	GetFriendList(ctx context.Context, id int) ([]entities.FriendsList, error) //
	GetGrammar(ctx context.Context, id int) (int, error)                       // amount, err
	GetVocabulary(ctx context.Context, id int) (int, error)                    // amount, err
	GetSpeaking(ctx context.Context, id int) (int, error)                      // amount, err
	GetLevel(ctx context.Context, id int) (int, error)                         // amount, err
	UpdatePasswordByID(ctx context.Context, id int, password string) error     // err
	//todo если надо UpdateNameByID(ctx context.Context, id int, name string) error //lol
	AddFriendByNick(ctx context.Context, nickname int, id int) (int, error) // сделано ради прикола
	AddFriendByID(ctx context.Context, friendID int, userID int) error      // err
	UpdateGrammarByID(ctx context.Context, id int, amount int) error        // err
	UpdateSpeakingByID(ctx context.Context, id int, amount int) error       // err
	UpdateVocabularyByID(ctx context.Context, id int, amount int) error     // err
	UpdateLevelByID(ctx context.Context, id int, amount int) error          //new value, err
	UpdateDaysByID(ctx context.Context, id int, amount int) error           //new value, err
	UpdateAchievementByID(ctx context.Context, id int, amount int) error    //new value, err
	Delete(ctx context.Context, id int) error                               // err
}

type TeachRepo interface {
	Create(ctx context.Context, teach entities.TeachCreate) (int, error)
	GetByID(ctx context.Context, id int) (*entities.Teach, error)
	GetPasswordByEmail(ctx context.Context, email string) (int, string, error)
	UpdatePasswordByID(ctx context.Context, id int, newPassword string) error
	Delete(ctx context.Context, id int) error
}

type AnswerRepo interface {
	Create(ctx context.Context, answer entities.AnswerBase) (int, error)
	GetByID(ctx context.Context, id int) (*entities.Answer, error)
	ChangeCorrect(ctx context.Context, id int, value bool) error
	Delete(ctx context.Context, idAns int, idQue int) error
}

type QuestionRepo interface {
	Create(ctx context.Context, question entities.QuestionBase) (int, error)
	GetByID(ctx context.Context, id int) (*entities.Question, error)
	AddAnswer(ctx context.Context, idAns int, idQue int) error
	DeleteAnswer(ctx context.Context, idAns int, idQue int) error
}

type TestRepo interface {
	Create(ctx context.Context, question entities.TestBase) (int, error)
	GetByID(ctx context.Context, id int) (*entities.Test, error)
	AddQ(ctx context.Context, idTest int, idQue int) error
	DeleteQ(ctx context.Context, idTest int, idQue int) error
	RetDB(ctx context.Context) *sqlx.DB
}

type CourseRepo interface {
	Create(ctx context.Context, question entities.CourseBase) (int, error)
	GetByID(ctx context.Context, id int) (*entities.Course, error)
	AddT(ctx context.Context, idTest int, idQue int) error
	DeleteT(ctx context.Context, idTest int, idQue int) error
}

type AudioRepo interface {
	Create(ctx context.Context, audio entities.AudioBase) (int, error)
	GetByID(ctx context.Context, id int) (*entities.Audio, error)
}

type FightRepo interface {
	SaveRes(ctx context.Context, value int) (*entities.Fight, error)
	GetByID(ctx context.Context, id int) (*entities.Test, int, error)
}
