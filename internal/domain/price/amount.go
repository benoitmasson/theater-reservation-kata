package price

import (
	"fmt"
	"math"
	"math/big"
)

type Amount struct {
	value *big.Float
}

func NewEmptyAmount() Amount {
	return Amount{
		value: big.NewFloat(0),
	}
}

func NewAmountFromBigFloat(value *big.Float) Amount {
	return Amount{
		value: value,
	}
}

func (a Amount) Add(b Amount) Amount {
	f := big.NewFloat(0)
	return NewAmountFromBigFloat(f.Add(a.value, b.value))
}

func (a Amount) Apply(r Rate) Amount {
	f := big.NewFloat(0)
	return NewAmountFromBigFloat(f.Mul(a.value, r.value))
}

func (a Amount) AsBigFloat() *big.Float {
	return a.value
}

func (a Amount) String() string {
	aFloat, _ := a.value.Float64()
	aFloat = math.Round(100*aFloat) / 100
	return fmt.Sprintf("%.2f", aFloat)
}
