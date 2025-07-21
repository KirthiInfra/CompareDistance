package compareDistance

import (
	"errors"
)

type unit string

const (
	m  unit = "m"
	km unit = "km"
	cm unit = "cm"
)

type distance struct {
	value int
	unit  unit
}

func (d *distance) InMeter() *distance {
	if d.unit == km {
		return &distance{value: d.value * 1000, unit: m}
	} else if d.unit == cm {
		return &distance{value: d.value / 100, unit: m}
	}
	return d
}

func (d *distance) IsDistanceEqual(d1 *distance) bool {
	return d.InMeter().value == d1.InMeter().value
}

func CreateDistancesStruct(value int, unit unit) (*distance, error) {
	if value <= 0 {
		return nil, errors.New("cannot create struct with zero or negative value")
	}
	if unit != m && unit != km && unit != cm {
		return nil, errors.New("invalid unit, supported units are 'm' or 'km' or 'cm'")
	}
	return &distance{value: value, unit: unit}, nil
}
