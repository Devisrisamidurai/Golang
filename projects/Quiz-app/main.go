package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	// Input the file name
	fName := flag.String("f", "Quiz.csv", "path of csv file")
	// Set timer
	timer := flag.Int("t", 30, "timer for the quiz")
	flag.Parse()

	// Pull the problems (calling problemPuller function)
	problems, err := problemPuller(*fName)
	if err != nil {
		exit(fmt.Sprintf("Something went wrong: %s", err.Error()))
	}

	// Create a variable to count correct answers
	correctAns := 0
	// Initialize timer
	tObj := time.NewTimer(time.Duration(*timer) * time.Second)
	ansC := make(chan string)

	// Loop through the problems, print the questions, accept answers
problemLoop:
	for i, p := range problems {
		var answer string
		fmt.Printf("Problem %d: %s= ", i+1, p.q)

		go func() {
			fmt.Scanln(&answer)
			ansC <- answer
		}()

		select {
		case <-tObj.C:
			fmt.Println()
			break problemLoop
		case iAns := <-ansC:
			if iAns == p.a {
				correctAns++
			}
		}
	}

	// Calculate the result
	fmt.Printf("Your result is %d out of %d\n", correctAns, len(problems))
	fmt.Println("Press enter to exit")
	fmt.Scanln() // Wait for user input before exiting
}

func problemPuller(fileName string) ([]problem, error) {
	// Open the file
	if fObj, err := os.Open(fileName); err == nil {
		// Reader instance
		csvR := csv.NewReader(fObj)
		// Read the file
		if cLines, err := csvR.ReadAll(); err == nil {
			return parseProblems(cLines), nil
		} else {
			return nil, fmt.Errorf("error in reading CSV format from %s file: %s", fileName, err.Error())
		}
	} else {
		return nil, fmt.Errorf("error in opening %s file: %s", fileName, err.Error())
	}
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func parseProblems(lines [][]string) []problem {
	// Parse lines into problem struct
	r := make([]problem, len(lines))
	for i := 0; i < len(lines); i++ {
		r[i] = problem{q: lines[i][0], a: lines[i][1]}
	}
	return r
}
