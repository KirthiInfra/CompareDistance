package main

import (
	"errors"
	"fmt"
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

func main() {
	var input1, input2 distance

	fmt.Println("Enter the input distance: ")
	fmt.Scan(&input1.value, &input1.unit)
	fmt.Println("Enter the input distance: ")
	fmt.Scan(&input2.value, &input2.unit)
	err := UnitSupportedOrNot(&input1, input2)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func UnitSupportedOrNot(d1 *distance, d2 distance) error {
	if (d1.unit != m && d1.unit != km) || (d2.unit != m && d2.unit != km) {
		return errors.New("invalid unit, supported units (m or km)")
	}
	return nil
}

func (d1 *distance) CompareDistances(d2 distance) bool {
	if d1.unit == d2.unit {
		return d1.value == d2.value
	} else {
		if (d1.unit == km && d1.value*1000 == d2.value) || (d2.unit == km && d2.value*1000 == d1.value) {
			return true
		}
		return false
	}
}
