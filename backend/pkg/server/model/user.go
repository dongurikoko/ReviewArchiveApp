package model

import (
	"database/sql"
	"errors"
	"fmt"
)

/* ユーザテーブルデータ
type User struct{
	ID int
	UID string
}*/

type UserRepository struct {
	Conn *sql.DB
}

func NewUserRepository(conn *sql.DB) *UserRepository{
	return &UserRepository{Conn : conn}
}

type UserRepositoryInterface interface{
	InsertUser(uid string)(int,error)
	SelectUserIDByUID(uid string)(int,error)
	SelectUserIDByUIDWithError(uid string)(int,error)
}

// ユーザテーブルにレコードを追加して、ユーザIDを返す
func (r *UserRepository)InsertUser(uid string)(int,error){
	result,err := r.Conn.Exec("INSERT INTO Users (uid) VALUES (?)",uid)
	if err != nil{
		return 0,fmt.Errorf("failed to InsertUser: %w",err)
	}

	// 最後に挿入された行のIDを取得する
	id,err := result.LastInsertId()
	if err != nil{
		return 0,fmt.Errorf("failed to retrieve last insert ID: %w",err)
	}

	return int(id),nil
}

// uidを元にユーザIDを取得する(無い場合は新規登録)
func (r *UserRepository)SelectUserIDByUID(uid string)(int,error){
	var userID int
	err := r.Conn.QueryRow("SELECT id FROM Users WHERE uid = ?",uid).Scan(&userID)
	if err != nil{
		if errors.Is(err,sql.ErrNoRows){
			// ユーザがいない場合は新規登録
			userID,err = r.InsertUser(uid)
			if err != nil{
				return 0,fmt.Errorf("failed to InsertUser in SelectUserIDByUID: %w",err)
			}
			return userID,nil
		}
		return 0,fmt.Errorf("failed to SelectUserIDByUID: %w",err)
	}

	return userID,nil
}

// uidを元にユーザIDを取得する(無い場合はエラーを返す)
func (r *UserRepository)SelectUserIDByUIDWithError(uid string)(int,error){
	var userID int
	err := r.Conn.QueryRow("SELECT id FROM Users WHERE uid = ?",uid).Scan(&userID)
	if err != nil{
		if err == sql.ErrNoRows {
            return 0, fmt.Errorf("no user found with uid: %s", uid)
        }
		return 0,fmt.Errorf("failed to SelectUserIDByUID: %w",err)
	}

	return userID,nil
}
