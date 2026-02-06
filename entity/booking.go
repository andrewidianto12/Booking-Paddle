package entity

import "time"

type Booking struct {
    BookingID   int
    UserID      int
    UserName    string
    CourtID     int
    CourtName   string
    TimeSlotID  int
    BookingDate time.Time
    StartTime   string
    EndTime     string
    TotalPrice  float64
    Status      string
    CreatedAt   time.Time
}
