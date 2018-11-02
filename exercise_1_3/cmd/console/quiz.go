package main

import "fmt"

type Quiz struct {
	Question   string
	Answer     string
	userAnswer string
	IsCorrect  bool
}

func (quiz *Quiz) askQuestion(number int) *Quiz {
	fmt.Printf("Problem#%d %s=", number, quiz.Question)
	return quiz
}

func (quiz *Quiz) receiveUserAnswer() *Quiz {
	var input string
	fmt.Scanf("%s\n", &input)
	quiz.userAnswer = input
	quiz.IsCorrect = quiz.Answer == input
	return quiz
}
