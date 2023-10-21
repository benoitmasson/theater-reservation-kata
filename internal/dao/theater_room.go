package dao

import (
	"slices"
	"sync"

	"github.com/benoitmasson/theater-reservation-kata/internal/types"
)

type TheaterRoomsDAO struct {
	theaterRoomMaps map[int64]types.TheaterRoom
	mutex           *sync.RWMutex
}

func NewTheaterRoomsDAO() TheaterRoomsDAO {
	dao := TheaterRoomsDAO{
		theaterRoomMaps: make(map[int64]types.TheaterRoom, 3),
		mutex:           &sync.RWMutex{},
	}

	dao.theaterRoomMaps[1] = fetchRoomForPerformance1()
	dao.theaterRoomMaps[2] = fetchRoomForPerformance1()
	dao.theaterRoomMaps[3] = fetchRoomForPerformance2()

	return dao
}

// FetchTheaterRoom simulates a room map/topology repository
func (dao *TheaterRoomsDAO) FetchTheaterRoom(performanceID int64) types.TheaterRoom {
	dao.mutex.RLock()
	defer dao.mutex.RUnlock()

	return dao.theaterRoomMaps[performanceID]
}

func (dao *TheaterRoomsDAO) SaveTheaterRoom(performanceID int64, room types.TheaterRoom) {
	dao.mutex.Lock()
	defer dao.mutex.Unlock()

	dao.theaterRoomMaps[performanceID] = room
}

func (dao *TheaterRoomsDAO) SaveSeats(performanceID int64, seatsIDs []string, status types.SeatStatus) {
	dao.mutex.Lock()
	defer dao.mutex.Unlock()

	room := dao.theaterRoomMaps[performanceID]
	for _, zone := range room.Zones {
		for _, row := range zone.Rows {
			for k := range row.Seats {
				seat := &row.Seats[k]
				if slices.ContainsFunc(seatsIDs, func(seatID string) bool {
					return seatID == seat.SeatID
				}) {
					seat.Status = status
				}
			}
		}
	}
}

