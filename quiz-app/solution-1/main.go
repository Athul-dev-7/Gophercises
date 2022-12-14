package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
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
	fmt.Println(records)

	// Converting 2D slice to Slice of struct's
	problems := parseRecords(records)
	fmt.Println(problems)

	// Logic for getting user inputs & calculating correct answers
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d : %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s", &answer)
		if answer == p.a {
			correct++
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
