package handler

import (
	"fmt"
	"net/http"
	"reviewArchive/pkg/server/controller"
	"strconv"

	"github.com/labstack/echo/v4"
)

type content struct {
	Content_id int      `json:"content_id"`
	Title      string   `json:"title"`
	Keywords   []string `json:"keywords"`
}

type alllist struct {
	Contents []content `json:"contents"`
}

type listbycontentid struct {
	Content_id  int      `json:"content_id"`
	Title       string   `json:"title"`
	Before_code string   `json:"before_code"`
	After_code  string   `json:"after_code"`
	Review      string   `json:"review"`
	Memo        string   `json:"memo"`
	Keywords    []string `json:"keywords"`
}

type ListHandler struct {
	ListController controller.ListControllerInterface
}

func NewListHandler(listController controller.ListControllerInterface) *ListHandler {
	return &ListHandler{
		ListController: listController,
	}
}

// コンテンツの一覧取得処理
func (h *ListHandler) HandleListGet() echo.HandlerFunc {
	return func(c echo.Context) error {
		lists, err := h.ListController.GetAllContents()
		if err != nil {
			return fmt.Errorf("failed to ListGet in HandleListGet: %w", err)
		}
		var response alllist
		for _, list := range lists {
			response.Contents = append(response.Contents, content{
				Content_id: list.Content_id,
				Title:      list.Title,
				Keywords:   list.Keywords,
			})
		}

		return c.JSON(http.StatusOK, &response)
	}
}

// 特定のコンテンツの詳細取得処理
func (h *ListHandler) HandleListGetByContentID() echo.HandlerFunc {
	return func(c echo.Context) error {
		// URLパラメータからcontent_idを取得
		content_id, err := strconv.Atoi(c.Param("content_id"))
		if err != nil {
			return fmt.Errorf("failed to get content_id in HandleListGetByContentID: %w", err)
		}

		//fmt.Printf("content_id: %v\n", content_id)

		content, err := h.ListController.GetContentsByContentID(content_id)
		if err != nil {
			return fmt.Errorf("failed to ListGetByContentID in HandleListGetByContentID: %w", err)
		}
		response := listbycontentid{
			Content_id:  content_id,
			Title:       content.Title,
			Before_code: content.Before_code,
			After_code:  content.After_code,
			Review:      content.Review,
			Memo:        content.Memo,
			Keywords:    content.Keywords,
		}

		return c.JSON(http.StatusOK, &response)
	}
}

// キーワードと一致するコンテンツの一覧取得処理
func (h *ListHandler) HandleListSearch() echo.HandlerFunc {
	return func(c echo.Context) error {
		// クエリパラメータからkeywordを取得
		keyword := c.QueryParam("keyword")
		// keywordが一致するコンテンツを取得
		lists, err := h.ListController.SearchContents(keyword)
		if err != nil {
			return fmt.Errorf("failed to SearchContents in HandleListSearch: %w", err)
		}

		var response alllist
		for _, list := range lists {
			response.Contents = append(response.Contents, content{
				Content_id: list.Content_id,
				Title:      list.Title,
				Keywords:   list.Keywords,
			})
		}
		return c.JSON(http.StatusOK, &response)
	}
}
