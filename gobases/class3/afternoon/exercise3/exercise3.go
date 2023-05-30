package exercise3

import "fmt"

func Run() {
	maintenances := MaintenanceCollection{
		{name: "Troca de óleo", price: 50},
		{name: "Troca de filtro de ar", price: 30},
		{name: "Pintura de escapamento", price: 120},
	}

	services := ServiceCollection{
		{name: "Troca de lâmpada", price: 20, workedMinutes: 25},
		{name: "Montagem de armário", price: 80, workedMinutes: 60},
		{name: "Instalação Elétrica", price: 150, workedMinutes: 120},
	}

	products := ProductCollection{
		{name: "Lata de Tinta", price: 250, amount: 1},
		{name: "Lixa de parede", price: 3, amount: 5},
		{name: "Massa Corrida", price: 50, amount: 1},
	}

	collections := []ICollection{
		products,
		services,
		maintenances,
	}

	channel := make(chan float64)
	result := 0.0

	for _, collection := range collections {
		go worker(collection, channel)
	}

	for index := 0; index < len(collections); index++ {
		result += <-channel
	}

	fmt.Println("Preço total de todos os serviçoes é:", result)
}

func worker(collection ICollection, channel chan float64) {
	result := collection.sum()
	fmt.Printf("Preço total de %T: %f\n", collection, result)
	channel <- collection.sum()
}
