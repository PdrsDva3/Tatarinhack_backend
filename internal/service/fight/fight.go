package fightserv

import (
	"Tatarinhack_backend/internal/entities"
	"Tatarinhack_backend/internal/repository"
	"Tatarinhack_backend/internal/service"
	"context"
)

func InitFightService(TestRepo repository.FightRepo) service.FightService {
	return FightService{TestRepo: TestRepo}
}

type FightService struct {
	TestRepo repository.FightRepo
}

func (f FightService) Get(ctx context.Context, id int) (*entities.Test, int, error) {
	out, pp, err := f.TestRepo.GetByID(ctx, id)
	if err != nil {
		return nil, 0, err
	}
	return out, pp, nil
}

func (f FightService) Post(ctx context.Context, answer *entities.FightStart) error {
	i := 0
	for _, el := range answer.AnswerID {
		if el {
			i += 1
		}
	}
	_, err := f.TestRepo.SaveRes(ctx, i)
	if err != nil {
		return err
	}
	return nil
}
