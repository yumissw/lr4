package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("введите количество элементов: ")
	fmt.Scan(&n)

	a := make([]int, n)
	b := make([]int, n)

	fmt.Println("введите элементы массива:")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	fmt.Println("\nмассив: ")

	for i := 0; i < n; i++ {
		fmt.Print(a[i], " ")
	}

	var j int = 0

	for i := n - 1; i >= 0; i-- {
		b[j] = a[i]
		j++
	}

	fmt.Println("\nперевернутый массив: ")

	for i := 0; i < n; i++ {
		fmt.Print(b[i], " ")
	}

}
