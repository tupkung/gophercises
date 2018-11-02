package main

import (
	"flag"
	"log"
)

var (
	filepath = flag.String("filepath", "./problems.csv", "Problem File path")
	overtime = flag.Int("overtime", 5, "Time to answer in each question")
)

func main() {
	flag.Parse()
	app := &App{}
	err := app.ReadCSVFile(*filepath)
	if err != nil {
		log.Fatal(err)
	}
	app.AskQuizzes(*overtime)
}