func fetchRoomForPerformance1() types.TheaterRoom {
	// Here, we can see the strong utility of a TestDataBuilder, to which we should pass for each zone:
	// - The list of row name prefixes
	// - The list of seat counts per row
	// - The list of reserved seat names
	//
	// For example:
	//
	// NewZoneBuilder()
	//     .WithRows("A", "B", "C", "D", "E", "F", "G")
	//     .WithSeatCountPerRow(7, 8, 9, 9, 10, 10, 10)
	//     .WithBookedSeats("A1", "A3", "A4", "B2")
	//     .Build();

	// Note: We are not the only application that can book seats, which explains the gaps
	// between reserved seats.
	standardZone := types.Zone{
		Category: "STANDARD",
		Rows: []types.Row{
			{
				Seats: []types.Seat{
					{SeatID: "A1", Status: types.SeatStatusBooked},
					{SeatID: "A2", Status: types.SeatStatusFree},
					{SeatID: "A3", Status: types.SeatStatusBooked},
					{SeatID: "A4", Status: types.SeatStatusBooked},
					{SeatID: "A5", Status: types.SeatStatusFree},
					{SeatID: "A6", Status: types.SeatStatusFree},
					{SeatID: "A7", Status: types.SeatStatusFree},
				},
			},
			{
				Seats: []types.Seat{
					{SeatID: "B1", Status: types.SeatStatusFree},
					{SeatID: "B2", Status: types.SeatStatusBooked},
					{SeatID: "B3", Status: types.SeatStatusFree},
					{SeatID: "B4", Status: types.SeatStatusFree},
					{SeatID: "B5", Status: types.SeatStatusFree},
					{SeatID: "B6", Status: types.SeatStatusFree},
					{SeatID: "B7", Status: types.SeatStatusFree},
					{SeatID: "B8", Status: types.SeatStatusFree},
				},
			},
			{
				Seats: []types.Seat{
					{SeatID: "C1", Status: types.SeatStatusFree},
					{SeatID: "C2", Status: types.SeatStatusFree},
					{SeatID: "C3", Status: types.SeatStatusFree},
					{SeatID: "C4", Status: types.SeatStatusFree},
					{SeatID: "C5", Status: types.SeatStatusFree},
					{SeatID: "C6", Status: types.SeatStatusFree},
					{SeatID: "C7", Status: types.SeatStatusFree},
					{SeatID: "C8", Status: types.SeatStatusFree},
					{SeatID: "C9", Status: types.SeatStatusFree},
				},
			},
			{
				Seats: []types.Seat{
					{SeatID: "D1", Status: types.SeatStatusFree},
					{SeatID: "D2", Status: types.SeatStatusFree},
					{SeatID: "D3", Status: types.SeatStatusFree},
					{SeatID: "D4", Status: types.SeatStatusFree},
					{SeatID: "D5", Status: types.SeatStatusFree},
					{SeatID: "D6", Status: types.SeatStatusFree},
					{SeatID: "D7", Status: types.SeatStatusFree},
					{SeatID: "D8", Status: types.SeatStatusFree},
					{SeatID: "D9", Status: types.SeatStatusFree},
				},
			},
			{
				Seats: []types.Seat{
					{SeatID: "E1", Status: types.SeatStatusFree},
					{SeatID: "E2", Status: types.SeatStatusFree},
					{SeatID: "E3", Status: types.SeatStatusFree},
					{SeatID: "E4", Status: types.SeatStatusFree},
					{SeatID: "E5", Status: types.SeatStatusFree},
					{SeatID: "E6", Status: types.SeatStatusFree},
					{SeatID: "E7", Status: types.SeatStatusFree},
					{SeatID: "E8", Status: types.SeatStatusFree},
					{SeatID: "E9", Status: types.SeatStatusFree},
					{SeatID: "E10", Status: types.SeatStatusFree},
				},
			},
			{
				Seats: []types.Seat{
					{SeatID: "F1", Status: types.SeatStatusFree},
					{SeatID: "F2", Status: types.SeatStatusFree},
					{SeatID: "F3", Status: types.SeatStatusFree},
					{SeatID: "F4", Status: types.SeatStatusFree},
					{SeatID: "F5", Status: types.SeatStatusFree},
					{SeatID: "F6", Status: types.SeatStatusFree},
					{SeatID: "F7", Status: types.SeatStatusFree},
					{SeatID: "F8", Status: types.SeatStatusFree},
					{SeatID: "F9", Status: types.SeatStatusFree},
					{SeatID: "F10", Status: types.SeatStatusFree},
				},
			},
			{
				Seats: []types.Seat{
					{SeatID: "G1", Status: types.SeatStatusFree},
					{SeatID: "G2", Status: types.SeatStatusFree},
					{SeatID: "G3", Status: types.SeatStatusFree},
					{SeatID: "G4", Status: types.SeatStatusFree},
					{SeatID: "G5", Status: types.SeatStatusFree},
					{SeatID: "G6", Status: types.SeatStatusFree},
					{SeatID: "G7", Status: types.SeatStatusFree},
					{SeatID: "G8", Status: types.SeatStatusFree},
					{SeatID: "G9", Status: types.SeatStatusFree},
					{SeatID: "G10", Status: types.SeatStatusFree},
				},
			},
		},
	}

	premiumZone := types.Zone{
		Category: "PREMIUM",
		Rows: []types.Row{
			{
				Seats: []types.Seat{
					{SeatID: "H1", Status: types.SeatStatusBooked},
					{SeatID: "H2", Status: types.SeatStatusFree},
					{SeatID: "H3", Status: types.SeatStatusBooked},
					{SeatID: "H4", Status: types.SeatStatusBooked},
					{SeatID: "H5", Status: types.SeatStatusFree},
					{SeatID: "H6", Status: types.SeatStatusFree},
					{SeatID: "H7", Status: types.SeatStatusFree},
				},
			},
			{
				Seats: []types.Seat{
					{SeatID: "I1", Status: types.SeatStatusFree},
					{SeatID: "I2", Status: types.SeatStatusBooked},
					{SeatID: "I3", Status: types.SeatStatusFree},
					{SeatID: "I4", Status: types.SeatStatusFree},
					{SeatID: "I5", Status: types.SeatStatusFree},
					{SeatID: "I6", Status: types.SeatStatusFree},
					{SeatID: "I7", Status: types.SeatStatusFree},
					{SeatID: "I8", Status: types.SeatStatusFree},
				},
			},
		},
	}

	room := types.TheaterRoom{
		Zones: []types.Zone{standardZone, premiumZone},
	}

	return room
}

func fetchRoomForPerformance2() types.TheaterRoom {
	standardZone := types.Zone{
		Category: "STANDARD",
		Rows: []types.Row{
			{
				Seats: []types.Seat{
					{SeatID: "R1-1", Status: types.SeatStatusBooked},
					{SeatID: "R1-2", Status: types.SeatStatusFree},
					{SeatID: "R1-3", Status: types.SeatStatusBooked},
					{SeatID: "R1-4", Status: types.SeatStatusBooked},
					{SeatID: "R1-5", Status: types.SeatStatusFree},
					{SeatID: "R1-6", Status: types.SeatStatusFree},
					{SeatID: "R1-7", Status: types.SeatStatusFree},
				},
			},
			{
				Seats: []types.Seat{
					{SeatID: "R2-1", Status: types.SeatStatusFree},
					{SeatID: "R2-2", Status: types.SeatStatusBooked},
					{SeatID: "R2-3", Status: types.SeatStatusFree},
					{SeatID: "R2-4", Status: types.SeatStatusFree},
					{SeatID: "R2-5", Status: types.SeatStatusFree},
					{SeatID: "R2-6", Status: types.SeatStatusFree},
					{SeatID: "R2-7", Status: types.SeatStatusFree},
				},
			},
		},
	}

	room := types.TheaterRoom{
		Zones: []types.Zone{standardZone},
	}

	return room
}
