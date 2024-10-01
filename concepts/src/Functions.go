package main

import "fmt"

func functions() {
	fmt.Println(adder(3, 4))
	fmt.Println(sentence("Hello, I am devi"))
	myresult, mylength, myname := total(1, 2, 3, 4)
	fmt.Println(myresult, mylength, myname)
}

func adder(a, b int) int {
	return a + b
}
func sentence(str string) string {
	return str
}
func total(values ...int) (int, int, string) {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	name := "Devi"
	return sum, len(values), name
}
