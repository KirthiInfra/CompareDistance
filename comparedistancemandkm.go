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
	var input distance

	fmt.Println("Enter the input distance: ")
	fmt.Scan(&input.value, &input.unit)
	err := UnitSupportedOrNot(input)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func UnitSupportedOrNot(d distance) error {
	if d.unit != "m" && d.unit != "km" {
		return errors.New("invalid unit, supported units (m or km)")
	}
	return nil
}
