package user

import (
	"Tatarinhack_backend/internal/entities"
	"Tatarinhack_backend/internal/repository"
	"Tatarinhack_backend/pkg/cerr"
	"context"
	"github.com/jmoiron/sqlx"
)

type RepositoryUser struct {
	db *sqlx.DB
}

func InitUserRepo(db *sqlx.DB) repository.UserRepo {
	return RepositoryUser{
		db}
}

func (usr RepositoryUser) GetLevel(ctx context.Context, id int) (int, error) {
	var level int
	rows := usr.db.QueryRowContext(ctx, `SELECT (lvl) FROM users WHERE id = $1;`, id)

	err := rows.Scan(&level)
	if err != nil {
		return 0, cerr.Err(cerr.User, cerr.Repository, cerr.Scan, err).Error()
	}

	return level, nil
}

func (usr RepositoryUser) GetGrammar(ctx context.Context, id int) (int, error) {
	var grammar int
	rows := usr.db.QueryRowContext(ctx, `SELECT (grammar) FROM users WHERE id = $1;`, id)

	err := rows.Scan(&grammar)
	if err != nil {
		return 0, cerr.Err(cerr.User, cerr.Repository, cerr.Scan, err).Error()
	}

	return grammar, nil
}

func (usr RepositoryUser) GetVocabulary(ctx context.Context, id int) (int, error) {
	var vocabulary int
	rows := usr.db.QueryRowContext(ctx, `SELECT vocabulary FROM users WHERE id = $1;`, id)

	err := rows.Scan(&vocabulary)
	if err != nil {
		return 0, cerr.Err(cerr.User, cerr.Repository, cerr.Scan, err).Error()
	}

	return vocabulary, nil
}

func (usr RepositoryUser) GetSpeaking(ctx context.Context, id int) (int, error) {
	var speaking int
	rows := usr.db.QueryRowContext(ctx, `SELECT (speaking) FROM users WHERE id = $1;`, id)

	err := rows.Scan(&speaking)
	if err != nil {
		return 0, cerr.Err(cerr.User, cerr.Repository, cerr.Scan, err).Error()
	}

	return speaking, nil
}

func (usr RepositoryUser) GetFriendList(ctx context.Context, id int) ([]entities.FriendsList, error) {
	var list []entities.FriendsList
	var ids []int
	err := usr.db.SelectContext(ctx, &ids, "SELECT id_second FROM friends_link where id_first=$1", id)
	if err != nil {
		return nil, cerr.Err(cerr.Test, cerr.Repository, cerr.Scan, err).Error()
	}
	for _, i := range ids {
		var friend entities.FriendsList
		rows := usr.db.QueryRowContext(ctx, `SELECT nick, sex from users WHERE id = $1;`, i)
		err := rows.Scan(&friend.Nick, &friend.Sex)
		if err != nil {
			return nil, cerr.Err(cerr.Test, cerr.Repository, cerr.Scan, err).Error()
		}
		friend.ID = i
		list = append(list, friend)
	}
	return list, nil
}

func (usr RepositoryUser) Create(ctx context.Context, user entities.UserCreate) (int, error) {
	var id int
	transaction, err := usr.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, cerr.Err(cerr.User, cerr.Repository, cerr.Transaction, err).Error()
	}
	row := transaction.QueryRowContext(ctx, `INSERT INTO users (nick, email, goal, sex, hashed_password) VALUES ($1, $2, $3, $4, $5) returning id;`,
		user.Nick, user.Email, user.Goal, user.Sex, user.Password)

	err = row.Scan(&id)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.User, cerr.Repository, cerr.Scan, err).Error()
	}
	if err := transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.User, cerr.Repository, cerr.Commit, err).Error()
	}
	return id, nil
}

