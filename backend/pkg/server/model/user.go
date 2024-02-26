package model

import (
	"database/sql"
	"fmt"
)

// ユーザテーブルの構造体
type User struct {
	UserID   int
	Username string
	Password string
	Email    string
}

type UserRepository struct {
	Conn *sql.DB
}

func NewUserRepository(conn *sql.DB) *UserRepository {
	return &UserRepository{
		Conn: conn,
	}
}

type UserRepositoryInterface interface {
	InsertUser(record *User) error
	SelectUserByUsername(username string) (*User, error)
}

// ユーザテーブルにレコードを登録する
func (r *UserRepository) InsertUser(record *User) error {
	// レコードを追加する
	_, err := r.Conn.Exec("INSERT INTO user (username, password, email) VALUES (?, ?, ?)",
		record.Username,
		record.Password,
		record.Email)
	if err != nil {
		return fmt.Errorf("failed to InsertUser: %w", err)
	}

	return nil
}

// ユーザ名からユーザ情報を取得する
func (r *UserRepository) SelectUserByUsername(username string) (*User, error) {
	row := r.Conn.QueryRow("SELECT * FROM user WHERE username = ?", username)

	user := &User{}
	err := row.Scan(&user.UserID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to scan in SelectUserByUsername: %w", err)
	}

	return user, nil
}
