package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	filename := "problems.csv"

	if len(os.Args) > 1 {
		filename = os.Args[2]
	}
	fmt.Println("Welcome to the Quiz!")
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	points, questions := 0, 0

	for {
		record, err := reader.Read()
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

	fmt.Printf("You scored: %v point(s) out of %v!\n", points, questions)
}
