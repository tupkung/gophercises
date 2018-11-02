package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type App struct {
	Problems []*Problem
	Total    int
	Correct  int
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

	app.Problems = changeRecordsToProblems(records)
	app.Total = len(app.Problems)

	return nil
}

func (app *App) ExecuteProblem(timeout int) {
	timer := time.NewTimer(time.Duration(timeout) * time.Second)
	for index, problem := range app.Problems {
		select {
		case <-timer.C:
			app.displayResult()
			return
		default:
			if problem.
				askQuestion(index + 1).
				readUserAnswer().
				checkUserAnswer().
				IsCorrect {
				app.Correct++
			}
		}
	}
	app.displayResult()
}

func (app *App) displayResult() {
	fmt.Printf("\nYou scored %d from %d\n", app.Correct, app.Total)
}
