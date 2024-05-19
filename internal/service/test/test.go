package test

import (
	"Tatarinhack_backend/internal/entities"
	"Tatarinhack_backend/internal/repository"
	"Tatarinhack_backend/internal/repository/user"
	"Tatarinhack_backend/internal/service"
	userserv "Tatarinhack_backend/internal/service/user"
	"context"
)

func InitTestService(TestRepo repository.TestRepo) service.TestService {
	return TestService{TestRepo: TestRepo}
}

type TestService struct {
	TestRepo repository.TestRepo
}

func (tst TestService) GetMe(ctx context.Context, testID int) (*entities.Test, error) {
	out, err := tst.TestRepo.GetByID(ctx, testID)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (tst TestService) Create(ctx context.Context, test entities.TestBase) (int, error) {
	id, err := tst.TestRepo.Create(ctx, test)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (tst TestService) AddTest(ctx context.Context, queID int, testID int) error {
	err := tst.TestRepo.AddQ(ctx, testID, queID)
	if err != nil {
		return err
	}
	return nil
}

type HandlerUser struct {
	service service.UserService
}

func InitUserHandler(service service.UserService) HandlerUser {
	return HandlerUser{
		service: service,
	}
}

func (tst TestService) TestAnswer(ctx context.Context, answer entities.TestAnswer) (int, error) {
	userRepo := user.InitUserRepo(tst.TestRepo.RetDB(ctx))
	userService := userserv.InitUserRepo(userRepo)
	flag := true
	i := 0
	for _, el := range answer.AnswerID {
		if !el {
			flag = false
		} else {
			i += 1
		}
	}
	if flag {
		err := userService.GrammarUp(ctx, answer.UserID)
		if err != nil {
			return 0, err
		}
		err = userService.SpeakingUp(ctx, answer.UserID)
		if err != nil {
			return 0, err
		}
		err = userService.VocabularyUp(ctx, answer.UserID)
		if err != nil {
			return 0, err
		}
		err = userService.LevelUp(ctx, answer.UserID)
		if err != nil {
			return 0, err
		}
	}
	return i, nil
}
