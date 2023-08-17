package main

import (
	"fmt"
)

func main(){
	slice := make([]int,3,6)
	slice[0]=1
	slice[1]=2
	slice[2]=3
	fmt.Println(len(slice),cap(slice))
	slice1 := slice
	slice = append(slice, 0)
	slice = append(slice, 5)
	slice = append(slice, 7)
	slice = append(slice, 11)
	fmt.Println(len(slice),cap(slice))
	slice[0]= 20
	fmt.Println(slice)
	fmt.Println(slice1)
}