package exercise3

import (
	"fmt"
)

func checkSalary(salary int) error {
	if salary < 1500 {
		return fmt.Errorf("erro: o mínimo tributável é 1500 e o salário informado é %d", salary)
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
