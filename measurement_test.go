package measurement

import (
	
	"testing"
)

func TestCheckEqualDistances(t *testing.T) {
	tests := []struct {
		name     string
		distance1Val    float64
		distance1Unit   Unit
		distance2Val    float64
		distance2Unit   Unit
		expected bool
	}{
		{
			name:  "1000 meters should be should be equal to 1000 meters",
			distance1Val: 1000, distance1Unit: meter,
			distance2Val: 1000, distance2Unit: meter,
			expected: true,
		},
		{
			name:  "1 meter should not be should be equal to 2 meters",
			distance1Val: 1, distance1Unit: meter,
			distance2Val: 2, distance2Unit: meter,
			expected: false,
		},
		{
			name:  "1000 meters should be equal to 1 kilometer",
			distance1Val: 1000, distance1Unit: meter,
			distance2Val: 1, distance2Unit: kilometer,
			expected: true,
		},
		{
			name:  "1 kilometer should be equal to 1000 meters",
			distance1Val: 1, distance1Unit: kilometer,
			distance2Val: 1000, distance2Unit: meter,
			expected: true,
		},
		{
			name:  "100 centimeter should be equal to 1 meter",
			distance1Val: 1, distance1Unit: meter,
			distance2Val: 100, distance2Unit: centimeter,
			expected: true,
		},
		{
			name:  "10 meters should be equal to 1000 centimeters",
			distance1Val: 10, distance1Unit: meter,
			distance2Val: 1000, distance2Unit: centimeter,
			expected: true,
		},
		{
			name:  "5 kilometers should be equal to 500000 centimeters",
			distance1Val: 5, distance1Unit: kilometer,
			distance2Val: 500000, distance2Unit: centimeter,
			expected: true,
		},
		{
			name:  "200000 centimeters should be equal to 2 kilometers",
			distance1Val: 2, distance1Unit: kilometer,
			distance2Val: 200000, distance2Unit: centimeter,
			expected: true,
		},
		
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			distance1, err1 := NewDistanceUnit(tt.distance1Val, tt.distance1Unit)
			distance2, err2 := NewDistanceUnit(tt.distance2Val, tt.distance2Unit)

			if err1 != nil || err2 != nil {
				t.Fatalf("error creating measurements: %v, %v", err1, err2)
			}

			got := distance1.IsEqual(distance2)
			if got != tt.expected {
				t.Errorf("%v", tt.name)
			}
		})
	}
}



func TestCannotCreateDistanceWithNegativeValue(t *testing.T) {
	_, err := NewDistanceUnit(-1, meter)
	if err == nil {
		t.Errorf("Expected error for negative value, got none")
	}
}

func TestAdd5MeterAnd1Kilometer(t *testing.T) {
	d1, _ := NewDistanceUnit(5, meter)
	distance2, _ := NewDistanceUnit(1, kilometer)
	add := d1.Add(distance2)

	if add.measurement.value != 1005 {
		t.Errorf("cannot expected error for invalid unit, got none")
	}
	if add.unit != meter {
		t.Errorf("cannot expected error for invalid unit, got none")
	}

}

func TestAdd5KilometerAnd1000Meter(t *testing.T) {

	val1, _ := NewDistanceUnit(5, kilometer)

	val2, _ := NewDistanceUnit(1000, meter)
	add:= val1.Add(val2)

	if add.value != 6 {
		t.Errorf("cannot expected error for invalid unit, got none")
	}
	if add.unit != kilometer {
		t.Errorf("cannot expected error for invalid unit, got none")
	}

}

func TestAdd5CentimeterAnd1Meter(t *testing.T) {
	val1, _ := NewDistanceUnit(5, centimeter)
	

	val2, _ := NewDistanceUnit(1, meter)
	
	add := val1.Add(val2)

	if add.value != 105 {
		t.Errorf("cannot expected error for invalid unit, got none")
	}
	if add.unit != centimeter {
		t.Errorf("cannot expected error for invalid unit, got none")
	}

}

func TestCannotCreateWeightWithNegativeValue(t *testing.T) {
	_, err := NewWeightUnit(-1, kilogram)
	if err == nil {
		t.Errorf("Expected error for negative value, got none")
	}
}

