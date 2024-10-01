package main

import "fmt"

func Loops() {
	start := 1
	things := []string{"maleta", "bolso", "ropa", "roljo"}
	fmt.Println(things)

	for i := 0; i < 10; i++ {
		fmt.Println(i + start)
	}

	for i := 0; i < len(things); i++ {
		fmt.Print(things[i] + " ")
	}

	for i := range things {
		fmt.Print(things[i] + " ")
	}
	for start < 100 {
		start += start
		if start == 32 {
			goto label
		}
		fmt.Println(start)
	}
label:
	fmt.Println("End of loop")

}
