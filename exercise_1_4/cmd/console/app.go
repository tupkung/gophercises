package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type App struct {
	Quizzes []*Quiz
	Total   int
	Correct int
}

func (app *App) ReadCSVFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}
	app.Quizzes = changeRecordsToQuizzes(records)
	app.Total = len(app.Quizzes)
	return nil
}

func (app *App) AskQuizzes(timeout int) {
	signal := make(chan string)
	for index, quiz := range app.Quizzes {
		timer := time.NewTimer(time.Duration(timeout) * time.Second)
		go func() {
			<-timer.C
			signal <- "OVERTIME"
		}()
		go func() {
			quiz.ExecuteCommand(
				ask(index+1),
				getAnswer,
			)
			if quiz.IsCorrect {
				app.Correct++
			}
			signal <- "ANSWER"
		}()
		switch result := <-signal; result {
		case "OVERTIME":
			app.showResult()
			return
		default:
			timer.Stop()
		}
	}
	app.showResult()
}

func (app *App) showResult() {
	fmt.Printf("\nYou score %d from %d\n", app.Correct, app.Total)
}

func ask(number int) func(quiz *Quiz) {
	return func(quiz *Quiz) {
		quiz.AskQuestion(number)
	}
}

func getAnswer(quiz *Quiz) {
	quiz.ReceiveUserAnswer()
}

func changeRecordsToQuizzes(records [][]string) []*Quiz {
	quizzes := make([]*Quiz, 0)

	for _, record := range records {
		quiz := &Quiz{
			Question: record[0],
			Answer:   record[1],
		}
		quizzes = append(quizzes, quiz)
	}

	return quizzes
}
