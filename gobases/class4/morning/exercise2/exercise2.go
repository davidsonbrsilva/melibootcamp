package exercise2

import (
	"errors"
	"fmt"
)

func checkSalary(salary int) error {
	if salary < 1500 {
		return errors.New("o salário digitado não está dentro do valor mínimo")
	}

	fmt.Println("necessário pagamento de imposto")
	return nil
}

func Run() {
	salary := 1000
	err := checkSalary(salary)

	if err != nil {
		fmt.Println(err)
	}
}
