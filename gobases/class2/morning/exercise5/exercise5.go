package exercise5

import (
	"fmt"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func Run() {
	getDogFoodAmount := animal("dog")
	getCatFoodAmount := animal("cat")
	getHamsterFoodAmount := animal("hamster")
	getTarantulaFoodAmount := animal("tarantula")

	fmt.Println(getDogFoodAmount(5))
	fmt.Println(getCatFoodAmount(5))
	fmt.Println(getHamsterFoodAmount(5))
	fmt.Println(getTarantulaFoodAmount(5))
}

func animal(animalType string) func(animalAmount float64) string {
	if animalType == dog {
		return getDogFoodAmount
	}

	if animalType == cat {
		return getCatFoodAmount
	}

	if animalType == hamster {
		return getHamsterFoodAmount
	}

	if animalType == tarantula {
		return getTarantulaFoodAmount
	}

	return func(animalAmount float64) string {
		return "Unexistent animal"
	}
}

func getDogFoodAmount(animalAmount float64) string {
	return fmt.Sprint(animalAmount*10.0, "kg")
}

func getCatFoodAmount(animalAmount float64) string {
	return fmt.Sprint(animalAmount*5.0, "kg")
}

func getHamsterFoodAmount(animalAmount float64) string {
	return fmt.Sprint(animalAmount*0.25, "kg")
}

func getTarantulaFoodAmount(animalAmount float64) string {
	return fmt.Sprint(animalAmount*0.15, "kg")
}
