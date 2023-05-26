package main

import (
	"fmt"
	"time"
)

func main() {
	word := "batata"
	fmt.Println(getWordSize(word))
	printEachCharacterOf(word)
	grantLoan(22, true, 2, 80000)
	printDate()

	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	getBenjaminAge(employees)

	for _, employee := range getAllHaveMoreThan21Years(employees) {
		fmt.Println(employee)
	}

	employees["Batata"] = 21
	delete(employees, "Pedro")
}

func getWordSize(word string) int {
	return len(word)
}

func printEachCharacterOf(word string) {
	for _, character := range word {
		fmt.Printf("%c", character)
	}

	fmt.Println()
}

func grantLoan(age int, isEmployee bool, yearsOfExperience float64, salary float64) {
	if age > 22 && isEmployee && yearsOfExperience > 1 {
		if salary > 100000 {
			fmt.Println("Empréstimo concedido sem juros.")
			return
		}

		fmt.Println("Empréstimo concedido com juros.")
		return
	}

	fmt.Println("Empréstimo não concedido.")
}

func printDate() {
	currentDate := time.Now()
	months := []string{"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho", "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"}
	fmt.Println(currentDate.Day(), "de", months[currentDate.Month()-1])
}

func getBenjaminAge(employees map[string]int) int {
	return employees["Benjamin"]
}

func getAllHaveMoreThan21Years(employees map[string]int) []string {
	result := []string{}

	for name, age := range employees {
		if age > 21 {
			result = append(result, name)
		}
	}

	return result
}
