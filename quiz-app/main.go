package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
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
	problems := parseRecords(records)
	fmt.Println(problems)
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
			a: record[1],
		}
	}
	return result
}

type Problem struct {
	q string
	a string
}
