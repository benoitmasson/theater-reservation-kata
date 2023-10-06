package dao

import (
	"sync"

	"github.com/benoitmasson/theater-reservation-kata/internal/types"
)

type ReservationDAO struct {
	reservationMap map[int64]*types.Reservation
	mutex          *sync.RWMutex
}

func NewReservationDAO() ReservationDAO {
	return ReservationDAO{
		reservationMap: make(map[int64]*types.Reservation),
		mutex:          &sync.RWMutex{},
	}
}

func (dao *ReservationDAO) Update(reservation types.Reservation) {
	dao.mutex.Lock()
	defer dao.mutex.Unlock()

	if dao.reservationMap == nil {
		dao.reservationMap = make(map[int64]*types.Reservation)
	}
	dao.reservationMap[reservation.ReservationID] = &reservation
}

func (dao *ReservationDAO) Find(reservationID int64) *types.Reservation {
	dao.mutex.RLock()
	defer dao.mutex.RUnlock()

	if dao.reservationMap == nil {
		return nil
	}
	return dao.reservationMap[reservationID]
}
