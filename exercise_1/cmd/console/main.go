package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type problem struct {
	question string
	answer   string
}

var (
	file         = flag.String("file", "./problems.csv", "File path")
	timeout      = flag.Int("timeout", 5, "Time out in each question")
	resultSignal = make(chan string)
)

func parseProblems(records [][]string) []problem {
	problems := make([]problem, 0)
	for _, record := range records {
		problem := problem{
			question: record[0],
			answer:   record[1],
		}
		problems = append(problems, problem)
	}
	return problems
}

func displayResult(correct int, total int) {
	fmt.Printf("\nYou scored %d out of %d. \n", correct, total)
}

func askQuestion(number int, problem problem) {
	fmt.Printf("Problem # %d : %s = ", number, problem.question)
	var answerText string
	fmt.Scanf("%s\n", &answerText)
	if answerText == problem.answer {
		resultSignal <- "CORRECT"
	} else {
		resultSignal <- "INCORRECT"
	}
}

func main() {
	flag.Parse()

	data, err := os.Open(*file)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	reader := csv.NewReader(data)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	problems := parseProblems(records)
	total := len(problems)
	correct := 0

Loop:
	for index, problem := range problems {
		timer := time.NewTimer(time.Duration(*timeout) * time.Second)
		go askQuestion(index+1, problem)
		go func() {
			<-timer.C
			resultSignal <- "OVERTIME"
		}()
		resultText := <-resultSignal
		switch resultText {
		case "CORRECT":
			correct++
		case "INCORRECT":
			//do nothing
		default:
			timer.Stop()
			break Loop
		}
		timer.Stop()
	}
	displayResult(correct, total)
}
