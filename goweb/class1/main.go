package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/davidsonbrsilva/ginq"
	"github.com/gin-gonic/gin"
)

type Product struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Color        string    `json:"color"`
	Price        float64   `json:"price"`
	Stock        int       `json:"stock"`
	Code         string    `json:"code"`
	IsPublished  bool      `json:"isPublished"`
	CreationDate time.Time `json:"creationDate"`
}

func main() {
	router := gin.Default()

	router.GET("/hello", sayHello)
	router.GET("/products", getAllProducts)
	router.GET("/products/:id", getProductById)

	router.Run()
}

func sayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello!",
	})
}

func getAllProducts(c *gin.Context) {
	products := getProductRepository(c)

	products = filterBy("name", products, c)
	products = filterBy("color", products, c)
	products = filterBy("price", products, c)
	products = filterBy("stock", products, c)
	products = filterBy("code", products, c)
	products = filterBy("isPublished", products, c)
	products = filterBy("creationDate", products, c)

	c.JSON(http.StatusOK, products)
}

func getProductById(c *gin.Context) {
	products := getProductRepository(c)

	var product *Product = nil
	for _, p := range products {
		if strconv.Itoa(p.Id) == c.Param("id") {
			product = &p
			break
		}
	}

	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "product not found",
		})
	}

	c.JSON(http.StatusOK, *product)
}

func getProductRepository(c *gin.Context) []Product {
	file, err := os.Open("products.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error recovering data",
		})
	}
	defer file.Close()

	data, _ := ioutil.ReadAll(file)
	var products []Product
	json.Unmarshal(data, &products)
	return products
}

func filterBy(fieldName string, products []Product, context *gin.Context) []Product {
	if context.Query(fieldName) == "" {
		return products
	}

	return ginq.
		From(products).
		Where(func(p Product) bool {
			return getFieldValueByField(fieldName, &p) == context.Query(fieldName)
		}).
		Select()
}

func getFieldValueByField(fieldName string, product *Product) string {
	fields := reflect.TypeOf(*product)

	for index := 0; index < fields.NumField(); index++ {
		field := fields.Field(index)
		tag := field.Tag.Get("json")

		if tag == fieldName {
			return reflect.ValueOf(*product).Field(index).String()
		}
	}

	return ""
}
