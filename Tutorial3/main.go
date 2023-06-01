package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sum() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func sumv2() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func sumv3() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func powv2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func Sqrt(x float64) float64 {
	z := float64(1)
	for math.Abs(x-(z*z)) > 0.1 {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func switch_func() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func switchv2() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today. ")
	case today + 1:
		fmt.Println("Tomorrow. ")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away. ")
	}
}

func switchv3() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func defer_func() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

func defer_multi() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
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
	// fmt.Println(Sqrt(200))
	// switch_func()
	// switchv2()
	// switchv3()
	// defer_func()
	defer_multi()
}
