package main

import (
	"fmt"
	"sync"
)

type Form interface {
	int | uint32 | uint64
}

func Foo1(c chan int) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		c <- i
	}
}

func Foo2(c chan int) {
	defer wg.Done()
	for i := 11; i <= 20; i++ {
		c <- i
	}
}

func Foo3(c chan int) {
	defer wg.Done()
	for i := 21; i <= 30; i++ {
		c <- i
	}
}

func main() {
	c := make(chan int)
	wg.Add(3)
	go Foo1(c)
	go Foo2(c)
	go Foo3(c)
	for nums := range c {
		fmt.Println(nums)
	}
	wg.Wait()
	
}
