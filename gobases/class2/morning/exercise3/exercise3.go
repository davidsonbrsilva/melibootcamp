package exercise3

import "fmt"

func Run() {
	fmt.Println(calculateSalary(50, ""))
	fmt.Println(calculateSalary(200, ""))
	fmt.Println(calculateSalary(200, "A"))
	fmt.Println(calculateSalary(200, "B"))
	fmt.Println(calculateSalary(200, "C"))
	fmt.Println(calculateSalary(50, "A"))
	fmt.Println(calculateSalary(50, "B"))
	fmt.Println(calculateSalary(50, "C"))

	fmt.Println(calculateSalary(172, "A"))
	fmt.Println(calculateSalary(176, "B"))
	fmt.Println(calculateSalary(162, "C"))
}

func calculateSalary(workedHours float64, category string) float64 {
	hasWorkedMoreThan160Hours := workedHours > 160
	multiplier := 0.0
	additionalRate := 1.0

	if category == "A" {
		multiplier = 3000

		if hasWorkedMoreThan160Hours {
			additionalRate = 1.5
		}
	}

	if category == "B" {
		multiplier = 1500

		if hasWorkedMoreThan160Hours {
			additionalRate = 1.2
		}
	}

	if category == "C" {
		multiplier = 1000
	}

	return workedHours * multiplier * additionalRate
}
