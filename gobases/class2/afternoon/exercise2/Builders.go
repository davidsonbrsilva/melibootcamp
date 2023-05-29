package exercise2

func buildProduct(kind string, name string, price float64) IProduct {
	return Product{kind, name, price}
}

func buildStore() IEcommerce {
	return &Store{[]IProduct{}}
}