func (usr RepositoryUser) Get(ctx context.Context, id int) (*entities.User, error) {
	var OldUser entities.User
	rows := usr.db.QueryRowContext(ctx, `SELECT nick, email, goal, sex, rating, grammar, vocabulary, speaking, lvl, days, achievement, cnt FROM users WHERE id = $1;`, id)

	err := rows.Scan(&OldUser.Nick, &OldUser.Email, &OldUser.Goal, &OldUser.Sex, &OldUser.Rating, &OldUser.Grammar, &OldUser.Vocabulary, &OldUser.Speaking, &OldUser.Level, &OldUser.Days, &OldUser.Achievement, &OldUser.Echp)
	if err != nil {
		return nil, cerr.Err(cerr.Test, cerr.Repository, cerr.Scan, err).Error()
	}
	var ids []int

	err = usr.db.SelectContext(ctx, &ids, "SELECT id_second FROM friends_link where id_first=$1", id)
	if err != nil {
		return nil, cerr.Err(cerr.Question, cerr.Repository, cerr.Scan, err).Error()
	}
	for _, i := range ids {
		var fr entities.FriendsList
		rows := usr.db.QueryRowContext(ctx, `SELECT nick, sex from users WHERE id = $1;`, i)
		err := rows.Scan(&fr.Nick, &fr.Sex)
		fr.ID = i
		if err != nil {
			return nil, cerr.Err(cerr.Question, cerr.Repository, cerr.Scan, err).Error()
		}
		OldUser.FriendsList = append(OldUser.FriendsList, fr)
	}

	OldUser.ID = id
	return &OldUser, nil
}

func (usr RepositoryUser) GetFriendByID(ctx context.Context, id int) (*entities.Friend, error) {
	var Friend entities.Friend

	rows := usr.db.QueryRowContext(ctx, `SELECT nick, sex, achievement, lvl, grammar, vocabulary, speaking FROM users WHERE id = $1;`, id)
	err := rows.Scan(&Friend.Nick, &Friend.Sex, &Friend.Achievement, &Friend.Level, &Friend.Grammar, &Friend.Vocabulary, &Friend.Speaking)
	if err != nil {
		return nil, cerr.Err(cerr.User, cerr.Repository, cerr.Scan, err).Error()
	}
	rowss := usr.db.QueryRowxContext(ctx, `SELECT id_second, nick, sex FROM friends_link WHERE id_first = $1;`, id)
	err = rowss.Scan(&Friend.FriendsList)
	if err != nil {
		return nil, cerr.Err(cerr.User, cerr.Repository, cerr.Scan, err).Error()
	}

	Friend.ID = id
	return &Friend, nil
}

func (usr RepositoryUser) GetManByID(ctx context.Context, id int) (*entities.Man, error) {
	var Man entities.Man

	rows := usr.db.QueryRowContext(ctx, `SELECT nick, achievement, lvl FROM users WHERE id = $1;`, id)
	err := rows.Scan(&Man.Nick, &Man.Achievement, &Man.Level)
	if err != nil {
		return nil, cerr.Err(cerr.User, cerr.Repository, cerr.Scan, err).Error()
	}

	Man.ID = id
	return &Man, nil
}

func (usr RepositoryUser) GetHshPwddByEmail(ctx context.Context, email string) (int, string, error) {
	var (
		id             int
		hashedPassword string
	)
	rows := usr.db.QueryRowContext(ctx, `SELECT id, hashed_password FROM users WHERE email = $1`, email)
	err := rows.Scan(&id, &hashedPassword)
	if err != nil {
		return 0, "", cerr.Err(cerr.User, cerr.Repository, cerr.Scan, err).Error()
	}
	return id, hashedPassword, nil
}

func (usr RepositoryUser) UpdatePasswordByID(ctx context.Context, id int, password string) error {
	transaction, err := usr.db.BeginTx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.User, cerr.Repository, cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `UPDATE users SET hashed_password=$2 WHERE id=$1;`, id, password)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.ExecCon, err).Error()
	}
	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.Rows, err).Error()
	}

	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.NoOneRow, err).Error()
	}

	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.Commit, err).Error()
	}
	return nil
}

