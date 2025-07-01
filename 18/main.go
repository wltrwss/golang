package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Book struct {
	Title  string `json:"Title"`
	Author string `json:"Author"`
	Year   int    `json:"Year"`
	Price  int    `json:"Price"`
}

func errFnMarshal(err *error) {
	if *err != nil {
		fmt.Println("=JSON MARSHALLING ERROR=")
	} else {
		fmt.Println("=JSON MARSHALLING COMPLETED SUCCESSFULLY=")
	}
}

func errFnWrite(err *error) {
	if *err != nil {
		fmt.Println("=JSON WRITING ERROR=")
	} else {
		fmt.Println("=JSON WRITING COMPLETED SUCCESSFULLY=")
	}
}

func errFnRead(err *error) {
	if *err != nil {
		fmt.Println("=JSON OPEN ERROR=")
	} else {
		fmt.Println("=JSON OPEN SUCCESSFULLY=")
	}
}

func errFnUnmarshal(err *error) {
	if *err != nil {
		fmt.Println("=JSON UNMARSHALING ERROR=")
	} else {
		fmt.Println("=JSON UNMARSHALING SUCCESSFULLY=")
	}
}

func main() {

	Books := []Book{
		{"Title-1", "Author-1", 1000, 100},
		{"Title-2", "Author-2", 2000, 200},
		{"Title-3", "Author-3", 3000, 300},
	}
	data, err := json.Marshal(Books)
	errFnMarshal(&err)
	err = os.WriteFile("library.json", data, 0644)
	errFnWrite(&err)

	raw, err := os.ReadFile("library.json")
	errFnRead(&err)
	err = json.Unmarshal(raw, &Books)
	errFnUnmarshal(&err)

	fmt.Printf("Library: %+v\n", Books)
}
