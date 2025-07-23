package measurement

import (
	"errors"
	
)

type Unit struct {
	name                 string
	baseConversionFactor float64
}

var (
	meter      = Unit{name: "meter", baseConversionFactor: 1}
	kilometer  = Unit{name: "kilometer", baseConversionFactor: 1000}
	centimeter = Unit{name: "centimeter", baseConversionFactor: 0.01}
)

type measurement struct {
	value float64
	unit  Unit
}

type Distance struct {
	measurement
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

func (d1 *Distance) IsEqual(d2 *Distance) bool { 
	return d1.measurement.IsEqual(&d2.measurement)
}

func (d1 *measurement) IsEqual(d2 *measurement) bool {
	return d1.InBase().value == d2.InBase().value
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