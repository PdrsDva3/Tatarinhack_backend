package test

import (
	"Tatarinhack_backend/internal/entities"
	"Tatarinhack_backend/internal/repository"
	"Tatarinhack_backend/pkg/cerr"
	"context"
	"github.com/jmoiron/sqlx"
)

type RepositoryTest struct {
	db *sqlx.DB
}

func InitTestRepository(db *sqlx.DB) repository.TestRepo {
	return RepositoryTest{
		db}
}

func (que RepositoryTest) Create(ctx context.Context, test entities.TestBase) (int, error) {
	var id int
	transaction, err := que.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, cerr.Err(cerr.Test, cerr.Repository, cerr.Transaction, err).Error()
	}
	row := transaction.QueryRowContext(ctx, `INSERT INTO tests (name, type, level, speed) VALUES ($1, $2, $3, $4) returning id;`,
		test.Name, test.Type, test.Level, test.Speed)
	err = row.Scan(&id)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Test, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Test, cerr.Repository, cerr.Scan, err).Error()
	}
	if err := transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Test, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Test, cerr.Repository, cerr.Commit, err).Error()
	}
	return id, nil
}

func (que RepositoryTest) GetByID(ctx context.Context, id int) (*entities.Test, error) {
	var test entities.Test
	rows := que.db.QueryRowContext(ctx, `SELECT (name, type, level, speed) from tests WHERE id = $1;`, id)
	err := rows.Scan(&test.Name, &test.Type, &test.Level, &test.Speed)
	if err != nil {
		return nil, cerr.Err(cerr.Test, cerr.Repository, cerr.Scan, err).Error()
	}
	test.ID = id
	var ids []int
	rowss := que.db.QueryRowxContext(ctx, `SELECT id_questions FROM questions_tests WHERE id_tests = $1;`, id)
	err = rowss.Scan(&ids)
	for _, i := range ids {
		var question entities.Question
		rows := que.db.QueryRowContext(ctx, `SELECT name, description from questions WHERE id = $1;`, i)
		err := rows.Scan(&question.Name, &question.Description)
		if err != nil {
			return nil, cerr.Err(cerr.Question, cerr.Repository, cerr.Scan, err).Error()
		}
		question.ID = id
		test.Questions = append(test.Questions, question)
	}
	if err != nil {
		return nil, cerr.Err(cerr.User, cerr.Repository, cerr.Scan, err).Error()
	}

	return &test, nil
}

func (que RepositoryTest) AddQ(ctx context.Context, idTest int, idQue int) error {
	transaction, err := que.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Test, cerr.Repository, cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `INSERT INTO questions_tests (id_questions, id_tests) VALUES ($1, $2);`, idQue, idTest)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Test, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Test, cerr.Repository, cerr.ExecCon, err).Error()
	}

	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Test, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Test, cerr.Repository, cerr.Rows, err).Error()
	}

	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Test, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Test, cerr.Repository, cerr.NoOneRow, err).Error()
	}

	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Test, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Test, cerr.Repository, cerr.Commit, err).Error()
	}
	return nil
}

func (que RepositoryTest) DeleteQ(ctx context.Context, idTest int, idQue int) error {
	transaction, err := que.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Test, cerr.Repository, cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `DELETE FROM questions_tests where id_questions=$1 and  id_tests = $2;`, idTest, idQue)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Test, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Test, cerr.Repository, cerr.ExecCon, err).Error()
	}

	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Test, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Test, cerr.Repository, cerr.Rows, err).Error()
	}

	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Test, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Test, cerr.Repository, cerr.NoOneRow, err).Error()
	}

	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Test, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Test, cerr.Repository, cerr.Commit, err).Error()
	}
	return nil
}
