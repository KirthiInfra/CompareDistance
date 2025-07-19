package main

import "testing"

func TestUnitSupportedOrNot(t *testing.T) {
	d := distance{
		value: 1000,
		unit:  "m",
	}
	err := UnitSupportedOrNot(d)
	if err != nil {
		t.Errorf("%v", err)
	}
}
