package trip

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/model/img"
	"github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/model/trip"
	fr "github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/repository/file"
	ir "github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/repository/img"
	txr "github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/repository/transaction"
	tr "github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/repository/trip"
)

const imgLayout = "2006:01:02 15:04:05"
const dbLayout = "2006-01-02 15:04:05"

type TripUsecase interface {
	RegisterTrip(userID string, img [][]byte) (string, error)
	GetImgsByTripID(tripID string) ([]*img.Img, error)
	GetDatesByTripID(tripID string) (*[2]string, error)
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
			if err := ExtractInfoAndSave(
				tu,
				tx,
				imgs[i],
				tripID.String(),
				userID,
			); err != nil {
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
	// DBから画像情報を取得
	imgs, err := tu.imgRepo.SelectImgsByTripID(tripID)
	if err != nil {
		return nil, err
	}
	if imgs == nil {
		errMsg := fmt.Errorf("img not found. tripID=%s", tripID)
		return nil, errMsg
	}
	// アクセス制限のためのpre-signed urlの発行
	// TODO: urlに関しては毎回urlを発行するのでDBに保存しなくても良いのでは？
	// 指定画像のpre-signed urlを発行しオブジェクトの情報を修正
	for _, img := range imgs {
		preSignedURL, err := tu.fileRepo.CreatepreSignedURL(img.ImgID)
		if err != nil {
			return nil, err
		}
		img.ImgUrl = preSignedURL
	}
	return imgs, nil
}

func (tu *tripUsecase) GetDatesByTripID(tripID string) (*[2]string, error) {
	dates, err := tu.imgRepo.SelectDatesByTripID(tripID)
	if err != nil {
		return nil, err
	}
	if dates == nil {
		errMsg := fmt.Errorf("img not found. tripID=%s", tripID)
		return nil, errMsg
	}
	// 出力用に型変換
	var dateArr [2]string
	for i, date := range dates {
		dateArr[i] = changeTypeAndexceptTime(*date)
	}
	return &dateArr, nil
}

// 画像から必要な情報抜き取り、外部ストレージとDBに保存
func ExtractInfoAndSave(
	tu *tripUsecase,
	tx *sql.Tx,
	imgData []byte,
	tripID string,
	userID string,
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
	filename := userID + "/" + imgID.String()
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
		ImgID:     filename,
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

// 型変換し日時情報をまとめる処理
func changeTypeAndexceptTime(t time.Time) string {
	dateString := t.Format(dbLayout)
	dateExceptTime := strings.Split(dateString, " ")[0]
	return dateExceptTime
}
