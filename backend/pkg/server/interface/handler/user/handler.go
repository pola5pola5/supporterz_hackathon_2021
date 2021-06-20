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
		fmt.Print("tstß")
		if userID == "" {
			errMsg := fmt.Errorf("userID is empty")
			c.JSON(
				http.StatusBadRequest,
				errMsg,
			)
			return errMsg
		}

		// ユーザ取得
		user, err := uh.userUsecase.GetUserByUserID(userID)
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				err,
			)
			return err
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
	Token string `json:"token"`
}

// TODO: tokenの扱い方検討
// HandleUserCreate　ユーザ作成
func (uh *userHandler) HandleUserCreate() echo.HandlerFunc {
	return func(c echo.Context) error {
		// リクエストBodyから更新情報を取得
		var requestBody userCreateRequest
		if err := c.Bind(&requestBody); err != nil {
			c.JSON(
				http.StatusBadRequest,
				err,
			)
			return err
		}

		authToken, err := uh.userUsecase.RegisterUser(
			requestBody.UserName,
			requestBody.Password,
		)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				err,
			)
			return err
		}
		res := userCreateResponse{
			Token: authToken,
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
			errMsg := fmt.Errorf("userID is empty")
			c.JSON(
				http.StatusBadRequest,
				errMsg,
			)
			return errMsg
		}

		// 旅情報の一覧取得
		trips, err := uh.userUsecase.GetTripsByUserID(userID)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				err,
			)
			return err
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
