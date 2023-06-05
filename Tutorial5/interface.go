package main

import (
	"fmt"
	"strings"
	"math"
	"time"
)

type I interface{
	M()
}

type T struct{
	S string
}

// func (t *T) M(){
// 	fmt.Println(t.S)
// }

type F float64

func (t *T) M(){
	if t == nil{
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i I){
	fmt.Printf("(%v, %T)\n",i,i)
}

func do(i interface{}){
	switch v:= i.(type){
	case int:
		 fmt.Printf("twice %v is %v\n",v,v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n",v,len(v))
	default:
		fmt.Printf("I don't know about type %T\n",v)
	}
}

type Person struct {
	Name string
	Age  int
}

func (i IPaddr) String() string {
	IPaddrString := []string{}
	for _,num:= range i{
		IPaddrString = append(IPaddrString, fmt.Sprint(int(num)))
	}
	return strings.Join(IPaddrString, ".")
}

type IPaddr [4]byte

type MyError struct{
	When time.Time
	What string
}

func (e ErrNegtiveSqrt) Error() string{
	return fmt.Sprintf("cannot Sqrt negative number: %v",float64(e))
}

// func run() error{
// 	return &MyError{
// 		time.Now(),
// 		"it didn't work",
// 	}
// }

type ErrNegtiveSqrt float64

func Sqrt(x float64)(float64,error){
	if x<0{
		return 0,ErrNegtiveSqrt(x)
	}
	return math.Sqrt(x),nil
}

func main(){
	// var i I

	// var t *T
	// i = &T{"Hello"}
	// describe(i)
	// i.M()

	// // i = F(math.Pi)
	// i=t
	// describe(i)
	// i.M()
	// var i interface{} = "hello"

	// s := i.(string)
	// fmt.Println(s)

	// s, ok:=i.(string)
	// fmt.Println(s, ok)

	// f, ok := i.(float64)
	// fmt.Println(f,ok)

	// f = i.(float64)
	// fmt.Println(f)
	// do(21)
	// do("hello")
	// do(true)
	// a := Person{"Arthur Dent", 42}
	// z :=Person{"Zaphod Beeblebrox",90010}
	// fmt.Println(a,z)
	// hosts := map[string]IPaddr{
	// 	"loopback": {127,0,0,1},
	// 	"googleDNS": {8,8,8,8},
	// }
	// for name, ip := range hosts{
	// 	fmt.Printf("%v: %v\n",name,ip)
	// }
	// if err := run(); err !=nil{
	// 	fmt.Println(err)
	// }
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}