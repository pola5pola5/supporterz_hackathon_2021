package trip

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/model/img"
	"github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/model/trip"
	"github.com/littletake/supporterz_hackathon_2021/pkg/server/infra/persistence/s3"
	fr "github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/repository/file"
	ir "github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/repository/img"
	txr "github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/repository/transaction"
	tr "github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/repository/trip"
)

const layout = "2006:01:02 15:04:05"

type TripUsecase interface {
	RegisterTrip(userID string, img [][]byte) (string, error)
	GetImgsByTripID(tripID string) ([]*img.Img, error)
	GetImgByUrl(imgUrl string) ([]byte, error)
}

type tripUsecase struct {
	fileRepo   fr.FileRepo
	imgRepo    ir.ImgRepo
	tripRepo   tr.TripRepo
	createUUID func() (uuid.UUID, error)
	txRepo     txr.TxRepo
}

func NewTripUsecase(
	fr fr.FileRepo,
	ir ir.ImgRepo,
	tr tr.TripRepo,
	f func() (uuid.UUID, error),
	txr txr.TxRepo,
) TripUsecase {
	return &tripUsecase{
		fileRepo:   fr,
		imgRepo:    ir,
		tripRepo:   tr,
		createUUID: f,
		txRepo:     txr,
	}
}

func (tu *tripUsecase) RegisterTrip(userID string, imgs [][]byte) (string, error) {
	tripID, err := tu.createUUID()
	if err != nil {
		return "", err
	}
	// ---
	// 旅の登録
	// ---
	// トランザクション
	if err := tu.txRepo.Transaction(func(tx *sql.Tx) error {
		// ユーザ確認
		trip := &trip.Trip{
			TripID: tripID.String(),
			UserID: userID,
		}
		if err := tu.tripRepo.InsertTrip(trip, tx); err != nil {
			return err
		}
		// ---
		// 画像から必要な情報抜き取り登録
		// ---
		for i := 0; i < len(imgs); i++ {
			if err := ExtractInfoAndSave(tu, tx, imgs[i], tripID.String()); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return "", err
	}
	return tripID.String(), nil
}

func (tu *tripUsecase) GetImgsByTripID(tripID string) ([]*img.Img, error) {
	imgs, err := tu.imgRepo.SelectImgsByTripID(tripID)
	if err != nil {
		return nil, err
	}
	if imgs == nil {
		errMsg := fmt.Errorf("img not found. tripID=%s", tripID)
		return nil, errMsg
	}
	return imgs, nil
}

// 画像のurlを受け取り、外部ストレージから画像を取得
func (tu *tripUsecase) GetImgByUrl(imgUrl string) ([]byte, error) {
	imgByte, err := s3.GetFile(imgUrl)
	if err != nil {
		return nil, err
	}
	return imgByte, nil
}


// 画像から必要な情報抜き取り、外部ストレージとDBに保存
func ExtractInfoAndSave(
	tu *tripUsecase,
	tx *sql.Tx,
	imgData []byte,
	tripID string,
) error {
	flag := 0
	imgID, err := tu.createUUID()
	if err != nil {
		return err
	}
	// ---
	// exif抜き取る
	// ---
	data, err := ExtractExifContent(imgData)
	if err != nil {
		return err
	}

	// ---
	// ストレージに保存
	// ---
	filename := imgID.String()
	imgUrl, err := tu.fileRepo.SaveFile(filename, imgData)
	if err != nil {
		return err
	}

	defer func() {
		if flag > 0 {
			// 外部ストレージに保存したファイルの削除
			if err := tu.fileRepo.DeleteFile(filename); err != nil {
				// TODO: エラーハンドリング考える
				log.Println(err)
			}
		}
	}()

	// ---
	// DB保存する
	// ---
	// 型アサーション
	lat, ok := data["gps_latitude"].(float64)
	if !ok {
		flag++
		return fmt.Errorf("extracted type is wrong. alt=%s", "latitude")
	}
	lng, ok := data["gps_longitude"].(float64)
	if !ok {
		flag++
		return fmt.Errorf("extracted type is wrong. alt=%s", "longitude")
	}

	time, ok := data["time"].(time.Time)
	if !ok {
		flag++
		return fmt.Errorf("extracted type is wrong. alt=%s", "time")
	}

	imgModel := &img.Img{
		ImgID:     imgID.String(),
		TripID:    tripID,
		ImgUrl:    imgUrl,
		Longitude: lng,
		Latitude:  lat,
		DataTime:  time,
	}
	if err := tu.imgRepo.InsertImg(imgModel, tx); err != nil {
		flag++
		return err
	}

	return nil
}
