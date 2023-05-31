package main

import (
	"fmt"
	// "time"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = l<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	// fmt.Println("Bay gio la: ",time.Now())
	// fmt.Println(i,j,k,c,python,java)
	fmt.Printf("Type : %T Value: %v\n", ToBe)
}
