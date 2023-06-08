package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Student struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func main() {
	file, err := os.Open("students.json")
	defer file.Close()

	data, _ := io.ReadAll(file)

	var students []Student

	err = json.Unmarshal(data, &students)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
	}

	for i := range students {
		fmt.Println(students[i])
	}
}
