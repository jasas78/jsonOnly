package main

import (
	"fmt"
	"os"
)

var (
	_jsonFile *os.File
    _filenameJson string = "1/1.json"
)

func _readJsonFile() {
	// Open our _jsonFile
	_jsonFile, ___err := os.Open( _filenameJson )

	// if we os.Open returns an error then handle it
	if ___err != nil {
		fmt.Printf("Opened <%s> failed\n" , _filenameJson )
		fmt.Println(___err)
		os.Exit(33)
	}

	fmt.Printf("Successfully Opened <%s>.\n" , _filenameJson )
	// defer the closing of our _jsonFile so that we can parse it later on
	defer _jsonFile.Close()
}

func main() {
	/*
	   fmt.Println("Hello World")
	*/
	_readJsonFile()
}
