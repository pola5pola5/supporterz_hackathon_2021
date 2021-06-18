package user

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	usecase "github.com/littletake/supporterz_hackathon_2021/pkg/server/usecase/user"
)

type UserHandler interface {
	HandleUserGet() echo.HandlerFunc
	HandleUserCreate() echo.HandlerFunc
	// HandleUserUpdate() echo.HandlerFunc
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(uu usecase.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: uu,
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
		userID := c.QueryParam("userID")
		if userID == "" {
			return c.JSON(
				http.StatusBadRequest,
				errors.New("userID is empty"),
			)
		}

		// ユーザ取得
		user, err := uh.userUsecase.GetUserByUserID(userID)
		if err != nil {
			return c.JSON(
				http.StatusBadRequest,
				fmt.Errorf("user not found. userID=%s", userID),
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
	Token string `json:"token"`
}

// TODO: tokenの扱い方検討
// HandleUserCreate　ユーザ作成
func (uh *userHandler) HandleUserCreate() echo.HandlerFunc {
	return func(c echo.Context) error {
		// リクエストBodyから更新情報を取得
		var requestBody userCreateRequest
		if err := c.Bind(&requestBody); err != nil {
			return c.JSON(
				http.StatusBadRequest,
				err.Error(),
			)
		}

		authToken, err := uh.userUsecase.RegisterUser(
			requestBody.UserName,
			requestBody.Password,
		)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				err.Error(),
			)
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
