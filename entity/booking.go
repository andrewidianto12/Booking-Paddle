package entity

import "time"

type Booking struct {
	BookingID  int
	UserID     int
	CourtID    int
	TimeSlotID int
	BookingDate time.Time
	TotalPrice float64
	Status     string
	CreatedAt time.Time
}