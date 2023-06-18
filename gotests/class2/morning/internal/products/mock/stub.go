package mock

import "encoding/json"

type Product struct {
	Id    int
	Name  string
	Price float64
}

type StoreStub struct{}

func (s *StoreStub) Read(data interface{}) error {
	products := []Product{
		{Id: 1, Name: "Xbox Series X", Price: 100},
		{Id: 2, Name: "Playstation 5", Price: 150},
	}

	encoded, _ := json.Marshal(products)
	return json.Unmarshal(encoded, &data)
}

func (s *StoreStub) Write(data interface{}) error {
	return nil
}
