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
	DeleteContentByContentID(id int,tx *sql.Tx) error
	SelectContent() ([]*ContentWithID, error)
	SelectContentByContentID(id int) (*Content, error)
	SelectContentByKeywordsAndUserID(keyword string,userID int)([]*ContentWithID,error)
	SelectContentByUserID(userID int) ([]*ContentWithID, error)
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
func (r *ContentRepository) DeleteContentByContentID(id int,tx *sql.Tx) error {
	_, err := tx.Exec("DELETE FROM Contents WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to DeleteContentByContentID: %w", err)
	}
	return nil
}

// contentテーブルを一覧取得する
func (r *ContentRepository) SelectContent() ([]*ContentWithID, error) {
	rows, err := r.Conn.Query("SELECT * FROM Contents")
	if err != nil {
		return nil, fmt.Errorf("failed to SelectContent: %w", err)
	}

	return ConverToContentWithID(rows)
}

// rowデータをContentWithIDデータに変換する
func ConverToContentWithID(rows *sql.Rows) ([]*ContentWithID, error) {
	defer rows.Close()

	var contents []*ContentWithID
	for rows.Next() {
		content := &ContentWithID{}
		if err := rows.Scan(&content.ContentID, &content.ContentValue.Title, &content.ContentValue.BeforeCode,
			&content.ContentValue.AfterCode, &content.ContentValue.Review, &content.ContentValue.Memo,&content.ContentValue.UserID); err != nil {
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

func (r *ContentRepository) SelectContentByUserID(userID int) ([]*ContentWithID, error) {
	rows, err := r.Conn.Query("SELECT * FROM Contents WHERE user_id = ?", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to SelectContentByUserID in SearchContents: %w", err)
	}

	return ConverToContentWithID(rows)
}

func (r *ContentRepository) SelectContentByKeywordsAndUserID(keyword string,userID int)([]*ContentWithID,error){
	// Contents,Tagging,Keywordテーブルを結合し、userID,keywordがそれぞれ一致するコンテンツを取得
	query := `
	SELECT 
		c.id AS content_id,
        c.title,
        c.before_code,
        c.after_code,
        c.review,
        c.memo,
        c.user_id
    FROM 
        Contents c
	JOIN 
        Tagging t ON c.id = t.content_id
    JOIN 
        Keywords k ON t.keyword_id = k.id
    WHERE 
        k.keyword LIKE ? AND c.user_id = ?;
    `
	partialKeyword := "%" + keyword + "%"
	rows, err := r.Conn.Query(query, partialKeyword, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to SelectContentByKeywordsAndUserID in SearchContents: %w", err)
	}

	return ConverToContentWithID(rows)
}

