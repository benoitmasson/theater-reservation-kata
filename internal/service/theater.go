package service

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/benoitmasson/theater-reservation-kata/internal/dao"
	"github.com/benoitmasson/theater-reservation-kata/internal/domain/price"
	"github.com/benoitmasson/theater-reservation-kata/internal/types"
)

type TheaterService struct {
	reservationService ReservationService

	theaterRoomsDAO     dao.TheaterRoomsDAO
	performancePriceDAO dao.PerformancePriceDAO
	voucherProgramDAO   dao.VoucherProgramDAO

	debug bool
}

func NewTheaterService(reservationDAO dao.ReservationDAO, theaterRoomsDAO dao.TheaterRoomsDAO, performancePriceDAO dao.PerformancePriceDAO, voucherProgramDAO dao.VoucherProgramDAO, debug bool) TheaterService {
	return TheaterService{
		reservationService:  NewReservationService(reservationDAO),
		theaterRoomsDAO:     theaterRoomsDAO,
		performancePriceDAO: performancePriceDAO,
		voucherProgramDAO:   voucherProgramDAO,
		debug:               debug,
	}
}

func (t *TheaterService) Reservation(customerID int64, reservationCount int, reservationCategory types.ZoneCategory, performance types.Performance) string {
	var reservation types.Reservation
	var sb strings.Builder
	var bookedSeats int
	var foundSeats []string
	seatsCategory := make(map[string]types.ZoneCategory)
	var zoneCategory types.ZoneCategory
	var remainingSeats int
	var totalSeats int
	var foundAllSeats bool

	sb.WriteString("<reservation>\n")
	sb.WriteString("\t<performance>\n")
	sb.WriteString("\t\t<play>")
	sb.WriteString(performance.Play)
	sb.WriteString("</play>\n")
	sb.WriteString("\t\t<date>")
	sb.WriteString(performance.StartTime.Format("2006-01-02"))
	sb.WriteString("</date>\n")
	sb.WriteString("\t\t<time>")
	sb.WriteString(performance.StartTime.Format("15:04:05"))
	sb.WriteString("</time>\n")
	sb.WriteString("\t</performance>\n")

	resID := t.reservationService.InitNewReservation()
	reservation.ReservationID = resID
	reservation.PerformanceID = performance.ID
	sb.WriteString("\t<reservationId>")
	sb.WriteString(strconv.FormatInt(resID, 10))
	sb.WriteString("</reservationId>\n")

	room := t.theaterRoomsDAO.FetchTheaterRoom(performance.ID)

	// find "reservationCount" first contiguous seats in any row
	for _, zone := range room.Zones {
		zoneCategory = zone.Category
		for _, row := range zone.Rows {
			seatsForRow := make([]string, 0, reservationCount)
			streakOfNotReservedSeats := 0
			for _, aSeat := range row.Seats {
				totalSeats++
				if aSeat.Status != types.SeatStatusBooked && aSeat.Status != types.SeatStatusBookingPending {
					remainingSeats++
					if reservationCategory != zoneCategory {
						continue
					}
					if !foundAllSeats {
						seatsForRow = append(seatsForRow, aSeat.SeatID)
						streakOfNotReservedSeats++
						if streakOfNotReservedSeats >= reservationCount {
							for _, seat := range seatsForRow {
								foundSeats = append(foundSeats, seat)
								seatsCategory[seat] = zoneCategory
							}
							foundAllSeats = true
							remainingSeats -= streakOfNotReservedSeats
						}
					}
				} else {
					seatsForRow = make([]string, 0, reservationCount)
					streakOfNotReservedSeats = 0
				}
			}
			if foundAllSeats {
				for _, seat := range row.Seats {
					bookedSeats++
					if slices.ContainsFunc(foundSeats, func(seatID string) bool {
						return seatID == seat.SeatID
					}) {
						if t.debug {
							fmt.Printf("MIAOU!!! : Seat %s will be saved as %s\n", seat.SeatID, types.SeatStatusBookingPending)
						}
					}
				}

				t.theaterRoomsDAO.SaveSeats(performance.ID, foundSeats, types.SeatStatusBookingPending)
			}
		}
	}
	reservation.Seats = foundSeats

	if foundAllSeats {
		reservation.Status = types.ReservationStatusPending
	} else {
		reservation.Status = types.ReservationStatusAborted
	}

	t.reservationService.Update(reservation)

	if performance.PerformanceNature == types.PerformanceNaturePremiere && remainingSeats < int(math.Floor(float64(totalSeats)*0.5)) {
		// keep 50% seats for VIP
		foundSeats = []string{}
		fmt.Println("Not enough VIP seats available for Premiere")
	} else if performance.PerformanceNature == types.PerformanceNaturePreview && remainingSeats < int(math.Floor(float64(totalSeats)*0.9)) {
		// keep 10% seats for VIP
		foundSeats = []string{}
		fmt.Println("Not enough VIP seats available for Preview")
	}

	if len(foundSeats) > 0 {
		sb.WriteString("\t<reservationStatus>FULFILLABLE</reservationStatus>\n")
		sb.WriteString("\t<seats>\n")
		for _, s := range foundSeats {
			sb.WriteString("\t\t<seat>\n")
			sb.WriteString("\t\t\t<id>")
			sb.WriteString(s)
			sb.WriteString("</id>\n")
			sb.WriteString("\t\t\t<category>")
			sb.WriteString(string(seatsCategory[s]))
			sb.WriteString("</category>\n")
			sb.WriteString("\t\t</seat>\n")
		}
		sb.WriteString("\t</seats>\n")
	} else {
		sb.WriteString("\t<reservationStatus>ABORTED</reservationStatus>\n")
	}

	// calculate raw price
	myPrice := price.NewAmountFromBigFloat(t.performancePriceDAO.FetchPerformancePrice(performance.ID))

	initialPrice := price.NewEmptyAmount()
	for _, foundSeat := range foundSeats {
		categoryRatio := price.NewRateFromFloat(1)
		if seatsCategory[foundSeat] == types.ZoneCategoryPremium {
			categoryRatio = price.NewRateFromFloat(1.5)
		}
		initialPrice = initialPrice.Add(myPrice.Apply(categoryRatio))
	}

	// check and apply discounts and fidelity program
	discountTime := price.NewRateFromBigFloat(t.voucherProgramDAO.FetchVoucherProgram(performance.StartTime.UTC()))

	// has he subscribed or not
	customerSubscriptionDAO := dao.CustomerSubscriptionDAO{}
	isSubscribed := customerSubscriptionDAO.FetchCustomerSubscription(customerID)

	totalBilling := initialPrice
	if isSubscribed {
		// apply a 25% discount when the user is subscribed
		totalBilling = initialPrice.Apply(price.NewDiscountPercentRateFromFloat(17.5))
	}
	totalBilling = totalBilling.Apply(price.NewRateFromFloat(1).Substract(discountTime))
	total := totalBilling.String() + "€"

	sb.WriteString("\t<seatCategory>")
	sb.WriteString(string(reservationCategory))
	sb.WriteString("</seatCategory>\n")
	sb.WriteString("\t<totalAmountDue>")
	sb.WriteString(total)
	sb.WriteString("</totalAmountDue>\n")
	sb.WriteString("</reservation>\n")
	return sb.String()
}

func (t *TheaterService) CancelReservation(reservationID int64, performanceID int64, seatsIDs []string) {
	t.theaterRoomsDAO.SaveSeats(performanceID, seatsIDs, types.SeatStatusFree)
	t.reservationService.Cancel(reservationID)
}
