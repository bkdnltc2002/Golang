package main

import (
	"errors"
	"fmt"
	"log"
)

func devidebyone1(num float64) (float64, error) {
	if num == 0 {
		err := errors.New("divide zero")
		return 0, err
	}
	return 1 / num, nil
}

func devidebyone2(num float32) (float32, error) {
	var (
		result float32
		err    error
	)

	defer func() {
		log.Println(result, err)
	}()

	if num == 0 {
		result = 0
		err = errors.New("divide zero")
		return result, err
	}
	if num < 0 {
		result = 0
		err = errors.New("do not support negative number")
		return result, err
	}
	result = 1 / num
	return result, nil
}

func devidebyone3(num float32) (float32, error) {
	var (
		result float32
		err    error
	)

	defer func() {
		log.Println(result, err)
	}()

	if num == 0 {
		result = 0
		err = errors.New("divide zero")
		return result, err
	}
	if num < 0 {
		result = 0
		err = errors.New("do not support negative number")
		return result, err
	}
	result = 1 / num
	return result, nil
}

func devidebyone4(num int32) (int32, error) {
	var (
		result int32
		err    error
	)

	defer func ()  {
		e:=recover()
		if e!=nil{
			log.Println(e)
		} 
	}()
	// defer func() {
	// 	log.Println(result, err)
	// }()

	if num < 0 {
		result = 0
		err = errors.New("do not support negative number")
		return result, err
	}
	result = 1/num
	return 1 / num, nil
}

func main() {
	fmt.Println(devidebyone4(0))
	devidebyone3(20)
	// devidebyone3(0)
	// devidebyone2(float32(math.Pow(10, -40)))
}
