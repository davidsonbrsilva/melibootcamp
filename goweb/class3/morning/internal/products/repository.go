package products

import "fmt"

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
	Update(id int, name string, price float64) (Product, error)
	PartialUpdate(id int, name string, price float64) (Product, error)
	Delete(id int) error
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

func (r *repository) Update(id int, name string, price float64) (Product, error) {
	p := Product{Name: name, Price: price}
	updated := false
	for i := range ps {
		if ps[i].Id == id {
			p.Id = id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("Produto %d não encontrado", id)
	}
	return p, nil
}

func (r *repository) PartialUpdate(id int, name string, price float64) (Product, error) {
	var p Product
	updated := false
	for i := range ps {
		if ps[i].Id == id {
			ps[i].Name = name
			ps[i].Price = price
			updated = true
			p = ps[i]
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("Produto %d não encontrado", id)
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range ps {
		if ps[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("Produto %d não encontrado", id)
	}
	ps = append(ps[:index], ps[index+1:]...)
	return nil
}
