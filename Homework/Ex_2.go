package main

import(
	"encoding/json"
	"fmt"
	"os"
)

type Student struct{
	Name string `json:"name"`
	Age int	`json:"age"`
	Gender string `json:"gender"`
	
}

func main(){
	students := []Student{
		{Name:"John",Age:15,Gender:"male"},
		{Name:"Daisy",Age:20,Gender:"female"},
		{Name:"Alex",Age:17,Gender:"female"},
	}
	
	file,err := os.Create("students.json")
	defer file.Close()

	encoder :=json.NewEncoder(file)
	err = encoder.Encode(students)
	if err!= nil{
		fmt.Println("Error encoding JSON:",err)
	}
	fmt.Println("Sucessfully converted")
}