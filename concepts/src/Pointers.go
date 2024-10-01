package main

import "fmt"

func Pointers() {

	var lifescore = 99
	lifeScorePointer := &lifescore

	lifescore *= 3

	fmt.Println(lifescore)
	fmt.Println(*lifeScorePointer)
}
