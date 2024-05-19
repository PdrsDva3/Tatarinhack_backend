package fight

import (
	"Tatarinhack_backend/internal/entities"
	"Tatarinhack_backend/internal/repository"
	"Tatarinhack_backend/pkg/cerr"
	"context"
	"github.com/jmoiron/sqlx"
)

type RepositoryFight struct {
	db *sqlx.DB
}

func (que RepositoryFight) SaveRes(ctx context.Context, value int) (*entities.Fight, error) {
	transaction, err := que.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, cerr.Err(cerr.Test, cerr.Repository, cerr.Transaction, err).Error()
	}
	_, err = transaction.ExecContext(ctx, `UPDATE fight set res_2=$1;`, value)
	transaction.Commit()
	if err != nil {
		return nil, cerr.Err(cerr.Test, cerr.Repository, cerr.Transaction, err).Error()
	}
	fight := entities.Fight{Session: 100, Test: 100, ID_1: 100, ID_2: 101, Res_1: 5, Res_2: value}
	return &fight, nil
}

func (tst RepositoryFight) GetByID(ctx context.Context, id int) (*entities.Test, int, error) {
	var val int
	rows := tst.db.QueryRowContext(ctx, `SELECT res_2 from fight WHERE id_1 = $1;`, 100)
	err := rows.Scan(&val)
	var test entities.Test
	rows = tst.db.QueryRowContext(ctx, `SELECT name, type, level, speed from tests WHERE id = $1;`, id)

	err = rows.Scan(&test.Name, &test.Type, &test.Level, &test.Speed)
	if err != nil {
		return nil, 0, cerr.Err(cerr.Test, cerr.Repository, cerr.Scan, err).Error()
	}
	test.ID = id
	var ids []int
	var ids2 []int
	var que entities.Question

	err = tst.db.SelectContext(ctx, &ids, "SELECT id_questions FROM questions_tests where id_tests=$1", id)
	if err != nil {
		return nil, 0, cerr.Err(cerr.Test, cerr.Repository, cerr.Scan, err).Error()
	}
	for _, i := range ids {
		rows := tst.db.QueryRowContext(ctx, `SELECT name, description from questions WHERE id = $1;`, i)
		err := rows.Scan(&que.Name, &que.Description)
		if err != nil {
			return nil, 0, cerr.Err(cerr.Test, cerr.Repository, cerr.Scan, err).Error()
		}
		que.ID = i
		err = tst.db.SelectContext(ctx, &ids2, "SELECT id_answer FROM answers_questions where id_question=$1", i)
		if err != nil {
			return nil, 0, cerr.Err(cerr.Test, cerr.Repository, cerr.Scan, err).Error()
		}
		for _, j := range ids2 {
			var ans entities.Ans
			rows := tst.db.QueryRowContext(ctx, `SELECT name, is_correct from answers WHERE id = $1;`, j)
			err := rows.Scan(&ans.Name, &ans.IsCorrect)
			ans.ID = i
			if err != nil {
				return nil, 0, cerr.Err(cerr.Question, cerr.Repository, cerr.Scan, err).Error()
			}
			que.Answers = append(que.Answers, ans)
		}
		test.Questions = append(test.Questions, que)
	}

	return &test, val, nil
}

func InitFightRepository(db *sqlx.DB) repository.FightRepo {
	return RepositoryFight{
		db}
}
