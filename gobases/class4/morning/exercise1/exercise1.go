package exercise1

import "fmt"

type MyCustomError struct{}

func (err *MyCustomError) Error() string {
	return "o salário digitado não está dentro do valor mínimo"
}

func checkSalary(salary int) error {
	if salary < 1500 {
		return &MyCustomError{}
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
