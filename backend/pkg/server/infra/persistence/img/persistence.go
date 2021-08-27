package img

import (
	"database/sql"
	"fmt"
	"log"
	"time"

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

func (ip imgPersistence) InsertImg(img *model.Img, tx *sql.Tx) error {
	stmt, err := tx.Prepare(
		"INSERT INTO img_table (img_id, trip_id, img_url, longitude, latitude, date_time) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(img.ImgID, img.TripID, img.ImgUrl, img.Longitude, img.Latitude, img.DataTime)
	return err
}

func (ip imgPersistence) SelectDatesByTripID(trip_id string) ([]*time.Time, error) {
	// 開始日時の取得
	row := ip.db.QueryRow("SELECT date_time FROM img_table WHERE trip_id = ? ORDER BY date_time LIMIT 1", trip_id)
	fmt.Print(row)
	start_date, err := convertToTime(row)
	if err != nil {
		return nil, err
	}
	// 終了日時の取得
	row = ip.db.QueryRow("SELECT date_time FROM img_table WHERE trip_id = ? ORDER BY date_time DESC LIMIT 1", trip_id)
	fmt.Print(row)
	end_date, err := convertToTime(row)
	if err != nil {
		return nil, err
	}
	dateSlice := [...]*time.Time{start_date, end_date}
	return dateSlice[:], nil
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
			&img.Longitude,
			&img.Latitude,
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

// rowデータをtime.Timeへ変換
func convertToTime(row *sql.Row) (*time.Time, error) {
	date := &time.Time{}
	err := row.Scan(date)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	return date, nil
}
