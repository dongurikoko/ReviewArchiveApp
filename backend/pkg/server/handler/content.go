package handler

import (
	"fmt"
	"net/http"
	"strconv"

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
			return fmt.Errorf("failed to bind request in HandleContentCreate: %w", err)
		}

		// contextからuserIDを取得
		userID := c.Get("userid").(string)

		if err := h.ContentController.ContentCreate(req,userID); err != nil {
			return fmt.Errorf("failed to ContentCreate in HandleContentCreate: %w", err)
		}
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Success to Create Content",
		})

	}
}

// コンテンツのアップデート処理
func (h *ContentHandler) HandleContentUpdate() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &controller.ContentRequest{}
		if err := c.Bind(req); err != nil {
			return fmt.Errorf("failed to bind request in HandleContentUpdate: %w", err)
		}
		// URLパラメータからcontent_idを取得
		contentID, err := strconv.Atoi(c.Param("content_id"))
		if err != nil {
			return fmt.Errorf("failed to get contentID in HandleContentUpdate: %w", err)
		}
		if err := h.ContentController.ContentUpdate(contentID, req); err != nil {
			return fmt.Errorf("failed to ContentUpdate in HandleContentUpdate: %w", err)
		}
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Success to Update Content",
		})
	}
}

// コンテンツの削除処理(contentテーブルを削除するとkeywordも自動で削除される)
func (h *ContentHandler) HandleContentDelete() echo.HandlerFunc {
	return func(c echo.Context) error {
		// URLパラメータからcontent_idを取得
		contentID, err := strconv.Atoi(c.Param("content_id"))

		//fmt.Printf("contentID: %v\n", contentID)

		if err != nil {
			return fmt.Errorf("failed to get contentID in HandleContentDelete: %w", err)
		}
		// コンテンツテーブルから削除
		if err := h.ContentController.ContentDelete(contentID); err != nil {
			return fmt.Errorf("failed to ContentDelete in HandleContentDelete: %w", err)
		}
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Success to Delete Content",
		})
	}

}
