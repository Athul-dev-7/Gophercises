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
}

func checkNilErr(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

type Problem struct {
	q string
	a string
}
