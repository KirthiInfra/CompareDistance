package Measurement

import (
	"testing"
)

func TestCompareDistance(t *testing.T) {
	tests := []struct {
		name     string
		d1       distance
		d2       distance
		expected bool
	}{
		{
			name:     "1000 meters equal to 1000 meters",
			d1:       distance{value: 1000, unit: m},
			d2:       distance{value: 1000, unit: m},
			expected: true,
		},
		{
			name:     "1 meter does not equal to 2 meters",
			d1:       distance{value: 1, unit: m},
			d2:       distance{value: 2, unit: m},
			expected: false,
		},
		{
			name:     "1000 meters equal to 1 kilometer",
			d1:       distance{value: 1000, unit: m},
			d2:       distance{value: 1, unit: km},
			expected: true,
		},
		{
			name:     "1 kilometer equals to 1000 meters",
			d1:       distance{value: 1, unit: km},
			d2:       distance{value: 1000, unit: m},
			expected: true,
		},
		{
			name:     "100 centimeter equals to 1 meter",
			d1:       distance{value: 1, unit: m},
			d2:       distance{value: 100, unit: cm},
			expected: true,
		},
		{
			name:     "10 meters equals to 1000 centimeters",
			d1:       distance{value: 1000, unit: cm},
			d2:       distance{value: 10, unit: m},
			expected: true,
		},
		{
			name:     "5 kilometers equals to 500000 centimeters",
			d1:       distance{value: 500000, unit: cm},
			d2:       distance{value: 5, unit: km},
			expected: true,
		},
		{
			name:     "200000 centimeters equals to 2 kilometers",
			d1:       distance{value: 2, unit: km},
			d2:       distance{value: 200000, unit: cm},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.d1.IsDistanceEqual(&tt.d2)
			if got != tt.expected {
				t.Errorf("IsDistanceEqual() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCreateDistanceWithValidParameters(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		unit  unit
	}{
		{"Valid Meter Unit", 1000, m},
		{"Valid Kilometer Unit", 2, km},
		{"Valid Centimeter Unit", 100, cm},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewDistance(tt.value, tt.unit)
			if err != nil {
				t.Errorf("Expected no error, got: %v", err)
			}
		})
	}
}

func TestCannotCreateDistanceWithNegativeValue(t *testing.T) {
	_, err := NewDistance(-1, m)
	if err == nil {
		t.Errorf("Expected error for negative value, got none")
	}
}

func TestCannotCreateDistanceWithInvalidUnit(t *testing.T) {
	_, err := NewDistance(1, "kmm")
	if err == nil {
		t.Errorf("Expected error for invalid unit, got none")
	}
}

func TestCreateDistanceWithValidUnitCm(t *testing.T) {
	_, err := NewDistance(1, "cm")
	if err != nil {
		t.Errorf("Expected error for invalid unit, got none")
	}
}

func TestAddTwoDistanceInMeter(t *testing.T) {
	d1 := distance{value: 5, unit: m}
	d2 := distance{value: 1, unit: km}
	add, unit := d1.Add(&d2)

	if add != 1005 {
		t.Errorf("cannot expected error for invalid unit, got none")
	}
	if unit != m {
		t.Errorf("cannot expected error for invalid unit, got none")
	}

}

func TestAddTwoDistanceInKilometer(t *testing.T) {
	d1 := distance{value: 5, unit: km}
	d2 := distance{value: 1000, unit: m}
	add, unit := d1.Add(&d2)

	if add != 6 {
		t.Errorf("cannot expected error for invalid unit, got none")
	}
	if unit != km {
		t.Errorf("cannot expected error for invalid unit, got none")
	}

}

func TestAddTwoDistanceInCentimeter(t *testing.T) {
	d1 := distance{value: 5, unit: cm}
	d2 := distance{value: 1, unit: m}
	add, unit := d1.Add(&d2)

	if add != 105 {
		t.Errorf("cannot expected error for invalid unit, got none")
	}
	if unit != cm {
		t.Errorf("cannot expected error for invalid unit, got none")
	}

}
