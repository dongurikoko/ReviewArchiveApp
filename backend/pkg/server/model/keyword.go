package model

import (
	"database/sql"
	"errors"
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
	InsertKeyword(keyword string,tx *sql.Tx)(int,error)
	SelectKeywordIDByKeyword(keyword string,tx *sql.Tx)(int,error)
	DeleteKeywordByID(id int) error
	SelectKeywordByID(id int) ([]*Keyword, error)
	SelectStringKeywordByID(contentID int) ([]string, error)
	SelectIDByContentKeyword(contentKeyword string) ([]int, error)
}

/* キーワードテーブルに挿入する(バルクインサートを使用する場合の実装)
func (r *KeywordRepository) InsertKeyword(id int, keywords []string) error {
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
*/

// キーワードテーブルにレコードを追加して、keywordIDを返す
func (r *KeywordRepository)InsertKeyword(keyword string,tx *sql.Tx)(int,error){
	result,err := tx.Exec("INSERT INTO Keywords (keyword) VALUES (?)",keyword)
	if err != nil{
		return 0,fmt.Errorf("failed to InsertKeyword in InsertKeyword: %w",err)
	}
	keywordID,err := result.LastInsertId()
	if err != nil{
		return 0,fmt.Errorf("failed to LastInsertId in InsertKeyword: %w",err)
	}
	return int(keywordID),nil
}

// キーワードを元にkeywordIDを取得する(無い場合は新規登録)
func (r *KeywordRepository)SelectKeywordIDByKeyword(keyword string,tx *sql.Tx)(int,error){
	var keywordID int
	err := tx.QueryRow("SELECT id FROM Keywords WHERE keyword = ?",keyword).Scan(&keywordID)
	if err != nil{
		if errors.Is(err,sql.ErrNoRows){
			// キーワードがない場合は新規登録
			keywordID,err = r.InsertKeyword(keyword,tx)
			if err != nil{
				return 0,fmt.Errorf("failed to InsertKeyword in SelectKeywordIDByKeyword: %w",err)
			}
			return keywordID,nil
		}
		return 0,fmt.Errorf("failed to SelectKeywordIDByKeyword: %w",err)
	}

	return keywordID,nil
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

// キーワードをidを条件に取得する(IDとkeywordを返す)
func (r *KeywordRepository) SelectKeywordByID(id int) ([]*Keyword, error) {
	rows, err := r.Conn.Query("SELECT * FROM keyword WHERE id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("failed to SELECT * in SelectKeywordByID: %w", err)
	}
	return ConverToKeyword(rows)
}

// キーワードをidを条件に取得する([]stringを返す場合)
func (r *KeywordRepository) SelectStringKeywordByID(contentID int) ([]string, error) {
	query := `
    SELECT 
        k.keyword
    FROM 
        Keywords k
    JOIN 
        Tagging t ON k.id = t.keyword_id
    WHERE 
        t.content_id = ?;
    `
	rows, err := r.Conn.Query(query, contentID)
	if err != nil {
		return nil, fmt.Errorf("failed to SELECT keyword in SelectStringKeywordByID: %w", err)
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

// キーワードテーブルからcontentKeywordを条件にidを取得する
func (r *KeywordRepository) SelectIDByContentKeyword(contentKeyword string) ([]int, error) {
	rows, err := r.Conn.Query("SELECT id FROM keyword WHERE contentKeyword LIKE ?", "%"+contentKeyword+"%")
	if err != nil {
		return nil, fmt.Errorf("failed to SELECT id in SelectIDByContentKeyword: %w", err)
	}
	return ConverToInt(rows)
}

// rowデータを[]intに変換する
func ConverToInt(rows *sql.Rows) ([]int, error) {
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("failed to Scan in ConverToInt: %w", err)
		}
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after row iteration in convertToInt: %w", err)
	}
	return ids, nil
}
