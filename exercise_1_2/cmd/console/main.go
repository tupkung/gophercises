package main

import (
	"flag"
	"log"
)

var (
	filepath = flag.String("filepath", "./problems.csv", "CSV File path")
	timeout  = flag.Int("timeout", 30, "Time to solve the problems")
)

func main() {
	app := &App{}
	err := app.ReadCSVfile(*filepath)
	if err != nil {
		log.Fatal(err)
	}
	app.ExecuteProblem(*timeout)
}
