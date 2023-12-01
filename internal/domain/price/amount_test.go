package price

import "testing"

func TestAddAmount(t *testing.T) {
	a := NewAmountFromFloat(25.99).Add(NewAmountFromFloat(59.99))
	b := NewAmountFromFloat(85.98)
	if !a.Equals(b) {
		t.Errorf("expected %v, got %v", b.AsBigFloat(), a.AsBigFloat())
	}
}

func TestAddNothingAmount(t *testing.T) {
	a := NewAmountFromFloat(25.99)
	b := a.Add(NewEmptyAmount())
	if !a.Equals(b) {
		t.Errorf("expected %v, got %v", a.AsBigFloat(), b.AsBigFloat())
	}
}

func TestApplyRateToAmount(t *testing.T) {
	a := NewAmountFromFloat(25.95).Apply(NewDiscountPercentRateFromFloat(20.))
	b := NewAmountFromFloat(20.76)
	if !a.Equals(b) {
		t.Errorf("expected %v, got %v", b.AsBigFloat(), a.AsBigFloat())
	}
}
