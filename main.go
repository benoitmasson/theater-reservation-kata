package main

import (
	"fmt"
	"time"

	"github.com/benoitmasson/theater-reservation-kata/internal/dao"
	"github.com/benoitmasson/theater-reservation-kata/internal/service"
	"github.com/benoitmasson/theater-reservation-kata/internal/types"
)

func main() {
	theaterService := service.NewTheaterService(
		dao.NewReservationDAO(),
		dao.NewTheaterRoomsDAO(),
		dao.NewPerformancePriceDAO(),
		dao.NewVoucherProgramDAO(),
		false, /* debug */
	)

	performance := types.Performance{
		ID:                1,
		Play:              "The CICD by Corneille",
		StartTime:         time.Date(2023, time.April, 22, 21, 0, 0, 0, time.UTC),
		PerformanceNature: types.PerformanceNaturePremiere,
	}
	fmt.Println(theaterService.Reservation(1, 4, types.ZoneCategoryStandard, performance))
	fmt.Println(theaterService.Reservation(1, 5, types.ZoneCategoryStandard, performance))

	performance2 := types.Performance{
		ID:                2,
		Play:              "Les fourberies de Scala - Molière",
		StartTime:         time.Date(2023, time.March, 21, 21, 0, 0, 0, time.UTC),
		PerformanceNature: types.PerformanceNaturePreview,
	}
	fmt.Println(theaterService.Reservation(2, 4, types.ZoneCategoryStandard, performance2))
}
