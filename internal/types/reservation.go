package types

import "fmt"

type ReservationStatus string

const (
	ReservationStatusPending     ReservationStatus = "PENDING"
	ReservationStatusFulfillable ReservationStatus = "FULFILLABLE"
	ReservationStatusAborted     ReservationStatus = "ABORTED"
	ReservationStatusCancelled   ReservationStatus = "CANCELLED"
)

type Reservation struct {
	ReservationID int64
	PerformanceID int64
	Status        ReservationStatus
	Seats         []string
}

func (r *Reservation) String() string {
	return fmt.Sprintf("Reservation{ReservationID=%q,Status=%q,Seats=%v}", r.ReservationID, r.Status, r.Seats)
}
