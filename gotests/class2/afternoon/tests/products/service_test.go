package tests

import (
	"class2/afternoon/internal/products"
	"class2/afternoon/tests/products/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceUpdate(t *testing.T) {
	storeMock := mock.StoreMock{}
	repository := products.NewRepository(&storeMock)
	service := products.NewService(repository)

	expected := products.Product{
		Id:    1,
		Name:  "After Update",
		Price: 150,
	}

	product, err := service.Update(1, "After Update", 150)

	assert.Equal(t, expected, product)
	assert.True(t, storeMock.ReadWasCalled)
	assert.Nil(t, err)
}

func TestServiceDeleteWhenProductExists(t *testing.T) {
	storeMock := mock.StoreMock{}
	repository := products.NewRepository(&storeMock)
	service := products.NewService(repository)

	err := service.Delete(1)

	assert.NoError(t, err)
	assert.True(t, storeMock.ReadWasCalled)
	assert.True(t, storeMock.WriteWasCalled)
}

func TestServiceDeleteWhenProductNotExists(t *testing.T) {
	storeMock := mock.StoreMock{}
	repository := products.NewRepository(&storeMock)
	service := products.NewService(repository)

	err := service.Delete(2)

	assert.Error(t, err)
	assert.True(t, storeMock.ReadWasCalled)
	assert.False(t, storeMock.WriteWasCalled)
}
