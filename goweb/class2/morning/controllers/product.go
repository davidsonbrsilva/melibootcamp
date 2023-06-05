package controllers

import (
	"class2/controllers/models"
	services "class2/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var request models.ProductRequest
	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	services.CreateProduct(request)
	c.JSON(http.StatusCreated, request)
}
