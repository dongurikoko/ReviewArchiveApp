package controller

import (
	"fmt"

	"reviewArchive/pkg/db"
	"reviewArchive/pkg/server/model"
)

type ContentRequest struct {
	Title       string
	BeforeCode string
	AfterCode  string
	Review      string
	Memo        string
	Keywords    []string
}

type ContentController struct {
	ContentRepository model.ContentRepositoryInterface
	KeywordRepository model.KeywordRepositoryInterface
	TaggingRepository model.TaggingRepositoryInterface
	UserRepository model.UserRepositoryInterface
}

func NewContentController(contentRepository model.ContentRepositoryInterface,
	keywordRepository model.KeywordRepositoryInterface,taggingRepository model.TaggingRepositoryInterface,
	userRepository model.UserRepositoryInterface) *ContentController {
	return &ContentController{
		ContentRepository: contentRepository,
		KeywordRepository: keywordRepository,
		TaggingRepository: taggingRepository,
		UserRepository: userRepository,
	}
}

type ContentControllerInterface interface {
	ContentCreate(record *ContentRequest) error
	ContentDelete(id int) error
	ContentUpdate(id int, record *ContentRequest) error
}

// コンテンツ作成ロジック
func (c *ContentController) ContentCreate(record *ContentRequest,uid string) error {
	// uidを元にuserIDを取得(無い場合は新規登録)
	userID, err := c.UserRepository.SelectUserIDByUID(uid)
	if err != nil {
		return fmt.Errorf("failed to SelectUserIDByUID in ContentCreate: %w", err)
	}
	
	conn,err := db.GetConn()
	if err != nil {
		return fmt.Errorf("failed to GetConn in ContentCreate: %w", err)
	}	    
	// トランザクション開始
	tx, err := conn.Begin()
	if err != nil {
		return fmt.Errorf("failed to Begin in ContentCreate: %w", err)
	}

	// ロールバックの準備
    defer func() {
        if err != nil {
            tx.Rollback()
        } else {
            err = tx.Commit()
            if err != nil {
                tx.Rollback()
            }
        }
    }()

	// コンテンツテーブルへの挿入
	recordContent := &model.Content{
		Title:       record.Title,
		BeforeCode:  record.BeforeCode,
		AfterCode:   record.AfterCode,
		Review:      record.Review,
		Memo:        record.Memo,
		UserID:      userID,
	}

	contentID, err := c.ContentRepository.InsertContent(recordContent,tx)
	if err != nil {
		return fmt.Errorf("failed to InsertContent in ContentCreate: %w", err)
	}

	// キーワードテーブルへの挿入
	for _, keyword := range record.Keywords {
		// keywordを元にkeywordIDを取得(無い場合は新規登録)
		keywordID, err := c.KeywordRepository.SelectKeywordIDByKeyword(keyword,tx)
		if err != nil {
			return fmt.Errorf("failed to InsertKeyword in ContentCreate: %w", err)
		}
		// タグテーブルへの挿入
		tagging := &model.Tagging{
			ContentID: contentID,
			KeywordID: keywordID,
		}
		if err := c.TaggingRepository.InsertTagging(tagging,tx); err != nil {
			return fmt.Errorf("failed to InsertTagging in ContentCreate: %w", err)
		}
	}

	return nil
}

// コンテンツ更新ロジック
func (c *ContentController) ContentUpdate(id int, record *ContentRequest) error {
	recordContent := &model.Content{
		Title:       record.Title,
		BeforeCode: record.BeforeCode,
		AfterCode:  record.AfterCode,
		Review:      record.Review,
		Memo:        record.Memo,
	}
	// コンテンツテーブルの更新
	if err := c.ContentRepository.UpdateContentByContentID(id, recordContent); err != nil {
		return fmt.Errorf("failed to UpdateContentByContentID in ContentUpdate: %w", err)
	}
	// キーワード削除
	if err := c.KeywordRepository.DeleteKeywordByID(id); err != nil {
		return fmt.Errorf("failed to DeleteKeywordByID in ContentUpdate: %w", err)
	}
	// キーワード追加
	if err := c.KeywordRepository.InsertKeyword(id, record.Keywords); err != nil {
		return fmt.Errorf("failed to InsertKeyword in ContentUpdate: %w", err)
	}
	return nil
}

// コンテンツ削除ロジック
func (c *ContentController) ContentDelete(id int) error {
	if err := c.ContentRepository.DeleteContentByContentID(id); err != nil {
		return fmt.Errorf("failed to DeleteContentByContentID in ContentDelete: %w", err)
	}
	return nil
}
