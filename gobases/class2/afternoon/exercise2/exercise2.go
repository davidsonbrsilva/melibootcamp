package exercise2

import "fmt"

const (
	SMALL  = "small"
	MEDIUM = "medium"
	LARGE  = "large"
)

func Run() {
	product1 := buildProduct(SMALL, "banana", 2.5)
	product2 := buildProduct(MEDIUM, "computer", 2500)
	product3 := buildProduct(LARGE, "table", 3000)

	store := buildStore()
	store.add(product1)
	store.add(product2)
	store.add(product3)

	fmt.Println(store.getTotal())
}
