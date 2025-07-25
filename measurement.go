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

	celcius = Unit{ name:"celcius",baseConversionFactor: 1}
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

type temperature struct{
	measurement
}

func (d *measurement) IsEqual(d1 *measurement) bool {
	return d.conversed == d1.conversed
}

func (d *distance) IsEqual(d1 *distance) bool {
	return d.measurement.IsEqual(&d1.measurement)
}

func (w1 *weight) IsEqual(w2 *weight) bool {
	return w1.measurement.IsEqual(&w2.measurement)
}

func (t1 *temperature)IsEqual(t2 *temperature) bool{
	return t1.measurement.IsEqual(&t2.measurement)
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

func NewWeight(i float64, unit Unit) (*weight, error) {
	return &weight{measurement{value: i, unit: unit, conversed: i * unit.baseConversionFactor}}, nil
}

func NewTemperature(i float64, unit Unit) (*temperature, error) {
	return &temperature{measurement{value: i, unit: unit, conversed: i * unit.baseConversionFactor}}, nil
}

func (m *measurement) Add(m1 *measurement) (*measurement) {
	
	result := m.conversed + m1.conversed
	baseFactor := m.unit.baseConversionFactor
	return &measurement{
		value:     result / baseFactor,
		unit:      m.unit,
		conversed: result,
	}
}

func (d1 *distance) Add(d2 *distance) (*distance){
	return &distance{*d1.measurement.Add(&d2.measurement)}
}

func (w1 *weight) Add(w2 *weight) (*weight){
	return &weight{*w1.measurement.Add(&w2.measurement)}
}

func (t1 *temperature) Add(t2 *temperature) (*temperature){
	return nil
}