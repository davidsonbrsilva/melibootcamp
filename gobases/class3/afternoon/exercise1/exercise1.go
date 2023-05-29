package exercise1

import "fmt"

func Run() {
	user1 := User{
		name:     "Example",
		age:      20,
		email:    "example@example.com",
		password: "my-super-secret-password",
	}

	user1.setName("Jonh")
	user1.setAge(30)

	fmt.Println(user1)
}
