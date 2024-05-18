package service

import (
	"Tatarinhack_backend/internal/entities"
	"context"
)


type UserService interface {
	Create(ctx context.Context, user entities.UserCreate) (int, error)          //done
	Login(ctx context.Context, user entities.UserLogin) (*entities.User, error) //done
	GetUser(ctx context.Context, id int) (*entities.User, error)                //done
	GetFriend(ctx context.Context, id int) (*entities.Friend, error)            //done
	GetMan(ctx context.Context, id int) (*entities.Man, error)                  //done
	AddFriend(ctx context.Context, userID int, friendID int) error              //done
	UpdatePassword(ctx context.Context, userID int, password string) error      // done
	GrammarUp(ctx context.Context, id int) error                                // done
	VocabularyUp(ctx context.Context, id int) error                             // done
	SpeakingUp(ctx context.Context, id int) error                               // done
	Delete(ctx context.Context, id int) error                                   //done
	DaysUp(ctx context.Context, id int, amount int) error
	LevelUp(ctx context.Context, id int) error
	AchievementUp(ctx context.Context, id int) error
	GetFriendsList(ctx context.Context, id int) ([]entities.FriendsList, error)
}

type TeachService interface {
	Create(ctx context.Context, teachCreate entities.TeachCreate) (int, error)
	Login(ctx context.Context, teachLogin entities.TeachLogin) (int, error)
	ChangePassword(ctx context.Context, teachID int, newPWD string) error
	GetMe(ctx context.Context, teachID int) (*entities.Teach, error)
	Delete(ctx context.Context, teachID int) error
}

type AnswerService interface {
	Create(ctx context.Context, answerCreate entities.AnswerBase) (int, error)
	GetMe(ctx context.Context, ansID int) (*entities.Answer, error)
	Change(ctx context.Context, ansID int, value bool) error
}

type QuestionService interface {
	Create(ctx context.Context, questCreate entities.QuestionBase) (int, error)
	GetMe(ctx context.Context, queID int) (*entities.Question, error)
	AddAnswer(ctx context.Context, queID int, ansID int) error
}

type TestService interface {
	Create(ctx context.Context, test entities.TestBase) (int, error)
	GetMe(ctx context.Context, testID int) (*entities.Test, error)
	AddTest(ctx context.Context, queID int, testID int) error
}
