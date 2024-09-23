package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	s := bufio.NewScanner(os.Stdin)
	sum := 0

	fmt.Println("введите числа: ")
	s.Scan()
	a := s.Text()

	words := strings.Split(a, " ")

	for _, word := range words {
		num, _ := strconv.Atoi(word)
		sum += num
	}

	fmt.Println("сумма:", sum)
}
