package services

import (
	"class2/controllers/models"
	repositories "class2/repositories/mocks"
)

func CreateProduct(request models.ProductRequest) {
	repositories.Save()
}
