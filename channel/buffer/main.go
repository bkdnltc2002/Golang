package main

import (
	"fmt"
)



func main() {
	defer func ()  {
		fmt.Println("Program end")
	}()
	ch := make(chan int)
	ch <- 14
	ch <- 20
}
