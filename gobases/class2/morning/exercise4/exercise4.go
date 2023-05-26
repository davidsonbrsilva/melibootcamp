package exercise4

import (
	"errors"
	"fmt"
	"math"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func min(values ...float64) float64 {
	result := math.Inf(1)

	for _, value := range values {
		if value < result {
			result = value
		}
	}

	return result
}

func avg(values ...float64) float64 {
	accumulator := 0.0

	for _, value := range values {
		accumulator += value
	}

	return accumulator / float64(len(values))
}

func max(values ...float64) float64 {
	result := math.Inf(-1)

	for _, value := range values {
		if value > result {
			result = value
		}
	}

	return result
}

func operation(operation string) (func(values ...float64) float64, error) {
	if operation == minimum {
		return min, nil
	}

	if operation == average {
		return avg, nil
	}

	if operation == maximum {
		return max, nil
	}

	return nil, errors.New("Operation not defined.")
}

func Run() {
	min, _ := operation(minimum)
	avg, _ := operation(average)
	max, _ := operation(maximum)

	fmt.Println(min(2, 3, 4))
	fmt.Println(avg(2, 3, 4))
	fmt.Println(max(2, 3, 4))

	undefined, err := operation("potato")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		undefined(2, 3, 4)
	}
}
