package answer

import (
	"Tatarinhack_backend/internal/entities"
	"Tatarinhack_backend/internal/repository"
	"Tatarinhack_backend/pkg/cerr"
	"context"
	"github.com/jmoiron/sqlx"
)

type RepositoryAnswer struct {
	db *sqlx.DB
}

func InitAnswerRepository(db *sqlx.DB) repository.AnswerRepo {
	return RepositoryAnswer{
		db}
}

func (ans RepositoryAnswer) Create(ctx context.Context, answer entities.AnswerBase) (int, error) {
	var id int
	transaction, err := ans.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, cerr.Err(cerr.Answer, cerr.Repository, cerr.Transaction, err).Error()
	}
	row := transaction.QueryRowContext(ctx, `INSERT INTO Answers (name, is_correct) VALUES ($1, $2) returning id;`,
		answer.Name, answer.IsCorrect)
	err = row.Scan(&id)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Answer, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Answer, cerr.Repository, cerr.Scan, err).Error()
	}
	if err := transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Answer, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Answer, cerr.Repository, cerr.Commit, err).Error()
	}
	return id, nil
}

func (ans RepositoryAnswer) GetByID(ctx context.Context, id int) (*entities.Answer, error) {
	var answer entities.Answer
	rows := ans.db.QueryRowContext(ctx, `SELECT name, is_correct from answers WHERE id = $1;`, id)
	err := rows.Scan(&answer.Name, &answer.IsCorrect)
	if err != nil {
		return nil, cerr.Err(cerr.Answer, cerr.Repository, cerr.Scan, err).Error()
	}
	answer.ID = id
	return &answer, nil
}

func (ans RepositoryAnswer) ChangeCorrect(ctx context.Context, id int, value bool) error {
	var correct bool
	rows := ans.db.QueryRowContext(ctx, `SELECT is_correct from answers WHERE id = $1;`, id)
	err := rows.Scan(&correct)
	if err != nil {
		return cerr.Err(cerr.Answer, cerr.Repository, cerr.Scan, err).Error()
	}
	transaction, err := ans.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Answer, cerr.Repository, cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `UPDATE answers SET is_correct=$2 WHERE id=$1;`, id, value)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Answer, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Answer, cerr.Repository, cerr.ExecCon, err).Error()
	}

	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Answer, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Answer, cerr.Repository, cerr.Rows, err).Error()
	}

	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Answer, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Answer, cerr.Repository, cerr.NoOneRow, err).Error()
	}

	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Answer, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Answer, cerr.Repository, cerr.Commit, err).Error()
	}
	return nil
}

func (ans RepositoryAnswer) Delete(ctx context.Context, idAns int, idQue int) error {
	transaction, err := ans.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Answer, cerr.Repository, cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `DELETE FROM answers_questions WHERE id_answer=$1 and id_question=$2;`, idAns, idQue)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Answer, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Answer, cerr.Repository, cerr.ExecCon, err).Error()
	}
	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Answer, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Answer, cerr.Repository, cerr.Rows, err).Error()
	}
	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Answer, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Answer, cerr.Repository, cerr.NoOneRow, err).Error()
	}
	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Answer, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Answer, cerr.Repository, cerr.Commit, err).Error()
	}

	return nil
}
