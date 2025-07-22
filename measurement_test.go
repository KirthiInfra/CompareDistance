package Measurement

import (
	"testing"
)
func TestCompareDistance(t *testing.T) {
	tests := []struct {
		name     string
		d1Val    float64
		d1Unit   unit
		d2Val    float64
		d2Unit   unit
		expected bool
	}{
		{
			name:     "1000 meters equal to 1000 meters",
			d1Val:    1000, d1Unit: m,
			d2Val:    1000, d2Unit: m,
			expected: true,
		},
		{
			name:     "1 meter does not equal to 2 meters",
			d1Val:    1, d1Unit: m,
			d2Val:    2, d2Unit: m,
			expected: false,
		},
		{
			name:     "1000 meters equal to 1 kilometer",
			d1Val:    1000, d1Unit: m,
			d2Val:    1, d2Unit: km,
			expected: true,
		},
		{
			name:     "1 kilometer equals to 1000 meters",
			d1Val:    1, d1Unit: km,
			d2Val:    1000, d2Unit: m,
			expected: true,
		},
		{
			name:     "100 centimeter equals to 1 meter",
			d1Val:    1, d1Unit: m,
			d2Val:    100, d2Unit: cm,
			expected: true,
		},
		{
			name:     "10 meters equals to 1000 centimeters",
			d1Val:    10, d1Unit: m,
			d2Val:    1000, d2Unit: cm,
			expected: true,
		},
		{
			name:     "5 kilometers equals to 500000 centimeters",
			d1Val:    5, d1Unit: km,
			d2Val:    500000, d2Unit: cm,
			expected: true,
		},
		{
			name:     "200000 centimeters equals to 2 kilometers",
			d1Val:    2, d1Unit: km,
			d2Val:    200000, d2Unit: cm,
			expected: true,
		},
		{
			name:     "1 Kilogram equals to 1000 gram",
			d1Val:    1, d1Unit: kg,
			d2Val:    1000, d2Unit: g,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d1, err1 := NewMeasurement(tt.d1Val, tt.d1Unit)
			d2, err2 := NewMeasurement(tt.d2Val, tt.d2Unit)

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
		unit  unit
	}{
		{"Valid Meter Unit", 1000, m},
		{"Valid Kilometer Unit", 2, km},
		{"Valid Centimeter Unit", 100, cm},
		{"Valid Kilogram Unit", 5, kg},
		{"Valid Gram Unit", 500, g},
		{"Valid Milligram Unit", 50000, mg},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewMeasurement(tt.value, tt.unit)
			if err != nil {
				t.Errorf("Expected no error, got: %v", err)
			}
		})
	}
}

func TestCannotCreateDistanceWithNegativeValue(t *testing.T) {
	_, err := NewMeasurement(-1, m)
	if err == nil {
		t.Errorf("Expected error for negative value, got none")
	}
}

func TestCannotCreateDistanceWithInvalidUnit(t *testing.T) {
	_, err := NewMeasurement(1, "kmm")
	if err == nil {
		t.Errorf("Expected error for invalid unit, got none")
	}
}

func TestCreateDistanceWithValidUnitCm(t *testing.T) {
	_, err := NewMeasurement(1, "cm")
	if err != nil {
		t.Errorf("Expected error for invalid unit, got none")
	}
}

func TestAddTwoDistanceInMeter(t *testing.T) {
	d1, _ := NewMeasurement(5, m)
	d2, _ := NewMeasurement(1, km)
	add := d1.Add(d2)

	if add.conversed != 1005 {
		t.Errorf("cannot expected error for invalid unit, got none")
	}
	if add.unit != m {
		t.Errorf("cannot expected error for invalid unit, got none")
	}

}

func TestAddTwoDistanceInKilometer(t *testing.T) {

	val1,_ := NewMeasurement(5,km)
	d1 := distance{*val1}

	val2,_ := NewMeasurement(1000,m)
	d2 := distance{*val2}
	add:= d1.Add(&d2.measurement)

	if add.value != 6 {
		t.Errorf("cannot expected error for invalid unit, got none")
	}
	if add.unit != km {
		t.Errorf("cannot expected error for invalid unit, got none")
	}

}

func TestAddTwoDistanceInCentimeter(t *testing.T) {
	val1,_ := NewMeasurement(5,cm)
	d1 := distance{*val1}

	val2,_ := NewMeasurement(1,m)
	d2 := distance{*val2}
	add := d1.measurement.Add(&d2.measurement)

	if add.value != 105 {
		t.Errorf("cannot expected error for invalid unit, got none")
	}
	if add.unit != cm {
		t.Errorf("cannot expected error for invalid unit, got none")
	}

}

func TestCannotCreateWeightWithInvalidUnit(t *testing.T) {
	_, err := NewMeasurement(1, "kgg")
	if err == nil {
		t.Errorf("Expected error for invalid unit, got none")
	}
}

