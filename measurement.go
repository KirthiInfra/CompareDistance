package Measurement

import (
	"errors"
	"math"
)

type Unit struct {
	name                 string
	baseConversionFactor float64
}

type TemperatureUnit struct {
	unit               Unit
	baseAdditionFactor float64
}

type DistanceUnit struct {
	name                 string
	baseConversionFactor float64
}

type WeightUnit struct {
	name                 string
	baseConversionFactor float64
}

var (
	Meter      = DistanceUnit{name: "m", baseConversionFactor: 1}
	Kilometer  = DistanceUnit{name: "km", baseConversionFactor: 1000}
	Centimeter = DistanceUnit{name: "cm", baseConversionFactor: 0.01}

	Gram      = WeightUnit{name: "g", baseConversionFactor: 1}
	Kilogram  = WeightUnit{name: "kg", baseConversionFactor: 1000}
	Milligram = WeightUnit{name: "mg", baseConversionFactor: 0.001}

	Celsius    = TemperatureUnit{unit: Unit{name: "celsius", baseConversionFactor: 1}, baseAdditionFactor: 0}
	Fahrenheit = TemperatureUnit{unit: Unit{name: "fahrenheit", baseConversionFactor: math.Ceil((5.0/9.0)*100) / 100}, baseAdditionFactor: -32}
	Kelvin     = TemperatureUnit{unit: Unit{name: "kelvin", baseConversionFactor: 1}, baseAdditionFactor: -273.15}
)

type measurement struct {
	value float64
	unit  Unit
}

type distance struct {
	value float64
	unit  DistanceUnit
}

type weight struct {
	value float64
	unit  WeightUnit
}

type temperature struct {
	value float64
	unit  TemperatureUnit
}

type EqualityChecker interface {
	IsEqual(equitableMeasurement EqualityChecker) bool
}

type Adder interface {
	Add(addMeasurement Adder) (Adder, error)
}

func NewDistance(i float64, unit DistanceUnit) (*distance, error) {
	if i <= 0 {
		return nil, errors.New("Cannot create distance with negative value")
	}
	return &distance{value: i, unit: unit}, nil
}

func NewWeight(i float64, unit WeightUnit) (*weight, error) {
	if i <= 0 {
		return nil, errors.New("Cannot create weight with negative value")
	}
	return &weight{value: i, unit: unit}, nil
}

func NewTemperature(i float64, unit TemperatureUnit) (*temperature, error) {
	if math.Floor(((i+unit.baseAdditionFactor)*unit.unit.baseConversionFactor)*100)/100 < (-273.15) {
		return nil, errors.New("Cannot create Temperature below range")
	}
	return &temperature{value: i, unit: unit}, nil
}

func (m *measurement) inBase() *measurement {
	convertedValue := m.value * m.unit.baseConversionFactor
	return &measurement{value: convertedValue, unit: m.unit}
}

func (m *distance) inBase() *distance {
	convertedValue := m.value * m.unit.baseConversionFactor
	return &distance{value: convertedValue, unit: m.unit}
}

func (w *weight) inBase() *weight {
	convertedValue := w.value * w.unit.baseConversionFactor
	return &weight{value: convertedValue, unit: w.unit}
}

func (m *temperature) inBase() *temperature {
	convertedValue := math.Floor(((m.value + m.unit.baseAdditionFactor) * m.unit.unit.baseConversionFactor))
	return &temperature{value: convertedValue, unit: m.unit}
}

func (d1 *distance) IsEqual(e EqualityChecker) bool {
	d2, ok := e.(*distance)
	if !ok {
		return false
	}
	return d1.inBase().value == d2.inBase().value
}

func (t1 *temperature) IsEqual(e EqualityChecker) bool {
	t2, ok := e.(*temperature)
	if !ok {
		return false
	}
	return math.Abs(t1.inBase().value-t2.inBase().value) < 1
}

func (w1 *weight) IsEqual(e EqualityChecker) bool {
	w2, ok := e.(*weight)
	if !ok {
		return false
	}
	return w1.inBase().value == w2.inBase().value
}

func (d1 *distance) Add(a Adder) (Adder, error) {
	d2, ok := a.(*distance)
	if !ok {
		return nil, errors.New("Operand types do not match")
	}
	result := d1.inBase().value + d2.inBase().value
	baseFactor := d1.unit.baseConversionFactor
	return &distance{
		value: result / baseFactor,
		unit:  d1.unit,
	}, nil
}

func (w1 *weight) Add(a Adder) (Adder, error) {
	w2, ok := a.(*weight)
	if !ok {
		return nil, errors.New("Operand types do not match")
	}
	result := w1.inBase().value + w2.inBase().value
	baseFactor := w1.unit.baseConversionFactor
	return &weight{
		value: result / baseFactor,
		unit:  w1.unit,
	}, nil
}
