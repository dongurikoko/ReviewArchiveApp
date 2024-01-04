package handler

import (
	"fmt"
	"log"
	"net/http"

	"reviewArchive/pkg/server/controller"

	"github.com/labstack/echo/v4"
)

type ContentHandler struct {
	ContentController controller.ContentControllerInterface
}

func NewContentHandler(contentController controller.ContentControllerInterface) *ContentHandler {
	return &ContentHandler{
		ContentController: contentController,
	}
}

// コンテンツ作成処理
func (h *ContentHandler) HandleContentCreate() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &controller.ContentRequest{}
		if err := c.Bind(req); err != nil {
			log.Printf("%v", err)
			return fmt.Errorf("failed to bind request in HandleContentCreate: %w", err)
		}
		if err := h.ContentController.ContentCreate(req); err != nil {
			log.Printf("%v", err)
			return fmt.Errorf("failed to ContentCreate in HandleContentCreate: %w", err)
		}
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Success to Create Content",
		})

	}
}

/* コンテンツのアップデート処理
func (h *ContentHandler) HandleContentUpdate() echo.HandlerFunc {
	return func(c echo.Context) error {
	}
}

// コンテンツの削除処理
func (h *ContentHandler) HandleContentDelete() echo.HandlerFunc {
	return func(c echo.Context) error {
	}
}*/
