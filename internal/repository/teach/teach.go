package teach

import (
	"Tatarinhack_backend/internal/entities"
	"Tatarinhack_backend/internal/repository"
	"Tatarinhack_backend/pkg/cerr"
	"context"
	"github.com/jmoiron/sqlx"
)

type RepositoryTeach struct {
	db *sqlx.DB
}

func InitTeachRepository(db *sqlx.DB) repository.TeachRepo {
	return RepositoryTeach{
		db}
}

func (tch RepositoryTeach) Create(ctx context.Context, teach entities.TeachCreate) (int, error) {
	var id int
	transaction, err := tch.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, cerr.Err(cerr.Teach, cerr.Repository, cerr.Transaction, err).Error()
	}
	row := transaction.QueryRowContext(ctx, `INSERT INTO teachers (nick, email, hashed_password) VALUES ($1, $2, $3) returning id;`,
		teach.Nick, teach.Email, teach.Password)
	err = row.Scan(&id)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Teach, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Teach, cerr.Repository, cerr.Scan, err).Error()
	}
	if err := transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Teach, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Teach, cerr.Repository, cerr.Commit, err).Error()
	}
	return id, nil
}

func (tch RepositoryTeach) GetByID(ctx context.Context, id int) (*entities.Teach, error) {
	var teach entities.Teach
	rows := tch.db.QueryRowContext(ctx, `SELECT nick, email from teachers WHERE id = $1;`, id)
	err := rows.Scan(&teach.Nick, &teach.Email)
	if err != nil {
		return nil, cerr.Err(cerr.Teach, cerr.Repository, cerr.Scan, err).Error()
	}
	teach.ID = id
	return &teach, nil
}

func (tch RepositoryTeach) GetPasswordByEmail(ctx context.Context, email string) (int, string, error) {
	var hshPassword string
	var id int
	row := tch.db.QueryRowContext(ctx, `SELECT id, hashed_password FROM teachers WHERE email = $1;`, email)
	err := row.Scan(&id, &hshPassword)
	if err != nil {
		return 0, "", cerr.Err(cerr.Teach, cerr.Repository, cerr.Scan, err).Error()
	}
	return id, hshPassword, nil
}

func (tch RepositoryTeach) UpdatePasswordByID(ctx context.Context, id int, newPassword string) error {
	transaction, err := tch.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Teach, cerr.Repository, cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `UPDATE teachers SET hashed_password=$2 WHERE id=$1;`, id, newPassword)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Teach, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Teach, cerr.Repository, cerr.ExecCon, err).Error()
	}

	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Teach, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Teach, cerr.Repository, cerr.Rows, err).Error()
	}

	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Teach, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Teach, cerr.Repository, cerr.NoOneRow, err).Error()
	}

	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Teach, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Teach, cerr.Repository, cerr.Commit, err).Error()
	}
	return nil
}

func (tch RepositoryTeach) Delete(ctx context.Context, id int) error {
	transaction, err := tch.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Teach, cerr.Repository, cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `DELETE FROM teachers WHERE id=$1;`, id)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Teach, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Teach, cerr.Repository, cerr.ExecCon, err).Error()
	}
	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Teach, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Teach, cerr.Repository, cerr.Rows, err).Error()
	}
	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Teach, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Teach, cerr.Repository, cerr.NoOneRow, err).Error()
	}
	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Teach, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Teach, cerr.Repository, cerr.Commit, err).Error()
	}

	return nil

}
