package main

import (
    "fmt"
    "sort"
    "strings"
)

// Student struct to hold the student's name.
type Student struct {
    Name string
}

func main() {
    // Creating a slice of students.
    students := []Student{
        {Name: "Charlie Brown"},
        {Name: "Alice Johnson"},
        {Name: "Bob Smith"},
    }

    // Sorting the students by their first name.
    // The sort.Slice function is flexible and allows for custom sorting logic.
    sort.Slice(students, func(i, j int) bool {
        // Splitting the names by space and comparing the first element (first name).
        return strings.Split(students[i].Name, " ")[0] < strings.Split(students[j].Name, " ")[0]
    })

    // Printing the sorted list of students.
    fmt.Println("Students sorted by first name:")
    for _, student := range students {
        fmt.Println(student.Name)
    }
}

