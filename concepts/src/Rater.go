package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Rater() {

	var name string
	var UserRating string

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your name: ")
	name, _ = reader.ReadString('\n')

	reader = bufio.NewReader(os.Stdin)
	fmt.Println("please rate our dosa center b/w 1 - 5: ")
	UserRating, _ = reader.ReadString('\n')
	rating, _ := strconv.ParseFloat(strings.TrimSpace(UserRating), 64)

	fmt.Printf("Hello,%v\n Thanks for rating our dosa center with %v star rating.\n Your rating was recorded in our system at %v.\n\n", name, rating, time.Now().Format(time.Stamp))

	if rating == 5 {
		fmt.Println("Thanks for rating us 5 star,we are glad you loved our dosas.")
	} else if rating == 4 || rating == 3 {
		fmt.Println("we are always improving our dosa center,we will be glad to serve you better next time.")
	} else if rating < 3 {
		fmt.Println("we are sorry for the inconvenience caused,we will get in touch with you soon.")
	}

}
