package model

import (
	"database/sql"
	"fmt"
)

// taggingテーブルデータ
type Tagging struct {
	ContentID int
	KeywordID int
}

type TaggingRepository struct {
	Conn *sql.DB
}

func NewTaggingRepository(conn *sql.DB) *TaggingRepository {
	return &TaggingRepository{Conn: conn}
}

type TaggingRepositoryInterface interface {
	InsertTagging(record *Tagging, tx *sql.Tx) error
	DeleteTaggingByContentID(contentID int, tx *sql.Tx) error
}

// taggingテーブルにレコードを追加する
func (r *TaggingRepository) InsertTagging(record *Tagging, tx *sql.Tx) error {
	_, err := tx.Exec("INSERT INTO Tagging (content_id,keyword_id) VALUES (?,?)", record.ContentID, record.KeywordID)
	if err != nil {
		return fmt.Errorf("failed to InsertTagging: %w", err)
	}
	return nil
}

// taggingテーブルのレコードをcontentIDを条件に削除する
func (r *TaggingRepository) DeleteTaggingByContentID(contentID int, tx *sql.Tx) error {
	_, err := tx.Exec("DELETE FROM Tagging WHERE content_id = ?", contentID)
	if err != nil {
		return fmt.Errorf("failed to DeleteTaggingByContentID: %w", err)
	}
	return nil
}
