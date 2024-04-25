package model

import (
	"database/sql"
	"fmt"
)

// taggingテーブルデータ
type Tagging struct{
	ContentID int
	KeywordID int
}

type TaggingRepository struct {
	Conn *sql.DB
}

func NewTaggingRepository(conn *sql.DB) *TaggingRepository{
	return &TaggingRepository{Conn : conn}
}

type TaggingRepositoryInterface interface{
	InsertTagging(record *Tagging) error
}

// taggingテーブルにレコードを追加する
func (r *TaggingRepository)InsertTagging(record *Tagging) error{
	// レコードを追加する
	_,err := r.Conn.Exec("INSERT INTO Tagging (content_id, keyword_id) VALUES (?, ?)",
		record.ContentID,
		record.KeywordID)
	if err != nil{
		return fmt.Errorf("failed to InsertTagging: %w",err)
	}

	return nil
}
