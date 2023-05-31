package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func packages() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

func imports() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

func exported_names() {
	fmt.Println(math.Pi)
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var i, j int = 1, 2

func zero() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func type_conversions() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

func constants() {
	const Pi = 3.14
	const World = "Vietnam"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// packages()
	// imports()
	// exported_names()
	// fmt.Println(add(42, 13))
	// a, b := swap("hello","world")
	// fmt.Println(a,b)
	// fmt.Println(split(17))
	// var i int
	// fmt.Println(i,c,python,java)
	// var c,python, java = true, false, "no!"
	// fmt.Println(i,j,c,python,java)
	// var i,j int = 1,2
	// k := 3
	// c, python, java := true, false, "no!"
	// fmt.Println(i,j,k,c,python,java)
	// fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	// fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	// fmt.Printf("Type: %T Value: %v\n", z, z)
	// zero()
	// type_conversions()
	// constants()
	// fmt.Println(needInt(Small))
	// fmt.Println(needFloat(Small))
	// fmt.Println(needFloat(Big))
}
