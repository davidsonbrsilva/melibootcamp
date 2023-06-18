package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivideWhenDenominatorIsNotZero(t *testing.T) {
	firstNumber := 2.0
	secondNumber := 1.0
	expectedResult := 2.0

	result, err := Divide(firstNumber, secondNumber)

	assert.Equal(t, expectedResult, result)
	assert.Nil(t, err)
}

func TestDivideWhenDenominatorIsZero(t *testing.T) {
	firstNumber := 2.0
	secondNumber := 0.0

	_, err := Divide(firstNumber, secondNumber)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrCannotDivideByZero)
}
