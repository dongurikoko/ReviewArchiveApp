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

		// contextからUIDを取得
		uid := c.Get("uid").(string)
		
		if err := h.ContentController.ContentCreate(req,uid); err != nil {
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
		// URLパラメータからcontentIDを取得
		contentID, err := strconv.Atoi(c.Param("content_id"))
		if err != nil {
			return fmt.Errorf("failed to get contentID in HandleContentUpdate: %w", err)
		}

		// contextからUIDを取得
		uid := c.Get("uid").(string)

		if err := h.ContentController.ContentUpdate(contentID, req, uid); err != nil {
			return fmt.Errorf("failed to ContentUpdate in HandleContentUpdate: %w", err)
		}
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Success to Update Content",
		})
	}
}

// コンテンツの削除処理
func (h *ContentHandler) HandleContentDelete() echo.HandlerFunc {
	return func(c echo.Context) error {
		// URLパラメータからcontent_idを取得
		contentID, err := strconv.Atoi(c.Param("content_id"))

		//fmt.Printf("contentID: %v\n", contentID)

		if err != nil {
			return fmt.Errorf("failed to get contentID in HandleContentDelete: %w", err)
		}
		// コンテンツテーブルとtaggingテーブルから削除
		if err := h.ContentController.DeleteByContentID(contentID); err != nil {
			return fmt.Errorf("failed to DeleteByContentID in HandleContentDelete: %w", err)
		}

		return c.JSON(http.StatusOK, echo.Map{
			"message": "Success to Delete Content",
		})
	}

}
