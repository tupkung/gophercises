package main

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
