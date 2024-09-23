package main

import (
	"fmt"
)

func main() {

	people := map[string]int{
		"даша": 19,
		"катя": 20,
	}

	for n, a := range people {
		fmt.Println(n, a)
	}

	var n string
	var a int

	fmt.Println("\nвведите имя: ")

	fmt.Scan(&n)

	fmt.Println("введите возраст: ")
	fmt.Scan(&a)

	people[n] = a

	fmt.Println("\nобновленная карта: ")

	for n, a := range people {
		fmt.Println(n, a)
	}

}