func TestCheckEqualWeights(t *testing.T) {
	tests := []struct {
		name     string
		weight1Val    float64
		weight1Unit   Unit
		weight2Val    float64
		weight2Unit   Unit
		expected bool
	}{
		{
			name:  "1000 grams should be to 1 kilogram",
			weight1Val: 1000, weight1Unit: gram,
			weight2Val: 1, weight2Unit: kilogram,
			expected: true,
		},
		{
			name:  "1 kilogram should not be equal to 2 kilograms",
			weight1Val: 1, weight1Unit: kilogram,
			weight2Val: 2, weight2Unit: kilogram,
			expected: false,
		},
		{
			name:  "2000000 milligrams should be equal to 2 kilograms",
			weight1Val: 2000000, weight1Unit: milligram,
			weight2Val: 2, weight2Unit: kilogram,
			expected: true,
		},
		{
			name:  "2 kilograms should be equal to 2000000 milligrams",
			weight1Val: 2, weight1Unit: kilogram,
			weight2Val: 2000000, weight2Unit: milligram,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			weight1, err1 := NewWeightUnit(tt.weight1Val, tt.weight1Unit)
			weight2, err2 := NewWeightUnit(tt.weight2Val, tt.weight2Unit)

			if err1 != nil || err2 != nil {
				t.Fatalf("error creating measurements: %v, %v", err1, err2)
			}

			got := weight1.IsEqual(weight2)
			if got != tt.expected {
				t.Errorf("%v", tt.name)
			}
		})
	}
}

func TestAdd1000GramsAnd1000Grams(t *testing.T) {
	weight1, _ := NewWeightUnit(1000, gram)
	

	weight2, _ := NewWeightUnit(1000, gram)
	
	add := weight1.Add(weight2)

	if add.value != 2000 {
		t.Errorf("expected value: 2000 but the actual value: %v", add.value)
	}
	if add.unit != gram {
		t.Errorf("expected unit: centimeter but the actual unit: %v", add.unit)
	}
}

func TestAdd1KilogramAnd300Grams(t *testing.T) {
	weight1, _ := NewWeightUnit(1, kilogram)
	weight2, _ := NewWeightUnit(300, gram)
	
	add := weight1.Add(weight2)

	if add.value != 1.3 {
		t.Errorf("expected value: 2000 but the actual value: %v", add.value)
	}
	if add.unit != kilogram {
		t.Errorf("expected unit: centimeter but the actual unit: %v", add.unit)
	}
}

func TestAdd500GramsAnd1Kilogram(t *testing.T) {
	weight1, _ := NewWeightUnit(500, gram)
	weight2, _ := NewWeightUnit(1, kilogram)
	
	add := weight1.Add(weight2)

	if add.value != 1500 {
		t.Errorf("expected value: 2000 but the actual value: %v", add.value)
	}
	if add.unit != gram {
		t.Errorf("expected unit: centimeter but the actual unit: %v", add.unit)
	}
}

func TestCreateTemperatureWithUnitCelsius(t *testing.T) {
	_, err := NewTemperatureUnit(100, celsius)
	if err != nil {
		t.Errorf("cannot create the temperature with celsius unit")
	}
}

func Test100CelsiusEqualsTo100Celsius(t *testing.T) {
	hundredcelsius1, _ := NewTemperatureUnit(100, celsius)
	hundredcelsius2, _ := NewTemperatureUnit(100, celsius)

	got := hundredcelsius1.IsEqual(hundredcelsius2)
	if got != true {
		t.Errorf("100 celsius should be equal to 100 celsius")
	}
}

func TestAdd100CelsiusWith100Celsius(t *testing.T) {
	hundredcelsius1, _ := NewTemperatureUnit(100, celsius)
	hundredcelsius2, _ := NewTemperatureUnit(100, celsius)

	err := hundredcelsius1.Add(hundredcelsius2)
	if err == nil {
		t.Errorf("Temperatures cannot be added")
	}
}

func TestCreateTemperatureWithUnitFahrenheit(t *testing.T) {
	_, err := NewTemperatureUnit(100, fahrenheit)
	if err != nil {
		t.Errorf("cannot create the temperature with fahrenheit unit")
	}
}