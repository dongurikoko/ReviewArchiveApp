package model

import (
	"database/sql"
	"fmt"
)

// キーワードテーブルデータ
type Keyword struct {
	ID             int
	ContentKeyword string
}

type KeywordRepository struct {
	Conn *sql.DB
}

func NewKeywordRepository(conn *sql.DB) *KeywordRepository {
	return &KeywordRepository{
		Conn: conn,
	}
}

type KeywordRepositoryInterface interface {
	InsertKeyword(id int, keywords []string) error
	DeleteKeywordByID(id int) error
	SelectKeywordByID(id int) ([]*Keyword, error)
	SelectStringKeywordByID(id int) ([]string, error)
}

// キーワードテーブルに挿入する
func (r *KeywordRepository) InsertKeyword(id int, keywords []string) error {
	// バルクインサートの実装
	insert := "INSERT INTO keyword (id, contentKeyword) VALUES "

	vals := make([]any, 0, len(keywords))
	for _, value := range keywords {
		insert += fmt.Sprintf("(%v,?),", id)
		vals = append(vals, value)
	}
	insert = insert[:len(insert)-1] // 最後のカンマを削除する

	stmt, err := r.Conn.Prepare(insert)

	if err != nil {
		return fmt.Errorf("failed to prepare insert in InsertKeyword: %w", err)
	}

	defer stmt.Close()

	if _, err = stmt.Exec(vals...); err != nil {
		return fmt.Errorf("failed to exec insert in InsertKeyword: %w", err)
	}
	return nil

}

/* キーワードを更新する
func (r *KeywordRepository) UpdateKeywordByID(id int, record *Keyword) error {
	if _, err := r.Conn.Exec("UPDATE keyword SET content_keyword = ? WHERE id = ?",
		record.ContentKeyword,
		id); err != nil {
		return fmt.Errorf("failed to UpdateKeywordByID: %w", err)
	}
	return nil
}*/

// キーワードを削除する
func (r *KeywordRepository) DeleteKeywordByID(id int) error {
	if _, err := r.Conn.Exec("DELETE FROM keyword WHERE id = ?", id); err != nil {
		return fmt.Errorf("failed to DeleteKeywordBy: %w", err)
	}
	return nil
}

// キーワードをidを条件に取得する
func (r *KeywordRepository) SelectKeywordByID(id int) ([]*Keyword, error) {
	rows, err := r.Conn.Query("SELECT * FROM keyword WHERE id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("failed to SELECT * in SelectKeywordByID: %w", err)
	}
	return ConverToKeyword(rows)
}

// キーワードをidを条件に取得する([]stringを返す場合)
func (r *KeywordRepository) SelectStringKeywordByID(id int) ([]string, error) {
	rows, err := r.Conn.Query("SELECT contentKeyword FROM keyword WHERE id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("failed to SELECT * in SelectKeywordByID: %w", err)
	}
	return ConverToString(rows)
}

// rowデータをKeywordデータに変換する
func ConverToKeyword(rows *sql.Rows) ([]*Keyword, error) {
	defer rows.Close()

	var keywords []*Keyword
	for rows.Next() {
		keyword := &Keyword{}
		if err := rows.Scan(&keyword.ID, &keyword.ContentKeyword); err != nil {
			return nil, fmt.Errorf("failed to Scan in ConverToKeyword: %w", err)
		}
		keywords = append(keywords, keyword)
	}
	return keywords, nil
}

// rowデータを[]stringに変換する
func ConverToString(rows *sql.Rows) ([]string, error) {
	defer rows.Close()

	var keywords []string
	for rows.Next() {
		var keyword string
		if err := rows.Scan(&keyword); err != nil {
			return nil, fmt.Errorf("failed to Scan in ConverToString: %w", err)
		}
		keywords = append(keywords, keyword)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after row iteration in convertTostring: %w", err)
	}
	return keywords, nil
}
