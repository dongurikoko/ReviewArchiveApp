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

	userRepository = model.NewUserRepository(sqlDB) //userテーブルの部分
	taggingRepository = model.NewTaggingRepository(sqlDB)
	//authMiddleware = middleware.NewMiddleware(userRepository)

	contentRepository = model.NewContentRepository(sqlDB)
	keywordRepository = model.NewKeywordRepository(sqlDB)

	contentController = controller.NewContentController(contentRepository, keywordRepository,taggingRepository,userRepository)
	listController    = controller.NewListController(contentRepository, keywordRepository)

	contentHandler = handler.NewContentHandler(contentController)
	listHandler    = handler.NewListHandler(listController)
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
		AllowHeaders: []string{"Content-Type,Accept,Origin,Authorization,x-token"},
	}))

	/* ===== URLマッピングを行う ===== */
	authAPI := e.Group("", middleware.AuthenticateMiddleware())
	// 認証を必要とするAPI
	authAPI.POST("/contents", contentHandler.HandleContentCreate())
	authAPI.POST("/contents/:content_id", contentHandler.HandleContentUpdate())
	authAPI.DELETE("/contents/:content_id", contentHandler.HandleContentDelete())
	authAPI.GET("/lists/:content_id", listHandler.HandleListGetByContentID())
	authAPI.GET("/lists/search", listHandler.HandleListSearch())

	/* ===== サーバの起動 ===== */
	log.Println("Server running...")
	if err := e.Start(addr); err != nil {
		log.Fatalf("failed to start server. %+v", err)
	}
}
