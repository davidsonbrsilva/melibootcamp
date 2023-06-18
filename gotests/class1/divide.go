package main

import "errors"

var (
	ErrCannotDivideByZero = errors.New("cannot divide by zero")
)

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrCannotDivideByZero
	}

	return a / b, nil
}
