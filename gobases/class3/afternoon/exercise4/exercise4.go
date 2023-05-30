package exercise4

import (
	"fmt"
	"math/rand"
	"time"
)

func Run() {
	insertionCollection1 := rand.Perm(100)
	insertionCollection2 := rand.Perm(1000)
	insertionCollection3 := rand.Perm(10000)

	bubleCollection1 := insertionCollection1
	bubleCollection2 := insertionCollection2
	bubleCollection3 := insertionCollection3

	selectionCollection1 := insertionCollection1
	selectionCollection2 := insertionCollection2
	selectionCollection3 := insertionCollection3

	execute(bubleCollection1, insertionCollection1, selectionCollection1)
	fmt.Println()
	execute(bubleCollection2, insertionCollection2, selectionCollection2)
	fmt.Println()
	execute(bubleCollection3, insertionCollection3, selectionCollection3)
}

func insertionSort(collection []int) {
	for i := 1; i < len(collection); i++ {
		for j := i; j > 0 && collection[j-1] > collection[j]; j-- {
			collection[j], collection[j-1] = collection[j-1], collection[j]
		}
	}
}

func selectionSort(collection []int) {
	for i := 0; i < len(collection); i++ {
		min := i
		for j := i + 1; j < len(collection); j++ {
			if collection[j] < collection[min] {
				min = j
			}
		}
		collection[i], collection[min] = collection[min], collection[i]
	}
}

func bubleSort(collection []int) {
	for i := 0; i < len(collection)-1; i++ {
		swapped := false

		for j := 0; j < len(collection)-i-1; j++ {
			if collection[j] > collection[j+1] {
				collection[j], collection[j+1] = collection[j+1], collection[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}

func swap(collection []int, position1 int, position2 int) {
	auxiliar := collection[position1]
	collection[position1] = collection[position2]
	collection[position2] = auxiliar
}

func worker(sortMethodName string, collection []int, sortMethod func([]int), channel chan time.Duration) {
	start := time.Now()
	sortMethod(collection)
	elapsed := time.Since(start)
	fmt.Printf("%s para coleção de %d itens: %v\n", sortMethodName, len(collection), elapsed)
	channel <- elapsed
}

func execute(bubleSortCollection []int, insertionSortCollection []int, selectionSortCollection []int) {
	channel := make(chan time.Duration)
	fmt.Printf("Iniciando algoritmos de ordenação para %d itens...\n", len(bubleSortCollection))
	go worker("Bubble Sort", bubleSortCollection, bubleSort, channel)
	go worker("Insertion Sort", insertionSortCollection, insertionSort, channel)
	go worker("Selection Sort", selectionSortCollection, selectionSort, channel)
	<-channel
	<-channel
	<-channel
	fmt.Printf("Fim da execução dos algoritmos de ordenação para %d itens...\n", len(bubleSortCollection))
}
