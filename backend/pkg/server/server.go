package server

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/littletake/supporterz_hackathon_2021/pkg/server/infra/db"
	ip "github.com/littletake/supporterz_hackathon_2021/pkg/server/infra/persistence/img"
	lp "github.com/littletake/supporterz_hackathon_2021/pkg/server/infra/persistence/local"
	tp "github.com/littletake/supporterz_hackathon_2021/pkg/server/infra/persistence/trip"
	up "github.com/littletake/supporterz_hackathon_2021/pkg/server/infra/persistence/user"
	sh "github.com/littletake/supporterz_hackathon_2021/pkg/server/interface/handler/setting"
	th "github.com/littletake/supporterz_hackathon_2021/pkg/server/interface/handler/trip"
	uh "github.com/littletake/supporterz_hackathon_2021/pkg/server/interface/handler/user"
	tu "github.com/littletake/supporterz_hackathon_2021/pkg/server/usecase/trip"
	uu "github.com/littletake/supporterz_hackathon_2021/pkg/server/usecase/user"
)

func Serve(addr string) {
	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// ---
	// DI
	// ---
	// infra
	userPersistence := up.NewPersistence(db.Conn)
	imgPersistence := ip.NewPersistence(db.Conn)
	tripPersistence := tp.NewPersistence(db.Conn)
	localPersistence := lp.NewPersistence()
	// usecase
	userUsecase := uu.NewUserUsecase(
		userPersistence,
		tripPersistence,
		uuid.NewRandom,
	)
	tripUsecase := tu.NewTripUsecase(
		localPersistence,
		imgPersistence,
		tripPersistence,
		uuid.NewRandom,
	)
	// interface
	settingHandler := sh.NewSettingHandler()
	userHandler := uh.NewUserHandler(userUsecase, tripUsecase)
	tripHandler := th.NewTripHandler(tripUsecase, userUsecase)
	// ---
	// URL マッピング
	// ---
	// setting
	e.GET("/setting/get", settingHandler.HandleSettingGet())
	// user
	e.GET("/user/get", userHandler.HandleUserGet())
	e.POST("/user/create", userHandler.HandleUserCreate())
	e.GET("/user/get_trip", userHandler.HandleUserTripGet())
	// img
	e.POST("/trip/save", tripHandler.HandleTripSave())
	e.GET("/trip/get", tripHandler.HandleTripGet())
	// html
	e.Static("/html", "static")

	// Start server
	e.Logger.Fatal(e.Start(addr))
}
