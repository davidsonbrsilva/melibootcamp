package exercise2

type Store struct {
	products []IProduct
}

func (store Store) getTotal() float64 {
	accumulator := 0.0

	for _, product := range store.products {
		accumulator, _ = product.calculateCost()
	}

	return accumulator
}

func (store *Store) add(product IProduct) {
	store.products = append(store.products, product)
}
