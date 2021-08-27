package img

import (
	"database/sql"
	"time"

	"github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/model/img"
)

type ImgRepo interface {
	SelectImgsByTripID(tripID string) ([]*img.Img, error)
	InsertImg(img *img.Img, tx *sql.Tx) error
	SelectDatesByTripID(tripID string) ([]*time.Time, error)
}
