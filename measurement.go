package Measurement

import (
	"errors"
)

type Unit struct {
	name                 string
	baseConversionFactor float64
}

var (
	meter      = Unit{name: "m", baseConversionFactor: 1}
	kilometer  = Unit{name: "km", baseConversionFactor: 1000}
	centimeter = Unit{name: "cm", baseConversionFactor: 0.01}

	gram      = Unit{name: "g", baseConversionFactor: 1}
	kilogram  = Unit{name: "kg", baseConversionFactor: 1000}
	milligram = Unit{name: "mg", baseConversionFactor: 0.001}
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

func (d *distance) IsEqual(d1 *distance) bool {
	return d.measurement.conversed == d1.measurement.conversed
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
	return &distance{measurement{value: i, unit: unit, conversed: i * unit.baseConversionFactor}}, nil
}

func (m *measurement) Add(m1 *measurement) (*measurement, error) {
	if m.unit != m1.unit {
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
