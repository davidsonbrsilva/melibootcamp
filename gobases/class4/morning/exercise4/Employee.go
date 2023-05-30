package exercise4

import "errors"

type Employee struct {
	workHourValue float64
	workedHours   int
}

func (employee Employee) getSalary() (float64, error) {
	salary := employee.workHourValue * float64(employee.workedHours)

	if salary >= 20000 {
		salary *= 0.9
	}

	if employee.workedHours < 80 {
		return -1, errors.New("erro: o trabalhador não pode ter trabalhado menos de 80 horas por mês")
	}

	return salary, nil
}
