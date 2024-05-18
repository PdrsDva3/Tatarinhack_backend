package user

import (
	"Tatarinhack_backend/internal/entities"
	"Tatarinhack_backend/internal/repository"
	"Tatarinhack_backend/internal/service"
	"Tatarinhack_backend/pkg/cerr"
	"context"
	"golang.org/x/crypto/bcrypt"
)

type RepositoryUser struct {
	UserRepo repository.UserRepo
}

func InitUserRepo(userRepo repository.UserRepo) service.UserService {
	return RepositoryUser{
		UserRepo: userRepo}
}

func (usr RepositoryUser) Create(ctx context.Context, user entities.UserCreate) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return 0, cerr.Err(cerr.User, cerr.Service, cerr.Hash, err).Error()
	}
	newUser := entities.UserCreate{
		UserBase: user.UserBase,
		Password: string(hashedPassword),
	}

	id, err := usr.UserRepo.Create(ctx, newUser)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (usr RepositoryUser) Login(ctx context.Context, user entities.UserLogin) (*entities.User, error) {
	id, hashedPassword, err := usr.UserRepo.GetHshPwddByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil {
		return nil, cerr.Err(cerr.User, cerr.Service, cerr.InvalidPWD, err).Error()
	}
	userr, err := usr.UserRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return userr, nil
}

func (usr RepositoryUser) AddFriend(ctx context.Context, userID int, friendID int) error {
	err := usr.UserRepo.AddFriendByID(ctx, friendID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (usr RepositoryUser) GetUser(ctx context.Context, id int) (*entities.User, error) {
	var user *entities.User
	user, err := usr.UserRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (usr RepositoryUser) GetFriend(ctx context.Context, id int) (*entities.Friend, error) {
	var friend *entities.Friend
	friend, err := usr.UserRepo.GetFriendByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return friend, nil
}

func (usr RepositoryUser) GetMan(ctx context.Context, id int) (*entities.Man, error) {
	var man *entities.Man
	man, err := usr.UserRepo.GetManByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return man, nil
}

func (usr RepositoryUser) UpdatePassword(ctx context.Context, userID int, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return cerr.Err(cerr.User, cerr.Service, cerr.Hash, err).Error()
	}
	err = usr.UserRepo.UpdatePasswordByID(ctx, userID, string(hashedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (usr RepositoryUser) GrammarUp(ctx context.Context, id int) error {
	amount, err := usr.UserRepo.GetGrammar(ctx, id)
	if err != nil {
		return err
	}
	if amount == 10 {
		return nil
	}
	amount += 1
	err = usr.UserRepo.UpdateGrammarByID(ctx, id, amount)
	if err != nil {
		return err
	}

	return nil
}

func (usr RepositoryUser) VocabularyUp(ctx context.Context, id int) error {
	amount, err := usr.UserRepo.GetVocabulary(ctx, id)
	if err != nil {
		return err
	}
	if amount == 10 {
		return nil
	}
	amount += 1
	err = usr.UserRepo.UpdateVocabularyByID(ctx, id, amount)
	if err != nil {
		return err
	}
	return nil
}

func (usr RepositoryUser) SpeakingUp(ctx context.Context, id int) error {
	amount, err := usr.UserRepo.GetVocabulary(ctx, id)
	if err != nil {
		return err
	}
	if amount == 10 {
		return nil
	}

	amount += 1
	err = usr.UserRepo.UpdateSpeakingByID(ctx, id, amount)
	if err != nil {
		return err
	}

	return nil
}

func (usr RepositoryUser) DaysUp(ctx context.Context, id int, amount int) error {
	//TODO implement me
	panic("implement me")
}

func (usr RepositoryUser) LevelUp(ctx context.Context, id int) error {
	amount, err := usr.UserRepo.GetLevel(ctx, id)
	if err != nil {
		return err
	}

	grAmount, err := usr.UserRepo.GetGrammar(ctx, id)
	if err != nil {
		return err
	}
	spAmount, err := usr.UserRepo.GetSpeaking(ctx, id)
	if err != nil {
		return err
	}
	vcAmount, err := usr.UserRepo.GetVocabulary(ctx, id)
	if err != nil {
		return err
	}
	sum := grAmount + spAmount + vcAmount
	if sum == 30 {
		amount += 1
		err := usr.UserRepo.UpdateLevelByID(ctx, id, amount)
		if err != nil {
			return err
		}
	}
	return nil
}

func (usr RepositoryUser) AchievementUp(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")

	//err := usr.UserRepo.UpdateAchievementByID(ctx, id)
}

func (usr RepositoryUser) GetFriendsList(ctx context.Context, id int) ([]entities.FriendsList, error) {
	friendsList, err := usr.UserRepo.GetFriendList(ctx, id)
	if err != nil {
		return nil, err
	}
	return friendsList, nil
}

func (usr RepositoryUser) Delete(ctx context.Context, id int) error {
	if err := usr.UserRepo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
