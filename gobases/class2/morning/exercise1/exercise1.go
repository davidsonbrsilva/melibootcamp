package exercise1

import "fmt"

func Run() {
	fmt.Println(getTax(50000))
	fmt.Println(getTax(150000))
	fmt.Println(getTax(25000))
}

func getTax(salary float64) float64 {
	if salary == 50000 {
		return 0.17
	}

	if salary == 150000 {
		return 0.27
	}

	return 0
}
