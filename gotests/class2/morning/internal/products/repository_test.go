package products

import (
	"testing"

	"class2/morning/internal/products/mock"

	"github.com/stretchr/testify/assert"
)

type StoreMock struct{}

func TestGetAll(t *testing.T) {
	storeStub := mock.StoreStub{}
	repository := NewRepository(&storeStub)

	expected := []Product{
		{Id: 1, Name: "Xbox Series X", Price: 100},
		{Id: 2, Name: "Playstation 5", Price: 150},
	}

	products, err := repository.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, products, expected)
}

func TestUpdate(t *testing.T) {
	storeMock := mock.StoreMock{}
	repository := NewRepository(&storeMock)

	expected := Product{
		Id:    1,
		Name:  "After Update",
		Price: 150,
	}

	product, err := repository.Update(1, "After Update", 150)

	assert.Equal(t, expected, product)
	assert.True(t, storeMock.ReadWasCalled)
	assert.Nil(t, err)
}
