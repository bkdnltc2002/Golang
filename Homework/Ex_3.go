package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	data, _ := ioutil.ReadAll(file)

	var students []Student

	err = json.Unmarshal(data, &students)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
	}

	for _, student := range students {
		fmt.Println(student)
	}
}
