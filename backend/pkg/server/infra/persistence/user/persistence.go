package user

import (
	"database/sql"
	"log"

	model "github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/model/user"
	repo "github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/repository/user"
)

type userPersistence struct {
	db *sql.DB
}

func NewPersistence(db *sql.DB) repo.UserRepo {
	return &userPersistence{
		db: db,
	}
}

func (up userPersistence) SelectUserByUserID(id string) (*model.User, error) {
	row := up.db.QueryRow("SELECT * FROM user_table WHERE user_id = ?", id)
	return convertToUser(row)
}

func (up userPersistence) SelectUserByUserNamePassword(name string, password string) (*model.User, error) {
	row := up.db.QueryRow("SELECT * FROM user_table WHERE user_name = ? AND password = ?", name, password)
	return convertToUser(row)
}

func (up userPersistence) InsertUser(record *model.User) error {
	// userテーブルへのレコードの登録を行うSQLを入力する
	stmt, err := up.db.Prepare("INSERT INTO user_table (user_id, user_name, password, auth_token) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		record.UserID,
		record.UserName,
		record.Password,
		record.AuthToken,
	)
	return err
}

// convertToUser rowデータをUserデータへ変換する
func convertToUser(row *sql.Row) (*model.User, error) {
	user := &model.User{}
	err := row.Scan(
		&user.UserID,
		&user.UserName,
		&user.Password,
		&user.AuthToken,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	return user, nil
}
