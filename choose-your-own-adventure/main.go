package main

import (
	"flag"
	"fmt"
)

func main() {
	fileName := flag.String("file", "gopher.json", "The JSON file with the story")
	flag.Parse()
	fmt.Printf("Using the story in %s\n", *fileName)
}
