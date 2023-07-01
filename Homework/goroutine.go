package main

import (
	"fmt"
	"time"
)

type Form interface {
	int | uint32 | uint64
}

func Foo1() {
	for i := 1; i <= 10; i++ {
		go fmt.Println(i)
	}
}

func Foo2() {
	for i := 11; i <= 20; i++ {
		go fmt.Println(i)
	}
}

func Foo3() {
	for i := 21; i <= 30; i++ {
		go fmt.Println(i)
	}
}

func main() {
	go Foo1()
	go Foo2()
	go Foo3()
	time.Sleep(time.Second)
}
