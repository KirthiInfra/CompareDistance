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
			d1Val: 1000, d1Unit: Meter,
			d2Val: 1000, d2Unit: Meter,
			expected: true,
		},
		{
			name:  "1 Meter does not equal to 2 meters",
			d1Val: 1, d1Unit: Meter,
			d2Val: 2, d2Unit: Meter,
			expected: false,
		},
		{
			name:  "1000 meters equal to 1 kilometer",
			d1Val: 1000, d1Unit: Meter,
			d2Val: 1, d2Unit: Kilometer,
			expected: true,
		},
		{
			name:  "1 kilometer equals to 1000 meters",
			d1Val: 1, d1Unit: Kilometer,
			d2Val: 1000, d2Unit: Meter,
			expected: true,
		},
		{
			name:  "100 Centimeter equals to 1 meter",
			d1Val: 1, d1Unit: Meter,
			d2Val: 100, d2Unit: Centimeter,
			expected: true,
		},
		{
			name:  "10 meters equals to 1000 Centimeters",
			d1Val: 10, d1Unit: Meter,
			d2Val: 1000, d2Unit: Centimeter,
			expected: true,
		},
		{
			name:  "5 Kilometers equals to 500000 Centimeters",
			d1Val: 5, d1Unit: Kilometer,
			d2Val: 500000, d2Unit: Centimeter,
			expected: true,
		},
		{
			name:  "200000 Centimeters equals to 2 Kilometers",
			d1Val: 2, d1Unit: Kilometer,
			d2Val: 200000, d2Unit: Centimeter,
			expected: true,
		},
		{
			name:  "1 kilogram equals to 1000 gram",
			d1Val: 1, d1Unit: Kilogram,
			d2Val: 1000, d2Unit: Gram,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d1, err1 := NewDistance(tt.d1Val, tt.d1Unit)
			d2, err2 := NewDistance(tt.d2Val, tt.d2Unit)

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
		{"Valid Meter Unit", 1000, Meter},
		{"Valid Kilometer Unit", 2, Kilogram},
		{"Valid Centimeter Unit", 100, Centimeter},
		{"Valid Kilogram Unit", 5, Kilogram},
		{"Valid Gram Unit", 500, Gram},
		{"Valid Milligram Unit", 50000, Milligram},
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
	_, err := NewDistance(-1, Meter)
	if err == nil {
		t.Errorf("Expected error for negative value, got none")
	}
}

func TestAddTwoDistanceInMeter(t *testing.T) {
	fiveMeter, _ := NewDistance(5, Meter)
	oneKilometer, _ := NewDistance(1, Kilometer)
	result := fiveMeter.Add(oneKilometer)

	if result.value != 1005 {
		t.Errorf("Expected distance in meter: 1005 but got: %f", result.value)
	}
	if result.unit != Meter {
		t.Errorf("cannot expected error for invalid unit, got none")
	}

}

func TestAddTwoDistanceInKilometer(t *testing.T) {

	d1, _ := NewDistance(5, Kilometer)
	d2, _ := NewDistance(1000, Meter)

	result := d1.Add(d2)

	if result.value != 6 {
		t.Errorf("cannot expected error for invalid unit, got none")
	}
	if result.unit != Kilometer {
		t.Errorf("cannot expected error for invalid unit, got none")
	}

}

func TestAddTwoDistanceInCentimeter(t *testing.T) {
	d1, _ := NewDistance(5, Centimeter)
	d2, _ := NewDistance(1, Meter)

	result := d1.Add(d2)

	if result.value != 105 {
		t.Errorf("cannot expected error for invalid unit, got none")
	}
	if result.unit != Centimeter {
		t.Errorf("cannot expected error for invalid unit, got none")
	}

}

func TestNewDistance(t *testing.T) {
	_, err := NewDistance(10, Meter)
	if err != nil {
		t.Errorf("Could not create distance with 10 meter units")
	}
}

func Test100MeterEquals100Meter(t *testing.T) {
	hundredMeter1, _ := NewDistance(100, Meter)
	hundredMeter2, _ := NewDistance(100, Meter)
	result := hundredMeter1.IsEqual(hundredMeter2)
	if result != true {
		t.Errorf("100 meter should be equal to 100 meter")
	}
}

