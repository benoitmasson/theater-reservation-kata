package service

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/andreyvit/diff"

	"github.com/benoitmasson/theater-reservation-kata/internal/dao"
	"github.com/benoitmasson/theater-reservation-kata/internal/types"
)

var (
	reservationDAO      = dao.NewReservationDAO()
	theaterRoomsDAO     = dao.NewTheaterRoomsDAO()
	performancePriceDAO = dao.NewPerformancePriceDAO()
	voucherProgramDAO   = dao.NewVoucherProgramDAO()

	theaterService     = NewTheaterService(reservationDAO, theaterRoomsDAO, performancePriceDAO, voucherProgramDAO, false)
	reservationService = NewReservationService(reservationDAO)

	performanceCICD = types.Performance{
		ID:                1,
		Play:              "The CICD by Corneille",
		StartTime:         time.Date(2023, time.April, 22, 21, 0, 0, 0, time.UTC),
		PerformanceNature: types.PerformanceNaturePremiere,
	}
	performanceScala = types.Performance{
		ID:                2,
		Play:              "Les fourberies de Scala - Molière",
		StartTime:         time.Date(2023, time.March, 21, 21, 0, 0, 0, time.UTC),
		PerformanceNature: types.PerformanceNaturePreview,
	}
	performanceJSON = types.Performance{
		ID:                3,
		Play:              "DOM JSON - Molière",
		StartTime:         time.Date(2023, time.March, 21, 21, 0, 0, 0, time.UTC),
		PerformanceNature: types.PerformanceNaturePremiere,
	}
)

func TestTheaterReservation(t *testing.T) {
	type cancelBefore struct {
		reservationID int64
		performanceID int64
		seatsIDs      []string
	}
	type reservation struct {
		nbSeats      int
		zoneCategory types.ZoneCategory
		expectedID   int64
	}
	type test struct {
		name         string
		customerID   int64
		performance  types.Performance
		reservations []reservation
		cancelBefore *cancelBefore
	}
	tests := []test{
		{
			name:        "reservation_failed_on_premiere_performance",
			customerID:  2,
			performance: performanceJSON,
			reservations: []reservation{{
				nbSeats:      4,
				zoneCategory: types.ZoneCategoryStandard,
				expectedID:   123456,
			}},
			cancelBefore: nil,
		},
		{
			name:        "cancel_then_reserve_on_premiere_performance_with_standard_category",
			customerID:  1,
			performance: performanceCICD,
			reservations: []reservation{{
				nbSeats:      4,
				zoneCategory: types.ZoneCategoryStandard,
				expectedID:   123457,
			}},
			cancelBefore: &cancelBefore{
				reservationID: 123456,
				performanceID: 1,
				seatsIDs:      []string{"B2"},
			},
		},
		{
			name:        "reservation_failed_on_preview_performance",
			customerID:  2,
			performance: performanceScala,
			reservations: []reservation{{
				nbSeats:      4,
				zoneCategory: types.ZoneCategoryStandard,
				expectedID:   123458,
			}},
			cancelBefore: nil,
		},
		{
			name:        "reserve_once_on_premiere_performance_with_premium_category",
			customerID:  1,
			performance: performanceCICD,
			reservations: []reservation{{
				nbSeats:      4,
				zoneCategory: types.ZoneCategoryPremium,
				expectedID:   123459,
			}},
			cancelBefore: nil,
		},
		{
			name:        "reserve_once_on_premiere_performance",
			customerID:  1,
			performance: performanceCICD,
			reservations: []reservation{{
				nbSeats:      4,
				zoneCategory: types.ZoneCategoryStandard,
				expectedID:   123460,
			}},
			cancelBefore: nil,
		},
		{
			name:        "reserve_twice_on_premiere_performance",
			customerID:  1,
			performance: performanceCICD,
			reservations: []reservation{
				{
					nbSeats:      4,
					zoneCategory: types.ZoneCategoryStandard,
					expectedID:   123461,
				},
				{
					nbSeats:      5,
					zoneCategory: types.ZoneCategoryStandard,
					expectedID:   123462,
				},
			},
			cancelBefore: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.cancelBefore != nil {
				theaterService.CancelReservation(test.cancelBefore.reservationID, test.cancelBefore.performanceID, test.cancelBefore.seatsIDs)
			}
			var actualXML string
			for _, reservation := range test.reservations {
				actualXML = theaterService.Reservation(test.customerID, reservation.nbSeats, reservation.zoneCategory, test.performance)

				// TODO: Add testing for reserved seat references
				if reservationService.Find(reservation.expectedID) == nil {
					t.Errorf("Reservation #%d not found", reservation.expectedID)
				}
			}

			verifyXML(t, actualXML, test.name)
		})
	}
}

func verifyXML(t *testing.T, actualXML string, referenceFilePrefix string) {
	referenceFolder := "testdata"
	err := os.MkdirAll(referenceFolder, 0o755)
	if err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}

	referenceFileName := filepath.Join(referenceFolder, fmt.Sprintf("%s.approved.xml", referenceFilePrefix))
	expected, err := os.ReadFile(referenceFileName)
	if err != nil && !os.IsNotExist(err) {
		t.Fatalf("Failed to read file: %v", err)
	}
	expectedXML := string(expected)

	updateApprovals, _ := strconv.ParseBool(os.Getenv("UPDATE_APPROVALS"))
	if os.IsNotExist(err) || updateApprovals {
		err = os.WriteFile(referenceFileName, []byte(actualXML), 0o644)
		if err != nil {
			t.Fatalf("Failed to write file: %v", err)
		}
		return
	}

	if actualXML != expectedXML {
		t.Errorf("Approval test failed, %s does not match:\n%v", referenceFileName, diffLinesToString(diff.LineDiffAsLines(expectedXML, actualXML)))
	}
}

func diffLinesToString(diff []string) string {
	sb := &strings.Builder{}
	i := 0
	for _, line := range diff {
		i++
		if len(line) > 0 && line[0] == '-' {
			fmt.Fprintf(sb, "%2d: %s\n", i, line)
		}
		if len(line) > 0 && line[0] == '+' {
			fmt.Fprintf(sb, "    %s\n", line)
			i--
		}
	}
	return sb.String()
}
