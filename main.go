package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type question struct {
	question string
	answer   string
}

func main() {
	filename := flag.String("file", "problems.csv", "file to read")
	timeInSeconds := flag.Int("time", 60, "time in seconds")
	flag.Parse()

	fmt.Println("Welcome to the Quiz!")

	questions, err := loadQuestions(*filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	score := runQuizWithTimer(questions, time.Duration(*timeInSeconds)*time.Second)

	fmt.Printf("You scored: %d point(s) out of %d!\n", score, len(questions))
}

func loadQuestions(filename string) ([]question, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	questions := make([]question, 0, len(records))

	for i, record := range records {
		if len(record) < 2 {
			return nil, fmt.Errorf("invalid record on line %d", i+1)
		}

		questions = append(questions, question{
			question: record[0],
			answer:   record[1],
		})
	}

	return questions, nil
}

func runQuiz(questions []question, score chan<- int) {
	correct := 0

	for _, q := range questions {
		fmt.Printf("Question: %s\n", q.question)

		var answer string
		fmt.Scanln(&answer)

		if answer == q.answer {
			correct++
		}
	}

	score <- correct
}

func runQuizWithTimer(questions []question, limit time.Duration) int {

	score := make(chan int, 1)

	go runQuiz(questions, score)

	timer := time.NewTimer(limit)
	defer timer.Stop()

	select {
	case score := <-score:
		fmt.Println("Quiz finished!")
		return score
	case <-timer.C:
		fmt.Println("Time is up!")
		return 0
	}
}
