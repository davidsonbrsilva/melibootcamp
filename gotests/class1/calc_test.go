package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSub(t *testing.T) {
	firstNumber := 3.0
	secondNumber := 1.5
	expectedResult := 1.5

	result := Sub(firstNumber, secondNumber)

	assert.Equal(t, expectedResult, result, "must be equals")
}