func Test1KilometerEquals1000Meter(t *testing.T) {
	oneKilometer, _ := NewDistance(1, Kilometer)
	thousandMeter, _ := NewDistance(1000, Meter)
	result := oneKilometer.IsEqual(thousandMeter)
	if result != true {
		t.Errorf("1 Kilometer should be equal to 1000 meter")
	}
}

func TestAdd100CentimeterAnd100Centimeter(t *testing.T) {
	hundredCentimeter1, _ := NewDistance(100, Centimeter)
	hundredCentimeter2, _ := NewDistance(100, Centimeter)
	result := hundredCentimeter1.Add(hundredCentimeter2)

	if result.value != 200 {
		t.Errorf("cannot expected error for invalid unit, got none")
	}
	if result.unit != Centimeter {
		t.Errorf("cannot expected error for invalid unit, got none")
	}
}

func TestNewWeight(t *testing.T) {
	_, err := NewWeight(100, Gram)
	if err != nil {
		t.Errorf("Could not create weight with 100 gram units")
	}
}

func Test1KilogramEquals1Kilogram(t *testing.T) {
	oneKilogram1, _ := NewWeight(1, Kilogram)
	oneKilogram2, _ := NewWeight(1, Kilogram)
	result := oneKilogram1.IsEqual(oneKilogram2)
	if result != true {
		t.Errorf("1  Kilogram should be equal to 1  Kilogram")
	}
}

func Test1000GramEquals1Kilogram(t *testing.T) {
	thousandGram, _ := NewWeight(1000, Gram)
	oneKilogram, _ := NewWeight(1, Kilogram)
	result := thousandGram.IsEqual(oneKilogram)
	if result != true {
		t.Errorf("1000 grams should be equal to 1  Kilogram")
	}
}

func TestAdd100GramAnd1Kilogram(t *testing.T) {
	hundredGram, _ := NewWeight(100, Gram)
	oneKilogram, _ := NewWeight(1, Kilogram)
	result := hundredGram.Add(oneKilogram)

	if result.value != 1100 {
		t.Errorf("expected value in Gram: 1100, but actual value in gram: %f", result.value)
	}
	if result.unit != Gram {
		t.Errorf("expected unit: Gram, but actual unit: %v", result.unit)
	}
}

func TestNewTemperature(t *testing.T) {
	_, err := NewTemperature(10, Celsius)
	if err != nil {
		t.Errorf("Could not create temperature with 10 degree Celsius")
	}
}

func Test0CelsiusEquals0Celsius(t *testing.T) {
	zeroCelsius1, _ := NewTemperature(0, Celsius)
	zeroCelsius2, _ := NewTemperature(0, Celsius)
	result := zeroCelsius1.IsEqual(zeroCelsius2)
	if result != true {
		t.Errorf("0 degree Celsius should be equal to 0 degree Celsius")
	}
}

func TestNewTemperatureWithFahrenheit(t *testing.T) {
	_, err := NewTemperature(10, Fahrenheit)
	if err != nil {
		t.Errorf("Could not create temperature with 10 degree Fahrenheit")
	}
}

func Test32FahrenheitEquals32Fahrenheit(t *testing.T) {
	thirtyTwoFahrenheit1, _ := NewTemperature(32, Fahrenheit)
	thirtyTwoFahrenheit2, _ := NewTemperature(32, Fahrenheit)
	result := thirtyTwoFahrenheit1.IsEqual(thirtyTwoFahrenheit2)
	if result != true {
		t.Errorf("32 degree Fahrenheit should be equal to 32 degree Fahrenheit")
	}
}

func Test0CelsiusEquals32Fahrenheit(t *testing.T) {
	zeroCelsius, _ := NewTemperature(0, Celsius)
	thirtyTwoFahrenheit, _ := NewTemperature(32, Fahrenheit)
	result := zeroCelsius.IsEqual(thirtyTwoFahrenheit)
	if result != true {
		t.Errorf("0 degree Celsius should be equal to 32 degree Fahrenheit")
	}
}

func Test100CelsiusEquals212Fahrenheit(t *testing.T) {
	hundredCelsius, _ := NewTemperature(100, Celsius)
	twoHundredAndTwelveFahrenheit, _ := NewTemperature(212, Fahrenheit)

	result := hundredCelsius.IsEqual(twoHundredAndTwelveFahrenheit)
	if result != true {
		t.Errorf("100 degree Celsius should be equal to 212 degree Fahrenheit")
	}
}

