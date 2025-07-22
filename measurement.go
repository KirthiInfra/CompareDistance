package Measurement

import (
	"errors"
)

type unit string

const (
	m  unit = "m"
	km unit = "km"
	cm unit = "cm"
	g  unit = "g"
	kg unit = "kg"
	mg unit = "mg"
)

type measurement struct {
	value float64
	unit  unit
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


func weightMap(Unit unit) float64 {
	m := map[unit]float64{
		kg: 1000,
		g:  1,
		mg: 0.001,
		
	}
	value,okay := m[Unit];if okay{
		return value
	}
	return 0
}

func distanceMap(Unit unit) float64 {
	m := map[unit]float64{
		km: 1000,
		m:  1,
		cm: 0.01,
	}

	value,okay := m[Unit];if okay{
		return value
	}
	return 0
}

func NewMeasurement(value float64, unit unit) (*measurement, error) {
	if value <= 0 {
		return nil, errors.New("cannot create struct with zero or negative value")
	}
	switch unit {
	case m, km, cm, g, kg, mg:
		return &measurement{value: value,
			unit:      unit,
			conversed: value * (weightMap(unit)+distanceMap(unit))}, nil
	default:
		return nil, errors.New("invalid unit, supported units are 'm' or 'km' or 'cm' or 'g' or 'kg' or 'mg'")
	}
}

func (m *measurement) Add(m1 *measurement)(*measurement) {
	result := m.conversed + m1.conversed
	div := distanceMap(m.unit)+ weightMap(m.unit)
	return &measurement{value: result/div, unit: m.unit, conversed: result}
}
