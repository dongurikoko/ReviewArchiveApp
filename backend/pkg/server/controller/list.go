package controller

import (
	"fmt"
	"reviewArchive/pkg/server/model"
)

type ListResponse struct {
	ContentID int
	Title      string
	Keywords   []string
}

type ListController struct {
	ContentRepository model.ContentRepositoryInterface
	KeywordRepository model.KeywordRepositoryInterface
	UserRepository model.UserRepositoryInterface
}

func NewListController(contentRepository model.ContentRepositoryInterface,
	keywordRepository model.KeywordRepositoryInterface,userRepository model.UserRepositoryInterface) *ListController {
	return &ListController{
		ContentRepository: contentRepository,
		KeywordRepository: keywordRepository,
		UserRepository: userRepository,
	}
}

type ListControllerInterface interface {
	GetAllContents() ([]*ListResponse, error)
	GetContentsByContentID(ID int) (*ContentRequest, error)
	SearchContents(keyword string,uid string) ([]*ListResponse, error)
}

// コンテンツの一覧取得ロジック
func (c *ListController) GetAllContents() ([]*ListResponse, error) {
	// コンテンツテーブルから全てのレコードを取得
	contentlists, err := c.ContentRepository.SelectContent()
	if err != nil {
		return nil, fmt.Errorf("failed to SelectContent in ListGet: %w", err)
	}

	var listResponses []*ListResponse
	for _, contentlist := range contentlists {
		keywords, err := c.KeywordRepository.SelectStringKeywordByID(contentlist.ContentID)
		if err != nil {
			return nil, fmt.Errorf("failed to SelectKeywordByID in ListGet: %w", err)
		}
		listResponse := &ListResponse{
			ContentID:  contentlist.ContentID,
			Title:      contentlist.ContentValue.Title,
			Keywords:   keywords,
		}
		listResponses = append(listResponses, listResponse)
	}

	return listResponses, nil

}

// 特定のコンテンツの詳細取得ロジック
func (c *ListController) GetContentsByContentID(ID int) (*ContentRequest, error) {
	// contentテーブルからIDを条件にレコードを取得
	content, err := c.ContentRepository.SelectContentByContentID(ID)

	if err != nil {
		return nil, fmt.Errorf("failed to SelectContentByContentID in ListGetByContentID: %w", err)
	}

	// keywordテーブルからIDを条件にレコードを取得
	keyword, err := c.KeywordRepository.SelectStringKeywordByID(ID)

	if err != nil {
		return nil, fmt.Errorf("failed to SelectStringKeywordByID in ListGetByContentID: %w", err)
	}

	// contentテーブルとkeywordテーブルのレコードを結合
	return &ContentRequest{
		Title:       content.Title,
		BeforeCode: content.BeforeCode,
		AfterCode:  content.AfterCode,
		Review:      content.Review,
		Memo:        content.Memo,
		Keywords:    keyword,
	}, nil
}

// コンテンツ検索ロジック
func (c *ListController) SearchContents(keyword string,uid string) ([]*ListResponse, error) {
	// uidを元にuserIDを取得
	userID, err := c.UserRepository.SelectUserIDByUIDWithError(uid)
	if err != nil {
		return nil, fmt.Errorf("failed to SelectUserIDByUIDWithError in SearchContents: %w", err)
	}

	// Contents,Tagging,Keywordテーブルを結合し、userID,keywordがそれぞれ一致するコンテンツを取得
	contentlists, err := c.ContentRepository.SelectContentByKeywordsAndUserID(keyword,userID)
	if err != nil {
		return nil, fmt.Errorf("failed to SelectContentByKeywordsAndUserID in SearchContents: %w", err)
	}

	var listResponses []*ListResponse
	for _, contentlist := range contentlists {
		// contentIDを元に紐づくキーワードを取得
		keywords, err := c.KeywordRepository.SelectStringKeywordByID(contentlist.ContentID)
		if err != nil {
			return nil, fmt.Errorf("failed to SelectKeywordByID in SearchContents: %w", err)
		}
		listResponse := &ListResponse{
			ContentID:  contentlist.ContentID,
			Title:      contentlist.ContentValue.Title,
			Keywords:   keywords,
		}
		listResponses = append(listResponses, listResponse)
	}

	return listResponses, nil
}
