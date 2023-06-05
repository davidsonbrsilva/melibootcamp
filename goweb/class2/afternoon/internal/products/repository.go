package products

type Product struct {
	Id    int
	Name  string
	Price float64
}

var ps []Product
var lastId int

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, name string, price float64) (Product, error)
	LastId() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Product, error) {
	return ps, nil
}

func (r *repository) LastId() (int, error) {
	return lastId, nil
}

func (r *repository) Store(id int, name string, price float64) (Product, error) {
	p := Product{id, name, price}
	ps = append(ps, p)
	lastId = p.Id
	return p, nil
}
