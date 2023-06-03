package main

import (
	"fmt"
	
)

type Vertex struct{
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

var n = map[string]Vertex{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func mutating_map(){
	m:= make(map[string]int)

	m["Answer"]= 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"]=48
	fmt.Println("The value:", m["Answer"])

	delete(m,"Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok:=m["Answer"]
	fmt.Println("The value:", v, "Present?",ok)
}

func WordCount(s string) map[string]int{
	return map[string]int{"x":1}
}

func compute(fn func(float64,float64) float64)float64{
	return fn(3,4)
}

func adder() func(int) int{
	sum:=0
	return func(x int) int {
		sum+=x
		return sum
	}
}

func fibonacci() func() int{
	prev, curr := 0,1
	return func() int {
		result := prev
		prev, curr = curr, prev+curr
		return result
	}
}

func main(){
	// fmt.Println(n)
	// mutating_map()
	// wc.Test(WordCount("Vietnam"))
	// hypot := func (x,y float64) float64  {
	// 	return math.Sqrt(x*x +y*y)
	// }
	// fmt.Println(hypot(5,12))

	// fmt.Println(compute(hypot))
	// fmt.Println(compute(math.Pow))
	// pos, neg := adder(),adder()
	f :=fibonacci()
	for  i := 0;  i < 10;  i++ {
		// fmt.Println(
		// 	pos(i),
		// 	neg(-2*i),
		// )
		fmt.Println(f())
	}
}