package exercise4

import "fmt"

func Run() {
	employee1 := Employee{40, 160}
	employee2 := Employee{40, 60}

	salary, err := employee1.getSalary()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(salary)

	salary, err = employee2.getSalary()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(salary)
}
