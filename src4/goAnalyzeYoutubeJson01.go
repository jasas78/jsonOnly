package main

import (
	"fmt"
	"os"
)

var (
	_jsonFile *os.File
)

func _readJsonFile() {
	// Open our _jsonFile
	_jsonFile, ___err := os.Open("1/1.json")

	// if we os.Open returns an error then handle it
	if ___err != nil {
		fmt.Println("Opened 1/1.json failed")
		fmt.Println(___err)
		os.Exit(33)
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our _jsonFile so that we can parse it later on
	defer _jsonFile.Close()
}

func main() {
	/*
	   fmt.Println("Hello World")
	*/
	_readJsonFile()
}
