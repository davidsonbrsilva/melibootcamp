package main

import (
	"class1/exercise1"
	"class1/exercise2"
	"class1/exercise3"
	"class1/exercise4"
	"fmt"
)

func main() {
	exercise1.Run()
	exercise2.Run()
	exercise3.Run()
	exercise4.Run()

	slice := [3]int{3, 3, 3}

	fmt.Println(slice[1:])
}
