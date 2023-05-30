package exercise2

type Product struct {
	id     int
	price  float64
	amount int
}

type ProductCollection []Product

func (productCollection ProductCollection) getTotal() float64 {
	accumulator := 0.0

	for _, product := range productCollection {
		accumulator += product.price * float64(product.amount)
	}

	return accumulator
}
