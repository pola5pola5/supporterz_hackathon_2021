package trip

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/littletake/supporterz_hackathon_2021/pkg/server/usecase/trip"
	usecase "github.com/littletake/supporterz_hackathon_2021/pkg/server/usecase/user"
)

type TripHandler interface {
	HandleTripSave() echo.HandlerFunc
	HandleTripGet() echo.HandlerFunc
}

type tripHandler struct {
	tripUsecase trip.TripUsecase
	userUsecase usecase.UserUsecase
}

func NewTripHandler(tu trip.TripUsecase, uu usecase.UserUsecase) TripHandler {
	return &tripHandler{
		tripUsecase: tu,
		userUsecase: uu,
	}
}

type tripSaveRequest struct {
	UserID string   `json:"user_id"`
	Imgs   []string `json:"imgs"`
}

type tripSaveResponse struct {
	TripID string `json:"trip_id"`
}

// TODO: 2回ログが出力される問題
func (th *tripHandler) HandleTripSave() echo.HandlerFunc {
	return func(c echo.Context) error {
		var requestBody tripSaveRequest
		if err := c.Bind(&requestBody); err != nil {
			c.JSON(
				http.StatusBadRequest,
				err,
			)
			return err
		}
		if requestBody.UserID == "" {
			errMsg := fmt.Errorf("userID is empty")
			c.JSON(
				http.StatusBadRequest,
				errMsg,
			)
			return errMsg
		}
		_, err := th.userUsecase.GetUserByUserID(requestBody.UserID)
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				err,
			)
			return err
		}

		// encode
		// TODO: 並列処理を実装する
		imgDecodedSlice := make([][]byte, len(requestBody.Imgs))
		for i := 0; i < len(requestBody.Imgs); i++ {
			imgEncoded, err := base64.StdEncoding.DecodeString(requestBody.Imgs[i])
			if err != nil {
				c.JSON(
					http.StatusInternalServerError,
					err,
				)
				return err
			}
			imgDecodedSlice[i] = imgEncoded
		}
		// usecase
		tripID, err := th.tripUsecase.RegisterTrip(requestBody.UserID, imgDecodedSlice)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				err,
			)
			return err
		}
		res := tripSaveResponse{
			TripID: tripID,
		}
		return c.JSON(
			http.StatusOK,
			res,
		)
	}
}

type tripGetResponse struct {
	ResponseType string     `json:"type"`
	Features     []features `json:"features"`
}

type features struct {
	Type       string     `json:"type"`
	Properties properties `json:"properties"`
	Geometry   geometry   `json:"geometry"`
}

type geometry struct {
	GeoType     string     `json:"type"`
	Coordinates [2]float64 `json:"coordinates"`
}

type properties struct {
	TripID string `json:"trip_id"`
	ImgURL string `json:"img_url"`
}

func (th *tripHandler) HandleTripGet() echo.HandlerFunc {
	return func(c echo.Context) error {
		// クエリパラメータからtripID取得
		tripID := c.QueryParam("trip_id")
		if tripID == "" {
			errMsg := fmt.Errorf("tripID is empty")
			c.JSON(
				http.StatusBadRequest,
				errMsg,
			)
			return errMsg
		}

		// Imgオブジェクト取得
		imgs, err := th.tripUsecase.GetImgsByTripID(tripID)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				err,
			)
			return err
		}

		// TODO: 並列処理にする
		resSlice := make([]features, len(imgs))
		for i := 0; i < len(imgs); i++ {
			properties := properties{
				TripID: imgs[i].TripID,
				ImgURL: imgs[i].ImgUrl,
			}
			geometry := geometry{
				GeoType: "Point",
				Coordinates: [...]float64{
					imgs[i].Latitude,
					imgs[i].Longitude,
				},
			}
			features := features{
				Type:       "Feature",
				Properties: properties,
				Geometry:   geometry,
			}
			resSlice[i] = features
		}

		res := tripGetResponse{
			ResponseType: "FeatureCollection",
			Features:     resSlice,
		}
		return c.JSON(
			http.StatusOK,
			res,
		)
	}
}
