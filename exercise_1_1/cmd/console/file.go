package main

import (
	"encoding/csv"
	"os"
)

type Problem struct {
	Question string
	Answer   string
}

func fromCSVrecordsToStruct(records [][]string) []Problem {
	problems := make([]Problem, 0)
	for _, record := range records {
		problem := Problem{
			Question: record[0],
			Answer:   record[1],
		}
		problems = append(problems, problem)
	}
	return problems
}

func readCSVfile(filepath string) ([]Problem, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return make([]Problem, 0), err
	}
	reader := csv.NewReader(file)

	data, err := reader.ReadAll()
	if err != nil {
		return make([]Problem, 0), err
	}

	return fromCSVrecordsToStruct(data), nil
}
