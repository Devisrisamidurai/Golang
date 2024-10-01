package main

import (
	"fmt"
	"sort"
)

func Slices() {
	var things = []string{"maleta", "ropa", "reloj"}
	fmt.Println(things)

	things = append(things, "bolso")
	fmt.Println(things)

	things = things[1:]
	fmt.Println(things)

	heros := make([]string, 3, 3)
	heros[0] = "thor"
	heros[1] = "ironman"
	heros[2] = "spiderman"

	fmt.Println(heros)
	heros = append(heros, "deadpool")

	fmt.Println(heros)

	arr := []int{3, 4, 5}
	issorted := sort.IntsAreSorted(arr)
	fmt.Println(issorted)
	sort.Ints(arr)
	fmt.Println(arr)

}
