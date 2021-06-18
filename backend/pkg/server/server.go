package server

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/littletake/supporterz_hackathon_2021/pkg/server/infra/db"
	up "github.com/littletake/supporterz_hackathon_2021/pkg/server/infra/persistence/user"
	"github.com/littletake/supporterz_hackathon_2021/pkg/server/interface/handler/setting"
	uh "github.com/littletake/supporterz_hackathon_2021/pkg/server/interface/handler/user"
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

	// usecase
	userUsecase := uu.NewUserUsecase(userPersistence, uuid.NewRandom)

	// interface
	settingHandler := setting.NewSettingHandler()
	userHandler := uh.NewUserHandler(userUsecase)

	// ---
	// URL マッピング
	// ---
	// setting
	e.GET("/setting/get", settingHandler.HandleSettingGet())
	// user
	e.GET("/user/get", userHandler.HandleUserGet())
	e.POST("/user/create", userHandler.HandleUserCreate())
	// img

	// Start server
	e.Logger.Fatal(e.Start(addr))
}
