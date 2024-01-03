package model

import (
	"database/sql"
	"fmt"
)

// コンテントテーブルデータ
type Content struct {
	Title       string
	Before_code string
	After_code  string
	Review      string
	Memo        string
}

type ContentRepository struct {
	Conn *sql.DB
}

func NewContentRepository(conn *sql.DB) *ContentRepository {
	return &ContentRepository{Conn: conn}
}

type ContentRepositoryInterface interface {
	InsertContent(record *Content) error
	UpdateContentByContentID(id int, record *Content) error
	DeleteContentByContentID(id int) error
	SelectContent() ([]*Content, error)
	SelectContentByContentID(id int) (*Content, error)
}

// contentテーブルにレコードを追加する
func (r *ContentRepository) InsertContent(record *Content) error {
	// レコードを追加する
	if _, err := r.Conn.Exec("INSERT INTO content (title, before_code, after_code, review, memo) VALUES (?, ?, ?, ?, ?)",
		record.Title,
		record.Before_code,
		record.After_code,
		record.Review,
		record.Memo); err != nil {
		return fmt.Errorf("failed to InsertContent: %w", err)
	}
	return nil
}

// contentテーブルのレコードをidを条件に更新する
func (r *ContentRepository) UpdateContentByContentID(id int, record *Content) error {
	if _, err := r.Conn.Exec("UPDATE content SET title = ?, before_code = ?, after_code = ?, review = ?, memo = ? WHERE id = ?",
		record.Title,
		record.Before_code,
		record.After_code,
		record.Review,
		record.Memo,
		id); err != nil {
		return fmt.Errorf("failed to UpdateContentByContentID: %w", err)
	}
	return nil
}

// contentテーブルのレコードをidを条件に削除する
func (r *ContentRepository) DeleteContentByContentID(id int) error {
	if _, err := r.Conn.Exec("DELETE FROM content WHERE id = ?", id); err != nil {
		return fmt.Errorf("failed to DeleteContentByContentID: %w", err)
	}
	return nil
}

// contentテーブルを一覧取得する
func (r *ContentRepository) SelectContent() ([]*Content, error) {
	rows, err := r.Conn.Query("SELECT * FROM content")
	if err != nil {
		return nil, fmt.Errorf("failed to SelectContent: %w", err)
	}

	return ConverToContent(rows)
}

// rowデータをContentデータに変換する
func ConverToContent(rows *sql.Rows) ([]*Content, error) {
	defer rows.Close()

	var contents []*Content
	for rows.Next() {
		content := &Content{}
		if err := rows.Scan(&content.Title, &content.Before_code, &content.After_code, &content.Review, &content.Memo); err != nil {
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
	row := r.Conn.QueryRow("SELECT * FROM content WHERE id = ?", id)

	content := &Content{}
	if err := row.Scan(&content.Title, &content.Before_code, &content.After_code, &content.Review, &content.Memo); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to Scan in SelectContentByContentID: %w", err)
	}

	return content, nil
}
