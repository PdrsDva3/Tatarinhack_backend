package answerserv

import (
	"Tatarinhack_backend/internal/entities"
	"Tatarinhack_backend/internal/repository"
	"Tatarinhack_backend/internal/service"
	"context"
)

func InitAnswerService(answerRepo repository.AnswerRepo) service.AnswerService {
	return AnswerService{AnswerRepo: answerRepo}
}

type AnswerService struct {
	AnswerRepo repository.AnswerRepo
}

func (ans AnswerService) GetMe(ctx context.Context, ansID int) (*entities.Answer, error) {
	out, err := ans.AnswerRepo.GetByID(ctx, ansID)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (ans AnswerService) Create(ctx context.Context, answerCreate entities.AnswerBase) (int, error) {
	id, err := ans.AnswerRepo.Create(ctx, answerCreate)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (ans AnswerService) Change(ctx context.Context, ansID int, value bool) error {
	err := ans.AnswerRepo.ChangeCorrect(ctx, ansID, value)
	if err != nil {
		return err
	}
	return nil
}
