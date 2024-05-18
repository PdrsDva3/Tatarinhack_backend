package teachserv

import (
	"Tatarinhack_backend/internal/entities"
	"Tatarinhack_backend/internal/repository"
	"Tatarinhack_backend/internal/service"
	"Tatarinhack_backend/pkg/cerr"
	"context"
	"golang.org/x/crypto/bcrypt"
)

func InitTeachService(teachRepo repository.TeachRepo) service.TeachService {
	return TeachService{TeachRepo: teachRepo}
}

type TeachService struct {
	TeachRepo repository.TeachRepo
}

func (tch TeachService) Create(ctx context.Context, teachCreate entities.TeachCreate) (int, error) {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(teachCreate.Password), 10)
	if err != nil {
		return 0, cerr.Err(cerr.Teach, cerr.Service, cerr.Hash, err).Error()
	}
	newTeach := entities.TeachCreate{
		TeachBase: teachCreate.TeachBase,
		Password:  string(hashed_password),
	}

	id, err := tch.TeachRepo.Create(ctx, newTeach)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (tch TeachService) Login(ctx context.Context, teachLogin entities.TeachLogin) (int, error) {
	id, pwd, err := tch.TeachRepo.GetPasswordByEmail(ctx, teachLogin.Email)
	if err != nil {
		return 0, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(pwd), []byte(teachLogin.Password))
	if err != nil {
		return 0, cerr.Err(cerr.Teach, cerr.Service, cerr.InvalidPWD, err).Error()
	}
	return id, nil
}

func (tch TeachService) ChangePassword(ctx context.Context, teachID int, newPWD string) error {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(newPWD), 10)
	if err != nil {
		return cerr.Err(cerr.Teach, cerr.Service, cerr.Hash, err).Error()
	}
	err = tch.TeachRepo.UpdatePasswordByID(ctx, teachID, string(hashed_password))
	if err != nil {
		return err
	}
	return nil
}

func (tch TeachService) GetMe(ctx context.Context, teachID int) (*entities.Teach, error) {
	teach, err := tch.TeachRepo.GetByID(ctx, teachID)
	if err != nil {
		return nil, err
	}
	if teach == nil {
		return nil, cerr.Err(cerr.Teach, cerr.Service, cerr.NotFound, nil).Error()
	}
	return teach, nil
}

func (tch TeachService) Delete(ctx context.Context, teachID int) error {
	err := tch.TeachRepo.Delete(ctx, teachID)
	if err != nil {
		return err
	}
	return nil
}
