package exercise2

type IProduct interface {
	calculateCost() (float64, error)
}
