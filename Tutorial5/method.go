package main

import(
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64{
	if f<0{
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct{
	X, Y float64
}

func (v Vertex) Abs()float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64){
	v.X=v.X*f
	v.Y=v.Y*f
}

func ScaleFunc(v *Vertex, f float64){
	v.X = v.X*f
	v.Y = v.Y*f
}

func AbsFunc(v Vertex) float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Abser interface {
	Abs() float64
}

type I interface{
	M()
}

type T struct{
	S string
}

func (t T) M(){
	fmt.Println(t.S)
}

func main(){
	// v:= &Vertex{3,4}
	// fmt.Printf("Before scaling: %+v, Abs: %v\n",v,v.Abs())
	// v.Scale(5)
	// fmt.Printf("After scaling: %+v, Abs: %v\n",v,v.Abs())
	// ScaleFunc(&v, 10)
	// fmt.Println(v.Abs())
	// fmt.Println(AbsFunc(v))

	// p:= &Vertex{4,3}
	// fmt.Println(p.Abs())
	// fmt.Println(AbsFunc(*p))
	// p.Scale(3)
	// ScaleFunc(p,8)
	// fmt.Println(v,p)
	// f := MyFloat(-math.Sqrt2)
	// fmt.Println(f.Abs())
	// var a Abser
	// f := MyFloat(-math.Sqrt2)
	// v := Vertex{3,4}

	// a = f
	// a = &v

	// a=v

	// fmt.Println(a.Abs())

	var i I = T{"hello"}
	i.M()
}