package user

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	tu "github.com/littletake/supporterz_hackathon_2021/pkg/server/usecase/trip"
	uu "github.com/littletake/supporterz_hackathon_2021/pkg/server/usecase/user"
)

type UserHandler interface {
	HandleUserGet() echo.HandlerFunc
	HandleUserCreate() echo.HandlerFunc
	HandleUserLogin() echo.HandlerFunc
	HandleUserTripGet() echo.HandlerFunc
}

type userHandler struct {
	userUsecase uu.UserUsecase
	tripUsecase tu.TripUsecase
}

func NewUserHandler(uu uu.UserUsecase, tu tu.TripUsecase) UserHandler {
	return &userHandler{
		userUsecase: uu,
		tripUsecase: tu,
	}
}

// userGetResponse ユーザ取得response
type userGetResponse struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
}

// HandleUserGet ユーザ情報取得
func (uh *userHandler) HandleUserGet() echo.HandlerFunc {
	return func(c echo.Context) error {
		// クエリパラメータからuserID取得
		userID := c.QueryParam("user_id")
		if userID == "" {
			errMsg := fmt.Errorf("userID is empty")
			return echo.NewHTTPError(
				http.StatusBadRequest,
				errMsg,
			)
		}

		// ユーザ取得
		user, err := uh.userUsecase.GetUserByUserID(userID)
		if err != nil {
			return echo.NewHTTPError(
				http.StatusBadRequest,
				err,
			)
		}

		// レスポンス
		res := userGetResponse{
			UserID:   user.UserID,
			UserName: user.UserName,
		}
		return c.JSON(
			http.StatusOK,
			res,
		)
	}
}

// userCreateRequest ユーザ作成request
type userCreateRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// userCreateResponse ユーザ作成response
type userCreateResponse struct {
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}

// HandleUserCreate　ユーザ作成
func (uh *userHandler) HandleUserCreate() echo.HandlerFunc {
	return func(c echo.Context) error {
		// リクエストBodyから更新情報を取得
		var requestBody userCreateRequest
		if err := c.Bind(&requestBody); err != nil {
			return echo.NewHTTPError(
				http.StatusBadRequest,
				err,
			)
		}

		// ユーザが重複しないように判定処理
		// TODO: DB側のエラー時に対応できない
		user, _ := uh.userUsecase.GetUserByUserNamePassword(
			requestBody.UserName,
			requestBody.Password,
		)
		if user != nil { // 同名かつ同パスワードのユーザが存在
			return echo.NewHTTPError(
				http.StatusBadRequest,
				fmt.Errorf(
					"user has already been registered. userName=%s, password=%s",
					requestBody.UserName,
					requestBody.Password,
				),
			)
		}

		authToken, userID, err := uh.userUsecase.RegisterUser(
			requestBody.UserName,
			requestBody.Password,
		)
		if err != nil {
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				err,
			)
		}
		res := userCreateResponse{
			Token:  authToken,
			UserID: userID,
		}
		return c.JSON(
			http.StatusOK,
			res,
		)
	}
}

// ログイン機能request
type userLoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// ログイン機能response
type userLoginResponse struct {
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}

// ログイン機能
func (uh *userHandler) HandleUserLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		// リクエストBodyからログイン情報を取得
		var requestBody userLoginRequest
		if err := c.Bind(&requestBody); err != nil {
			return echo.NewHTTPError(
				http.StatusBadRequest,
				err,
			)
		}

		// ユーザの確認
		user, err := uh.userUsecase.GetUserByUserNamePassword(
			requestBody.UserName,
			requestBody.Password,
		)
		// ユーザがいない時
		if user == nil && err != nil {
			return echo.NewHTTPError(
				http.StatusBadRequest,
				err,
			)
		} else if err != nil { // サーバ上のエラー
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				err,
			)
		}

		// レスポンス
		res := userLoginResponse{
			UserID: user.UserID,
			Token:  user.AuthToken,
		}
		return c.JSON(
			http.StatusOK,
			res,
		)
	}
}

type userTripGetResponse struct {
	TripID []string `json:"trip_id"`
}

func (uh *userHandler) HandleUserTripGet() echo.HandlerFunc {
	return func(c echo.Context) error {
		// クエリパラメータからuserID取得
		userID := c.QueryParam("user_id")
		if userID == "" {
			return echo.NewHTTPError(
				http.StatusBadRequest,
				fmt.Errorf("userID is empty"),
			)
		}

		// 旅情報の一覧取得
		trips, err := uh.userUsecase.GetTripsByUserID(userID)
		if err != nil {
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				err,
			)
		}
		tripIDSlice := make([]string, len(trips))
		for i := 0; i < len(trips); i++ {
			// TODO: ポインタ注意
			tripIDSlice[i] = trips[i].TripID
		}
		res := userTripGetResponse{
			TripID: tripIDSlice,
		}
		return c.JSON(
			http.StatusOK,
			res,
		)
	}
}
