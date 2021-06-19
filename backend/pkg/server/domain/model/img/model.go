package img

import "time"

type Img struct {
	ImgID     string
	TripID    string
	ImgUrl    string
	Latitude  float64
	Longitude float64
	DataTime  time.Time
}
