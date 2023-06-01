package main

import (
	"fmt"
	"math"
)

func sum(){
	sum := 0
	for i:=0; i<10;i++ {
		sum += i
	}
	fmt.Println(sum)
}

func sumv2(){
	sum:=1
	for ; sum<1000; {
		sum +=sum
	}
	fmt.Println(sum)
}

func sumv3(){
	sum :=1
	for sum <1000{
		sum+=sum
	}
	fmt.Println(sum)
}

func sqrt(x float64) string{
	if x< 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x,n,lim float64) float64{
	if v:= math.Pow(x,n); v<lim{
		return v
	}
	return lim
}

func powv2(x,n,lim float64) float64{
	if v:=math.Pow(x,n); v<lim{
		return v
	} else {
		fmt.Printf("%g >= %g\n",v,lim)
	}
	return lim
}

func Sqrt(x float64) float64{

}

func main(){
	// sum()
	// sumv2()
	// sumv3()
	// fmt.Println(sqrt(2),sqrt(-4))
	// fmt.Println(
	// 	pow(3,2,10),
	// 	pow(3,3,20),
	// )
	// fmt.Println(
	// 	powv2(3,2,10),
	// 	powv2(3,3,20),
	// )
	fmt.Println(Sqrt(2))
}