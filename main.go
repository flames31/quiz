package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type Quiz struct {
	total     int
	score     int
	questions []Question
}

type Question struct {
	question string
	answer   string
}

func main() {
	filename := flag.String("file", "problems.csv", "file to read")
	timeInSeconds := flag.Int("time", 60, "time in seconds")

	fmt.Println("Welcome to the Quiz!")

	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	quiz := getQuiz(csv.NewReader(file))

	runQuizWithTimer(quiz, *timeInSeconds)

	fmt.Printf("You scored: %v point(s) out of %v!\n", quiz.score, quiz.total)
}

func getFileName() string {
	filename := "problems.csv"
	if len(os.Args) > 2 && os.Args[1] == "-file" {
		filename = os.Args[2]
	}
	return filename
}

func getQuiz(reader *csv.Reader) *Quiz {
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	quiz := &Quiz{
		total:     len(records),
		score:     0,
		questions: make([]Question, len(records)),
	}

	for i, record := range records {
		quiz.questions[i] = Question{
			question: record[0],
			answer:   record[1],
		}
	}
	return quiz
}

func runQuiz(quiz *Quiz) {
	for _, question := range quiz.questions {
		fmt.Println("Question: ", question.question)
		var input string
		fmt.Scanln(&input)

		if question.answer == input {
			quiz.score++
		}
	}
}

func runQuizWithTimer(quiz *Quiz, timeInSeconds int) {

	done := make(chan bool, 1)
	go func() {
		runQuiz(quiz)
		done <- true
	}()
	select {
	case <-done:
		fmt.Println("Quiz finished!")
	case <-time.After(time.Duration(timeInSeconds) * time.Second):
		fmt.Println("Time is up!")
	}
}
