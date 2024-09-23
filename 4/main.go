package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("введите строку: ")
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Println(strings.ToUpper(s))
}