func (usr RepositoryUser) AddFriendByNick(ctx context.Context, nickname int, id int) (int, error) {
	var (
		friendID  int
		friendSex string
		name      string
		sex       string
	)
	//friend
	rows := usr.db.QueryRowContext(ctx, `SELECT id, sex FROM users WHERE name = $1;`, nickname)
	err := rows.Scan(&friendID, &friendSex)
	if err != nil {
		return 0, cerr.Err(cerr.User, cerr.Repository, cerr.Scan, err).Error()
	}

	//user
	rowss := usr.db.QueryRowContext(ctx, `SELECT nick, sex FROM users WHERE id = $1;`, id)
	err = rowss.Scan(&name, &sex)
	if err != nil {
		return 0, cerr.Err(cerr.User, cerr.Repository, cerr.Scan, err).Error()
	}

	// ----------friend
	transaction, err := usr.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, cerr.Err(cerr.User, cerr.Repository, cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `INSERT INTO friends_link id_first, id_second, nick, sex VALUES ($1, $2, $3, $4)`, friendID, id, name, sex)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.User, cerr.Repository, cerr.ExecCon, err).Error()
	}
	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.User, cerr.Repository, cerr.Rows, err).Error()
	}

	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.User, cerr.Repository, cerr.NoOneRow, err).Error()
	}

	//--------user
	result, err = transaction.ExecContext(ctx, `INSERT INTO friends_link id_first, id_second, nick, sex VALUES ($1, $2, $3, $4)`, id, friendID, nickname, friendSex)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.User, cerr.Repository, cerr.ExecCon, err).Error()
	}
	count, err = result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.User, cerr.Repository, cerr.Rows, err).Error()
	}

	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.User, cerr.Repository, cerr.NoOneRow, err).Error()
	}

	if err = transaction.Commit(); err != nil {
		return 0, cerr.Err(cerr.User, cerr.Repository, cerr.Commit, err).Error()
	}

	return friendID, nil
}

func (usr RepositoryUser) AddFriendByID(ctx context.Context, friendID int, userID int) error {
	var (
		friendSex  string
		friendNick string
		userSex    string
		userNick   string
	)

	//friend
	rows := usr.db.QueryRowContext(ctx, `SELECT nick, sex FROM users WHERE id = $1;`, friendID)
	err := rows.Scan(&friendNick, &friendSex)
	if err != nil {
		return cerr.Err(cerr.User, cerr.Repository, cerr.Scan, err).Error()
	}

	//user
	rowss := usr.db.QueryRowContext(ctx, `SELECT nick, sex FROM users WHERE id = $1;`, userID)
	err = rowss.Scan(&userNick, &userSex)
	if err != nil {
		return cerr.Err(cerr.User, cerr.Repository, cerr.Scan, err).Error()
	}

	// ----------friend
	transaction, err := usr.db.BeginTx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.User, cerr.Repository, cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `INSERT INTO friends_link id_first, id_second, nick, sex VALUES ($1, $2, $3, $4)`, friendID, userID, userNick, userSex)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.ExecCon, err).Error()
	}
	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.Rows, err).Error()
	}

	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.NoOneRow, err).Error()
	}

	//--------user
	result, err = transaction.ExecContext(ctx, `INSERT INTO friends_link id_first, id_second, nick, sex VALUES ($1, $2, $3, $4)`, userID, friendID, friendNick, friendSex)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.ExecCon, err).Error()
	}
	count, err = result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.Rows, err).Error()
	}

	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.NoOneRow, err).Error()
	}

	if err = transaction.Commit(); err != nil {
		return cerr.Err(cerr.User, cerr.Repository, cerr.Commit, err).Error()
	}

	return nil
}

func (usr RepositoryUser) UpdateGrammarByID(ctx context.Context, id int, amount int) error {
	transaction, err := usr.db.BeginTx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.User, cerr.Repository, cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `UPDATE users SET grammar=$2 WHERE id=$1;`, id, amount)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.ExecCon, err).Error()
	}
	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.Rows, err).Error()
	}

	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.NoOneRow, err).Error()
	}

	if err := transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.Commit, err).Error()
	}

	return nil
}

