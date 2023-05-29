package exercise2

type IEcommerce interface {
	getTotal() float64
	add(product IProduct)
}
