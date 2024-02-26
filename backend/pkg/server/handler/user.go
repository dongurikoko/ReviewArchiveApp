package handler

import (
	"fmt"
	"log"
	"net/http"
	"reviewArchive/pkg/auth"
	"reviewArchive/pkg/server/model"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserCreateRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// UserHandler ユーザー関連のリクエストを処理するハンドラー
type UserHandler struct {
	UserRepository model.UserRepositoryInterface
}

// NewUserHandler 新しいUserHandlerを作成する
func NewUserHandler(userRepository model.UserRepositoryInterface) *UserHandler {
	return &UserHandler{
		UserRepository: userRepository,
	}
}

// HandleUserCreate ユーザー作成処理
func (h *UserHandler) HandleUserCreate() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &UserCreateRequest{}
		if err := c.Bind(req); err != nil {
			return fmt.Errorf("failed to bind request in HandleUserCreate: %w", err)
		}

		// パスワードのハッシュ化
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("failed to GenerateFromPassword in HandleUserCreate: %w", err)
		}

		// ユーザーテーブルに登録
		if err := h.UserRepository.InsertUser(&model.User{
			Username: req.UserName,
			Password: string(hashedPassword),
			Email:    req.Email,
		}); err != nil {
			return fmt.Errorf("failed to InsertUser in HandleUserCreate: %w", err)
		}

		return c.JSON(http.StatusOK, echo.Map{
			"message": "Success to Create User",
		})
	}
}

type UserLoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// ユーザログイン機能(ユーザ名とパスワードを受け取って認証トークンを返す)
func (h *UserHandler) HandleUserLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &UserLoginRequest{}
		if err := c.Bind(req); err != nil {
			return fmt.Errorf("failed to bind request in HandleUserLogin: %w", err)
		}

		// ユーザ名からユーザ情報を取得
		user, err := h.UserRepository.SelectUserByUsername(req.UserName)
		if err != nil {
			return fmt.Errorf("failed to SelectUserByUsername in HandleUserLogin: %w", err)
		}
		if user == nil {
			return fmt.Errorf("user not found. username=%s", req.UserName)
		}

		// パスワードの照合
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			return fmt.Errorf("password is wrong: %w", err)
		}

		log.Println("login success: username = ", req.UserName)

		//JWT認証を用いて認証トークンを生成
		token, err := auth.GetTokenHandler(user)

	}
}
