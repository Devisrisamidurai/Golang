package main

import "fmt"

func Arrays() {

	var numbers [3]string
	numbers[0] = "uno"
	numbers[1] = "dos"
	numbers[2] = "tris"

	fmt.Println(numbers)

	var colors = [4]string{"azul", "verde", "rojo", "gris"}
	fmt.Println(colors)
	fmt.Println(colors[3])
	fmt.Println(len(colors))
}
