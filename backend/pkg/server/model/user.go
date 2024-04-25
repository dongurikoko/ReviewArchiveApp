package model

import (
	"database/sql"
	"fmt"
)

// ユーザテーブルデータ
type User struct{
	userID int
	UUID string
}

type UserRepository struct {
	Conn *sql.DB
}

func NewUserRepository(conn *sql.DB) *UserRepository{
	return &UserRepository{Conn : conn}
}

type UserRepositoryInterface interface{
	InsertUser(record *User)(int,error)
	SelectUserIDByUUID(uuid string)(int,error)
}

func (r *UserRepository)InsertUser(record *User)(int,error){
	result,err := r.Conn.Exec("INSERT INTO Users (uuid) VALUES (?)",
		record.UUID)
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

// uuidを元にユーザIDを取得する
func (r *UserRepository)SelectUserIDByUUID(uuid string)(int,error){
	var userID int
	err := r.Conn.QueryRow("SELECT id FROM Users WHERE uuid = ?",uuid).Scan(&userID)
	if err != nil{
		return 0,fmt.Errorf("failed to SelectUserIDByUUID: %w",err)
	}

	return userID,nil
}
