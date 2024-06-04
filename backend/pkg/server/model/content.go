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
	InsertContent(record *Content) (int, error)
	UpdateContentByContentID(id int, record *Content) error
	DeleteContentByContentID(id int) error
	SelectContent() ([]*ContentWithID, error)
	SelectContentByContentID(id int) (*Content, error)
}

// contentテーブルにレコードを追加する
func (r *ContentRepository) InsertContent(record *Content) (int, error) {
	// レコードを追加する
	result, err := r.Conn.Exec("INSERT INTO Contents (title, before_code, after_code, review, memo) VALUES (?, ?, ?, ?, ?)",
		record.Title,
		record.BeforeCode,
		record.AfterCode,
		record.Review,
		record.Memo)
	if err != nil {
		return 0, fmt.Errorf("failed to InsertContent: %w", err)
	}

	// 最後に挿入された行のIDを取得する
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve last insert ID: %w", err)
	}

	return int(id), nil
}

// contentテーブルのレコードをidを条件に更新する
func (r *ContentRepository) UpdateContentByContentID(id int, record *Content) error {
	if _, err := r.Conn.Exec("UPDATE content SET title = ?, before_code = ?, after_code = ?, review = ?, memo = ? WHERE content_id = ?",
		record.Title,
		record.BeforeCode,
		record.AfterCode,
		record.Review,
		record.Memo,
		id); err != nil {
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
	row := r.Conn.QueryRow("SELECT title,before_code,after_code,review,memo FROM content WHERE content_id = ?", id)

	if err := row.Err(); err != nil {
		return nil, fmt.Errorf("failed to QueryRow in SelectContentByContentID: %w", err)
	}
	content := &Content{}
	err := row.Scan(&content.Title, &content.BeforeCode, &content.AfterCode, &content.Review, &content.Memo)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("failed to Scan in SelectContentByContentID: %w", err)
	}

	return content, nil
}
