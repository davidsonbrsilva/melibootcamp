package exercise2

type Ecommerce struct{}

func (ecommerce Ecommerce) newProduct(name string, price float64) Product {
	return Product{
		name:  name,
		price: price,
	}
}

func (ecommerce Ecommerce) addProduct(user *User, product *Product, amount int) {
	product.amount = amount
	user.products = append(user.products, *product)
}

func (ecommerce Ecommerce) deleteAllProducts(user *User) {
	user.products = make([]Product, 0)
}
