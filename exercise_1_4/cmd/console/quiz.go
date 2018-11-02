package main

import "fmt"

type Quiz struct {
	Question   string
	Answer     string
	userAnswer string
	IsCorrect  bool
}

func (quiz *Quiz) AskQuestion(number int) {
	fmt.Printf("Problem#%d %s=", number, quiz.Question)
}

func (quiz *Quiz) ReceiveUserAnswer() {
	var input string
	fmt.Scanf("%s\n", &input)
	quiz.userAnswer = input
	quiz.IsCorrect = quiz.Answer == quiz.userAnswer
}

func (quiz *Quiz) ExecuteCommand(options ...func(*Quiz)) {
	for _, option := range options {
		option(quiz)
	}
}
