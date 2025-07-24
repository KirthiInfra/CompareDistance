package Measurement

import (
	"testing"
)

func TestCompareDistance(t *testing.T) {
	tests := []struct {
		name     string
		d1Val    float64
		d1Unit   Unit
		d2Val    float64
		d2Unit   Unit
		expected bool
	}{
		{
			name:  "1000 meters equal to 1000 meters",
			d1Val: 1000, d1Unit: meter,
			d2Val: 1000, d2Unit: meter,
			expected: true,
		},
		{
			name:  "1 meter does not equal to 2 meters",
			d1Val: 1, d1Unit: meter,
			d2Val: 2, d2Unit: meter,
			expected: false,
		},
		{
			name:  "1000 meters equal to 1 kilometer",
			d1Val: 1000, d1Unit: meter,
			d2Val: 1, d2Unit: kilometer,
			expected: true,
		},
		{
			name:  "1 kilometer equals to 1000 meters",
			d1Val: 1, d1Unit: kilometer,
			d2Val: 1000, d2Unit: meter,
			expected: true,
		},
		{
			name:  "100 centimeter equals to 1 meter",
			d1Val: 1, d1Unit: meter,
			d2Val: 100, d2Unit: centimeter,
			expected: true,
		},
		{
			name:  "10 meters equals to 1000 centimeters",
			d1Val: 10, d1Unit: meter,
			d2Val: 1000, d2Unit: centimeter,
			expected: true,
		},
		{
			name:  "5 kilometers equals to 500000 centimeters",
			d1Val: 5, d1Unit: kilometer,
			d2Val: 500000, d2Unit: centimeter,
			expected: true,
		},
		{
			name:  "200000 centimeters equals to 2 kilometers",
			d1Val: 2, d1Unit: kilometer,
			d2Val: 200000, d2Unit: centimeter,
			expected: true,
		},
		{
			name:  "1 Kilogram equals to 1000 gram",
			d1Val: 1, d1Unit: kilogram,
			d2Val: 1000, d2Unit: gram,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d1, err1 := newMeasurement(tt.d1Val, tt.d1Unit)
			d2, err2 := newMeasurement(tt.d2Val, tt.d2Unit)

			if err1 != nil || err2 != nil {
				t.Fatalf("error creating measurements: %v, %v", err1, err2)
			}

			got := d1.IsEqual(d2)
			if got != tt.expected {
				t.Errorf("IsEqual() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCreateNewMeasurementWithValidParameters(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		unit  Unit
	}{
		{"Valid Meter Unit", 1000, meter},
		{"Valid Kilometer Unit", 2, kilogram},
		{"Valid Centimeter Unit", 100, centimeter},
		{"Valid Kilogram Unit", 5, kilogram},
		{"Valid Gram Unit", 500, gram},
		{"Valid Milligram Unit", 50000, milligram},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := newMeasurement(tt.value, tt.unit)
			if err != nil {
				t.Errorf("Expected no error, got: %v", err)
			}
		})
	}
}

func TestCannotCreateDistanceWithNegativeValue(t *testing.T) {
	_, err := newMeasurement(-1, meter)
	if err == nil {
		t.Errorf("Expected error for negative value, got none")
	}
}

func TestAddTwoDistanceInMeter(t *testing.T) {
	d1, _ := newMeasurement(5, meter)
	d2, _ := newMeasurement(1, kilometer)
	add, _ := d1.Add(d2)

	if add.conversed != 1005 {
		t.Errorf("cannot expected error for invalid unit, got none")
	}
	if add.unit != meter {
		t.Errorf("cannot expected error for invalid unit, got none")
	}

}

func TestAddTwoDistanceInKilometer(t *testing.T) {

	val1, _ := newMeasurement(5, kilometer)
	d1 := distance{*val1}

	val2, _ := newMeasurement(1000, meter)
	d2 := distance{*val2}
	add, _ := d1.Add(&d2.measurement)

	if add.value != 6 {
		t.Errorf("cannot expected error for invalid unit, got none")
	}
	if add.unit != kilometer {
		t.Errorf("cannot expected error for invalid unit, got none")
	}

}

func TestAddTwoDistanceInCentimeter(t *testing.T) {
	val1, _ := newMeasurement(5, centimeter)
	d1 := distance{*val1}

	val2, _ := newMeasurement(1, meter)
	d2 := distance{*val2}
	add, _ := d1.measurement.Add(&d2.measurement)

	if add.value != 105 {
		t.Errorf("cannot expected error for invalid unit, got none")
	}
	if add.unit != centimeter {
		t.Errorf("cannot expected error for invalid unit, got none")
	}

}
