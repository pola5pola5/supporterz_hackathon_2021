package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	uu "github.com/littletake/supporterz_hackathon_2021/pkg/server/usecase/user"
)

type MyMiddleware interface {
	Authenticate(echo.HandlerFunc) echo.HandlerFunc
}

type myMiddleware struct {
	userUseCase uu.UserUsecase
}

// NewMyMiddleware userUseCaseと疎通
func NewMyMiddleware(uu uu.UserUsecase) MyMiddleware {
	return &myMiddleware{
		userUseCase: uu,
	}
}

func (m *myMiddleware) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// tokenの確認
		token := c.Request().Header["X-Token"][0]
		if token == "" {
			return echo.NewHTTPError(
				http.StatusBadRequest,
				fmt.Errorf("X-Token is empty"),
			)
		}
		_, err := m.userUseCase.GetUserByAuthToken(token)
		if err != nil {
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				err,
			)
		}
		return next(c)
	}
}
