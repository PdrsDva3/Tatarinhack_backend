package service

import (
	"Tatarinhack_backend/internal/entities"
	"context"
)

type TeachService interface {
	Create(ctx context.Context, teachCreate entities.TeachCreate) (int, error)
	Login(ctx context.Context, teachLogin entities.TeachLogin) (int, error)
	ChangePassword(ctx context.Context, teachID int, newPWD string) error
	GetMe(ctx context.Context, teachID int) (*entities.Teach, error)
	Delete(ctx context.Context, teachID int) error
}
