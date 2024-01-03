package server

import (
	"log"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"

	"go-college/pkg/db"
	"go-college/pkg/http/middleware"
	"go-college/pkg/server/handler"
	"go-college/pkg/server/model"
	"go-college/pkg/server/service"

	"github.com/redis/go-redis/v9"
)

var (
	sqlDB, _ = db.GetConn()
	//httpResponse = response.NewHttpResponse()
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redisサーバーのアドレス
		Password: "",               // パスワードがある場合は設定
		DB:       0,                // 使用するDB番号
	})
	redisCtx = db.GetRedisCtx()

	userRepository = model.NewUserRepository(sqlDB) //userテーブルの部分
	authMiddleware = middleware.NewMiddleware(userRepository)

	gachaProbabilityRepository   = model.NewGachaProbabilityRepository(sqlDB) //gachaProbabilityテーブルの部分
	userCollectionItemRepository = model.NewUserCollectionItemRepository(sqlDB)
	collectionItemRepository     = model.NewCollectionItemRepository(sqlDB)
	redisRepository              = model.NewRedisRepository(redisClient, redisCtx)

	gameService       = service.NewGameService(userRepository)
	gachaService      = service.NewGachaService(userRepository, gachaProbabilityRepository, userCollectionItemRepository, collectionItemRepository)
	rankingService    = service.NewRankingService(userRepository, redisRepository)
	collectionService = service.NewCollectionService(userCollectionItemRepository, collectionItemRepository, redisRepository)

	userHandler       = handler.NewUserHandler(userRepository)
	settingHandler    = handler.NewSettingHandler()
	gameHandler       = handler.NewGameHandler(gameService)
	gachaHandler      = handler.NewGachaHandler(gachaService)
	rankingHandler    = handler.NewRankingHandler(rankingService)
	collectionHandler = handler.NewCollectionHandler(collectionService)
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
	e.GET("/setting/get", settingHandler.HandleSettingGet())
	e.POST("/user/create", userHandler.HandleUserCreate())

	// 認証を必要とするAPI
	// AuthenticateMiddlewareによってx-tokenヘッダのチェックとユーザーの特定が行われる
	authAPI := e.Group("", authMiddleware.AuthenticateMiddleware())
	authAPI.GET("/user/get", userHandler.HandleUserGet())
	authAPI.POST("/user/update", userHandler.HandleUserUpdate())
	authAPI.GET("/collection/list", collectionHandler.HandleCollectionList())

	authAPI.GET("/ranking/list", rankingHandler.HandleRankingGet())

	authAPI.POST("/game/finish", gameHandler.HandleGameFinish())

	authAPI.POST("/gacha/draw", gachaHandler.HandleGachaDraw())

	/* ===== サーバの起動 ===== */
	log.Println("Server running...")
	if err := e.Start(addr); err != nil {
		log.Fatalf("failed to start server. %+v", err)
	}
}
