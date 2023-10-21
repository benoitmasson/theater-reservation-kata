package dao

import "math/big"

type PerformancePriceDAO struct{}

func NewPerformancePriceDAO() PerformancePriceDAO {
	return PerformancePriceDAO{}
}

// FetchPerformancePrice simulates a performance pricing repository
func (dao *PerformancePriceDAO) FetchPerformancePrice(performanceID int64) *big.Float {
	if performanceID == 1 {
		return big.NewFloat(35.00)
	} else {
		return big.NewFloat(28.50)
	}
}
