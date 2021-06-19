package trip

import (
	"fmt"
	"time"

	"github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	"github.com/google/uuid"
	"github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/model/img"
	"github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/model/trip"
	fr "github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/repository/file"
	ir "github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/repository/img"
	tr "github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/repository/trip"
)

const path = "./save/"
const layout = "2006:01:02 15:04:05"

type TripUsecase interface {
	RegisterTrip(userID string, img [][]byte) (string, error)
	GetImgsByTripID(tripID string) ([]*img.Img, error)
}

type tripUsecase struct {
	fileRepo   fr.FileRepo
	imgRepo    ir.ImgRepo
	tripRepo   tr.TripRepo
	createUUID func() (uuid.UUID, error)
}

func NewTripUsecase(
	fr fr.FileRepo,
	ir ir.ImgRepo,
	tr tr.TripRepo,
	f func() (uuid.UUID, error),
) TripUsecase {
	return &tripUsecase{
		fileRepo:   fr,
		imgRepo:    ir,
		tripRepo:   tr,
		createUUID: f,
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
	// ユーザ確認
	trip := &trip.Trip{
		TripID: tripID.String(),
		UserID: userID,
	}
	if err := tu.tripRepo.InsertTrip(trip); err != nil {
		return "", err
	}

	// ---
	// 画像から必要な情報抜き取り登録
	// ---
	for i := 0; i < len(imgs); i++ {
		imgID, err := tu.createUUID()
		if err != nil {
			return "", err
		}
		// ---
		// exif抜き取る
		// ---
		data, err := extractExifContent(imgs[i])
		if err != nil {
			return "", err
		}

		// ---
		// s3に保存
		// TODO: 現状はサーバ上に保存
		// ---
		filename := path + imgID.String()
		if err := tu.fileRepo.SaveFile(filename, imgs[i]); err != nil {
			return "", err
		}

		// ---
		// DB保存する
		// ---
		// 型アサーション
		lat, ok := data["gps_latitude"].(float64)
		if !ok {
			return "", fmt.Errorf("extracted type is wrong. alt=%s", "latitude")
		}
		lng, ok := data["gps_longitude"].(float64)
		if !ok {
			return "", fmt.Errorf("extracted type is wrong. alt=%s", "longitude")
		}

		time, ok := data["time"].(time.Time)
		if !ok {
			return "", fmt.Errorf("extracted type is wrong. alt=%s", "time")
		}

		img := &img.Img{
			ImgID:     imgID.String(),
			TripID:    tripID.String(),
			ImgUrl:    filename,
			Latitude:  lat,
			Longitude: lng,
			DataTime:  time,
		}
		if err := tu.imgRepo.InsertImg(img); err != nil {
			return "", err
		}
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

// GPS情報を抜き取る
func extractExifContent(img []byte) (map[string]interface{}, error) {
	// ---
	// exif抽出
	// ---
	rawExif, err := exif.SearchAndExtractExif(img)
	if err != nil {
		return nil, err
	}
	im, err := exifcommon.NewIfdMappingWithStandard()
	if err != nil {
		return nil, err
	}
	ti := exif.NewTagIndex()
	_, index, err := exif.Collect(im, ti, rawExif)
	if err != nil {
		return nil, err
	}
	data := make(map[string]interface{})

	// ---
	// gps
	// ---
	ifd, err := index.RootIfd.ChildWithIfdPath(exifcommon.IfdGpsInfoStandardIfdIdentity)
	if err != nil {
		return nil, err
	}
	gi, err := ifd.GpsInfo()
	if err != nil {
		return nil, err
	}
	data["gps_latitude"] = gi.Latitude.Decimal()
	data["gps_longitude"] = gi.Longitude.Decimal()

	// ---
	// time
	// ---
	results, err := index.Tree[0].FindTagWithName("DateTime")
	if err != nil {
		return nil, err
	}
	if len(results) != 1 {
		errMsg := fmt.Errorf("there wasn't exactly one result. len=%d", len(results))
		return nil, errMsg
	}
	ite := results[0]
	time, err := ite.Value()
	if err != nil {
		return nil, err
	}
	// 型アサーション
	t, ok := time.(string)
	if !ok {
		return nil, fmt.Errorf("extracted type is wrong. alt=%s", "time")
	}

	timeTime, err := stringToTime(t)
	if err != nil {
		return nil, err
	}
	data["time"] = *timeTime

	return data, nil
}

func stringToTime(str string) (*time.Time, error) {
	t, err := time.Parse(layout, str)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
