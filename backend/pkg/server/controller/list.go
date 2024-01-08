package controller

import (
	"fmt"
	"reviewArchive/pkg/server/model"
)

type ListResponse struct {
	Content_id int
	Title      string
	Keywords   []string
}

type ListController struct {
	ContentRepository model.ContentRepositoryInterface
	KeywordRepository model.KeywordRepositoryInterface
}

func NewListContoroller(contentRepository model.ContentRepositoryInterface,
	keywordRepository model.KeywordRepositoryInterface) *ListController {
	return &ListController{
		ContentRepository: contentRepository,
		KeywordRepository: keywordRepository,
	}
}

type ListControllerInterface interface {
	ListGet() ([]*ListResponse, error)
	ListGetByContentID(ID int) (*ContentRequest, error)
}

// コンテンツの一覧取得ロジック
func (c *ListController) ListGet() ([]*ListResponse, error) {
	// コンテンツテーブルから全てのレコードを取得
	contentlists, err := c.ContentRepository.SelectContent()
	if err != nil {
		return nil, fmt.Errorf("failed to SelectContent in ListGet: %w", err)
	}

	var listResponses []*ListResponse
	for _, contentlist := range contentlists {
		// コンテンツテーブルから取得したレコードのIDを元にキーワードテーブルからレコードを取得
		keywords, err := c.KeywordRepository.SelectStringKeywordByID(contentlist.Content_id)
		if err != nil {
			return nil, fmt.Errorf("failed to SelectKeywordByID in ListGet: %w", err)
		}
		// コンテンツテーブルから取得したレコードとキーワードテーブルから取得したレコードを結合
		listResponse := &ListResponse{
			Content_id: contentlist.Content_id,
			Title:      contentlist.Content_value.Title,
			Keywords:   keywords,
		}
		listResponses = append(listResponses, listResponse)
	}

	return listResponses, nil

}

// 特定のコンテンツの詳細取得ロジック
func (c *ListController) ListGetByContentID(ID int) (*ContentRequest, error) {
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
		Before_code: content.Before_code,
		After_code:  content.After_code,
		Review:      content.Review,
		Memo:        content.Memo,
		Keywords:    keyword,
	}, nil
}
