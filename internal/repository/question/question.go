package question

import (
	"Tatarinhack_backend/internal/entities"
	"Tatarinhack_backend/internal/repository"
	"Tatarinhack_backend/pkg/cerr"
	"context"
	"github.com/jmoiron/sqlx"
)

type RepositoryQuestion struct {
	db *sqlx.DB
}

func InitQuestionRepository(db *sqlx.DB) repository.QuestionRepo {
	return RepositoryQuestion{
		db}
}

func (que RepositoryQuestion) Create(ctx context.Context, question entities.QuestionBase) (int, error) {
	var id int
	transaction, err := que.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, cerr.Err(cerr.Question, cerr.Repository, cerr.Transaction, err).Error()
	}
	row := transaction.QueryRowContext(ctx, `INSERT INTO questions (name, description) VALUES ($1, $2) returning id;`,
		question.Name, question.Description)
	err = row.Scan(&id)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Question, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Question, cerr.Repository, cerr.Scan, err).Error()
	}
	if err := transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Question, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Question, cerr.Repository, cerr.Commit, err).Error()
	}
	return id, nil
}

func (que RepositoryQuestion) GetByID(ctx context.Context, id int) (*entities.Question, error) {
	var question entities.Question
	rows := que.db.QueryRowContext(ctx, `SELECT name, description from questions WHERE id = $1;`, id)
	err := rows.Scan(&question.Name, &question.Description)
	if err != nil {
		return nil, cerr.Err(cerr.Question, cerr.Repository, cerr.Scan, err).Error()
	}
	question.ID = id
	return &question, nil
}

func (que RepositoryQuestion) AddAnswer(ctx context.Context, id_ans int, id_que int) error {
	transaction, err := que.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Question, cerr.Repository, cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `INSERT INTO answers_questions (id_answer, id_question) VALUES ($1, $2);`, id_ans, id_que)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Question, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Question, cerr.Repository, cerr.ExecCon, err).Error()
	}

	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Question, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Question, cerr.Repository, cerr.Rows, err).Error()
	}

	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Question, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Question, cerr.Repository, cerr.NoOneRow, err).Error()
	}

	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Question, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Question, cerr.Repository, cerr.Commit, err).Error()
	}
	return nil
}

func (que RepositoryQuestion) DeleteAnswer(ctx context.Context, id_ans int, id_que int) error {
	transaction, err := que.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Question, cerr.Repository, cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `DELETE FROM answers_questions where id_answer=$1 and  id_question = $2;`, id_ans, id_que)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Question, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Question, cerr.Repository, cerr.ExecCon, err).Error()
	}

	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Question, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Question, cerr.Repository, cerr.Rows, err).Error()
	}

	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Question, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Question, cerr.Repository, cerr.NoOneRow, err).Error()
	}

	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Question, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Question, cerr.Repository, cerr.Commit, err).Error()
	}
	return nil
}
