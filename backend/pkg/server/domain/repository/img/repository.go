package img

import "github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/model/img"

type ImgRepo interface {
	SelectImgsByTripID(tripID string) ([]*img.Img, error)
	InsertImg(img *img.Img) error
}
