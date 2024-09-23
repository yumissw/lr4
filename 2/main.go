package main

import "fmt"

func func1(m map[string]int) float64 {

	sum := 0
	i := 0

	for _, a := range m {
		sum += a
		i++
	}

	return float64(sum) / float64(i)

	//return s
}

func main() {

	people := map[string]int{
		"даша":   19,
		"катя":   20,
		"лариса": 34,
	}

	fmt.Println("карта: ")

	for n, a := range people {
		fmt.Println(n, a)
	}

	fmt.Println("\nсредний возраст: ", func1(people))

}
