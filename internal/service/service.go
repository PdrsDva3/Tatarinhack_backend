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
	AddQue(ctx context.Context, queID int, testID int) error
}

type CourseService interface {
	Create(ctx context.Context, course entities.CourseBase) (int, error)
	GetMe(ctx context.Context, courseID int) (*entities.Course, error)
	AddTest(ctx context.Context, courseID int, testID int) error
}
