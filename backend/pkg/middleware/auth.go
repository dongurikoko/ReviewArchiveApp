package middleware

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

func AuthenticateMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 環境変数CREDENTIALSの値をログに出力
			credentialsPath := os.Getenv("CREDENTIALS")
			log.Printf("Using CREDENTIALS file: %s\n", credentialsPath)

			// Firebase SDKのセットアップ
			opt := option.WithCredentialsFile(credentialsPath)
			// firebaseアプリケーションを初期化
			app, err := firebase.NewApp(context.Background(), nil, opt)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "failed to initialize firebase app")
			}

			// firebase authを初期化
			auth,err := app.Auth(context.Background())
			if err != nil{
				return echo.NewHTTPError(http.StatusInternalServerError, "failed to initialize firebase auth")
			}

			// クライアントから送られてきたJWT取得
			authHeader := c.Request().Header.Get("Authorization") // 認証に使用されるトークンは、"Bearer [token]" の形式で格納
			idToken := strings.Replace(authHeader, "Bearer ", "", 1)

			// JWTの検証
			token, err := auth.VerifyIDToken(context.Background(), idToken)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
			}

			log.Printf("Verified ID Token: %v\n\n",token)

			// JWTからUIDを取得
			// uid := token.Claims["sub"]

			//log.Printf("uid: %v\n\n",uid)

			uid := token.Claims["user_id"]

			// UIDをコンテキストにセット
			c.Set("uid", uid)

			//log.Printf("uid: %v\n",uid)

			// 次の処理
			return next(c)
		}
	}
}
