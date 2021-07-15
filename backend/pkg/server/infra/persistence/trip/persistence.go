package trip

import (
	"database/sql"
	"log"

	model "github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/model/trip"
	repo "github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/repository/trip"
)

type tripPersistence struct {
	db *sql.DB
}

func NewPersistence(db *sql.DB) repo.TripRepo {
	return &tripPersistence{
		db: db,
	}
}

func (tp tripPersistence) SelectTripsByUserID(user_id string) ([]*model.Trip, error) {
	rows, err := tp.db.Query("SELECT * FROM trip_table WHERE user_id = ?", user_id)
	if err != nil {
		return nil, err
	}
	return convertToTrips(rows)
}

func (tp tripPersistence) InsertTrip(record *model.Trip, tx *sql.Tx) error {
	stmt, err := tx.Prepare("INSERT INTO trip_table (trip_id, user_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		record.TripID,
		record.UserID,
	)
	return err
}

// rowデータをTripデータへ変換する
func convertToTrips(rows *sql.Rows) ([]*model.Trip, error) {
	trips := []*model.Trip{}
	for rows.Next() {
		trip := model.Trip{}
		err := rows.Scan(
			&trip.TripID,
			&trip.UserID,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			log.Println(err)
			return nil, err
		}
		trips = append(trips, &trip)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	return trips, nil
}
