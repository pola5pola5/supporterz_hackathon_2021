package img

import (
	"database/sql"
	"log"

	model "github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/model/img"
	repo "github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/repository/img"
)

type imgPersistence struct {
	db *sql.DB
}

func NewPersistence(db *sql.DB) repo.ImgRepo {
	return &imgPersistence{
		db: db,
	}
}

// TODO: datatime順に取得してくる
func (ip imgPersistence) SelectImgsByTripID(tripID string) ([]*model.Img, error) {
	rows, err := ip.db.Query("SELECT * FROM img_table WHERE trip_id = ? ORDER BY date_time ASC", tripID)
	if err != nil {
		return nil, err
	}
	return convertToImgs(rows)
}

// TODO: バルクインサートに変更する
func (ip imgPersistence) InsertImg(img *model.Img) error {
	stmt, err := ip.db.Prepare(
		"INSERT INTO img_table (img_id, trip_id, img_url, latitude, longitude, date_time) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(img.ImgID, img.TripID, img.ImgUrl, img.Latitude, img.Longitude, img.DataTime)
	return err
}

// rowデータをImgデータへ変換する
func convertToImgs(rows *sql.Rows) ([]*model.Img, error) {
	imgs := []*model.Img{}
	for rows.Next() {
		img := model.Img{}
		err := rows.Scan(
			&img.ImgID,
			&img.TripID,
			&img.ImgUrl,
			&img.Latitude,
			&img.Longitude,
			&img.DataTime,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			log.Println(err)
			return nil, err
		}
		imgs = append(imgs, &img)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	return imgs, nil
}
