package main

import (
	"flag"
	"log"
)

var (
	filepath = flag.String("filepath", "./problems.csv", "CSV File path")
	timeout  = flag.Int("timeout", 5, "Time to solve in each quiz")
)

func main() {
	app := &App{}
	err := app.ReadCSVfile(*filepath)
	if err != nil {
		log.Fatal(err)
	}
	app.AskQuizzes(*timeout)
}
