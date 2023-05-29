package exercise1

type User struct {
	name     string
	age      int
	email    string
	password string
}

func (user *User) setName(name string) {
	user.name = name
}

func (user *User) setAge(age int) {
	user.age = age
}

func (user *User) setEmail(email string) {
	user.email = email
}

func (user *User) setPassword(password string) {
	user.password = password
}
