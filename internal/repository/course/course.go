package course

import (
	"Tatarinhack_backend/internal/entities"
	"Tatarinhack_backend/internal/repository"
	"Tatarinhack_backend/pkg/cerr"
	"context"
	"github.com/jmoiron/sqlx"
)

type RepositoryCourse struct {
	db *sqlx.DB
}

func InitCourseRepository(db *sqlx.DB) repository.CourseRepo {
	return RepositoryCourse{
		db}
}

func (cou RepositoryCourse) Create(ctx context.Context, cour entities.CourseBase) (int, error) {
	var id int
	transaction, err := cou.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, cerr.Err(cerr.Test, cerr.Repository, cerr.Transaction, err).Error()
	}
	row := transaction.QueryRowContext(ctx, `INSERT INTO courses (name, level, description) VALUES ($1, $2, $3) returning id;`,
		cour.Name, cour.Level, cour.Description)
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

func (cou RepositoryCourse) GetByID(ctx context.Context, id int) (*entities.Course, error) {
	var course entities.Course
	rows := cou.db.QueryRowContext(ctx, `SELECT (name, level, description) from courses WHERE id = $1;`, id)
	err := rows.Scan(&course.Name, &course.Level, &course.Description)
	if err != nil {
		return nil, cerr.Err(cerr.Test, cerr.Repository, cerr.Scan, err).Error()
	}
	course.ID = id
	var ids []int
	rowss := cou.db.QueryRowxContext(ctx, `SELECT id_tests FROM courses_tests WHERE id_courses = $1;`, id)
	err = rowss.Scan(&ids)
	if err != nil {
		return nil, cerr.Err(cerr.User, cerr.Repository, cerr.Scan, err).Error()
	}

	for _, i := range ids {
		var question entities.TestGet
		rows := cou.db.QueryRowContext(ctx, `SELECT name from tests WHERE id = $1;`, i)
		err := rows.Scan(&question.Name)
		if err != nil {
			return nil, cerr.Err(cerr.Test, cerr.Repository, cerr.Scan, err).Error()
		}
		question.ID = id
		course.Tests = append(course.Tests, question)
	}
	return &course, nil
}

func (cou RepositoryCourse) AddT(ctx context.Context, idTest int, idCourse int) error {
	transaction, err := cou.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Course, cerr.Repository, cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `INSERT INTO courses_tests (id_courses, id_tests) VALUES ($1, $2);`, idCourse, idTest)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Course, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Course, cerr.Repository, cerr.ExecCon, err).Error()
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

func (cou RepositoryCourse) DeleteT(ctx context.Context, idCourse int, idQue int) error {
	transaction, err := cou.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Test, cerr.Repository, cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `DELETE FROM courses_tests where id_courses=$1 and  id_tests = $2;`, idCourse, idQue)
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
