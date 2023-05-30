package exercise2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	data, err := os.ReadFile("products.csv")

	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}

	products := convertToProducts(data)
	printAsTable(products)
}

func convertToProducts(text []byte) []Product {
	lines := strings.Split(string(text[:]), ";")
	products := make([]Product, 0)

	for _, line := range lines {
		if line == "\n" {
			break
		}

		fields := strings.Split(line, ",")

		id, _ := strconv.Atoi(fields[0])
		price, _ := strconv.ParseFloat(fields[1], 64)
		amount, _ := strconv.Atoi(fields[2])
		products = append(products, Product{
			id:     id,
			price:  price,
			amount: amount,
		})
	}

	return products
}

func printAsTable(products ProductCollection) {
	fmt.Printf("%-6s%10s%10s\n", "ID", "Price", "Amount")

	for _, product := range products {
		fmt.Printf("%-6d%10f%10d\n", product.id, product.price, product.amount)
	}

	fmt.Printf("%-6s%10f%10s\n", "", products.getTotal(), "")
}
