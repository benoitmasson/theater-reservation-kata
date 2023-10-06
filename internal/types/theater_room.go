package types

import "fmt"

type TheaterRoom struct {
	Zones []Zone
}

type ZoneCategory string

const (
	ZoneCategoryStandard ZoneCategory = "STANDARD"
	ZoneCategoryPremium  ZoneCategory = "PREMIUM"
)

type Zone struct {
	Rows     []Row
	Category ZoneCategory
}

type Row struct {
	Seats []Seat
}

type SeatStatus string

const (
	SeatStatusFree           = "FREE"
	SeatStatusBooked         = "BOOKED"
	SeatStatusBookingPending = "BOOKING_PENDING"
)

type Seat struct {
	SeatID string
	Status SeatStatus
}

func (s Seat) String() string {
	return fmt.Sprintf("Seat{SeatID=%q,Status=%q}", s.SeatID, s.Status)
}
