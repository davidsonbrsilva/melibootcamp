package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortNumbers(t *testing.T) {
	numbers := []int{5, 2, 3, 6, 1, 9, 13}

	expected := []int{1, 2, 3, 5, 6, 9, 13}
	result := SortNumbers(numbers)

	assert.Equal(t, expected, result)
}
