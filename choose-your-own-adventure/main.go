package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fileName := flag.String("file", "gopher.json", "The JSON file with the story")
	flag.Parse()
	fmt.Printf("Using the story in %s\n", *fileName)

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
}
