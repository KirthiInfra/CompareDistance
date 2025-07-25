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
	result := d1.Add(d2)

	if result.conversed != 1005 {
		t.Errorf("cannot expected error for invalid unit, got none")
	}
	if result.unit != meter {
		t.Errorf("cannot expected error for invalid unit, got none")
	}

}

func TestAddTwoDistanceInKilometer(t *testing.T) {

	val1, _ := newMeasurement(5, kilometer)
	d1 := distance{*val1}

	val2, _ := newMeasurement(1000, meter)
	d2 := distance{*val2}
	result := d1.Add(&d2)

	if result.value != 6 {
		t.Errorf("cannot expected error for invalid unit, got none")
	}
	if result.unit != kilometer {
		t.Errorf("cannot expected error for invalid unit, got none")
	}

}

func TestAddTwoDistanceInCentimeter(t *testing.T) {
	val1, _ := newMeasurement(5, centimeter)
	d1 := distance{*val1}

	val2, _ := newMeasurement(1, meter)
	d2 := distance{*val2}
	result:= d1.measurement.Add(&d2.measurement)

	if result.value != 105 {
		t.Errorf("cannot expected error for invalid unit, got none")
	}
	if result.unit != centimeter {
		t.Errorf("cannot expected error for invalid unit, got none")
	}

}

func TestNewDistance(t *testing.T) {
	_, err := NewDistance(10, meter)
	if err != nil {
		t.Errorf("Could not create distance with 10 meter units")
	}
}

func Test100MeterEquals100Meter(t *testing.T) {
	hundredMeter1, _ := NewDistance(100, meter)
	hundredMeter2, _ := NewDistance(100, meter)
	result := hundredMeter1.IsEqual(hundredMeter2)
	if result != true {
		t.Errorf("100 meter should be equal to 100 meter")
	}
}

func Test1KilometerEquals1000Meter(t *testing.T) {
	oneKilometer, _:=NewDistance(1, kilometer)
	thousandMeter, _:=NewDistance(1000, meter)
	result := oneKilometer.IsEqual(thousandMeter)
	if result != true {
		t.Errorf("1 kilometer should be equal to 1000 meter")
	}
}

func TestAdd100CentimeterAnd100Centimeter(t *testing.T){
	hundredCentimeter1, _ := NewDistance(100, centimeter)
	hundredCentimeter2, _ := NewDistance(100, centimeter)
	result := hundredCentimeter1.Add(hundredCentimeter2)

	if result.conversed != 2 {
		t.Errorf("cannot expected error for invalid unit, got none")
	}
	if result.unit != centimeter {
		t.Errorf("cannot expected error for invalid unit, got none")
	}
}

func TestNewWeight(t *testing.T) {
	_, err := NewWeight(100, gram)
	if err != nil {
		t.Errorf("Could not create weight with 100 gram units")
	}
}

func Test1KilogramEquals1Kilogram(t *testing.T) {
	oneKilogram1, _ := NewWeight(1, kilogram)
	oneKilogram2, _ := NewWeight(1, kilogram)
	result := oneKilogram1.IsEqual(oneKilogram2)
	if result != true {
		t.Errorf("1 kilogram should be equal to 1 kilogram")
	}
}

func Test1000GramEquals1Kilogram(t *testing.T) {
	thousandGram, _:=NewWeight(1000, gram)
	oneKilogram, _:=NewWeight(1, kilogram)
	result := thousandGram.IsEqual(oneKilogram)
	if result != true {
		t.Errorf("1000 grams should be equal to 1 kilogram")
	}
}

func TestAdd100GramAnd1Kilogram(t *testing.T){
	hundredGram, _ := NewWeight(100, gram)
	oneKilogram, _ := NewWeight(1, kilogram)
	result := hundredGram.Add(oneKilogram)

	if result.conversed != 1100 {
		t.Errorf("expected value in gram: 1100, but actual value in gram: %f", result.conversed)
	}
	if result.unit != gram{
		t.Errorf("expected unit: gram, but actual unit: %v", result.unit)
	}
}

func TestNewTemperature(t *testing.T) {
	_, err := NewTemperature(10, celsius)
	if err != nil {
		t.Errorf("Could not create distance with 10 meter units")
	}
}

func Test0CelsiusEquals0Celsius(t *testing.T) {
	zeroCelsius1, _ := NewTemperature(0, celsius)
	zeroCelsius2, _ := NewTemperature(0, celsius)
	result := zeroCelsius1.IsEqual(zeroCelsius2)
	if result != true {
		t.Errorf("100 meter should be equal to 100 meter")
	}
}

func TestAdd1CelsiusAnd1Celsius(t *testing.T){
	oneCelsius1, _ := NewTemperature(1, celsius)
	oneCelsius2, _ := NewTemperature(2, celsius)
	result := oneCelsius1.Add(oneCelsius2)

	if result != nil {
		t.Errorf("Temperature cannot be added")
	}
	
}