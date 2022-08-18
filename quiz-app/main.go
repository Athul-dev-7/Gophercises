package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	csvFile, err := os.Open("./problems.csv")
	checkNilErr(err)
	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	checkNilErr(err)
	fmt.Println(records)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
