package main

import "fmt"

var i, j int = 1,2

func main(){
	var c, python, java = true, false, "no!"
	fmt.Println(i,j,c,python,java)
	Print("100",1)
}

func Print(x,y interface{}){
	fmt.Println(x,y)
}