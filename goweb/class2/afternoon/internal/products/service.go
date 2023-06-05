package products

type Service interface {
	GetAll() ([]Product, error)
	Store(name string, price float64) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) Store(name string, price float64) (Product, error) {
	lastId, err := s.repository.LastId()
	if err != nil {
		return Product{}, err
	}

	lastId++

	product, err := s.repository.Store(lastId, name, price)

	if err != nil {
		return Product{}, err
	}

	return product, nil
}
