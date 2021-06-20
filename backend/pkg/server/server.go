package server

import (
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/littletake/supporterz_hackathon_2021/pkg/server/infra/db"
	ip "github.com/littletake/supporterz_hackathon_2021/pkg/server/infra/persistence/img"
	lp "github.com/littletake/supporterz_hackathon_2021/pkg/server/infra/persistence/local"
	"github.com/littletake/supporterz_hackathon_2021/pkg/server/infra/persistence/s3"
	txp "github.com/littletake/supporterz_hackathon_2021/pkg/server/infra/persistence/transaction"
	tp "github.com/littletake/supporterz_hackathon_2021/pkg/server/infra/persistence/trip"
	up "github.com/littletake/supporterz_hackathon_2021/pkg/server/infra/persistence/user"
	sh "github.com/littletake/supporterz_hackathon_2021/pkg/server/interface/handler/setting"
	th "github.com/littletake/supporterz_hackathon_2021/pkg/server/interface/handler/trip"
	uh "github.com/littletake/supporterz_hackathon_2021/pkg/server/interface/handler/user"
	tu "github.com/littletake/supporterz_hackathon_2021/pkg/server/usecase/trip"
	uu "github.com/littletake/supporterz_hackathon_2021/pkg/server/usecase/user"
)

func Serve(addr string) {
	flag := os.Getenv("LOCAL")
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
	s3Persistence := s3.NewPersistence()
	txpPersistence := txp.NewPersistence(db.Conn)

	// usecase
	userUsecase := uu.NewUserUsecase(
		userPersistence,
		tripPersistence,
		uuid.NewRandom,
	)
	var tripUsecase tu.TripUsecase
	if flag == "local" {
		tripUsecase = tu.NewTripUsecase(
			localPersistence,
			imgPersistence,
			tripPersistence,
			uuid.NewRandom,
			txpPersistence,
		)
		log.Println("local")
	} else {
		tripUsecase = tu.NewTripUsecase(
			s3Persistence,
			imgPersistence,
			tripPersistence,
			uuid.NewRandom,
			txpPersistence,
		)
		log.Println("aws")
	}

	// interface
	settingHandler := sh.NewSettingHandler()
	userHandler := uh.NewUserHandler(userUsecase, tripUsecase)
	tripHandler := th.NewTripHandler(tripUsecase, userUsecase)
	// ---
	// URL マッピング
	// ---
	api := e.Group("/api")
	// setting
	api.GET("/setting/get", settingHandler.HandleSettingGet())
	// user
	api.GET("/user/get", userHandler.HandleUserGet())
	api.POST("/user/create", userHandler.HandleUserCreate())
	api.GET("/user/get_trip", userHandler.HandleUserTripGet())
	// img
	api.POST("/trip/save", tripHandler.HandleTripSave())
	api.GET("/trip/get", tripHandler.HandleTripGet())
	// html
	api.Static("/html", "static")

	// Start server
	e.Logger.Fatal(e.Start(addr))
}
