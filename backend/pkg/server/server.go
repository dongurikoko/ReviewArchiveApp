package server

import (
	"log"

	"reviewArchive/pkg/db"
	"reviewArchive/pkg/server/controller"
	"reviewArchive/pkg/server/handler"
	"reviewArchive/pkg/server/model"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

var (
	sqlDB, _ = db.GetConn()

	//userRepository = model.NewUserRepository(sqlDB) //userテーブルの部分
	//authMiddleware = middleware.NewMiddleware(userRepository)

	contentRepository = model.NewContentRepository(sqlDB)
	keywordRepository = model.NewKeywordRepository(sqlDB)

	contentController = controller.NewContentController(contentRepository, keywordRepository)
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
		AllowHeaders: []string{"Content-Type,Accept,Origin,x-token"},
	}))

	/* ===== URLマッピングを行う ===== */
	// 認証を必要としないAPI
	e.POST("/contents", contentHandler.HandleContentCreate())
	e.POST("/contents/:content_id", contentHandler.HandleContentUpdate())
	e.DELETE("/contents/:content_id", contentHandler.HandleContentDelete())
	e.GET("/lists", listHandler.HandleListGet())
	e.GET("/lists/:content_id", listHandler.HandleListGetByContentID())
	e.GET("/lists/search", listHandler.HandleListSearch())

	/* ===== サーバの起動 ===== */
	log.Println("Server running...")
	if err := e.Start(addr); err != nil {
		log.Fatalf("failed to start server. %+v", err)
	}
}
