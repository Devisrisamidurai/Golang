package main

import (
	"fmt"
)

func main() {
	slice := make([]string, 4)

	slice[0] = "apples"
	slice[1] = "oranges"
	slice[2] = "grapes"
	slice[3] = "mango"

	fmt.Println(slice)
}
