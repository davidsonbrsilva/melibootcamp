package exercise1

import "fmt"

type Student struct {
	firstName     string
	lastName      string
	rg            string
	admissionDate string
}

func (student Student) toDetails() string {
	return fmt.Sprintf(`
		Nome: %s
		Sobrenome: %s
		RG: %s
		Data de admissão: %s`,
		student.firstName,
		student.lastName,
		student.rg,
		student.admissionDate,
	)
}

func Run() {
	student1 := Student{
		firstName:     "Davidson",
		lastName:      "Bruno",
		rg:            "12.345.678",
		admissionDate: "2023-05-22",
	}

	student2 := Student{
		firstName:     "Foo",
		lastName:      "Bar",
		rg:            "98.765.432",
		admissionDate: "2023-05-22",
	}

	fmt.Println(student1.toDetails())
	fmt.Println(student2.toDetails())
}
