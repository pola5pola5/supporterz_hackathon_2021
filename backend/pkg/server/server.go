package server

import (
	"log"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/littletake/supporterz_hackathon_2021/pkg/server/infra/db"
	ip "github.com/littletake/supporterz_hackathon_2021/pkg/server/infra/persistence/img"
	"github.com/littletake/supporterz_hackathon_2021/pkg/server/infra/persistence/s3"
	txp "github.com/littletake/supporterz_hackathon_2021/pkg/server/infra/persistence/transaction"
	tp "github.com/littletake/supporterz_hackathon_2021/pkg/server/infra/persistence/trip"
	up "github.com/littletake/supporterz_hackathon_2021/pkg/server/infra/persistence/user"
	sh "github.com/littletake/supporterz_hackathon_2021/pkg/server/interface/handler/setting"
	th "github.com/littletake/supporterz_hackathon_2021/pkg/server/interface/handler/trip"
	uh "github.com/littletake/supporterz_hackathon_2021/pkg/server/interface/handler/user"
	m "github.com/littletake/supporterz_hackathon_2021/pkg/server/interface/middleware"
	tu "github.com/littletake/supporterz_hackathon_2021/pkg/server/usecase/trip"
	uu "github.com/littletake/supporterz_hackathon_2021/pkg/server/usecase/user"
)

func Serve(addr string) {

	// ---
	// DI
	// ---
	// infra
	userPersistence := up.NewPersistence(db.Conn)
	imgPersistence := ip.NewPersistence(db.Conn)
	tripPersistence := tp.NewPersistence(db.Conn)
	s3Persistence := s3.NewPersistence()
	txpPersistence := txp.NewPersistence(db.Conn)

	// usecase
	userUsecase := uu.NewUserUsecase(
		userPersistence,
		tripPersistence,
		uuid.NewRandom,
	)
	tripUsecase := tu.NewTripUsecase(
		s3Persistence,
		imgPersistence,
		tripPersistence,
		uuid.NewRandom,
		txpPersistence,
	)
	log.Println("aws")

	// interface
	settingHandler := sh.NewSettingHandler()
	userHandler := uh.NewUserHandler(userUsecase, tripUsecase)
	tripHandler := th.NewTripHandler(tripUsecase, userUsecase)
	mw := m.NewMyMiddleware(userUsecase)
	// ---
	// URL マッピング
	// ---
	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	api := e.Group("/api")
	api.POST("/user/create", userHandler.HandleUserCreate())
	api.POST("/user/login", userHandler.HandleUserLogin())

	apiAuth := api.Group("/auth")
	apiAuth.Use(mw.Authenticate)
	// setting
	apiAuth.GET("/setting/get", settingHandler.HandleSettingGet())
	// user
	apiAuth.GET("/user/get", userHandler.HandleUserGet())
	apiAuth.GET("/user/get_trip", userHandler.HandleUserTripGet())
	// img
	apiAuth.POST("/trip/save", tripHandler.HandleTripSave())
	apiAuth.GET("/trip/get", tripHandler.HandleTripGet())
	apiAuth.GET("/trip/get_date", tripHandler.HandleDateGet())
	// html
	apiAuth.Static("/html", "static")

	// Start server
	e.Logger.Fatal(e.Start(addr))
}
