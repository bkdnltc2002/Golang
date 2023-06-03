package main

import (
	"fmt"
	"math"
	"strings"
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
	q := []int{2,3,5,7,11,13}
	fmt.Println(q)

	r := []bool{true,false,true,true,false,true}
	fmt.Println(r)

	s := []struct{
		i int
		b bool
	}{
		{2,true},
		{3, false},
		{5,true},
		{7,true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func slice_bounds(){
	s := []int{2,3,5,7,11,13}

	s = s[1:4]
	fmt.Println(s)
	
	s = s[:2]
	fmt.Println(s)

	s= s[1:]
	fmt.Println(s)
}

func slice_len_cap(){
	s := []int{2,3,5,7,11,13}
	// printSlice(s)

	s = s[:4]
	// printSlice(s)

	s = s[2:]
	// printSlice(s)
}

func printSlice(s []int){
	fmt.Printf("%s len=%d cap=%d %v\n",len(s),cap(s),s)
}

func nil_slices(){
	var s[]int
	fmt.Println(s, len(s), cap(s))
	if s == nil{
		fmt.Println("nil!")
	}
}

// func making_slices(){
// 	a := make([]int,5)
// 	// printSlice("a",a)

// 	b:= make([]int,0,5)
// 	// printSlice("b",b)

// 	c:=b[:2]
// 	// printSlice("c",c)

// 	d:= c[2:5]
// 	// printSlice("d",d)
// }



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

func add_slice(v []int) {
	v = append(v, 200)
}

func slice_slice(){
	board := [][]string{
		[]string{"_","_","_"},
		[]string{"_","_","_"},
		[]string{"_","_","_"},
	}

	board[0][0]="X"
	board[2][2]="O"
	board[1][2]="X"
	board[1][0]="O"
	board[0][2]="X"

	for i :=0; i<len(board); i++{
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appends(){
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2,3,4)
	printSlice(s)
}

func ranges(){
	var pow = []int{1,2,4,8,16,32,64,128}
	for i, v:= range pow{
		fmt.Printf("2**%d = %d\n",i,v)
	}
}

func range_continued(){
	pow := make([]int,10)
	for i:= range pow{
		pow[i] = 1 << uint(i) 
	}
	for _, value := range pow{
		fmt.Printf("%d\n", value)
	}
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
	// v := []int{1,3,5}
	// slice_2nd(v)
	// fmt.Println(v)
	// add_slice(v)
	// fmt.Println(v)
	// slice_liter()
	// slice_bounds()
	// slice_len_cap()
	// nil_slices()
	// making_slices()
	// slice_slice()
	// appends()
	// ranges()
	// range_continued()
	
}