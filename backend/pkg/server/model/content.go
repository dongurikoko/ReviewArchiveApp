package model

import (
	"database/sql"
	"errors"
	"fmt"
)

// コンテントテーブルデータ
type Content struct {
	Title       string
	BeforeCode 	string
	AfterCode  	string
	Review      string
	Memo        string
	UserID	  	int
}

type ContentWithID struct {
	ContentID    int
	ContentValue Content
}

type ContentRepository struct {
	Conn *sql.DB
}

func NewContentRepository(conn *sql.DB) *ContentRepository {
	return &ContentRepository{Conn: conn}
}

type ContentRepositoryInterface interface {
	InsertContent(record *Content,tx *sql.Tx) (int, error)
	UpdateContentByContentID(id int, record *Content,tx *sql.Tx) error
	DeleteContentByContentID(id int) error
	SelectContent() ([]*ContentWithID, error)
	SelectContentByContentID(id int) (*Content, error)
}

// contentテーブルにレコードを追加し、追加したcontentIDを返す
func (r *ContentRepository) InsertContent(record *Content,tx *sql.Tx) (int, error) {
	result, err := tx.Exec("INSERT INTO Contents (title,before_code,after_code,review,memo,user_id) VALUES (?,?,?,?,?,?)",
		record.Title, record.BeforeCode, record.AfterCode, record.Review, record.Memo, record.UserID)
	if err != nil {
		return 0, fmt.Errorf("failed to InsertContent: %w", err)
	}
	contentID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to LastInsertId in InsertContent: %w", err)
	}
	return int(contentID), nil
}

// contentテーブルのレコードをidを条件に更新する
func (r *ContentRepository) UpdateContentByContentID(id int, record *Content,tx *sql.Tx) error {
	if _, err := tx.Exec("UPDATE Contents SET title = ?, before_code = ?, after_code = ?, review = ?, memo = ?, user_id = ? WHERE id = ?",
		record.Title, record.BeforeCode, record.AfterCode, record.Review, record.Memo, record.UserID, id); err != nil {
		return fmt.Errorf("failed to UpdateContentByContentID: %w", err)
	}
	return nil
}

// contentテーブルのレコードをidを条件に削除する
func (r *ContentRepository) DeleteContentByContentID(id int) error {
	if _, err := r.Conn.Exec("DELETE FROM content WHERE content_id = ?", id); err != nil {
		return fmt.Errorf("failed to DeleteContentByContentID: %w", err)
	}
	return nil
}

// contentテーブルを一覧取得する
func (r *ContentRepository) SelectContent() ([]*ContentWithID, error) {
	rows, err := r.Conn.Query("SELECT * FROM content")
	if err != nil {
		return nil, fmt.Errorf("failed to SelectContent: %w", err)
	}

	return ConverToContent(rows)
}

// rowデータをContentデータに変換する
func ConverToContent(rows *sql.Rows) ([]*ContentWithID, error) {
	defer rows.Close()

	var contents []*ContentWithID
	for rows.Next() {
		content := &ContentWithID{}
		if err := rows.Scan(&content.ContentID, &content.ContentValue.Title, &content.ContentValue.BeforeCode,
			&content.ContentValue.AfterCode, &content.ContentValue.Review, &content.ContentValue.Memo); err != nil {
			return nil, fmt.Errorf("error scanning row in ConverToContent: %w", err)
		}
		contents = append(contents, content)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after row iteration in ConverToContent: %w", err)
	}
	return contents, nil
}

// contentテーブルをidを条件に取得する
func (r *ContentRepository) SelectContentByContentID(id int) (*Content, error) {
	row := r.Conn.QueryRow("SELECT title,before_code,after_code,review,memo,user_id FROM Contents WHERE id = ?", id)

	if err := row.Err(); err != nil {
		return nil, fmt.Errorf("failed to QueryRow in SelectContentByContentID: %w", err)
	}
	content := &Content{}
	err := row.Scan(&content.Title, &content.BeforeCode, &content.AfterCode, &content.Review, &content.Memo, &content.UserID)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("failed to Scan in SelectContentByContentID: %w", err)
	}

	return content, nil
}
