package exercise1

import (
	"fmt"
	"os"
)

func Run() {
	_, err := os.Open("arquivo.txt")

	defer handlePanic()

	if os.IsNotExist(err) {
		panic("o arquivo indicado não foi encontrado ou está danificado")
	}
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("pânico recuperado:", r)
	}
}