func (usr RepositoryUser) UpdateSpeakingByID(ctx context.Context, id int, amount int) error {
	transaction, err := usr.db.BeginTx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.User, cerr.Repository, cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `UPDATE users SET speaking=$2 WHERE id=$1;`, id, amount)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.ExecCon, err).Error()
	}
	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.Rows, err).Error()
	}

	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.NoOneRow, err).Error()
	}

	if err := transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.Commit, err).Error()
	}

	return nil
}

func (usr RepositoryUser) UpdateVocabularyByID(ctx context.Context, id int, amount int) error {
	transaction, err := usr.db.BeginTx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.User, cerr.Repository, cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `UPDATE users SET vocabulary=$2 WHERE id=$1;`, id, amount)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.ExecCon, err).Error()
	}
	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.Rows, err).Error()
	}

	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.NoOneRow, err).Error()
	}

	if err := transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.Commit, err).Error()
	}

	return nil
}

func (usr RepositoryUser) UpdateLevelByID(ctx context.Context, id int, amount int) error {
	transaction, err := usr.db.BeginTx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.User, cerr.Repository, cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `UPDATE users SET lvl=$2 WHERE id=$1;`, id, amount)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.ExecCon, err).Error()
	}
	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.Rows, err).Error()
	}

	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.NoOneRow, err).Error()
	}

	if err := transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.Commit, err).Error()
	}

	return nil
}

func (usr RepositoryUser) UpdateDaysByID(ctx context.Context, id int, amount int) error {
	transaction, err := usr.db.BeginTx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.User, cerr.Repository, cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `UPDATE users SET days=$2 WHERE id=$1;`, id, amount)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.ExecCon, err).Error()
	}
	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.Rows, err).Error()
	}

	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.NoOneRow, err).Error()
	}

	if err := transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.Commit, err).Error()
	}

	return nil
}

func (usr RepositoryUser) UpdateAchievementByID(ctx context.Context, id int, amount int) error {
	transaction, err := usr.db.BeginTx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.User, cerr.Repository, cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `UPDATE users SET achievment=$2 WHERE id=$1;`, id, amount)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.ExecCon, err).Error()
	}
	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.Rows, err).Error()
	}

	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.NoOneRow, err).Error()
	}

	if err := transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.Commit, err).Error()
	}

	return nil
}

func (usr RepositoryUser) Delete(ctx context.Context, id int) error {
	transaction, err := usr.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.User, cerr.Repository, cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `DELETE FROM users WHERE id=$1;`, id)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.ExecCon, err).Error()
	}
	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.Rows, err).Error()
	}
	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.NoOneRow, err).Error()
	}
	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.Commit, err).Error()
	}

	return nil
}

func (usr RepositoryUser) GetEchp(ctx context.Context, id int) (int, error) {
	var cnt int
	rows := usr.db.QueryRowContext(ctx, `SELECT cnt from users where id=$1;`, id)
	err := rows.Scan(&cnt)
	if err != nil {
		return 0, err
	}

	return cnt, nil
}

func (usr RepositoryUser) UpdEchp(ctx context.Context, id int, p_cnt int) error {
	var cnt int
	rows := usr.db.QueryRowContext(ctx, `SELECT cnt from users where id=$1;`, id)
	err := rows.Scan(&cnt)
	if err != nil {
		return err
	}

	transaction, err := usr.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	result, err := transaction.ExecContext(ctx, `UPDATE users set cnt=$2 where id=$1`, id, p_cnt)
	if err != nil {
		if err := transaction.Rollback(); err != nil {
			return err
		}
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.Rows, err).Error()
	}

	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.NoOneRow, err).Error()
	}

	if err := transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.User, cerr.Repository, cerr.Commit, err).Error()
	}
	return nil
}
