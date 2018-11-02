package main

func changeRecordsToProblems(records [][]string) []*Problem {
	problems := make([]*Problem, 0)
	for _, record := range records {
		problem := &Problem{
			Question: record[0],
			Answer:   record[1],
		}
		problems = append(problems, problem)
	}
	return problems
}
