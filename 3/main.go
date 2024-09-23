package main

import (
	"fmt"
)

func main() {

	people := map[string]int{
		"даша": 19,
		"катя": 20,
		"соня": 19,
	}

	for n, a := range people {
		fmt.Println(n, a)
	}

	var n string
	var i int

	fmt.Println("\nвведите имя: ")

	fmt.Scan(&n)

	for key := range people {
		if key == n {
			i++
		}
	}

	if i == 0 {
		fmt.Println("записи ", n, " нет в карте ")
	} else {
		delete(people, n)
		fmt.Println("\nобновленная карта: ")

		for n, a := range people {
			fmt.Println(n, a)
		}
	}

}
