package img

import "time"

type Img struct {
	ImgID     string
	TripID    string
	ImgUrl    string
	Longitude float64
	Latitude  float64
	DataTime  time.Time
}
