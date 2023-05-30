package exercise2

import (
	"fmt"
	"os"
)

func Run() {

	defer handlePanic()

	// Task 1
	idGenerator := Id{}
	idGenerator.nextValue = -2

	fmt.Println(idGenerator.New())

	idGenerator.nextValue = 0

	fmt.Println(idGenerator.New())

	// Task 2
	_, err := os.Open("arquivo.txt")

	if os.IsNotExist(err) {
		panic("o arquivo indicado não foi encontrado ou está danificado")
	}
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("erro:", r)
	}
}
