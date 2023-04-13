// This programm read in a quiz provided via a CSV file and give the quiz to a user,
// keeping track of how many questions they get right and how many they get incorrect
package main

import (
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]

	quiz, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error loading quizzes at %s!\n", filename)
		os.Exit(1)
	}

	fmt.Println(string(quiz))

	// for _, line := range strings.Split(string(data), "\n")
}
