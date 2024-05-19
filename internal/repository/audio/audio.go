package audio

import (
	"Tatarinhack_backend/internal/entities"
	"Tatarinhack_backend/internal/repository"
	"Tatarinhack_backend/pkg/cerr"
	"context"
	"github.com/jmoiron/sqlx"
)

type RepositoryAudio struct {
	db *sqlx.DB
}

func (a RepositoryAudio) Create(ctx context.Context, audio entities.AudioBase) (int, error) {
	var id int
	transaction, err := a.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, cerr.Err(cerr.Question, cerr.Repository, cerr.Transaction, err).Error()
	}
	row := transaction.QueryRowContext(ctx, `INSERT INTO audios (correct_answer) VALUES ($1) returning id;`,
		audio.Text)
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

func (a RepositoryAudio) GetByID(ctx context.Context, id int) (*entities.Audio, error) {
	var audio entities.Audio
	rows := a.db.QueryRowContext(ctx, `SELECT correct_answer from audios WHERE id = $1;`, id)
	err := rows.Scan(&audio.Text)
	if err != nil {
		return nil, cerr.Err(cerr.Teach, cerr.Repository, cerr.Scan, err).Error()
	}
	audio.ID = id
	return &audio, nil
}

func InitAudioRepository(db *sqlx.DB) repository.AudioRepo {
	return RepositoryAudio{
		db}
}
