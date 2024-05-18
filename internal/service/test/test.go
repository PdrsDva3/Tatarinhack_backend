package test

import (
	"Tatarinhack_backend/internal/entities"
	"Tatarinhack_backend/internal/repository"
	"Tatarinhack_backend/internal/service"
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

func (tst TestService) AddQue(ctx context.Context, queID int, testID int) error {
	err := tst.TestRepo.AddQ(ctx, testID, queID)
	if err != nil {
		return err
	}
	return nil
}
