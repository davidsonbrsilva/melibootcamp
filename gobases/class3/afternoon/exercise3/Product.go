package exercise3

type Product struct {
	name   string
	price  float64
	amount int
}

type ProductCollection []Product

func (products ProductCollection) sum() float64 {
	accumulator := 0.0

	for _, product := range products {
		accumulator += product.price * float64(product.amount)
	}

	return accumulator
}
