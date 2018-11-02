package main

import (
	"flag"
	"log"
	"time"
)

var (
	filePath = flag.String("path", "./problems.csv", "CSV Filepath")
	timeout  = flag.Int("timeout", 30, "Timeout to solve the questions")
)

func main() {
	problems, err := readCSVfile(*filePath)
	if err != nil {
		log.Fatal(err)
	}

	timer := time.NewTimer(time.Duration(*timeout) * time.Second)
	total := len(problems)
	correct := 0
	for index, problem := range problems {
		select {
		case <-timer.C:
			displayResult(correct, total)
			return
		default:
			displayQuestion((index + 1), problem.Question)
			if checkAnswer(problem.Answer, readUserAnswer()) {
				correct++
			}
		}
	}
	displayResult(correct, total)
}
