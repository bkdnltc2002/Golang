package main

import (
	"fmt"
	"slices"
 
)

func contain(v int, s []int)bool{
	return slices.Contains(s,v)
}

func containv2(a []int, b []int)[]int{
	result := []int{}
	bmap :=make(map[int]struct{})

	for _, val := range b {
		bmap[val] = struct{}{}
	}

	for _, val := range a{
		if _, ok := bmap[val]; ok{
			result = append(result, val)
		}
	}
	return result
}

func main(){
	s := []int {1,2,3,4,5}
	s2:= []int {4,6,3,8,9}
	v:= 6
	fmt.Println(contain(v,s))
	fmt.Println(containv2(s,s2))
}