package main

import(
	"fmt"
	"sort"
)

func sort_slice(s []int) []int  {

	sort.Ints(s)
	// n:= len(s)
	// for i:=0;i<n-1;i++{
	// 	for j:=0;j<n-i-1;j++{
	// 		if s[j]>s[j+1]{
	// 			s[j],s[j+1]=s[j+1],s[j]
	// 		}
	// 	}
	// }
	return s
}

func main(){
	numbers :=[]uint{5,2,7,1,9,19,34,23,54,24,64,67,53}
	fmt.Println(sort_slice(numbers))
}