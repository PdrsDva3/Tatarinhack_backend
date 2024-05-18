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
