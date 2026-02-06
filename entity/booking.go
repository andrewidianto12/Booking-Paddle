package entity

import "time"

type Booking struct {
	ID         int
	UserID     int
	UserName   string
	CourtID    int
	CourtName  string
	Date       time.Time
	TimeSlotID int
	StartTime  string
	EndTime    string
	TotalPrice float64
	Status     string
}
