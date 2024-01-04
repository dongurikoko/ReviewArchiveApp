package controller

import (
	"fmt"

	"reviewArchive/pkg/server/model"
)

type ContentRequest struct {
	Title       string
	Before_code string
	After_code  string
	Review      string
	Memo        string
	Keywords    []string
}

type ContentController struct {
	ContentRepository model.ContentRepositoryInterface
	KeywordRepository model.KeywordRepositoryInterface
}

func NewContentContoroller(contentRepository model.ContentRepositoryInterface,
	keywordRepository model.KeywordRepositoryInterface) *ContentController {
	return &ContentController{
		ContentRepository: contentRepository,
		KeywordRepository: keywordRepository,
	}
}

type ContentControllerInterface interface {
	ContentCreate(record *ContentRequest) error
	ContentDelete(id int) error
}

// コンテンツ作成ロジック
func (c *ContentController) ContentCreate(record *ContentRequest) error {
	recordContent := &model.Content{
		Title:       record.Title,
		Before_code: record.Before_code,
		After_code:  record.After_code,
		Review:      record.Review,
		Memo:        record.Memo,
	}
	// コンテンツテーブルへの挿入
	id, err := c.ContentRepository.InsertContent(recordContent)
	if err != nil {
		return fmt.Errorf("failed to InsertContent in ContentCreate: %w", err)
	}

	// キーワードテーブルへの挿入
	if err := c.KeywordRepository.InsertKeyword(id, record.Keywords); err != nil {
		return fmt.Errorf("failed to InsertKeyword in ContentCreate: %w", err)
	}

	return nil
}

/* コンテンツ更新ロジック
func (c *ContentController) ContentUpdate(record *Content_keywords) error {
	// コンテンツテーブルの更新

	// キーワード削除
	// キーワード追加
}*/

// コンテンツ削除ロジック
func (c *ContentController) ContentDelete(id int) error {
	if err := c.ContentRepository.DeleteContentByContentID(id); err != nil {
		return fmt.Errorf("failed to DeleteContentByContentID in ContentDelete: %w", err)
	}
	return nil
}
