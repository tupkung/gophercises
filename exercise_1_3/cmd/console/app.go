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

func (app *App) ReadCSVfile(filepath string) error {
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
			if quiz.askQuestion(index + 1).
				receiveUserAnswer().
				IsCorrect {
				app.Correct++
			}

			signal <- "ANSWERED"
		}()
		go func() {
			<-timer.C
			signal <- "TIMEOUT"
		}()
		switch <-signal {
		case "TIMEOUT":
			app.displayResult()
			return
		default:
			timer.Stop()
		}
	}
	app.displayResult()
}

func (app *App) displayResult() {
	fmt.Printf("\nYou score %d from %d", app.Correct, app.Total)
}
