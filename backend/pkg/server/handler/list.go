package handler

import (
	"fmt"
	"net/http"
	"reviewArchive/pkg/server/controller"

	"github.com/labstack/echo/v4"
)

type content struct {
	Content_id int      `json:"content_id"`
	Title      string   `json:"title"`
	Keywords   []string `json:"keywords"`
}

type contentlist struct {
	Contents []content `json:"contents"`
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
		lists, err := h.ListController.ListGet()
		if err != nil {
			return fmt.Errorf("failed to ListGet in HandleListGet: %w", err)
		}
		var response contentlist
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
