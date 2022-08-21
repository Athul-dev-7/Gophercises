package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Reading csv file
	csvFile, err := os.Open("./problems.csv")
	if err != nil {
		checkNilErr(fmt.Sprintf("Failed to open the csv file %v\n", err))
	}
	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		checkNilErr("Failed to read the csv file")
	}

	// Converting 2D slice to Slice of struct's
	problems := parseRecords(records)

	// Setting timer
	timer := time.NewTimer(time.Second * time.Duration(3))

	// Logic for getting user inputs & calculating correct answers
	correct := 0
problemLogic:
	for i, p := range problems {
		fmt.Printf("Problem #%d : %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Println("")
			break problemLogic
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}

	}
	fmt.Printf("You scored %d out of %d\n", correct, len(problems))

}

func checkNilErr(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func parseRecords(records [][]string) []Problem {
	result := make([]Problem, len(records))
	for i, record := range records {
		result[i] = Problem{
			q: record[0],
			a: strings.TrimSpace(record[1]),
		}
	}
	return result
}

type Problem struct {
	q string
	a string
}
