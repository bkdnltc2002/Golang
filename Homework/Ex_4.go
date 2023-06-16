package main

import (
	"fmt"
)

type Student struct {
	ID   string
	Name string
}

var students []Student

func check_ID(s []Student) []string {
	idSet := make(map[string]bool)
	var result []string

	for _,student := range students {
		idSet[student.ID]=true
	}
	
	for _, student := range s {
		if idSet[student.ID] {
			result = append(result, student.ID)
			}
	}
	return result
}

func main() {

	students = append(students, Student{ID: "1", Name: "John"})
	students = append(students, Student{ID: "2", Name: "Daisy"})
	students = append(students, Student{ID: "3", Name: "Alex"})

	class := []Student{
		{ID: "1", Name: "James"},
		{ID: "4", Name: "Josh"},
		{ID: "3", Name: "Jennie"},
	}

	id_list := check_ID(class)
	if len(id_list) != 0 {
		fmt.Println("ID in Students list: ", id_list)
	} else {
		fmt.Println("There is no IDs in Students list: ", id_list)
	}

}
