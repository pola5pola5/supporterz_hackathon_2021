package trip

import "github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/model/trip"

type TripRepo interface {
	SelectTripByTripID(tripID string) (*trip.Trip, error)
	InsertTrip(trip *trip.Trip) error
	UpdateTrip(trip *trip.Trip) error
}
