package handler

import (
	"fmt"
	"net/http"
	"reviewArchive/pkg/server/controller"
	"strconv"

	"github.com/labstack/echo/v4"
)

type content struct {
	ContentID  int      `json:"content_id"`
	Title      string   `json:"title"`
	Keywords   []string `json:"keywords"`
}

type alllist struct {
	Contents []content `json:"contents"`
}

type listbycontentid struct {
	ContentID  int      `json:"content_id"`
	Title       string   `json:"title"`
	BeforeCode string   `json:"before_code"`
	AfterCode  string   `json:"after_code"`
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
				ContentID: list.ContentID,
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
		contentID, err := strconv.Atoi(c.Param("content_id"))
		if err != nil {
			return fmt.Errorf("failed to get contentID in HandleListGetByContentID: %w", err)
		}

		//fmt.Printf("contentID: %v\n", contentID)

		content, err := h.ListController.GetContentsByContentID(contentID)
		if err != nil {
			return fmt.Errorf("failed to ListGetByContentID in HandleListGetByContentID: %w", err)
		}
		response := listbycontentid{
			ContentID:  contentID,
			Title:       content.Title,
			BeforeCode: content.BeforeCode,
			AfterCode:  content.AfterCode,
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
		
		// contextからUIDを取得
		uidwithnil := c.Get("uid")

		// uidがnilなら何も表示しない
		if uidwithnil == nil {
			return c.JSON(http.StatusOK, &alllist{})
		}

		uid := uidwithnil.(string)

		// keywordが一致するコンテンツを一覧取得
		lists, err := h.ListController.SearchContents(keyword,uid)
		if err != nil {
			return fmt.Errorf("failed to SearchContents in HandleListSearch: %w", err)
		}

		var response alllist
		for _, list := range lists {
			response.Contents = append(response.Contents, content{
				ContentID: list.ContentID,
				Title:     list.Title,
				Keywords:  list.Keywords,
			})
		}
		return c.JSON(http.StatusOK, &response)
	}
}
