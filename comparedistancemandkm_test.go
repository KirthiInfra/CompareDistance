package main

import "testing"

func TestUnitSupportedOrNot(t *testing.T) {
	d1 := distance{
		value: 1000,
		unit:  "m",
	}
	d2 := distance{
		value: 1,
		unit:  "km",
	}
	err := UnitSupportedOrNot(&d1, d2)
	if err != nil {
		t.Errorf("%v", err)
	}
}
