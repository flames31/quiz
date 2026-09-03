package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	filename := getFileName()

	fmt.Println("Welcome to the Quiz!")

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	score, total := runQuiz(csv.NewReader(file))

	fmt.Printf("You scored: %v point(s) out of %v!\n", score, total)
}

func getFileName() string {
	filename := "problems.csv"
	if len(os.Args) > 1 {
		filename = os.Args[2]
	}
	return filename
}

func runQuiz(csvReader *csv.Reader) (int, int) {
	points, questions := 0, 0

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		q, a := record[0], record[1]
		fmt.Println("Question: ", q)
		var input string
		fmt.Scanln(&input)

		if a == input {
			points++
		}
		questions++
	}

	return points, questions
}
