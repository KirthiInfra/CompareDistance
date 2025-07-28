package Measurement

import (
	"errors"
	"math"
)

type EqualityChecker interface {
	IsEqual() bool
}

type Adder interface {
	Add() interface{}
}

type Unit struct {
	name                 string
	baseConversionFactor float64
}

type TemperatureUnit struct {
	unit               Unit
	baseAdditionFactor float64
}

var (
	meter      = Unit{name: "m", baseConversionFactor: 1}
	kilometer  = Unit{name: "km", baseConversionFactor: 1000}
	centimeter = Unit{name: "cm", baseConversionFactor: 0.01}

	gram      = Unit{name: "g", baseConversionFactor: 1}
	kilogram  = Unit{name: "kg", baseConversionFactor: 1000}
	milligram = Unit{name: "mg", baseConversionFactor: 0.001}

	celsius    = TemperatureUnit{unit: Unit{name: "celsius", baseConversionFactor: 1}, baseAdditionFactor: 0}
	fahrenheit = TemperatureUnit{unit: Unit{name: "fahrenheit", baseConversionFactor: math.Ceil((5.0/9.0)*100) / 100}, baseAdditionFactor: -32}
	kelvin     = TemperatureUnit{unit: Unit{name: "kelvin", baseConversionFactor: 1}, baseAdditionFactor: -273.15}
)

type measurement struct {
	value float64
	unit  Unit
}

type distance struct {
	measurement
}

type weight struct {
	measurement
}

type temperature struct {
	value float64
	unit  TemperatureUnit
}

func (d *distance) IsEqual(d1 *distance) bool {
	return d.inBase().value == d1.inBase().value
}

func (w1 *weight) IsEqual(w2 *weight) bool {
	return w1.inBase().value == w2.inBase().value
}

func (t1 *temperature) IsEqual(t2 *temperature) bool {
	return math.Abs(t1.inBase().value-t2.inBase().value) < 1
}

func NewDistance(i float64, unit Unit) (*distance, error) {
	if i <= 0 {
		return nil, errors.New("Cannot create distance with negative value")
	}
	return &distance{measurement{value: i, unit: unit}}, nil
}

func NewWeight(i float64, unit Unit) (*weight, error) {
	if i <= 0 {
		return nil, errors.New("Cannot create weight with negative value")
	}
	return &weight{measurement{value: i, unit: unit}}, nil
}

func NewTemperature(i float64, unit TemperatureUnit) (*temperature, error) {
	if i < (-273.15) {
		return nil, errors.New("Cannot create Temperature below -273.15 Celsius")
	}
	return &temperature{value: i, unit: unit}, nil
}

func (m *measurement) inBase() *measurement {
	convertedValue := m.value * m.unit.baseConversionFactor
	return &measurement{value: convertedValue, unit: m.unit}
}

func (m *temperature) inBase() *temperature {
	convertedValue := math.Floor(((m.value + m.unit.baseAdditionFactor) * m.unit.unit.baseConversionFactor))
	return &temperature{value: convertedValue, unit: m.unit}
}

func (m *measurement) Add(m1 *measurement) *measurement {

	result := m.inBase().value + m1.inBase().value
	baseFactor := m.unit.baseConversionFactor
	return &measurement{
		value: result / baseFactor,
		unit:  m.unit,
	}
}

func (d1 *distance) Add(d2 *distance) *distance {
	result := d1.inBase().value + d2.inBase().value
	baseFactor := d1.unit.baseConversionFactor
	return &distance{measurement{
		value: result / baseFactor,
		unit:  d1.unit,
	}}
}

func (w1 *weight) Add(w2 *weight) *weight {
	return &weight{*w1.measurement.Add(&w2.measurement)}
}
