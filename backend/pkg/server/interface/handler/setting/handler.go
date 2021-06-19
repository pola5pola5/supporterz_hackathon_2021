package setting

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type SettingHandler interface {
	HandleSettingGet() echo.HandlerFunc
}

type settingHandler struct {
}

func NewSettingHandler() SettingHandler {
	return &settingHandler{}
}

// HandleUserGet ユーザ情報取得
func (sh *settingHandler) HandleSettingGet() echo.HandlerFunc {
	return func(c echo.Context) error {
		// レスポンス
		// settingGetResponse レスポンス形式
		type settingGetResponse struct {
			Hello string `json:"hello"`
		}
		res := settingGetResponse{
			Hello: "hello",
		}
		return c.JSON(
			http.StatusOK,
			res,
		)
	}
}
