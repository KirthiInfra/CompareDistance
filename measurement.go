package Measurement

import (
	"errors"
)

type Unit struct {
	name                 string
	baseConversionFactor float64
	unitType             string
}

var (
	meter      = Unit{name: "m", baseConversionFactor: 1, unitType: "distance"}
	kilometer  = Unit{name: "km", baseConversionFactor: 1000, unitType: "distance"}
	centimeter = Unit{name: "cm", baseConversionFactor: 0.01, unitType: "distance"}

	gram      = Unit{name: "g", baseConversionFactor: 1, unitType: "weight"}
	kilogram  = Unit{name: "kg", baseConversionFactor: 1000, unitType: "weight"}
	milligram = Unit{name: "mg", baseConversionFactor: 0.001, unitType: "weight"}
)

type measurement struct {
	value     float64
	unit      Unit
	conversed float64
}

type distance struct {
	measurement
}

type weight struct {
	measurement
}

func (d *measurement) IsEqual(d1 *measurement) bool {
	return d.conversed == d1.conversed
}

func newMeasurement(value float64, unit Unit) (*measurement, error) {
	if value <= 0 {
		return nil, errors.New("cannot create struct with zero or negative value")
	}

	return &measurement{
		value:     value,
		unit:      unit,
		conversed: value * unit.baseConversionFactor,
	}, nil
}

func NewDistance(i float64, unit Unit) (*distance, error) {
	return &distance{measurement{value:i, unit:unit, conversed:i*unit.baseConversionFactor}}, nil
}

func (m *measurement) Add(m1 *measurement) (*measurement, error) {
	if m.unit.unitType != m1.unit.unitType {
		return nil, errors.New("cannot add different unit types (e.g. weight and distance)")
	}

	result := m.conversed + m1.conversed
	baseFactor := m.unit.baseConversionFactor
	return &measurement{
		value:     result / baseFactor,
		unit:      m.unit,
		conversed: result,
	}, nil
}
