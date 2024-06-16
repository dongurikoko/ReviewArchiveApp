package controller

import (
	"fmt"
	"log"

	"reviewArchive/pkg/db"
	"reviewArchive/pkg/server/model"
)

type ContentRequest struct {
	Title      string
	BeforeCode string
	AfterCode  string
	Review     string
	Memo       string
	Keywords   []string
}

type ContentController struct {
	ContentRepository model.ContentRepositoryInterface
	KeywordRepository model.KeywordRepositoryInterface
	TaggingRepository model.TaggingRepositoryInterface
	UserRepository    model.UserRepositoryInterface
}

func NewContentController(contentRepository model.ContentRepositoryInterface,
	keywordRepository model.KeywordRepositoryInterface, taggingRepository model.TaggingRepositoryInterface,
	userRepository model.UserRepositoryInterface) *ContentController {
	return &ContentController{
		ContentRepository: contentRepository,
		KeywordRepository: keywordRepository,
		TaggingRepository: taggingRepository,
		UserRepository:    userRepository,
	}
}

type ContentControllerInterface interface {
	ContentCreate(record *ContentRequest, uid string) error
	DeleteByContentID(contentID int) error
	ContentUpdate(contentID int, record *ContentRequest, uid string) error
}

// コンテンツ作成ロジック
func (c *ContentController) ContentCreate(record *ContentRequest, uid string) error {
	// uidを元にuserIDを取得(無い場合は新規登録)
	userID, err := c.UserRepository.SelectUserIDByUID(uid)
	if err != nil {
		return fmt.Errorf("failed to SelectUserIDByUID in ContentCreate: %w", err)
	}

	conn, err := db.GetConn()
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
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("rollback error: %v", rbErr)
			}
		} else {
			if err = tx.Commit(); err != nil {
				if rbErr := tx.Rollback(); rbErr != nil {
					log.Printf("rollback error: %v", rbErr)
				}
			}
		}
	}()

	// コンテンツテーブルへの挿入
	recordContent := &model.Content{
		Title:      record.Title,
		BeforeCode: record.BeforeCode,
		AfterCode:  record.AfterCode,
		Review:     record.Review,
		Memo:       record.Memo,
		UserID:     userID,
	}

	contentID, err := c.ContentRepository.InsertContent(recordContent, tx)
	if err != nil {
		return fmt.Errorf("failed to InsertContent in ContentCreate: %w", err)
	}

	// キーワードテーブルへの挿入
	for _, keyword := range record.Keywords {
		// keywordを元にkeywordIDを取得(無い場合は新規登録)
		keywordID, err := c.KeywordRepository.SelectKeywordIDByKeyword(keyword, tx)
		if err != nil {
			return fmt.Errorf("failed to InsertKeyword in ContentCreate: %w", err)
		}
		// タグテーブルへの挿入
		tagging := &model.Tagging{
			ContentID: contentID,
			KeywordID: keywordID,
		}
		if err := c.TaggingRepository.InsertTagging(tagging, tx); err != nil {
			return fmt.Errorf("failed to InsertTagging in ContentCreate: %w", err)
		}
	}

	return nil
}

// コンテンツ更新ロジック
func (c *ContentController) ContentUpdate(contentID int, record *ContentRequest, uid string) error {
	// uidを元にuserIDを取得(無い場合はエラー)
	userID, err := c.UserRepository.SelectUserIDByUIDWithError(uid)
	if err != nil {
		return fmt.Errorf("failed to SelectUserIDByUIDWithError in ContentUpdate: %w", err)
	}

	// コンテンツIDに対応するコンテンツが存在するか確認
	content, err := c.ContentRepository.SelectContentByContentID(contentID)
	if err != nil {
		return fmt.Errorf("failed to SelectContentByContentID in ContentUpdate: %w", err)
	}
	// contentのUserIDとuidのUserIDが一致するか確認
	if content.UserID != userID {
		return fmt.Errorf("userID is not matched in ContentUpdate: %w", err)
	}

	conn, err := db.GetConn()
	if err != nil {
		return fmt.Errorf("failed to GetConn in ContentCreate: %w", err)
	}
	// トランザクション開始
	tx, err := conn.Begin()
	if err != nil {
		return fmt.Errorf("failed to Begin in ContentCreate: %w", err)
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("rollback error: %v", rbErr)
			}
		} else {
			if err = tx.Commit(); err != nil {
				if rbErr := tx.Rollback(); rbErr != nil {
					log.Printf("rollback error: %v", rbErr)
				}
			}
		}
	}()

	// コンテンツテーブルの更新
	recordContent := &model.Content{
		Title:      record.Title,
		BeforeCode: record.BeforeCode,
		AfterCode:  record.AfterCode,
		Review:     record.Review,
		Memo:       record.Memo,
		UserID:     userID,
	}

	if err := c.ContentRepository.UpdateContentByContentID(contentID, recordContent, tx); err != nil {
		return fmt.Errorf("failed to UpdateContentByContentID in ContentUpdate: %w", err)
	}

	// タグテーブルの削除
	if err := c.TaggingRepository.DeleteTaggingByContentID(contentID, tx); err != nil {
		return fmt.Errorf("failed to DeleteTaggingByContentID in ContentUpdate: %w", err)
	}

	// キーワードテーブルへの挿入
	for _, keyword := range record.Keywords {
		// keywordを元にkeywordIDを取得(無い場合は新規登録)
		keywordID, err := c.KeywordRepository.SelectKeywordIDByKeyword(keyword, tx)
		if err != nil {
			return fmt.Errorf("failed to InsertKeyword in ContentUpdate: %w", err)
		}
		// タグテーブルへの挿入
		tagging := &model.Tagging{
			ContentID: contentID,
			KeywordID: keywordID,
		}
		if err := c.TaggingRepository.InsertTagging(tagging, tx); err != nil {
			return fmt.Errorf("failed to InsertTagging in ContentUpdate: %w", err)
		}
	}
	return nil
}

// コンテンツ削除ロジック
func (c *ContentController) DeleteByContentID(contentID int) error {
	conn, err := db.GetConn()
	if err != nil {
		return fmt.Errorf("failed to GetConn in ContentCreate: %w", err)
	}
	// トランザクション開始
	tx, err := conn.Begin()
	if err != nil {
		return fmt.Errorf("failed to Begin in ContentCreate: %w", err)
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("rollback error: %v", rbErr)
			}
		} else {
			if err = tx.Commit(); err != nil {
				if rbErr := tx.Rollback(); rbErr != nil {
					log.Printf("rollback error: %v", rbErr)
				}
			}
		}
	}()

	// タグテーブルの削除
	if err := c.TaggingRepository.DeleteTaggingByContentID(contentID, tx); err != nil {
		return fmt.Errorf("failed to DeleteTaggingByContentID in DeleteByContentID: %w", err)
	}

	// コンテンツテーブルの削除
	if err := c.ContentRepository.DeleteContentByContentID(contentID, tx); err != nil {
		return fmt.Errorf("failed to DeleteContentByContentID in DeleteByContentID: %w", err)
	}

	return nil
}
