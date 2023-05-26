package exercise2

import (
	"errors"
	"fmt"
)

func Run() {
	fmt.Println(calculateAverage())
	fmt.Println(calculateAverage(2))
	fmt.Println(calculateAverage(2, 3))
	fmt.Println(calculateAverage(2, 3, 5))
	fmt.Println(calculateAverage(2, 3, 5, 6, 7, 8, 9, 54))
}

func calculateAverage(numbers ...float64) (float64, error) {
	var err error

	if len(numbers) <= 0 {
		err = errors.New("É necessário pelo menos um número para calcular a média.")
		return -1, err
	}

	accumulator := numbers[0]

	for _, number := range numbers {
		accumulator += number
	}

	return accumulator / float64(len(numbers)), err
}