func Test212FahrenheitEqual100Celsius(t *testing.T) {
	twoHundredAndTwelveFahrenheit, _ := NewTemperature(212, Fahrenheit)
	hundredCelsius, _ := NewTemperature(100, Celsius)

	result := twoHundredAndTwelveFahrenheit.IsEqual(hundredCelsius)
	if result != true {
		t.Errorf("212 degree Fahrenheit should be equal to 100 degree Celsius")
	}
}

func Test100CelsiusEqual212Fahrenheit(t *testing.T) {
	hundredCelsius, _ := NewTemperature(100, Celsius)
	twoHundredAndTwelveFahrenheit, _ := NewTemperature(212, Fahrenheit)

	result := hundredCelsius.IsEqual(twoHundredAndTwelveFahrenheit)
	if result != true {
		t.Errorf("212 degree Fahrenheit should be equal to 100 degree Celsius")
	}
}

func TestNewTemperatureWithKelvin(t *testing.T) {
	_, err := NewTemperature(10, Kelvin)
	if err != nil {
		t.Errorf("Could not create temperature with 10 Kelvin")
	}
}

func Test0KelvinEquals0Kelvin(t *testing.T) {
	zeroKelvin1, _ := NewTemperature(0, Kelvin)
	zeroKelvin2, _ := NewTemperature(0, Kelvin)
	result := zeroKelvin1.IsEqual(zeroKelvin2)
	if result != true {
		t.Errorf("0 Kelvin should be equal to 0 Kelvin")
	}
}

func Test0CelsiusEquals273Kelvin(t *testing.T) {
	zeroCelsius, _ := NewTemperature(0, Celsius)
	twoSeventyThreeKelvin, _ := NewTemperature(273.15, Kelvin)
	result := zeroCelsius.IsEqual(twoSeventyThreeKelvin)
	if result != true {
		t.Errorf("0 degree Celsius should be equal to 273.15 Kelvin")
	}
}

func Test273KelvinEquals32Fahrenheit(t *testing.T) {
	twoSeventyThreeKelvin, _ := NewTemperature(273.15, Kelvin)
	thirtyTwoFahrenheit, _ := NewTemperature(32, Fahrenheit)
	result := twoSeventyThreeKelvin.IsEqual(thirtyTwoFahrenheit)

	if result != true {
		t.Errorf("273.15 Kelvin should be equal to 32 degree Fahrenheit")
	}
}

func TestCannotCreateTemperatureBelowMinus273Celsius(t *testing.T) {
	_, err := NewTemperature(-300, Celsius)
	if err == nil {
		t.Errorf("Cannot create Temperature below -273.15 Celcius")
	}

}

func TestCannotCreateTemperatureBelow0Kelvin(t *testing.T) {
	_, err := NewTemperature(-1, Kelvin)
	if err == nil {
		t.Errorf("Cannot create Temperature below 0 Kelvin")
	}

}

func TestCannotCreateTemperatureBelowMinus459Fahrenheit(t *testing.T) {
	_, err := NewTemperature(-460, Fahrenheit)
	if err == nil {
		t.Errorf("Cannot create Temperature below 0 Kelvin")
	}

}

func TestCreateEqualityCheckerInterface(t *testing.T) {
	var equitableMeasurement EqualityChecker
	temp, _ := NewTemperature(30, Celsius)
	equitableMeasurement = temp
	equitableMeasurement.IsEqual(temp)

}

func Test100MeterEquals100MeterUsingEqualityChecker(t *testing.T) {
	hundredMeter1, _ := NewDistance(100, Meter)
	var e EqualityChecker
	e = hundredMeter1
	result := e.IsEqual(hundredMeter1)
	if result != true {
		t.Errorf("100 meter should be equal to 100 meter")
	}
}

func Test1KilogramEquals1KilogramUsingEqualityChecker(t *testing.T) {
	oneKilogram, _ := NewWeight(1, Kilogram)
	var e EqualityChecker
	e = oneKilogram
	result := e.IsEqual(oneKilogram)
	if result != true {
		t.Errorf("1 Kilogram should be equal to 1 Kilogram")
	}
}
