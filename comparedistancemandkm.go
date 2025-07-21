package compareDistance

import (
	"errors"
)

type unit string

const (
	m  unit = "m"
	km unit = "km"
)

type distance struct {
	value int
	unit  unit
}

func (d1 *distance) IsDistanceEqual(d2 distance) bool {
	if d1.unit == d2.unit {
		return d1.value == d2.value
	}
	if (d1.unit == km && d1.value*1000 == d2.value) || (d2.unit == km && d2.value*1000 == d1.value) {
		return true
	}
	return false
}

func CreateDistancesStruct(value int, unit unit) (*distance, error) {
	if value <= 0 {
		return nil, errors.New("cannot create struct with zero or negative value")
	}
	if unit != m && unit != km {
		return nil, errors.New("invalid unit, supported units are 'm' or 'km'")
	}
	return &distance{value: value, unit: unit}, nil
}
