package service

import (
	"sync"

	"github.com/benoitmasson/theater-reservation-kata/internal/dao"
	"github.com/benoitmasson/theater-reservation-kata/internal/types"
)

type ReservationService struct {
	currentID int64
	idMutex   sync.Mutex

	reservationDAO dao.ReservationDAO
}

func NewReservationService(reservationDAO dao.ReservationDAO) ReservationService {
	return ReservationService{
		currentID:      123455,
		reservationDAO: reservationDAO,
	}
}

func (r *ReservationService) InitNewReservation() int64 {
	r.idMutex.Lock()
	defer r.idMutex.Unlock()

	r.currentID++
	return r.currentID
}

func (r *ReservationService) Update(reservation types.Reservation) {
	r.reservationDAO.Update(reservation)
}

func (r *ReservationService) Find(reservationID int64) *types.Reservation {
	return r.reservationDAO.Find(reservationID)
}

func (r *ReservationService) Cancel(reservationID int64) {
	reservation := r.Find(reservationID)
	if reservation != nil {
		reservation.Status = types.ReservationStatusCancelled
		reservation.Seats = []string{}
		r.Update(*reservation)
	}
}
