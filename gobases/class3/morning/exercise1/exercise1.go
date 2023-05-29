package exercise1

import (
	"fmt"
	"log"
	"os"
)

type Product struct {
	id     int
	price  float64
	amount int
}

func (product Product) toString() string {
	return fmt.Sprint(product.id, ",", product.price, ",", product.amount)
}

func Run() {
	products := []Product{
		{1, 4.05, 5},
		{2, 5.05, 10},
		{3, 6.05, 15},
	}

	file, err := os.Create("products.csv")

	if err != nil {
		log.Fatal(err)
	}

	for _, product := range products {
		file.WriteString(fmt.Sprint(product.toString(), ";\n"))
	}

	file.Close()
}
