package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	CYOA "choose-your-own-adventure/story"
)

func main() {
	fileName := flag.String("file", "gopher.json", "The JSON file with the story")
	flag.Parse()
	fmt.Printf("Using the story in %s\n", *fileName)

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	story, err := CYOA.JSONStory(file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", story)
	// PrettyPrint(story)
}

func PrettyPrint(v interface{}) (err error) {
	result, err := json.MarshalIndent(v, "", " ")
	if err == nil {
		fmt.Println(string(result))
	}
	return
}
