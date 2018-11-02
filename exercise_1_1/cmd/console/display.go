package main

import "fmt"

func displayResult(correct int, total int) {
	fmt.Printf("\nYou scored %d from %d\n", correct, total)
}

func displayQuestion(number int, question string) {
	fmt.Printf("Problem # %d %s = ", number, question)
}

func readUserAnswer() string {
	var input string
	fmt.Scanf("%s\n", &input)
	return input
}

func checkAnswer(answer string, userAnswer string) bool {
	return answer == userAnswer
}
