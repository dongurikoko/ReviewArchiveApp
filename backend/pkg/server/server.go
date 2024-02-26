package server

import (
	"log"

	"reviewArchive/pkg/db"
	"reviewArchive/pkg/middleware"
	"reviewArchive/pkg/server/controller"
	"reviewArchive/pkg/server/handler"
	"reviewArchive/pkg/server/model"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

var (
	sqlDB, _ = db.GetConn()

	userRepository = model.NewUserRepository(sqlDB)
	authMiddleware = middleware.NewMiddleware(userRepository)

	contentRepository = model.NewContentRepository(sqlDB)
	keywordRepository = model.NewKeywordRepository(sqlDB)

	contentController = controller.NewContentController(contentRepository, keywordRepository)
	listController    = controller.NewListController(contentRepository, keywordRepository)

	contentHandler = handler.NewContentHandler(contentController)
	listHandler    = handler.NewListHandler(listController)
	userHandler    = handler.NewUserHandler(userRepository)
)

// Serve HTTPサーバを起動する
func Serve(addr string) {
	e := echo.New()
	// panicが発生した場合の処理
	e.Use(echomiddleware.Recover())
	// CORSの設定
	e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		Skipper:      echomiddleware.DefaultCORSConfig.Skipper,
		AllowOrigins: echomiddleware.DefaultCORSConfig.AllowOrigins,
		AllowMethods: echomiddleware.DefaultCORSConfig.AllowMethods,
		AllowHeaders: []string{"Content-Type,Accept,Origin,x-token"},
	}))

	/* ===== URLマッピングを行う ===== */
	// 認証を必要としないAPI
	e.POST("/signup", userHandler.HandleUserCreate())
	e.POST("/login", userHandler.HandleUserLogin())

	// 認証を必要とするAPI
	authAPI := e.Group("", authMiddleware.AuthenticateMiddleware())
	authAPI.GET("user/get", userHandler.HandleUserGet())
	authAPI.POST("/content/create", contentHandler.HandleContentCreate())
	authAPI.POST("/content/update/:content_id", contentHandler.HandleContentUpdate())
	authAPI.DELETE("/content/delete/:content_id", contentHandler.HandleContentDelete())
	authAPI.GET("/list/get", listHandler.HandleListGet())
	authAPI.GET("/list/get/:content_id", listHandler.HandleListGetByContentID())
	authAPI.GET("/list/search", listHandler.HandleListSearch())

	/* ===== サーバの起動 ===== */
	log.Println("Server running...")
	if err := e.Start(addr); err != nil {
		log.Fatalf("failed to start server. %+v", err)
	}
}
