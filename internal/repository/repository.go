package repository

import (
	"Tatarinhack_backend/internal/entities"
	"context"
)

type TeachRepo interface {
	Create(ctx context.Context, teach entities.TeachCreate) (int, error)
	GetByID(ctx context.Context, id int) (*entities.Teach, error)
	GetPasswordByEmail(ctx context.Context, email string) (int, string, error)
	UpdatePasswordByID(ctx context.Context, id int, newPassword string) error
	Delete(ctx context.Context, id int) error
}

type AnswerRepo interface {
	Create(ctx context.Context, answer entities.AnswerBase) (int, error)
	GetByID(ctx context.Context, id int) (*entities.Answer, error)
	ChangeCorrect(ctx context.Context, id int, value bool) error
	Delete(ctx context.Context, idAns int, idQue int) error
}

type QuestionRepo interface {
	Create(ctx context.Context, question entities.QuestionBase) (int, error)
	GetByID(ctx context.Context, id int) (*entities.Question, error)
	AddAnswer(ctx context.Context, idAns int, idQue int) error
	DeleteAnswer(ctx context.Context, idAns int, idQue int) error
}

type TestRepo interface {
	Create(ctx context.Context, question entities.TestBase) (int, error)
	GetByID(ctx context.Context, id int) (*entities.Test, error)
	AddQ(ctx context.Context, idTest int, idQue int) error
	DeleteQ(ctx context.Context, idTest int, idQue int) error
}

type CourseRepo interface {
	Create(ctx context.Context, question entities.CourseBase) (int, error)
	GetByID(ctx context.Context, id int) (*entities.Course, error)
	AddT(ctx context.Context, idTest int, idQue int) error
	DeleteT(ctx context.Context, idTest int, idQue int) error
}
