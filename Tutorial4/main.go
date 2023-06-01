package main

import (
	"fmt"
	"math"
)

func pointers(){
	i,j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p /37
	fmt.Println(j)
}

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1,2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p = &Vertex{1,2}
)

func array(){
	var a[2]string
	a[0] ="hello"
	a[1]= "World"
	fmt.Println(a[0],a[1])
	fmt.Println(a)

	primes :=[6]int{2,3,5,7,11,13}
	fmt.Println(primes)
	var s []int = primes[:][1:4]
	fmt.Println(s)
}

func slices_pointers(){
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a:= names[0:2]
	b:= names[1:3]
	fmt.Println(a,b)

	b[0] = "XXX"
	fmt.Println(a,b)
	fmt.Println(names)
}

func slice_liter(){

}

func (v Vertex) GetLength() float64{
	var vX, vY = float64(v.X),float64(v.Y)
	return math.Round(math.Sqrt(vX*vX+vY*vY)*100)/100
}

func (v Vertex) updateX (x float64) float64{
	v.X = int(x)
	return float64(v.X)
}

func (v Vertex) updateY (y float64) float64{
	v.Y = int(y)
	return float64(v.Y)
}

func (v *Vertex) printResult () {
	fmt.Println("X: ",v.X)
	fmt.Println("Y: ",v.Y)
}

func slice_2nd(v []int) {
	v[1]=100
}
update
func add_slice(v []int) {
	v = append(v, 200)
}

func sqrt_slice(v []int64){
	for i:
}

func main() {
	// pointers()
	// fmt.Println(Vertex{1,2})
	// v := Vertex{1,2}af
	// v.X = 4
	// fmt.Println(v.X)
	// v := Vertex{1,2}
	// p := &v
	// p.X = 1e9
	// fmt.Println(v)
	// fmt.Println(v1,p,v2,v3)
	// array()
	// slices_pointers()
	// v := Vertex{3,5}
	// fmt.Println(v.GetLength())
	// v.printResult()
	// fmt.Println()
	// v.updateX(2)
	// fmt.Println(v.GetLength())
	// v.printResult()
	// fmt.Println()
	// v.updateY(6)
	// fmt.Println(v.GetLength())
	// v.printResult()
	// fmt.Println()
	v := []int{1,3,5}
	slice_2nd(v)
	fmt.Println(v)
	add_slice(v)
	fmt.Println(v)
}