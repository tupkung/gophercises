package main

import "fmt"

type Problem struct {
	Question   string
	Answer     string
	userAnswer string
	IsCorrect  bool
}

func (problem *Problem) askQuestion(number int) *Problem {
	fmt.Printf("Problem#%d %s=", number, problem.Question)
	return problem
}

func (problem *Problem) readUserAnswer() *Problem {
	var input string
	fmt.Scanf("%s\n", &input)
	problem.userAnswer = input
	return problem
}

func (problem *Problem) checkUserAnswer() *Problem {
	problem.IsCorrect = problem.userAnswer == problem.Answer
	return problem
}
