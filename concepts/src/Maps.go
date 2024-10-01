package main

import "fmt"

func Maps() {

	// new - only allocates memory
	// make - allocate memory and initializes it

	score := make(map[string]int)
	score["Kayal"] = 34
	score["sri"] = 59
	score["Deva"] = 65
	score["Devi"] = 43
	fmt.Println(score)

	getKayalscore := score["Kayal"]
	fmt.Println(getKayalscore)

	delete(score, "Deva")
	fmt.Println(score)

	for a, b := range score {
		fmt.Printf("score of %v is %v\n", a, b)
	}

}
