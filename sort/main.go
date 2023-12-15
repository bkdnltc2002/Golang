package main

import (
	"fmt"
	"slices"
)

func sortInt(slice []int) []int {
	slices.Sort(slice)
	return slice
}

func sortStr(slice []string) []string {
	slices.Sort(slice)
	return slice
}

func main() {
	intslice := []int{5, 6, 2, 4, 8}
	strslice := []string{"b", "e", "x", "i", "l"}
	fmt.Println(sortInt(intslice))
	fmt.Println(sortStr(strslice))
}
