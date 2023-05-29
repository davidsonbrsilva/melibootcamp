package exercise2

import "fmt"

func Run() {
	user1 := User{
		firstName: "Jonh",
		lastName:  "Doe",
		email:     "jonhdoe@example.com",
		products:  make([]Product, 0),
	}

	fmt.Println(user1)

	product1 := Product{
		name:  "banana",
		price: 4.5,
	}

	fmt.Println(product1)

	ecommerce := Ecommerce{}
	ecommerce.addProduct(&user1, &product1, 5)

	fmt.Println(user1.products)

	product2 := ecommerce.newProduct("potato", 3.0)
	ecommerce.addProduct(&user1, &product2, 10)

	fmt.Println(user1.products)

	ecommerce.deleteAllProducts(&user1)

	fmt.Println(user1.products)
}
