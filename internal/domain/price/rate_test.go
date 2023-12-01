package price

import "testing"

func TestAddRate(t *testing.T) {
	r := NewRateFromFloat(0.15).Add(NewRateFromFloat(0.2))
	s := NewRateFromFloat(0.35)
	if !r.Equals(s) {
		t.Errorf("expected %v, got %v", s.AsBigFloat(), r.AsBigFloat())
	}
}

func TestMultiplyRate(t *testing.T) {
	r := NewRateFromFloat(0.15).Multiply(NewRateFromFloat(0.5))
	s := NewRateFromFloat(0.075)
	if !r.Equals(s) {
		t.Errorf("expected %v, got %v", s.AsBigFloat(), r.AsBigFloat())
	}
}

func TestDiscountPercentRate(t *testing.T) {
	r := NewDiscountPercentRateFromFloat(25.)
	s := NewRateFromFloat(0.75)
	if !r.Equals(s) {
		t.Errorf("expected %v, got %v", s.AsBigFloat(), r.AsBigFloat())
	}
}
