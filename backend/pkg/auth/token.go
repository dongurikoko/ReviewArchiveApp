package auth

import (
	"fmt"
	"os"
	"reviewArchive/pkg/server/model"
	"time"

	"github.com/golang-jwt/jwt"
)

// GetTokenHandler JWTトークンを生成
func GetTokenHandler(user *model.User) (string, error) {
	// headerのセット
	token := jwt.New(jwt.SigningMethodHS256)

	// claimsのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.UserID
	claims["name"] = user.Username
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //有効期限を指定

	// 署名の生成
	tokenString, err := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))
	if err != nil {
		return "", fmt.Errorf("failed to SignedString in GetTokenHandler: %w", err)
	}

	return tokenString, nil
}
