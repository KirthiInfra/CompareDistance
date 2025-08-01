package measurement

import (
	"errors"
	
)

type Unit struct {
	name                 string
	baseConversionFactor float64
}

var (
	//distance units
	meter      = Unit{name: "meter", baseConversionFactor: 1}
	kilometer  = Unit{name: "kilometer", baseConversionFactor: 1000}
	centimeter = Unit{name: "centimeter", baseConversionFactor: 0.01}
	//weight units
	gram       = Unit{name: "gram", baseConversionFactor: 1}
	kilogram   = Unit{name: "kilogram", baseConversionFactor: 1000}
	milligram  = Unit{name: "milligram", baseConversionFactor: 0.001}
	//temperature unit
	celsius    = Unit{name: "celsius", baseConversionFactor: 1} 
	fahrenheit = Unit{name: "fahrenheit", baseConversionFactor: 1.8}
)

type measurement struct {
	value float64
	unit  Unit
}

type Distance struct {
	measurement
}

type Weight struct {
	measurement
}

type Temperature struct {
	measurement
	baseAdditionFactor float64
}

func baseAdditionFactorForTemperature(unit Unit) float64 {
    m := map[Unit]float64{
        celsius: 0,
        fahrenheit: 32,
    }
    return m[unit]
}



func NewDistanceUnit(value float64, unit Unit) (*Distance, error) { 
	if value < 0 {
		return nil, errors.New("distance cannot be negative")
	}
	if unit == meter || unit == kilometer || unit == centimeter {
		return &Distance{measurement{value: value, unit: unit}}, nil
	}
	return nil, errors.New("invalid unit")
}

func NewWeightUnit(value float64, unit Unit) (*Weight, error) { //creating new Distance struct
	if value < 0 {
		return nil, errors.New("distance cannot be negative")
	}
	if unit == gram || unit == kilogram || unit == milligram {
		return &Weight{measurement{value: value, unit: unit}}, nil
	}
	return nil, errors.New("invalid unit")
}

func NewTemperatureUnit(value float64, unit Unit) (*Temperature, error) { 
	if unit == celsius || unit == fahrenheit {
		return &Temperature{measurement: measurement{value: value, unit: unit}, 
		baseAdditionFactor: baseAdditionFactorForTemperature(unit)}, nil
	}
	return nil, errors.New("invalid unit")
}

func (d1 *Distance) IsEqual(d2 *Distance) bool { 
	return d1.measurement.IsEqual(&d2.measurement)
}

func (w1 *Weight) IsEqual(w2 *Weight) bool { 
	return w1.measurement.IsEqual(&w2.measurement)
}

func (t1 *Temperature) IsEqual(t2 *Temperature) bool { 
	return t1.measurement.IsEqual(&t2.measurement)
}

func (m1 *measurement) IsEqual(m2 *measurement) bool {
	return m1.InBase().value == m2.InBase().value
}

func (d *measurement) InBase() *measurement { 
	return &measurement{value: d.value * d.unit.baseConversionFactor, unit: d.unit}
}

func (d1 *measurement) Add(d2 *measurement) *measurement {

	baseResult := d1.InBase().value + d2.InBase().value

	resultInSelfUnit := baseResult / d1.unit.baseConversionFactor

	return &measurement{value: resultInSelfUnit, unit: d1.unit} 
}

func (d1 *Distance) Add(d2 *Distance) *Distance {
	return &Distance{*(d1.measurement.Add(&d2.measurement))}
}

func (d1 *Weight) Add(d2 *Weight) *Weight {
	return &Weight{*(d1.measurement.Add(&d2.measurement))}
}

func (d1 *Temperature) Add(d2 *Temperature) error {
	return errors.New("temperatures cannot be added")
}