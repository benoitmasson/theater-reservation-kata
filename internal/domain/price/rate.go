package price

import "math/big"

type Rate struct {
	value *big.Float
}

func NewRateFromFloat(f float64) Rate {
	return Rate{
		value: big.NewFloat(f),
	}
}

func NewRateFromBigFloat(value *big.Float) Rate {
	return Rate{
		value: value,
	}
}

func NewDiscountPercentRateFromFloat(percent float64) Rate {
	return NewRateFromFloat(1).Substract(NewRateFromFloat(percent / 100.))
}

func (r Rate) Add(s Rate) Rate {
	f := big.NewFloat(0)
	return NewRateFromBigFloat(f.Add(r.value, s.value))
}

func (r Rate) Substract(s Rate) Rate {
	f := big.NewFloat(0)
	return NewRateFromBigFloat(f.Sub(r.value, s.value))
}

func (r Rate) Multiply(s Rate) Rate {
	f := big.NewFloat(0)
	return NewRateFromBigFloat(f.Mul(r.value, s.value))
}

func (r Rate) Equals(s Rate) bool {
	return r.value.Cmp(s.value) == 0
}

func (r Rate) AsBigFloat() *big.Float {
	return r.value
}
