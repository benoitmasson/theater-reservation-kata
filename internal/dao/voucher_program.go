package dao

import (
	"math/big"
	"time"
)

type VoucherProgramDAO struct{}

func NewVoucherProgramDAO() VoucherProgramDAO {
	return VoucherProgramDAO{}
}

// FetchVoucherProgram simulates a voucher program repository
func (dao *VoucherProgramDAO) FetchVoucherProgram(reservationDate time.Time) *big.Float {
	voucher := big.NewFloat(0)

	// applies from reservation date, not performance date
	if reservationDate.Before(time.Date(2023, time.April, 30, 0, 0, 0, 0, time.UTC)) {
		voucher.SetFloat64(0.20)
	}

	return voucher
}
