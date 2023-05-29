package exercise2

import "errors"

type Product struct {
	kind  string
	name  string
	price float64
}

func (product Product) calculateCost() (float64, error) {
	if product.kind == SMALL {
		return product.price, nil
	}

	if product.kind == MEDIUM {
		return product.price * 1.03, nil
	}

	if product.kind == LARGE {
		return product.price*1.06 + 2500, nil
	}

	return -1, errors.New("Product type not recognized")
}
