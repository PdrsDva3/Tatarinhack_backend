package questionserv

import (
	"Tatarinhack_backend/internal/entities"
	"Tatarinhack_backend/internal/repository"
	"Tatarinhack_backend/internal/service"
	"context"
)

func InitQuestService(QuestRepo repository.QuestionRepo) service.QuestionService {
	return QuestService{QuestRepo: QuestRepo}
}

type QuestService struct {
	QuestRepo repository.QuestionRepo
}

func (ans QuestService) GetMe(ctx context.Context, ansID int) (*entities.Question, error) {
	out, err := ans.QuestRepo.GetByID(ctx, ansID)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (ans QuestService) Create(ctx context.Context, QuestCreate entities.QuestionBase) (int, error) {
	id, err := ans.QuestRepo.Create(ctx, QuestCreate)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (ans QuestService) AddAnswer(ctx context.Context, queID int, ansID int) error {
	err := ans.QuestRepo.AddAnswer(ctx, ansID, queID)
	if err != nil {
		return err
	}
	return nil
}
